// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { Middleware } from "@reduxjs/toolkit";
import { MAIN_WINDOW } from "@synnaxlabs/drift";
import { AsyncKV, Deep, UnknownRecord, DeepKey } from "@synnaxlabs/x";
import { getVersion } from "@tauri-apps/api/app";
import { appWindow } from "@tauri-apps/api/window";

import { VersionStoreState } from "../version";

const PERSISTED_STATE_KEY = "delta-persisted-state";

export interface RequiredState extends VersionStoreState {}

/**
 * Returns a function that preloads the state from the given key-value store on the main
 * window.
 *
 * @param db - the key-value store to load the state from.
 * @returns a redux middleware.
 */
export const newPreloadState =
  <S extends RequiredState>(db: AsyncKV<string, S>) =>
  async (): Promise<S | undefined> => {
    if (appWindow.label !== MAIN_WINDOW) return undefined;
    const state = await db.get(PERSISTED_STATE_KEY);
    if (state == null) return undefined;
    return await reconcileVersions(state);
  };

export interface PersistStateMiddlewareConfig<S extends UnknownRecord<S>> {
  /** The key-value store to persist to. */
  db: AsyncKV<string, S>;
  /** The keys to exclude from persistence. */
  exclude?: Array<DeepKey<S>>;
}

/**
 * Returns a redux middleware that persists the state to the given key-value store on
 * the main window. NOTE: this key-value store does not encrypt sensitive data! BE CAREFUL!
 *
 * @param db - the key-value store to persist to.
 * @returns a redux middleware.
 */
export const newPersistStateMiddleware =
  <S>({ db, exclude = [] }: PersistStateMiddlewareConfig<S>): Middleware<{}, S> =>
  (store) =>
  (next) =>
  (action) => {
    const result = next(action);
    if (appWindow.label !== MAIN_WINDOW) return result;
    const state = Deep.delete(JSON.parse(JSON.stringify(store.getState())), ...exclude);
    void db.set(PERSISTED_STATE_KEY, state);
    return result;
  };

const reconcileVersions = async <S extends RequiredState>(
  state: S
): Promise<S | undefined> => {
  const storedVersion = state.version.version;
  const tauriVersion = await getVersion();
  return storedVersion === tauriVersion ? state : undefined;
};
