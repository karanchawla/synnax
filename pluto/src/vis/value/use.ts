// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { useLayoutEffect } from "react";

import { type z } from "zod";

import { Aether } from "@/aether";
import { useMemoDeepEqualProps } from "@/memo";
import { text } from "@/text/core";
import { Value } from "@/vis/value/aether/value";

export const corePropsZ = Value.z
  .omit({ font: true })
  .partial({ color: true })
  .extend({ level: text.levelZ.optional() });

export interface UseProps extends z.input<typeof corePropsZ> {
  aetherKey: string;
}

export interface UseReturn {
  width: number;
}

export const use = ({
  aetherKey,
  box,
  telem,
  color,
  precision,
  minWidth,
  level = "small",
}: UseProps): UseReturn => {
  const memoProps = useMemoDeepEqualProps({
    box,
    telem,
    color,
    precision,
    level,
    minWidth,
  });

  const [, state, setState] = Aether.use({
    aetherKey,
    type: Value.TYPE,
    schema: Value.z,
    initialState: memoProps,
  });

  useLayoutEffect(() => setState((prev) => ({ ...prev, ...memoProps })), [memoProps]);
  return { width: state.width ?? state.minWidth };
};
