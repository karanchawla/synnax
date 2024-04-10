// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { describe, test, expect } from "vitest";

import { newClient } from "@/setupspecs";

const client = newClient();

const ZERO_UUID = "00000000-0000-0000-0000-000000000000";

describe("Vis", () => {
  describe("create", () => {
    test("create one", async () => {
      const ws = await client.workspaces.create({
        name: "PID",
        layout: { one: 1 },
      });
      const vis = await client.workspaces.vis.create(ws.key, {
        name: "PID",
        type: "pid",
        data: { one: 1 },
      });
      expect(vis.name).toEqual("PID");
      expect(vis.key).not.toEqual(ZERO_UUID);
      expect(vis.data.one).toEqual(1);
    });
  });
  describe("rename", () => {
    test("rename one", async () => {
      const ws = await client.workspaces.create({
        name: "PID",
        layout: { one: 1 },
      });
      const linePlot = await client.workspaces.vis.create(ws.key, {
        name: "PID",
        type: "pid",
        data: { one: 1 },
      });
      await client.workspaces.vis.rename(linePlot.key, "PID2");
      const res = await client.workspaces.vis.retrieve(linePlot.key);
      expect(res.name).toEqual("PID2");
    });
  });
  describe("setData", () => {
    test("set data", async () => {
      const ws = await client.workspaces.create({
        name: "PID",
        layout: { one: 1 },
      });
      const linePlot = await client.workspaces.vis.create(ws.key, {
        name: "PID",
        type: "pid",
        data: { one: 1 },
      });
      await client.workspaces.vis.setData(linePlot.key, { two: 2 });
      const res = await client.workspaces.vis.retrieve(linePlot.key);
      expect(res.data.two).toEqual(2);
    });
  });
  describe("delete", () => {
    test("delete one", async () => {
      const ws = await client.workspaces.create({
        name: "PID",
        layout: { one: 1 },
      });
      const linePlot = await client.workspaces.vis.create(ws.key, {
        name: "PID",
        type: "pid",
        data: { one: 1 },
      });
      await client.workspaces.vis.delete(linePlot.key);
      await expect(client.workspaces.vis.retrieve(linePlot.key)).rejects.toThrow();
    });
  });
});
