// Copyright 2024 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

#pragma once

#include "opc.h"
#include "util.h"
#include "driver/config/config.h"
#include "driver/task/task.h"
#include "driver/pipeline/acquisition.h"

namespace opc {
struct ReaderChannelConfig {
    /// @brief the node id.
    std::string node_id;
    UA_NodeId node;
    /// @brief the corresponding channel key to write the variable for the node from.
    ChannelKey channel;
    /// @brief the channel fetched from the Synnax server. This does not need to
    /// be provided via the JSON configuration.
    Channel ch;
    bool enabled;

    ReaderChannelConfig() = default;

    explicit ReaderChannelConfig(
        config::Parser &parser
    ): node_id(parser.required<std::string>("node_id")),
       node(parseNodeId("node_id", parser)),
       channel(parser.required<ChannelKey>("channel")),
        enabled(parser.optional<bool>("enabled", true)) {
    }
};

struct ReaderConfig {
    /// @brief the device representing the OPC UA server to read from.
    std::string device;
    /// @brief sets the acquisition rate.
    Rate sample_rate;
    /// @brief sets the stream rate.
    Rate stream_rate;
    /// @brief array_size;
    size_t array_size;

    /// @brief the list of channels to read from the server.
    std::vector<ReaderChannelConfig> channels;

    ReaderConfig() = default;

    explicit ReaderConfig(config::Parser &parser);

    std::vector<ChannelKey> channelKeys() const {
        auto keys = std::vector<ChannelKey>(channels.size());
        for (std::size_t i = 0; i < channels.size(); i++) keys[i] = channels[i].channel;
        return keys;
    }
};

/// @brief a task that reads values from an OPC UA server.
class Reader final : public task::Task {
public:
    explicit Reader(
        const std::shared_ptr<task::Context> &ctx,
        synnax::Task task,
        ReaderConfig cfg,
        const breaker::Config &breaker_config,
        std::shared_ptr<pipeline::Source> source,
        synnax::WriterConfig writer_config
    ): ctx(ctx),
       task(std::move(task)),
       cfg(std::move(cfg)),
       breaker(breaker::Breaker(breaker)),
       pipe(pipeline::Acquisition(ctx, writer_config, source, breaker_config)) {
    }

    static std::unique_ptr<task::Task> configure(
        const std::shared_ptr<task::Context> &ctx,
        const synnax::Task &task
    );

    void exec(task::Command &cmd) override;
    
    void stop();

    void start();


private:
    std::shared_ptr<task::Context> ctx;
    synnax::Task task;
    ReaderConfig cfg;
    breaker::Breaker breaker;
    pipeline::Acquisition pipe;
};
}
