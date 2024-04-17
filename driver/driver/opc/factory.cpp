// Copyright 2024 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

#include "glog/logging.h"
#include "driver/driver/opc/opc.h"
#include "driver/driver/opc/scanner.h"
#include "driver/driver/opc/reader.h"

std::pair<std::unique_ptr<task::Task>, bool> opc::Factory::configureTask(
    const std::shared_ptr<task::Context> &ctx,
    const synnax::Task &task
) {
    if (task.type == "opcScanner")
        return {std::make_unique<Scanner>(ctx, task), true};
    if (task.type == "opcReader")
        return {std::make_unique<Reader>(ctx, task), true};
    return {nullptr, false};
}

std::vector<std::pair<synnax::Task, std::unique_ptr<task::Task> > >
opc::Factory::configureInitialTasks(
    const std::shared_ptr<task::Context> &ctx,
    const synnax::Rack &rack
) {
    std::vector<std::pair<synnax::Task, std::unique_ptr<task::Task> > > tasks;
    auto [existing, err] = rack.tasks.list();
    if (err) {
        LOG(ERROR) << "[OPC] Failed to list existing tasks: " << err;
        return tasks;
    }
    // check if a task with the same type and name already exists
    bool hasScanner = false;
    for (const auto &t: existing) {
        if (t.type == "opcScanner") {
            LOG(INFO) << "[OPC] found existing scanner task with key: " << t.key << "skipping creation." << std::endl;
            hasScanner = true;
        }
    }

    if (!hasScanner) {
        auto sy_task = synnax::Task(
            rack.key,
            "opc Scanner",
            "opcScanner",
            ""
        );
        auto err= rack.tasks.create(sy_task);
        LOG(INFO) << "[OPC] created scanner task with key: " << sy_task.key;
        if (err) {
            LOG(ERROR) << "[OPC] Failed to create scanner task: " << err;
            return tasks;
        }
        auto [task, ok] = configureTask(ctx, sy_task);
        auto pair = std::make_pair(sy_task, std::move(task));
        if (ok) tasks.emplace_back(std::move(pair));
    }
    return tasks;
}