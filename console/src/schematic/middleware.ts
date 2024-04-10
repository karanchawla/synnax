// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { Layout } from "@/layout";
import { effectMiddleware, type MiddlewareEffect } from "@/middleware";
import { selectSliceState } from "@/schematic/selectors";
import {
  remove,
  type RemovePayload,
  type StoreState,
  fixThemeContrast,
  type FixThemeContrastPayload,
} from "@/schematic/slice";

export const deleteEffect: MiddlewareEffect<
  Layout.StoreState & StoreState,
  Layout.RemovePayload | Layout.SetSlicePayload,
  RemovePayload
> = ({ action, dispatch, getState }) => {
  const state = getState();
  const schematicSLice = selectSliceState(state);
  const layoutSlice = Layout.selectSliceState(state);
  // This is the case where the action does an explicit removal.
  const keys = "keys" in action.payload ? action.payload.keys : [];
  // We also just do a genera purpose garbage collection if necessary.
  const toRemove = Object.keys(schematicSLice.schematics).filter(
    (p) => keys.includes(p) || layoutSlice.layouts[p] == null,
  );
  if (toRemove.length > 0) dispatch(remove({ layoutKeys: toRemove }));
};

export const themeChangeEffect: MiddlewareEffect<
  Layout.StoreState & StoreState,
  Layout.SetActiveThemePayload,
  FixThemeContrastPayload
> = ({ dispatch, getState }) => {
  const theme = Layout.selectRawTheme(getState());
  dispatch(fixThemeContrast({ theme }));
};

export const MIDDLEWARE = [
  effectMiddleware([Layout.remove.type, Layout.setWorkspace.type], [deleteEffect]),
  effectMiddleware([Layout.setActiveTheme.type], [themeChangeEffect]),
];
