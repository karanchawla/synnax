// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { PIDNode } from "@synnaxlabs/pluto";

import { PIDSliceState, PIDState, PIDStoreState } from "./slice";

import { useMemoSelect } from "@/hooks";

export const selectPIDState = (state: PIDStoreState): PIDSliceState => state.pid;

export const selectPID = (state: PIDStoreState, key: string): PIDState =>
  selectPIDState(state).pids[key];

export const useSelectPID = (key: string): PIDState =>
  useMemoSelect((state: PIDStoreState) => selectPID(state, key), [key]);

export const selectSelectedPIDElementsProps = (
  state: PIDStoreState,
  layoutKey: string
): PIDElementInfo[] => {
  const pid = selectPID(state, layoutKey);
  const selected = pid.nodes.filter((node) => node.selected);
  return selected.map((node) => ({
    key: node.key,
    node,
    props: pid.props[node.key],
  }));
};

export interface PIDElementInfo {
  key: string;
  node: PIDNode;
  props: unknown;
}

export const useSelectSelectedPIDElementsProps = (
  layoutKey: string
): PIDElementInfo[] =>
  useMemoSelect(
    (state: PIDStoreState) => selectSelectedPIDElementsProps(state, layoutKey),
    [layoutKey]
  );

export const selectPIDElementProps = (
  state: PIDStoreState,
  layoutKey: string,
  key: string
): unknown => {
  const pid = selectPID(state, layoutKey);
  return pid.props[key];
};

export const useSelectPIDElementProps = (layoutKey: string, key: string): unknown =>
  useMemoSelect(
    (state: PIDStoreState) => selectPIDElementProps(state, layoutKey, key),
    [layoutKey, key]
  );
