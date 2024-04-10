// Copyright 2024 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import {
  useCallback,
  type PropsWithChildren,
  type ReactElement,
  type MouseEventHandler,
  type MouseEvent,
} from "react";

import { Icon } from "@synnaxlabs/media";
import { Button, Dropdown as Core, Align, Synnax } from "@synnaxlabs/pluto";
import { List as CoreList } from "@synnaxlabs/pluto/list";
import { Menu as PMenu } from "@synnaxlabs/pluto/menu";
import { Text } from "@synnaxlabs/pluto/text";
import { useDispatch } from "react-redux";

import { connectWindowLayout } from "@/cluster/Connect";
import { type Cluster } from "@/cluster/core";
import { LOCAL_KEY } from "@/cluster/local";
import { useSelect, useSelectLocalState, useSelectMany } from "@/cluster/selectors";
import { remove, setActive, setLocalState } from "@/cluster/slice";
import { CSS } from "@/css";
import { Layout } from "@/layout";

import "@/cluster/Dropdown.css";

export const List = (): ReactElement => {
  const menuProps = PMenu.useContextMenu();
  const dispatch = useDispatch();
  const data = Object.values(useSelectMany());
  const active = useSelect();
  const openWindow = Layout.usePlacer();

  const selected = active != null ? [active?.key] : [];

  const handleConnect = (key: string | null): void => {
    dispatch(setActive(key));
  };

  const handleRemove = (keys: string[]): void => {
    dispatch(remove({ keys }));
  };

  const contextMenu = useCallback(
    ({ keys: [key] }: PMenu.ContextMenuMenuProps): ReactElement | null => {
      if (key == null) return null;
      const handleSelect = (menuKey: string): void => {
        if (key == null) return;
        switch (menuKey) {
          case "remove":
            return handleRemove([key]);
          case "connect":
            return handleConnect(key);
          case "disconnect":
            return handleConnect(null);
        }
      };

      return (
        <PMenu.Menu level="small" onChange={handleSelect}>
          {key !== LOCAL_KEY && (
            <PMenu.Item startIcon={<Icon.Delete />} size="small" itemKey="remove">
              Remove
            </PMenu.Item>
          )}
          {key === active?.key ? (
            <PMenu.Item
              startIcon={<Icon.Disconnect />}
              size="small"
              itemKey="disconnect"
            >
              Disconnect
            </PMenu.Item>
          ) : (
            <PMenu.Item startIcon={<Icon.Connect />} size="small" itemKey="connect">
              Connect
            </PMenu.Item>
          )}
        </PMenu.Menu>
      );
    },
    [active?.key, handleConnect, handleRemove],
  );

  return (
    <Align.Pack className={CSS.B("cluster-list")} direction="y">
      <Align.Pack direction="x" justify="spaceBetween" size="large" grow>
        <Align.Space
          className={CSS.B("cluster-list-title")}
          direction="y"
          justify="center"
          grow
        >
          <Text.WithIcon level="h5" startIcon={<Icon.Cluster />}>
            Clusters
          </Text.WithIcon>
        </Align.Space>
        <Button.Button
          variant="outlined"
          size="medium"
          startIcon={<Icon.Add />}
          onClick={() => openWindow(connectWindowLayout)}
          className={CSS.B("cluster-list-add")}
        >
          Add
        </Button.Button>
      </Align.Pack>
      <PMenu.ContextMenu
        style={{ width: "100%", height: 300 }}
        menu={contextMenu}
        {...menuProps}
      >
        <CoreList.List<string, Cluster> data={data} emptyContent={<NoneConnected />}>
          <CoreList.Selector
            value={selected}
            onChange={handleConnect}
            allowMultiple={false}
          >
            <CoreList.Core<string, Cluster> style={{ height: "100%", width: "100%" }}>
              {(p) => <ListItem {...p} />}
            </CoreList.Core>
          </CoreList.Selector>
        </CoreList.List>
      </PMenu.ContextMenu>
    </Align.Pack>
  );
};

const ListItem = (props: CoreList.ItemProps<string, Cluster>): ReactElement => {
  const dispatch = useDispatch();
  const { status, schematic } = useSelectLocalState();
  const isLocal = props.entry.key === LOCAL_KEY;
  let icon: ReactElement | null = null;
  let loading = false;
  if (isLocal) {
    switch (status) {
      case "starting":
        icon = <Icon.Loading />;
        loading = true;
        break;
      case "running":
        icon = <Icon.Pause />;
        break;
      case "stopping":
        icon = <Icon.Loading />;
        loading = true;
        break;
      case "stopped":
        icon = <Icon.Play />;
        break;
    }
  }
  const handleClick: MouseEventHandler = (e): void => {
    e.stopPropagation();
    if (!isLocal) return;
    if (status === "running") dispatch(setLocalState({ command: "stop" }));
    if (status === "stopped") dispatch(setLocalState({ command: "start" }));
  };
  return (
    <CoreList.ItemFrame
      className={CSS(CSS.B("cluster-list-item"), isLocal && "local")}
      direction="x"
      align="center"
      {...props}
    >
      <Align.Space direction="y" justify="spaceBetween" size={0.5} grow>
        <Text.Text level="p" weight={450}>
          {props.entry.name}
        </Text.Text>
        <Text.Text level="p" shade={6}>
          {props.entry.props.host}:{props.entry.props.port}
        </Text.Text>
      </Align.Space>
      {isLocal && (
        <Align.Space direction="y" align="end" size="small">
          {icon != null && (
            <Button.Icon
              disabled={status === "starting" || status === "stopping"}
              onClick={handleClick}
              variant="outlined"
              loading={loading}
            >
              {icon}
            </Button.Icon>
          )}
          <Text.Text level="p" shade={6}>
            Schematic {isLocal ? schematic : "N/A"}
          </Text.Text>
        </Align.Space>
      )}
    </CoreList.ItemFrame>
  );
};

export interface NoneConnectedProps extends PropsWithChildren {}

export const NoneConnectedBoundary = ({
  children,
}: NoneConnectedProps): ReactElement => {
  const client = Synnax.use();
  if (client != null) return <>{children}</>;
  return <NoneConnected />;
};

export const NoneConnected = (): ReactElement => {
  const placer = Layout.usePlacer();

  const handleCluster: Text.TextProps["onClick"] = (e: MouseEvent) => {
    e.stopPropagation();
    placer(connectWindowLayout);
  };

  return (
    <Align.Space empty style={{ height: "100%", position: "relative" }}>
      <Align.Center direction="y" style={{ height: "100%" }} size="small">
        <Text.Text level="p">No cluster connected.</Text.Text>
        <Text.Link level="p" onClick={handleCluster}>
          Connect a cluster
        </Text.Link>
      </Align.Center>
    </Align.Space>
  );
};

export const Dropdown = (): ReactElement => {
  const dropProps = Core.use();
  const cluster = useSelect();

  return (
    <Core.Dialog
      {...dropProps}
      variant="floating"
      bordered={false}
      location={{ x: "left", y: "top" }}
      className={CSS.B("cluster-dropdown")}
    >
      <Button.Button
        onClick={dropProps.toggle}
        variant="text"
        startIcon={<Icon.Cluster />}
        justify="center"
      >
        {cluster?.name ?? "No Active Cluster"}
      </Button.Button>
      <List />
    </Core.Dialog>
  );
};
