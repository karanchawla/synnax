// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { type FC, type ReactElement } from "react";

import { Icon } from "@synnaxlabs/media";
import { Align, Status, Text } from "@synnaxlabs/pluto";

import { ToolbarHeader, ToolbarTitle } from "@/components";
import { Layout } from "@/layout";
import { LinePlot } from "@/lineplot";
import { Schematic } from "@/schematic";
// import { Table } from "@/table";
import { create } from "@/vis/create";

import { LayoutSelector } from "./LayoutSelector";
import { type LayoutType } from "./types";

export const VisToolbarTitle = (): ReactElement => (
  <ToolbarTitle icon={<Icon.Visualize />}>Visualization</ToolbarTitle>
);

interface ToolbarProps {
  layoutKey: string;
}

const SelectVis = ({ layoutKey }: ToolbarProps): ReactElement => (
  <Align.Space justify="spaceBetween" style={{ height: "100%" }} empty>
    <ToolbarHeader>
      <VisToolbarTitle />
    </ToolbarHeader>
    <LayoutSelector layoutKey={layoutKey} />
  </Align.Space>
);

const TOOLBARS: Record<LayoutType | "vis", FC<ToolbarProps>> = {
  schematic: Schematic.Toolbar,
  lineplot: LinePlot.Toolbar,
  vis: SelectVis,
};
const NoVis = (): ReactElement => {
  const placer = Layout.usePlacer();
  return (
    <Align.Space justify="spaceBetween" style={{ height: "100%" }} empty>
      <ToolbarHeader>
        <VisToolbarTitle />
      </ToolbarHeader>
      <Align.Center direction="x" size="small">
        <Status.Text level="p" variant="disabled" hideIcon>
          No visualization selected. Select a visualization or
        </Status.Text>
        <Text.Link level="p" onClick={() => placer(create({}))}>
          create a new one.
        </Text.Link>
      </Align.Center>
    </Align.Space>
  );
};

const Content = (): ReactElement => {
  const layout = Layout.useSelectActiveMosaicLayout();
  if (layout == null) return <NoVis />;
  const Toolbar = TOOLBARS[layout.type as LayoutType];
  if (Toolbar == null) return <NoVis />;
  return <Toolbar layoutKey={layout?.key} />;
};

export const Toolbar: Layout.NavDrawerItem = {
  key: "visualization",
  content: <Content />,
  tooltip: "Visualize",
  icon: <Icon.Visualize />,
  minSize: 125,
  maxSize: 250,
};
