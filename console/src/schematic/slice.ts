// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { type PayloadAction, createSlice } from "@reduxjs/toolkit";
import {
  type Control,
  type Viewport,
  type Diagram,
  type Schematic,
} from "@synnaxlabs/pluto";
import { Color } from "@synnaxlabs/pluto/color";
import { type Theming } from "@synnaxlabs/pluto/theming";
import { box, scale, xy, deep, migrate } from "@synnaxlabs/x";
import { nanoid } from "nanoid/non-secure";
import { v4 as uuidV4 } from "uuid";

import { type Layout } from "@/layout";

export type NodeProps = object & {
  key: Schematic.Variant;
  color?: Color.Crude;
};

export interface State extends migrate.Migratable {
  editable: boolean;
  snapshot: boolean;
  remoteCreated: boolean;
  viewport: Diagram.Viewport;
  nodes: Diagram.Node[];
  edges: Diagram.Edge[];
  props: Record<string, NodeProps>;
  control: Control.Status;
  controlAcquireTrigger: number;
}

interface CopyBuffer {
  pos: xy.Crude;
  nodes: Diagram.Node[];
  edges: Diagram.Edge[];
  props: Record<string, NodeProps>;
}

const ZERO_COPY_BUFFER: CopyBuffer = {
  pos: xy.ZERO,
  nodes: [],
  edges: [],
  props: {},
};

// ||||| TOOLBAR |||||

const TOOLBAR_TABS = ["symbols", "properties"] as const;
export type ToolbarTab = (typeof TOOLBAR_TABS)[number];

export interface ToolbarState {
  activeTab: ToolbarTab;
}

export interface SliceState extends migrate.Migratable {
  mode: Viewport.Mode;
  copy: CopyBuffer;
  toolbar: ToolbarState;
  schematics: Record<string, State>;
}

export const SLICE_NAME = "schematic";

export interface StoreState {
  [SLICE_NAME]: SliceState;
}

export const ZERO_STATE: State = {
  version: "0.0.0",
  snapshot: false,
  nodes: [],
  edges: [],
  props: {},
  remoteCreated: false,
  viewport: { position: xy.ZERO, zoom: 1 },
  editable: true,
  control: "released",
  controlAcquireTrigger: 0,
};

export const ZERO_SLICE_STATE: SliceState = {
  version: "0.0.0",
  mode: "select",
  copy: { ...ZERO_COPY_BUFFER },
  toolbar: { activeTab: "symbols" },
  schematics: {},
};

export interface SetViewportPayload {
  layoutKey: string;
  viewport: Diagram.Viewport;
}

export interface AddElementPayload {
  layoutKey: string;
  key: string;
  props: NodeProps;
  node?: Partial<Diagram.Node>;
}

export interface SetElementPropsPayload {
  layoutKey: string;
  key: string;
  props: NodeProps;
}

export interface FixThemeContrastPayload {
  theme: Theming.ThemeSpec;
}

export interface SetNodesPayload {
  layoutKey: string;
  mode?: "replace" | "update";
  nodes: Diagram.Node[];
}

export interface SetNodePositionsPayload {
  layoutKey: string;
  positions: Record<string, xy.XY>;
}

export interface SetEdgesPayload {
  layoutKey: string;
  edges: Diagram.Edge[];
}

export interface CreatePayload extends State {
  key: string;
}

export interface RemovePayload {
  layoutKeys: string[];
}

export interface SetEditablePayload {
  layoutKey: string;
  editable: boolean;
}

export interface SetControlStatusPayload {
  layoutKey: string;
  control: Control.Status;
}

export interface ToggleControlPayload {
  layoutKey: string;
  status: Control.Status;
}

export interface SetActiveToolbarTabPayload {
  tab: ToolbarTab;
}

export interface CopySelectionPayload {}

export interface PasteSelectionPayload {
  layoutKey: string;
  pos: xy.XY;
}

export interface ClearSelectionPayload {
  layoutKey: string;
}

export interface SetViewportModePayload {
  mode: Viewport.Mode;
}

