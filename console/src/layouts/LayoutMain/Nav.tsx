// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { type ReactElement } from "react";

import { Icon, Logo } from "@synnaxlabs/media";
import { Divider, Nav, Button, OS, Text } from "@synnaxlabs/pluto";

import { Cluster } from "@/cluster";
import { Controls } from "@/components";
import { NAV_DRAWERS, NavMenu } from "@/components/nav/Nav";
import { CSS } from "@/css";
import { Docs } from "@/docs";
import { Layout } from "@/layout";
import { NAV_SIZES } from "@/layouts/LayoutMain/constants";
import { LinePlot } from "@/lineplot";
import { Palette } from "@/palette/Palette";
import { type TriggerConfig } from "@/palette/types";
import { Schematic } from "@/schematic";
import { Range } from "@/range";
import { SERVICES } from "@/services";
import { Version } from "@/version";
import { Vis } from "@/vis";
import { Workspace } from "@/workspace";

import "@/layouts/LayoutMain/Nav.css";

const DEFAULT_TRIGGER: TriggerConfig = {
  defaultMode: "command",
  resource: [["Control", "P"]],
  command: [["Control", "Shift", "P"]],
};

const COMMANDS = [
  ...LinePlot.COMMANDS,
  ...Layout.COMMANDS,
  ...Schematic.COMMANDS,
  ...Docs.COMMANDS,
  ...Workspace.COMMANDS,
  ...Cluster.COMMANDS,
  ...Range.COMMANDS,
];

const NavTopPalette = (): ReactElement => {
  return (
    <Palette
      commands={COMMANDS}
      triggers={DEFAULT_TRIGGER}
      services={SERVICES}
      commandSymbol=">"
    />
  );
};

/**
 * NavTop is the top navigation bar for the Synnax Console. Try to keep this component
 * presentational.
 */
export const NavTop = (): ReactElement => {
  const placer = Layout.usePlacer();

  const os = OS.use();
  const handleDocs = (): void => {
    placer(Docs.createLayout());
  };

  return (
    <Nav.Bar
      data-tauri-drag-region
      location="top"
      size={NAV_SIZES.top}
      className={CSS(CSS.B("main-nav"), CSS.B("main-nav-top"))}
    >
      <Nav.Bar.Start className="console-main-nav-top__start" data-tauri-drag-region>
        <Controls className="console-controls--macos" visibleIfOS="MacOS" />
        {os === "Windows" && (
          <Logo
            className="console-main-nav-top__logo"
            variant="icon"
            data-tauri-drag-region
          />
        )}
        <Workspace.Selector />
      </Nav.Bar.Start>
      <Nav.Bar.Content
        grow
        justify="center"
        className="console-main-nav-top__center"
        data-tauri-drag-region
      >
        <NavTopPalette />
      </Nav.Bar.Content>
      <Nav.Bar.End
        className="console-main-nav-top__end"
        justify="end"
        data-tauri-drag-region
      >
        <Button.Icon
          size="medium"
          onClick={handleDocs}
          tooltip={<Text.Text level="small">Documentation</Text.Text>}
        >
          <Icon.QuestionMark />
        </Button.Icon>
        <Controls className="console-controls--windows" visibleIfOS="Windows" />
      </Nav.Bar.End>
    </Nav.Bar>
  );
};

/**
 * NavLeft is the left navigation drawer for the Synnax Console. Try to keep this component
 * presentational.
 */
export const NavLeft = (): ReactElement => {
  const { onSelect, menuItems } = Layout.useNavDrawer("left", NAV_DRAWERS);
  const os = OS.use();
  return (
    <Nav.Bar className={CSS.B("main-nav")} location="left" size={NAV_SIZES.side}>
      {os !== "Windows" && (
        <Nav.Bar.Start className="console-main-nav-left__start" bordered>
          <Logo className="console-main-nav-left__logo" />
        </Nav.Bar.Start>
      )}
      <Nav.Bar.Content className="console-main-nav__content">
        <NavMenu onChange={onSelect}>{menuItems}</NavMenu>
      </Nav.Bar.Content>
    </Nav.Bar>
  );
};

/**
 * NavRight is the right navigation bar for the Synnax Console. Try to keep this component
 * presentational.
 */
export const NavRight = (): ReactElement | null => {
  const { menuItems, onSelect } = Layout.useNavDrawer("right", NAV_DRAWERS);
  const { menuItems: bottomMenuItems, onSelect: onBottomSelect } = Layout.useNavDrawer(
    "bottom",
    NAV_DRAWERS,
  );
  return (
    <Nav.Bar className={CSS.B("main-nav")} location="right" size={NAV_SIZES.side}>
      <Nav.Bar.Content className="console-main-nav__content" size="small">
        <NavMenu onChange={onSelect}>{menuItems}</NavMenu>
      </Nav.Bar.Content>
      {bottomMenuItems.length > 0 && (
        <Nav.Bar.End className="console-main-nav__content" bordered>
          <NavMenu onChange={onBottomSelect}>{bottomMenuItems}</NavMenu>
        </Nav.Bar.End>
      )}
    </Nav.Bar>
  );
};

/**
 * NavBottom is the bottom navigation bar for the Synnax Console. Try to keep this component
 * presentational.
 */
export const NavBottom = (): ReactElement => {
  return (
    <Nav.Bar className={CSS.B("main-nav")} location="bottom" size={NAV_SIZES.bottom}>
      <Nav.Bar.Start>
        <Vis.NavControls />
      </Nav.Bar.Start>
      <Nav.Bar.End className="console-main-nav-bottom__end" empty>
        <Divider.Divider />
        <Version.Badge level="p" />
        <Divider.Divider />
        <Cluster.Dropdown />
        <Divider.Divider />
        <Cluster.ConnectionBadge />
      </Nav.Bar.End>
    </Nav.Bar>
  );
};
