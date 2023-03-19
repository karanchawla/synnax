// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { Space, Tab, Tabs } from "@synnaxlabs/pluto";

import { ToolbarHeader } from "@/components";
import { VisToolbarTitle } from "@/vis/components";
import { LinePlotAxisControls } from "@/vis/line/controls/LinePlotAxisControls";
import { LinePlotChannelControls } from "@/vis/line/controls/LinePlotChannelControls";

export interface LinePlotToolbarProps {
  layoutKey: string;
}

export const LinePlotToolBar = ({ layoutKey }: LinePlotToolbarProps): JSX.Element => {
  const content = ({ tabKey }: Tab): JSX.Element => {
    switch (tabKey) {
      case "axes":
        return <LinePlotAxisControls layoutKey={layoutKey} />;
      default:
        return <LinePlotChannelControls layoutKey={layoutKey} />;
    }
  };

  const tabProps = Tabs.useStatic({
    tabs: [
      {
        tabKey: "channels",
        name: "Channels",
      },
    ],
    content,
  });

  return (
    <Space empty>
      <Tabs.Provider value={tabProps}>
        <ToolbarHeader>
          <VisToolbarTitle />
          <Tabs.Selector style={{ borderBottom: "none" }} size="large" />
        </ToolbarHeader>
        <Tabs.Content />
      </Tabs.Provider>
    </Space>
  );
};