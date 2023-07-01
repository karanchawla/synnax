// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { nanoid } from "nanoid";

import { Layout } from "@/layout";

export interface VisMeta {
  variant: string;
  key: string;
}

export const createVis = (props: Omit<Partial<Layout>, "type">): Layout => {
  const {
    location = "mosaic",
    name = "Visualizaton",
    key = nanoid(),
    window,
    tab,
  } = props;
  return {
    type: "vis",
    location,
    name,
    key,
    window,
    tab,
  };
};
