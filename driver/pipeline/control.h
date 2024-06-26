// Copyright 2024 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

#pragma once

#include <memory>
#include <thread>
#include "client/cpp/synnax.h"
#include "driver/task/task.h"
#include "driver/breaker/breaker.h"

namespace pipeline {
class Sink {
public:
    virtual ~Sink() = default;
    virtual freighter::Error write(synnax::Frame frame) = 0;
    virtual freighter::Error start() = 0;
    virtual freighter::Error stop() = 0;
};

class Control {
public:

    Control() = default;

    Control(
        std::shared_ptr<task::Context> ctx,
        synnax::StreamerConfig streamer_config,
        std::unique_ptr<Sink> sink,
        const breaker::Config &breaker_config
    );

    ~Control(){
        LOG(INFO) << "Control destructor called";
    }

    
    void start();
    void stop();

private:
    std::shared_ptr<task::Context> ctx;

    /// @brief writer thread.
    std::unique_ptr<std::thread> thread;

    /// @brief synnax writer
    std::unique_ptr<synnax::Streamer> streamer;
    synnax::StreamerConfig streamer_config;

    /// @brief synnax writer
    synnax::WriterConfig writer_config;

    /// @brief daq interface
    std::unique_ptr<Sink> sink;

    /// @brief breaker
    breaker::Breaker breaker;

    void runInternal();

    void run();
};
}