export interface SetRemoteCreatedPayload {
  layoutKey: string;
}

export const calculatePos = (
  region: box.Box,
  cursor: xy.XY,
  viewport: Diagram.Viewport,
): xy.XY => {
  const zoomXY = xy.construct(viewport.zoom);
  const s = scale.XY.translate(xy.scale(box.topLeft(region), -1))
    .translate(xy.scale(viewport.position, -1))
    .magnify({
      x: 1 / zoomXY.x,
      y: 1 / zoomXY.y,
    });
  return s.pos(cursor);
};

const MIGRATIONS: migrate.Migrations = {};

export const migrateSlice = migrate.migrator<SliceState, SliceState>(MIGRATIONS);

export const { actions, reducer } = createSlice({
  name: SLICE_NAME,
  initialState: ZERO_SLICE_STATE,
  reducers: {
    copySelection: (state, _: PayloadAction<CopySelectionPayload>) => {
      // for each schematic, find the keys of the selected nodes and edges
      // and add them to the copy buffer. Then get the props of each
      // selected node and edge and add them to the copy buffer.
      const { schematics } = state;
      const copyBuffer: CopyBuffer = {
        nodes: [],
        edges: [],
        props: {},
        pos: xy.ZERO,
      };
      Object.values(schematics).forEach((schematic) => {
        const { nodes, edges, props } = schematic;
        const selectedNodes = nodes.filter((node) => node.selected);
        const selectedEdges = edges.filter((edge) => edge.selected);
        copyBuffer.nodes = [...copyBuffer.nodes, ...selectedNodes];
        copyBuffer.edges = [...copyBuffer.edges, ...selectedEdges];
        selectedNodes.forEach((node) => {
          copyBuffer.props[node.key] = props[node.key];
        });
        selectedEdges.forEach((edge) => {
          copyBuffer.props[edge.key] = props[edge.key];
        });
      });
      const { nodes } = copyBuffer;
      if (nodes.length > 0) {
        const pos = nodes.reduce(
          (acc, node) => xy.translate(acc, node.position),
          xy.ZERO,
        );
        copyBuffer.pos = xy.scale(pos, 1 / nodes.length);
      }
      state.copy = copyBuffer;
    },
    pasteSelection: (state, { payload }: PayloadAction<PasteSelectionPayload>) => {
      const { pos, layoutKey } = payload;
      const console = xy.translation(state.copy.pos, pos);
      const schematic = state.schematics[layoutKey];
      const keys: Record<string, string> = {};
      const nextNodes = state.copy.nodes.map((node) => {
        const key: string = nanoid();
        schematic.props[key] = state.copy.props[node.key];
        keys[node.key] = key;
        return {
          ...node,
          position: xy.translate(node.position, console),
          key,
          selected: true,
        };
      });
      const nextEdges = state.copy.edges.map((edge) => {
        const key: string = nanoid();
        return {
          ...edge,
          key,
          source: keys[edge.source],
          target: keys[edge.target],
          selected: true,
        };
      });
      schematic.edges = [
        ...schematic.edges.map((edge) => ({ ...edge, selected: false })),
        ...nextEdges,
      ];
      schematic.nodes = [
        ...schematic.nodes.map((node) => ({ ...node, selected: false })),
        ...nextNodes,
      ];
    },
    create: (state, { payload }: PayloadAction<CreatePayload>) => {
      const { key: layoutKey } = payload;
      const schematic = { ...ZERO_STATE, ...payload };
      if (schematic.snapshot) {
        schematic.editable = false;
        clearSelections(schematic);
      }
      state.schematics[layoutKey] = schematic;
      state.toolbar.activeTab = "symbols";
    },
    clearSelection: (state, { payload }: PayloadAction<ClearSelectionPayload>) => {
      const { layoutKey } = payload;
      const schematic = state.schematics[layoutKey];
      schematic.nodes.forEach((node) => {
        node.selected = false;
      });
      schematic.edges.forEach((edge) => {
        edge.selected = false;
      });
      state.toolbar.activeTab = "symbols";
    },
    remove: (state, { payload }: PayloadAction<RemovePayload>) => {
      const { layoutKeys } = payload;
      layoutKeys.forEach((layoutKey) => {
        const schematic = state.schematics[layoutKey];
        if (schematic.control === "acquired") schematic.controlAcquireTrigger -= 1;
        // eslint-disable-next-line @typescript-eslint/no-dynamic-delete
        delete state.schematics[layoutKey];
      });
    },
    addElement: (state, { payload }: PayloadAction<AddElementPayload>) => {
      const { layoutKey, key, props, node } = payload;
      const schematic = state.schematics[layoutKey];
      if (!schematic.editable) return;
      schematic.nodes.push({
        key,
        selected: false,
        position: xy.ZERO,
        ...node,
      });
      schematic.props[key] = props;
    },
    setElementProps: (state, { payload }: PayloadAction<SetElementPropsPayload>) => {
      const { layoutKey, key, props } = payload;
      const schematic = state.schematics[layoutKey];
      if (!schematic.editable) return;
      if (key in schematic.props) {
        schematic.props[key] = { ...schematic.props[key], ...props };
      } else {
        const edge = schematic.edges.findIndex((edge) => edge.key === key);
        if (edge !== -1) {
          schematic.edges[edge] = { ...schematic.edges[edge], ...props };
        }
      }
    },
    setNodes: (state, { payload }: PayloadAction<SetNodesPayload>) => {
      const { layoutKey, nodes, mode = "replace" } = payload;
      const schematic = state.schematics[layoutKey];
      if (mode === "replace") schematic.nodes = nodes;
      else {
        const keys = nodes.map((node) => node.key);
        schematic.nodes = [
          ...schematic.nodes.filter((node) => !keys.includes(node.key)),
          ...nodes,
        ];
      }
      const anySelected =
        nodes.some((node) => node.selected) ||
        schematic.edges.some((edge) => edge.selected);
      if (anySelected) {
        if (state.toolbar.activeTab !== "properties")
          clearOtherSelections(state, layoutKey);
        state.toolbar.activeTab = "properties";
      } else state.toolbar.activeTab = "symbols";
    },
    setNodePositions: (state, { payload }: PayloadAction<SetNodePositionsPayload>) => {
      const { layoutKey, positions } = payload;
      const schematic = state.schematics[layoutKey];
      Object.entries(positions).forEach(([key, position]) => {
        const node = schematic.nodes.find((node) => node.key === key);
        if (node == null) return;
        node.position = position;
      });
    },
    setEdges: (state, { payload }: PayloadAction<SetEdgesPayload>) => {
      const { layoutKey, edges } = payload;
      const schematic = state.schematics[layoutKey];
      // check for new edges
      const prevKeys = schematic.edges.map((edge) => edge.key);
      const newEdges = edges.filter((edge) => !prevKeys.includes(edge.key));
      newEdges.forEach((edge) => {
        const source = schematic.nodes.find((node) => node.key === edge.source);
        const target = schematic.nodes.find((node) => node.key === edge.target);
        if (source == null || target == null) return;
        const sourceProps = schematic.props[source.key];
        const targetProps = schematic.props[target.key];
        if (sourceProps.color === targetProps.color && sourceProps.color != null)
          edge.color = sourceProps.color;
      });
      schematic.edges = edges;
      const anySelected =
        edges.some((edge) => edge.selected) ||
        schematic.nodes.some((node) => node.selected);
      if (anySelected) {
        if (state.toolbar.activeTab !== "properties")
          clearOtherSelections(state, layoutKey);
        state.toolbar.activeTab = "properties";
      } else state.toolbar.activeTab = "symbols";
    },
    setActiveToolbarTab: (
      state,
      { payload }: PayloadAction<SetActiveToolbarTabPayload>,
    ) => {
      const { tab } = payload;
      state.toolbar.activeTab = tab;
    },
    setViewport: (state, { payload }: PayloadAction<SetViewportPayload>) => {
      const { layoutKey, viewport } = payload;
      const schematic = state.schematics[layoutKey];
      schematic.viewport = viewport;
    },
    setEditable: (state, { payload }: PayloadAction<SetEditablePayload>) => {
      const { layoutKey, editable } = payload;
      const schematic = state.schematics[layoutKey];
      clearSelections(schematic);
      if (schematic.control === "acquired") {
        schematic.controlAcquireTrigger -= 1;
      }
      if (schematic.snapshot) return;
      schematic.editable = editable;
    },
    toggleControl: (state, { payload }: PayloadAction<ToggleControlPayload>) => {
      let { layoutKey, status } = payload;
      const schematic = state.schematics[layoutKey];
      if (status == null)
        status = schematic.control === "released" ? "acquired" : "released";
      if (status === "released") schematic.controlAcquireTrigger -= 1;
      else schematic.controlAcquireTrigger += 1;
    },
    setControlStatus: (state, { payload }: PayloadAction<SetControlStatusPayload>) => {
      const { layoutKey, control } = payload;
      const schematic = state.schematics[layoutKey];
      if (schematic == null) return;
      schematic.control = control;
      if (control === "acquired") schematic.editable = false;
    },
    setViewportMode: (
      state,
      { payload: { mode } }: PayloadAction<SetViewportModePayload>,
    ) => {
      state.mode = mode;
    },
    setRemoteCreated: (state, { payload }: PayloadAction<SetRemoteCreatedPayload>) => {
      const { layoutKey } = payload;
      const schematic = state.schematics[layoutKey];
      schematic.remoteCreated = true;
    },
    fixThemeContrast: (state, { payload }: PayloadAction<FixThemeContrastPayload>) => {
      const { theme } = payload;
      const bgColor = new Color.Color(theme.colors.gray.l0);
      Object.values(state.schematics).forEach((schematic) => {
        const { nodes, edges, props } = schematic;
        nodes.forEach((node) => {
          const nodeProps = props[node.key];
          if ("color" in nodeProps) {
            const c = new Color.Color(nodeProps.color as string);
            // check the contrast of the color
            if (c.contrast(bgColor) < 1.1) {
              // if the contrast is too low, change the color to the contrast color
              nodeProps.color = theme.colors.gray.l9;
            }
          }
        });
        edges.forEach((edge) => {
          if (
            edge.color != null &&
            new Color.Color(edge.color as string).contrast(bgColor) < 1.1
          ) {
            edge.color = theme.colors.gray.l9;
          } else if (edge.color == null) {
            edge.color = theme.colors.gray.l9;
          }
        });
      });
    },
  },
});

