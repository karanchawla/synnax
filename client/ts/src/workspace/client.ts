// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { sendRequired, type UnaryClient } from "@synnaxlabs/freighter";
import {
  type UnknownRecord,
  type AsyncTermSearcher,
  toArray,
  unknownRecordZ,
} from "@synnaxlabs/x";
import { z } from "zod";

import { vis } from "@/workspace/vis";

export const workspaceKeyZ = z.string().uuid();

export type Key = z.infer<typeof workspaceKeyZ>;

export type Params = Key | Key[];

export const workspaceZ = z.object({
  name: z.string(),
  key: workspaceKeyZ,
  layout: unknownRecordZ.or(
    z.string().transform((s) => JSON.parse(s) as UnknownRecord),
  ),
});

export const workspaceRemoteZ = workspaceZ.omit({ layout: true }).extend({
  layout: z.string().transform((s) => JSON.parse(s) as UnknownRecord),
});

export type Workspace = z.infer<typeof workspaceZ>;

const retrieveReqZ = z.object({
  keys: workspaceKeyZ.array().optional(),
  search: z.string().optional(),
  author: z.string().uuid().optional(),
  offset: z.number().optional(),
  limit: z.number().optional(),
});

const retrieveResZ = z.object({
  workspaces: workspaceRemoteZ.array(),
});

const newWorkspaceZ = workspaceZ.partial({ key: true }).transform((w) => ({
  ...w,
  layout: JSON.stringify(w.layout),
}));

export type NewWorkspace = z.input<typeof newWorkspaceZ>;

const createReqZ = z.object({
  workspaces: newWorkspaceZ.array(),
});

const createResZ = z.object({
  workspaces: workspaceRemoteZ.array(),
});

const deleteReqZ = z.object({
  keys: workspaceKeyZ.array(),
});

const deleteResZ = z.object({});

const renameReqZ = z.object({
  key: workspaceKeyZ,
  name: z.string(),
});

const renameResZ = z.object({});

const setLayoutReqZ = z.object({
  key: workspaceKeyZ,
  layout: z.unknown().transform((u) => JSON.stringify(u)),
});

const setLayoutResZ = z.object({});

export type CreateResponse = z.infer<typeof createResZ>;

const CREATE_ENDPOINT = "/workspace/create";
const DELETE_ENDPOINT = "/workspace/delete";
const RENAME_ENDPOINT = "/workspace/rename";
const SET_LAYOUT_ENDPOINT = "/workspace/set-layout";
const RETRIEVE_ENDPOINT = "/workspace/retrieve";

export class Client implements AsyncTermSearcher<string, Key, Workspace> {
  readonly vis: vis.Client;
  private readonly client: UnaryClient;

  constructor(client: UnaryClient) {
    this.client = client;
    this.vis = new vis.Client(client);
  }

  async search(term: string): Promise<Workspace[]> {
    return (
      await sendRequired(
        this.client,
        RETRIEVE_ENDPOINT,
        { search: term },
        retrieveReqZ,
        retrieveResZ,
      )
    ).workspaces;
  }

  async retrieve(key: Key): Promise<Workspace>;

  async retrieve(keys: Key[]): Promise<Workspace[]>;

  async retrieve(keys: Key | Key[]): Promise<Workspace | Workspace[]> {
    const isMany = Array.isArray(keys);
    const res = await sendRequired(
      this.client,
      RETRIEVE_ENDPOINT,
      { keys: toArray(keys) },
      retrieveReqZ,
      retrieveResZ,
    );
    return isMany ? res.workspaces : res.workspaces[0];
  }

  async page(offset: number, limit: number): Promise<Workspace[]> {
    return (
      await sendRequired(
        this.client,
        RETRIEVE_ENDPOINT,
        { offset, limit },
        retrieveReqZ,
        retrieveResZ,
      )
    ).workspaces;
  }

  async create(workspace: NewWorkspace): Promise<Workspace>;

  async create(
    workspaces: NewWorkspace | NewWorkspace[],
  ): Promise<Workspace | Workspace[]> {
    const isMany = Array.isArray(workspaces);
    const res = await sendRequired<typeof createReqZ, typeof createResZ>(
      this.client,
      CREATE_ENDPOINT,
      { workspaces: toArray(workspaces) },
      createReqZ,
      createResZ,
    );
    return isMany ? res.workspaces : res.workspaces[0];
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

  async setLayout(key: Key, layout: UnknownRecord): Promise<void> {
    await sendRequired<typeof setLayoutReqZ, typeof setLayoutResZ>(
      this.client,
      SET_LAYOUT_ENDPOINT,
      { key, layout },
      setLayoutReqZ,
      setLayoutResZ,
    );
  }

  async delete(...keys: Key[]): Promise<void> {
    await sendRequired<typeof deleteReqZ, typeof deleteResZ>(
      this.client,
      DELETE_ENDPOINT,
      { keys: toArray(keys) },
      deleteReqZ,
      deleteResZ,
    );
  }
}
