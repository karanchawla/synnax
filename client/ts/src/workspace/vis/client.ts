// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { sendRequired, type UnaryClient } from "@synnaxlabs/freighter";
import { toArray, unknownRecordZ, type UnknownRecord } from "@synnaxlabs/x";
import { z } from "zod";

const RETRIEVE_ENDPOINT = "/workspace/vis/retrieve";
const CREATE_ENDPOINT = "/workspace/vis/create";
const DELETE_ENDPOINT = "/workspace/vis/delete";
const RENAME_ENDPOINT = "/workspace/vis/rename";
const SET_DATA_ENDPOINT = "/workspace/vis/set-data";

export const keyZ = z.string().uuid();
export type Key = z.infer<typeof keyZ>;
export type Params = Key | Key[];

export const visZ = z.object({
  key: z.string(),
  name: z.string(),
  type: z.string(),
  data: unknownRecordZ.or(z.string().transform((s) => JSON.parse(s) as UnknownRecord)),
});

export type Vis = z.infer<typeof visZ>;

const retrieveReqZ = z.object({ keys: z.string().array() });

const resZ = z.object({ vis: visZ.array() });

export const newVisZ = visZ.partial({ key: true }).transform((p) => ({
  ...p,
  data: JSON.stringify(p.data),
}));

export type NewVisZ = z.input<typeof newVisZ>;

const createReqZ = z.object({
  workspace: z.string(),
  vis: newVisZ.array(),
});

const createResZ = z.object({
  vis: visZ.array(),
});

const deleteReqZ = z.object({
  keys: keyZ.array(),
});

const deleteResZ = z.object({});

const renameReqZ = z.object({
  key: keyZ,
  name: z.string(),
});

const renameResZ = z.object({});

const setDataReqZ = z.object({
  key: keyZ,
  data: z.string(),
});

const setDataResZ = z.object({});

export class Client {
  private readonly client: UnaryClient;

  constructor(client: UnaryClient) {
    this.client = client;
  }

  async create(workspace: string, vis: NewVisZ): Promise<Vis>;

  async create(workspace: string, vis: NewVisZ[]): Promise<Vis[]>;

  async create(workspace: string, vis: NewVisZ | NewVisZ[]): Promise<Vis | Vis[]> {
    const isSingle = !Array.isArray(vis);
    const res = await sendRequired<typeof createReqZ, typeof createResZ>(
      this.client,
      CREATE_ENDPOINT,
      { workspace, vis: toArray(vis) },
      createReqZ,
      createResZ,
    );
    return isSingle ? res.vis[0] : res.vis;
  }

  async rename(key: Key, name: string): Promise<void> {
    await sendRequired<typeof renameReqZ, typeof renameResZ>(
      this.client,
      RENAME_ENDPOINT,
      { key, name },
      renameReqZ,
      renameResZ,
    );
  }

  async setData(key: Key, data: UnknownRecord): Promise<void> {
    await sendRequired<typeof setDataReqZ, typeof setDataResZ>(
      this.client,
      SET_DATA_ENDPOINT,
      { key, data: JSON.stringify(data) },
      setDataReqZ,
      setDataResZ,
    );
  }

  async retrieve(keys: Key): Promise<Vis>;

  async retrieve(keys: Key[]): Promise<Vis[]>;

  async retrieve(keys: Params): Promise<Vis | Vis[]> {
    const isSingle = !Array.isArray(keys);
    const res = await sendRequired(
      this.client,
      RETRIEVE_ENDPOINT,
      { keys: toArray(keys) },
      retrieveReqZ,
      resZ,
    );
    return isSingle ? res.vis[0] : res.vis;
  }

  async delete(keys: Params): Promise<void> {
    await sendRequired<typeof deleteReqZ, typeof deleteResZ>(
      this.client,
      DELETE_ENDPOINT,
      { keys: toArray(keys) },
      deleteReqZ,
      deleteResZ,
    );
  }
}
