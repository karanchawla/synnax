// Copyright 2024 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

#include <set>
#include "include/open62541/client_highlevel.h"
#include "glog/logging.h"
#include "driver/driver/opc/reader.h"
#include "driver/driver/opc/util.h"
#include "driver/driver/config/config.h"
#include "include/open62541/client_config_default.h"
#include "include/open62541/client_subscriptions.h"
#include "include/open62541/types.h"
#include "include/open62541/plugin/log_stdout.h"
#include "driver/driver/pipeline/acquisition.h"


using namespace opc;

ReaderConfig::ReaderConfig(
    config::Parser &parser
): device(parser.required<std::string>("device")) {
    sample_rate = Rate(parser.required<std::float_t>("sample_rate"));
    stream_rate = Rate(parser.required<std::float_t>("stream_rate"));
    parser.iter("channels", [&](config::Parser &channel_builder) {
        auto ch = ReaderChannelConfig(channel_builder);
        if (ch.enabled) channels.push_back(ch);
    });
}

class ReaderSource final : public pipeline::Source {
public:
    ReaderConfig cfg;
    std::shared_ptr<UA_Client> client;
    std::set<ChannelKey> indexes;

    ReaderSource(
        ReaderConfig cfg,
        const std::shared_ptr<UA_Client> &client,
        std::set<ChannelKey> indexes
    )
        : cfg(std::move(cfg)), client(client), indexes(std::move(indexes)) {
    }

    std::pair<Frame, freighter::Error> read() override {
        auto fr = Frame(cfg.channels.size() + indexes.size());
        std::this_thread::sleep_for(cfg.sample_rate.period().nanoseconds());
        for (const auto &ch: cfg.channels) {
            if (!ch.enabled) continue;
            UA_Variant *value = UA_Variant_new();
            const UA_StatusCode status = UA_Client_readValueAttribute(
                client.get(), ch.node, value);
            if (status != UA_STATUSCODE_GOOD)
                LOG(ERROR) << "Unable to read value from opc server";
            else {
                const auto val = val_to_series(value, ch.ch.data_type);
                fr.add(ch.channel, val);
            }
            UA_Variant_delete(value);
        }
        const auto now = synnax::TimeStamp::now();
        for (const auto &idx: indexes) fr.add(idx, Series(now));
        return std::make_pair(std::move(fr), freighter::NIL);
    }
};


Reader::Reader(
    const std::shared_ptr<task::Context> &ctx,
    synnax::Task task
): ctx(ctx), task(task) {
    // Step 1. Parse the configuration to ensure that it is valid.
    auto config_parser = config::Parser(task.config);
    cfg = ReaderConfig(config_parser);
    if (!config_parser.ok()) {
        LOG(ERROR) << "[OPC Reader] failed to parse configuration for " << task.name;
        ctx->setState({
            .task = task.key,
            .variant = "error",
            .details = config_parser.error_json(),
        });
        return;
    }

    LOG(INFO) << "[OPC Reader] successfully parsed configuration for " << task.name;

    auto [device, dev_err] = ctx->client->hardware.retrieveDevice(cfg.device);
    if (dev_err) {
        LOG(ERROR) << "[OPC Reader] failed to retrieve device " << cfg.device <<
                " error: " << dev_err.message();
        ctx->setState({
            .task = task.key,
            .variant = "error",
            .details = json{
                {"message", dev_err.message()}
            }
        });
        return;
    }

    auto properties_parser = config::Parser(device.properties);
    auto properties = DeviceProperties(properties_parser);


    auto breaker_config = breaker::Config{
        .name = task.name,
        .base_interval = 1 * SECOND,
        .max_retries = 20,
        .scale = 1.2,
    };
    breaker = breaker::Breaker(breaker_config);

    // Fetch additional index channels we also need as part of the configuration.
    auto [res, err] = retrieveAdditionalChannelInfo();
    if (err) {
        ctx->setState({
            .task = task.key,
            .variant = "error",
            .details = json{{"message", err.message()}}
        });
        return;
    }
    auto [channelKeys, indexes] = res;

    // Connect to the OPC UA server.
    auto [ua_client, conn_err] = opc::connect(properties.connection);
    if (conn_err) {
        ctx->setState({
            .task = task.key,
            .variant = "error",
            .details = {{"message", conn_err.message()}}
        });
        return;
    }

    for (auto i = 0; i < cfg.channels.size(); i++) {
        auto ch = cfg.channels[i];
        UA_Variant *value = UA_Variant_new();
        const UA_StatusCode status = UA_Client_readValueAttribute(
            ua_client.get(),
            ch.node,
            value
        );
        if (status != UA_STATUSCODE_GOOD) {
            if (status == UA_STATUSCODE_BADNODEIDUNKNOWN) {
                config_parser.field_err("channels." + std::to_string(i),
                                        "opc node not found");
            } else {
                config_parser.field_err("channels." + std::to_string(i),
                                        "failed to read value" + std::string(
                                            UA_StatusCode_name(status)));
            }
            LOG(ERROR) << "failed to read value for channel " << ch.node_id;
        }
        UA_Variant_delete(value);
    }

    if (!config_parser.ok()) {
        ctx->setState({
            .task = task.key,
            .variant = "error",
            .details = config_parser.error_json(),
        });
        return;
    }

    auto source = std::make_unique<ReaderSource>(
        cfg,
        ua_client,
        indexes
    );

    auto writer_cfg = synnax::WriterConfig{
        .channels = channelKeys,
        .start = TimeStamp::now(),
        .mode = synnax::WriterPersistStream
    };

    pipe = pipeline::Acquisition(
        ctx,
        writer_cfg,
        std::move(source),
        breaker_config
    );
    ctx->setState({
        .task = task.key,
        .variant = "success",
        .details = {
            {"running", false}
        }
    });
}

void Reader::exec(task::Command &cmd) {
    if (cmd.type == "start") {
        pipe.start();
        return ctx->setState({
            .task = task.key,
            .variant = "success",
            .details = {
                {"running", true}
            }
        });
    }
    if (cmd.type == "stop") {
        pipe.stop();
        return ctx->setState({
            .task = task.key,
            .variant = "success",
            .details = {
                {"running", false}
            }
        });
    }
    LOG(ERROR) << "unknown command type: " << cmd.type;
}
