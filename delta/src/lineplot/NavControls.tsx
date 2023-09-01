// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { ReactElement, ReactNode, useMemo } from "react";

import { Icon } from "@synnaxlabs/media";
import { Align, Button, Divider, Select, Text, Viewport } from "@synnaxlabs/pluto";
import { useDispatch } from "react-redux";

import { Layout } from "@/layout";
import { useSelectControlState, useSelectViewportMode } from "@/lineplot/selectors";
import {
  ClickMode,
  setControlState,
  setViewport,
  setViewportMode,
} from "@/lineplot/slice";

export const NavControls = (): ReactElement => {
  const control = useSelectControlState();
  const vis = Layout.useSelectActiveMosaicTabKey();
  const mode = useSelectViewportMode();
  const d = useDispatch();

  const handleModeChange = (mode: Viewport.Mode): void => {
    d(setViewportMode({ mode }));
  };

  const handleClickModeChange = (clickMode: ClickMode): void => {
    d(setControlState({ state: { clickMode } }));
  };

  const handleTooltipChange = (tooltip: boolean): void => {
    d(setControlState({ state: { enableTooltip: tooltip } }));
  };

  const handleZoomReset = (): void => {
    if (vis != null) d(setViewport({ layoutKey: vis }));
  };

  const triggers = useMemo(() => Viewport.DEFAULT_TRIGGERS[mode], [mode]);

  return (
    <>
      <Viewport.SelectMode
        value={mode}
        bordered={false}
        rounded={false}
        onChange={handleModeChange}
        triggers={triggers}
        size="medium"
      />
      <Button.Icon
        onClick={handleZoomReset}
        variant="text"
        tooltipLocation={{ x: "right", y: "top" }}
        tooltip={
          <Align.Space direction="x" align="center">
            <Text.Text level="small">Reset Zoom</Text.Text>
            <Text.Keyboard level="small">
              <Text.Symbols.Meta />
            </Text.Keyboard>
            <Text.Keyboard level="small">Click</Text.Keyboard>
          </Align.Space>
        }
        size="medium"
      >
        <Icon.Expand />
      </Button.Icon>

      <Divider.Divider />
      <Button.ToggleIcon
        value={control.enableTooltip}
        onChange={handleTooltipChange}
        checkedVariant="filled"
        uncheckedVariant="text"
        sharp
        tooltip={
          <Align.Space direction="x" align="center">
            <Text.Text level="small">Show Tooltip on Hover</Text.Text>
          </Align.Space>
        }
      >
        <Icon.Tooltip />
      </Button.ToggleIcon>
      <Divider.Divider />
      <Select.Button<
        ClickMode,
        { key: ClickMode; icon: ReactElement; tooltip: ReactNode }
      >
        value={control.clickMode as ClickMode}
        onChange={handleClickModeChange}
        size="medium"
        bordered={false}
        rounded={false}
        entryRenderKey="icon"
        allowNone
        data={[
          {
            key: "measure",
            icon: <Icon.Rule />,
            tooltip: (
              <Align.Space direction="x" align="center">
                <Text.Text level="small">Slope</Text.Text>
                <Text.Keyboard level="small">Alt</Text.Keyboard>
                <Text.Keyboard level="small">Drag</Text.Keyboard>
              </Align.Space>
            ),
          },
          {
            key: "annotate",
            icon: <Icon.Annotate />,
            tooltip: (
              <Align.Space direction="x" align="center">
                <Text.Text level="small">Annotate</Text.Text>
                <Text.Keyboard level="small">Alt</Text.Keyboard>
                <Text.Keyboard level="small">Click</Text.Keyboard>
              </Align.Space>
            ),
          },
        ]}
      >
        {({ title: _, entry, ...props }) => (
          <Button.Icon
            {...props}
            key={entry.key}
            variant={props.selected ? "filled" : "text"}
            style={{}}
            size="medium"
            tooltip={entry.tooltip}
            tooltipLocation={{ x: "right", y: "top" }}
          >
            {entry.icon}
          </Button.Icon>
        )}
      </Select.Button>
    </>
  );
};
