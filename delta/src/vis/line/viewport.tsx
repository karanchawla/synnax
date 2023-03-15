// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { useCallback, useEffect, useState, useTransition } from "react";

import {
  UseViewportReturn as PUseViewportReturn,
  Menu as PMenu,
  Viewport as PViewport,
  UseViewportHandler,
  UseContextMenuReturn,
} from "@synnaxlabs/pluto";
import { Box, DECIMAL_BOX, XY } from "@synnaxlabs/x";
import { useDispatch } from "react-redux";

import { selectRequiredVis, updateVis, VisualizationStoreState } from "../store";

import { LineVis } from "./core";

import { useMemoSelect } from "@/hooks";
import { LayoutStoreState } from "@/layout";

export interface HoverState {
  cursor: XY;
  box: Box;
}

export interface UseViewportReturn {
  viewportProps: PUseViewportReturn;
  menuProps: UseContextMenuReturn;
  viewport: Box;
  selection: Box | null;
  hover: HoverState | null;
}

export const use = (key: string): UseViewportReturn => {
  const [viewport, setViewport] = useState<Box>(DECIMAL_BOX);
  const [selection, setSelection] = useState<Box | null>(null);
  const [hover, setHover] = useState<HoverState | null>(null);
  const [, startTransition] = useTransition();

  const core = useMemoSelect(
    (state: VisualizationStoreState & LayoutStoreState) =>
      selectRequiredVis<LineVis>(state, key, "line").viewport,
    []
  );

  const dispatch = useDispatch();

  const updateViewport = useCallback(
    (box: Box) =>
      dispatch(updateVis({ key, viewport: { pan: box.bottomLeft, zoom: box.dims } })),
    [key]
  );

  // We're just using this to persist the viewport state
  // and only care about loading it back on mount.
  useEffect(() => {
    const viewport = new Box(core.pan, core.zoom).reRoot("bottomLeft");
    setViewport(viewport);
  }, [core]);

  const menuProps = PMenu.useContextMenu();

  const handleViewport: UseViewportHandler = useCallback(
    (props) => {
      const { box, mode, cursor, stage } = props;
      if (mode === "hover") return setHover({ cursor, box });
      if (mode === "select") {
        setSelection(box);
        return menuProps.open(cursor);
      }
      startTransition(() => {
        setSelection(null);
        setViewport(box);
      });
      if (stage === "end") updateViewport(box);
    },
    [updateViewport]
  );

  const viewportProps = PViewport.use({ onChange: handleViewport });

  return { viewportProps, menuProps, viewport, selection, hover };
};

export const Viewport = {
  use,
};
