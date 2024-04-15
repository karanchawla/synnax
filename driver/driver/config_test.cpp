// Copyright 2024 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

#include <gtest/gtest.h>

#include "driver.h"
#include "driver/driver/driver.h"
#include "nlohmann/json.hpp"

using json = nlohmann::json;

/// @brief it should correctly apply defaults for an empty configuration.
TEST(TestDriverConfig, parseEmptyConfig) {
    auto [cfg, err] = driver::parseConfig(json::object());
    ASSERT_FALSE(err) << err;
    ASSERT_EQ(cfg.client_config.host, "localhost");
    ASSERT_NEAR(cfg.breaker_config.scale, 1.2, 0.0001);
}

TEST(TestDriverConfig, testValidConfig) {
    json config = {
        {"connection", {
            {"host", "demo.synnaxlabs.com"},
            {"port", 80},
            {"username", "admin"},
            {"password", "admin"},
            {"ca_cert_file", "ca.pem"},
            {"client_cert_file", "client.pem"},
            {"client_key_file", "client.key"}
        }},
        {"retry", {
            {"base_interval", 2},
            {"max_retries", 100},
            {"scale", 1.5}
        }},
        {"rack", {
            {"key", 1},
            {"name", "rack_1"}
        }},
        {"integrations", {"opc"}}
    };
    auto [cfg, err] = driver::parseConfig(config);
    ASSERT_FALSE(err) << err;
    ASSERT_EQ(cfg.client_config.host, "demo.synnaxlabs.com");
    ASSERT_EQ(cfg.client_config.port, 80);
    ASSERT_EQ(cfg.client_config.username, "admin");
    ASSERT_EQ(cfg.client_config.password, "admin");
    ASSERT_EQ(cfg.client_config.ca_cert_file, "ca.pem");
    ASSERT_EQ(cfg.client_config.client_cert_file, "client.pem");
    ASSERT_EQ(cfg.client_config.client_key_file, "client.key");
    ASSERT_EQ(cfg.breaker_config.base_interval, synnax::SECOND * 2);
    ASSERT_EQ(cfg.breaker_config.max_retries, 100);
    ASSERT_NEAR(cfg.breaker_config.scale, 1.5, 0.0001);
    ASSERT_EQ(cfg.rack_key, 1);
    ASSERT_EQ(cfg.rack_name, "rack_1");
    ASSERT_EQ(cfg.integrations.size(), 1);
    ASSERT_EQ(cfg.integrations[0], "opc");
}
