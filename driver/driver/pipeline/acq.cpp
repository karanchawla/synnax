
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

#include <thread>
#include "driver/pipeline/acq.h"
#include "driver/errors/errors.h"
#include "nlohmann/json.hpp"
#include "acq.h"

#pragma once

using json = nlohmann::json;

//using namespace pipeline;

Acq::Acq(){}

Acq::Acq(synnax::WriterConfig writer_config,
         std::shared_ptr<synnax::Synnax> client,
         std::unique_ptr<daq::AcqReader> daq_reader):
         writer_config(writer_config),
         client(client),
         daq_reader(std::move(daq_reader)) {}

void Acq::start() {
    running = true;
    acq_thread = std::thread(&Acq::run, this);
}

void Acq::stop() {
    running = false;
    acq_thread.join(); // FIXME: I dont want to call join im p sure (elham)
}

void Acq::run() {
    // start daq read
    auto dq_err = daq_reader->start();
    if (!dq_err.ok()) { // daq read error
        if (dq_err.type == TYPE_TRANSIENT_HARDWARE_ERROR && breaker->wait()) run();
        return;
    }

    // start synnax writer
    auto [writer, wo_err] = client->telem.openWriter(writer_config);
    if (!wo_err.ok()) { // synnax write error
        daq_reader->stop();
        if (wo_err.type == freighter::TYPE_UNREACHABLE && breaker->wait()) run();
        return;
    }

    bool retry = false;
    while (running) {
        // perform a daq read
        auto [frame, error] = daq_reader->read();
        if (!error.ok()) {
            // Any other type means we've encountered a critical hardware failure
            // or configuration error and can't proceed.
            retry = error.type == TYPE_TRANSIENT_HARDWARE_ERROR;
        }

//        std::cout << "holy fucking shit fuckface fuck" << std::endl;
//        auto s =  frame.series->at(0).uint64();

//        for (int j = 0; j < s.size(); j++){
//            std::cout << s[j] << ", ";
//        }

//        std::cout << "Writing" << std::endl;
        if (!writer.write(std::move(frame))) { // write frame to channel

            auto err = writer.error();
            if (!err.ok()) {
                retry = error.type == freighter::TYPE_UNREACHABLE;
                std::cout << "failed to write" << std::endl;
                break;
            }
        }
        // synnax commit
        auto now = synnax::TimeStamp::now();

//        std::cout << "committing" << std::endl;
        if (now - last_commit > commit_interval) {
            auto [end, ok] = writer.commit();
            auto err = writer.error();
            if (!ok) {
                retry = error.type == freighter::TYPE_UNREACHABLE;
                std::cout << "committing failed" << std::endl;
                break;
            }
//            std::cout << "end: " << end.value << std::endl;
            last_commit = now;
        }
    }
//    std::cout  << "shutting things down" << std::endl;
    daq_reader->stop();
    auto err = writer.close();
//    std::cout << "Acq run error: " << err.message() << std::endl;
    if (retry && breaker->wait()) run();
}



