// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { Layout } from "@/layout";
import { createSlice } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";
import { useDispatch } from "react-redux";

export const SLICE_NAME = "version";

export const SLICE_UPDATE = "updateAvailable";


export interface SliceState {
  version: string;
  updateAvailable:boolean;
}

export interface StoreState {
  [SLICE_NAME]: SliceState;
}

const initialState: SliceState = {
  version: "0.0.0",
  updateAvailable: false,
};

export type SetVersionAction = PayloadAction<string>;
export type SetUpdate= PayloadAction<boolean>;


export const { actions, reducer } = createSlice({
  name: SLICE_NAME,
  initialState,
  reducers: {
    set: (state, { payload: version }: SetVersionAction) => {
      state.version = version;
    },
    setUpdateAvailable: (state, { payload: updateAvailable }: SetUpdate) => {
      state.updateAvailable = updateAvailable;
    },
  },
});

export const { set, setUpdateAvailable } = actions;


export type Action = ReturnType<(typeof actions)[keyof typeof actions]>;
export type Payload = Action["payload"];