const clearOtherSelections = (state: SliceState, layoutKey: string): void => {
  Object.keys(state.schematics).forEach((key) => {
    // If any of the nodes or edges in other Diagram slices are selected, deselect them.
    if (key === layoutKey) return;
    clearSelections(state.schematics[key]);
  });
};

const clearSelections = (state: State): void => {
  state.nodes.forEach((node) => {
    node.selected = false;
  });
  state.edges.forEach((edge) => {
    edge.selected = false;
  });
};

export const {
  setNodePositions,
  toggleControl,
  setControlStatus,
  addElement,
  setEdges,
  setNodes,
  remove,
  clearSelection,
  create: internalCreate,
  setElementProps,
  setActiveToolbarTab,
  setViewport,
  setEditable,
  copySelection,
  pasteSelection,
  setViewportMode,
  setRemoteCreated,
  fixThemeContrast,
} = actions;

export type Action = ReturnType<(typeof actions)[keyof typeof actions]>;
export type Payload = Action["payload"];

export type LayoutType = "schematic";
export const LAYOUT_TYPE = "schematic";

export const create =
  (
    initial: Partial<State> & Omit<Partial<Layout.LayoutState>, "type">,
  ): Layout.Creator =>
  ({ dispatch }) => {
    const { name = "Schematic", location = "mosaic", window, tab, ...rest } = initial;
    const key = initial.key ?? uuidV4();
    dispatch(actions.create({ ...deep.copy(ZERO_STATE), key, ...rest }));
    return {
      key,
      location,
      name,
      type: LAYOUT_TYPE,
      window,
      tab,
    };
  };
