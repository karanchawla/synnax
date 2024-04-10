// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { useCallback } from "react";

import { type SynnaxProps } from "@synnaxlabs/client";
import { Drift } from "@synnaxlabs/drift";
import { useAsyncWindowLifecycle, useSelectWindowKey } from "@synnaxlabs/drift/react";
import { Status, useAsyncEffect, useSyncedRef } from "@synnaxlabs/pluto";
import { TimeStamp } from "@synnaxlabs/x";
import { path } from "@tauri-apps/api";
import { Child, Command } from "@tauri-apps/api/shell";
import { useDispatch } from "react-redux";

import { useSelectLocalState } from "@/cluster/selectors";
import {
  setLocalState,
  set,
  LOCAL_CLUSTER_KEY,
  setActive,
  LOCAL_PROPS,
  LOCAL,
} from "@/cluster/slice";
import { testConnection } from "@/cluster/testConnection";

// The name of the sidecar binary.
const BINARY_NAME = "bin/sy";
export const LOCAL_KEY = "local";

export const useLocalServer = (): void => {
  const win = useSelectWindowKey();

  const d = useDispatch();
  const { schematic, command } = useSelectLocalState();
  const status = Status.useAggregator();
  const schematicRef = useSyncedRef(schematic);

  const startLocalServer = async (): Promise<void> => {
    if (win !== Drift.MAIN_WINDOW) return;
    // The only case where we'll run into a stranded Schematic is if the user closes the
    // application or hard reloads the page. This means that we only need to kill
    // stranded schematics on application load, so we don't pass the Schematic in as a dependency.
    if (schematic !== 0) {
      console.log("Killing stranded local server", schematic);
      await new Child(schematic).kill();
    }

    const dataPath = (await path.homeDir()) + "/.synnax/console/synnax-data";
    const command = Command.sidecar(BINARY_NAME, [
      "start",
      "-i",
      "-l",
      "localhost:9090",
      "-d",
      dataPath,
    ]);

    const handleLog = (v: string): void => {
      // All of our logs are JSON parseable.
      const { level, msg, error } = JSON.parse(v);

      const isInfo = level === "info";
      // This means the server has booted up.
      if (isInfo && msg === "starting server") {
        // Set the Schematic in local state so we can kill it later fi we need to.

        // Test the connection to the local server.
        testConnection(LOCAL_PROPS)
          .then(() => {
            d(set(LOCAL));
            d(setLocalState({ schematic: serverProcess.schematic, status: "running" }));
            d(setActive(LOCAL_KEY));
          })
          .catch(console.error);
      } else if (isInfo && msg === "shutdown successful") {
        // If the server has shut down, we'll set the Schematic to 0.
        d(setLocalState({ schematic: 0, status: "stopped" }));
      }

      // If the server fails to boot up, we'll get a fatal error.
      if (level === "fatal")
        status({
          time: TimeStamp.now(),
          variant: "error",
          message: error,
          key: "local-server",
        });
    };

    const handleClose = (): void => {
      d(setLocalState({ schematic: 0, status: "stopped" }));
    };

    command.stderr.on("data", handleLog);
    command.on("close", handleClose);
    const serverProcess = await command.spawn();

    d(setLocalState({ schematic: serverProcess.schematic, status: "starting" }));
  };

  const stopLocalServer = useCallback(async (): Promise<void> => {
    if (schematicRef.current === 0) return;
    d(setLocalState({ schematic, status: "stopping" }));
    const serverProcess = new Child(schematicRef.current);
    await serverProcess.write("stop\n");
    d(setActive(null));
  }, []);

  useAsyncEffect(async () => {
    if (win !== Drift.MAIN_WINDOW) return;
    if (command === "start") return await startLocalServer();
    if (command === "stop") return await stopLocalServer();
  }, [win, command]);

  useAsyncWindowLifecycle(async () => {
    return async () => await stopLocalServer();
  });
};
