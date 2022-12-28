import { createSlice } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";
import { Mosaic, Theming } from "@synnaxlabs/pluto";
import type { MosaicLeaf, Location, Theme } from "@synnaxlabs/pluto";

import { Layout } from "../types";

/** The state of the layout slice */
export interface LayoutState {
  /** The current theme. */
  activeTheme: string;
  /**
   * A record of theme keys to themes. The active theme is guaranteed to be present
   * in this record. */
  themes: Record<string, Theme>;
  /**
   * A record of layout keys to layouts. These represent the properties of all layouts
   * currently rendered in the mosaic or in external windows.
   */
  layouts: Record<string, Layout>;
  mosaic: MosaicState;
  nav: NavState;
}

export interface MosaicState {
  activeTab: string | null;
  root: MosaicLeaf;
}

export interface NavState {
  drawer: NavDrawerState;
}

export type NavdrawerLocation = "right" | "left" | "bottom";

export interface NavDrawerState {
  left: NavdrawerEntryState;
  right: NavdrawerEntryState;
  bottom: NavdrawerEntryState;
}

export interface NavdrawerEntryState {
  activeItem: string | null;
  menuItems: string[];
  size: number;
}

/**
 * The name of the layout slice in a larger store.
 * NOTE: This must be the name of the slice in the store, or else all selectors will fail.
 */
export const LAYOUT_SLICE_NAME = "layout";

/**
 * Represents a partial view of a larger store that contains the layout slice. This is
 * typically used for hooks that accept the entire store state as a parameter but only
 * need access to the layout slice.
 */
export interface LayoutStoreState {
  [LAYOUT_SLICE_NAME]: LayoutState;
}

const initialState: LayoutState = {
  activeTheme: "synnaxDark",
  themes: Theming.themes,
  layouts: {
    main: {
      title: "Main",
      key: "main",
      type: "main",
      location: "window",
      window: {
        navTop: false,
      },
    },
  },
  mosaic: {
    activeTab: null,
    root: {
      key: 1,
      tabs: [],
    },
  },
  nav: {
    drawer: {
      left: {
        activeItem: null,
        menuItems: ["clusters", "resources"],
        size: 0,
      },
      right: {
        activeItem: null,
        menuItems: ["workspace"],
        size: 0,
      },
      bottom: {
        activeItem: null,
        menuItems: ["visualization"],
        size: 0,
      },
    },
  },
};

/** Signature for the placeLayut action. */
export type PlaceLayoutAction = PayloadAction<Layout>;
/** Signature for the removeLayout action. */
export type RemoveLayoutAction = PayloadAction<string>;

/** Signature for the setTheme action. */
export type SetActiveTheme = PayloadAction<string>;
/** Signature for the toggleTheme action. */
export type ToggleActiveThemeAction = PayloadAction<void>;

type DeleteLayoutMosaicTabAction = PayloadAction<{ tabKey: string }>;
type MoveLayoutMosaicTabAction = PayloadAction<{
  tabKey: string;
  key: number;
  loc: Location;
}>;
type ResizeLayoutMosaicTabAction = PayloadAction<{ key: number; size: number }>;
type SelectLayoutMosaicTabAction = PayloadAction<{ tabKey: string }>;
type RenameLayoutMosaicTabAction = PayloadAction<{ tabKey: string; title: string }>;

type SetNavdrawerEntryState = PayloadAction<{
  location: NavdrawerLocation;
  state: Partial<NavdrawerEntryState>;
}>;

export const {
  actions: {
    placeLayout,
    removeLayout,
    toggleActiveTheme,
    setActiveTheme,
    deleteLayoutMosaicTab,
    moveLayoutMosaicTab,
    selectLayoutMosaicTab,
    resizeLayoutMosaicTab,
    renameLayoutMosaicTab,
    setNavdrawerEntryState,
  },
  reducer: layoutReducer,
} = createSlice({
  name: LAYOUT_SLICE_NAME,
  initialState,
  reducers: {
    placeLayout: (state, { payload: layout }: PlaceLayoutAction) => {
      const { key, location, title } = layout;

      const prev = state.layouts[key];

      // If we're moving from a mosaic, remove the tab.
      if (prev != null && prev.location === "mosaic" && location !== "mosaic")
        state.mosaic.root = Mosaic.removeTab(state.mosaic.root, key);

      // If we're moving to a mosaic, insert a tab.
      if (location === "mosaic") {
        state.mosaic.root = Mosaic.insertTab(state.mosaic.root, { tabKey: key, title });
        state.mosaic.activeTab = key;
      }

      state.layouts[key] = layout;
    },
    removeLayout: (state, { payload: contentKey }: RemoveLayoutAction) => {
      const layout = state.layouts[contentKey];
      if (layout == null) return;
      const { location } = layout;

      if (location === "mosaic")
        state.mosaic.root = Mosaic.removeTab(state.mosaic.root, contentKey);

      // eslint-disable-next-line @typescript-eslint/no-dynamic-delete
      delete state.layouts[contentKey];
    },
    deleteLayoutMosaicTab: (
      state,
      { payload: { tabKey } }: DeleteLayoutMosaicTabAction
    ) => {
      state.mosaic.root = Mosaic.removeTab(state.mosaic.root, tabKey);
      // eslint-disable-next-line @typescript-eslint/no-dynamic-delete
      delete state.layouts[tabKey];
    },
    moveLayoutMosaicTab: (
      state,
      { payload: { tabKey, key, loc } }: MoveLayoutMosaicTabAction
    ) => {
      state.mosaic.root = Mosaic.moveTab(state.mosaic.root, tabKey, loc, key);
    },
    selectLayoutMosaicTab: (
      state,
      { payload: { tabKey } }: SelectLayoutMosaicTabAction
    ) => {
      state.mosaic.root = Mosaic.selectTab(state.mosaic.root, tabKey);
      state.mosaic.activeTab = tabKey;
    },
    resizeLayoutMosaicTab: (
      state,
      { payload: { key, size } }: ResizeLayoutMosaicTabAction
    ) => {
      state.mosaic.root = Mosaic.resizeLeaf(state.mosaic.root, key, size);
    },
    renameLayoutMosaicTab: (
      state,
      { payload: { tabKey, title } }: RenameLayoutMosaicTabAction
    ) => {
      const layout = state.layouts[tabKey];
      if (layout != null) layout.title = title;
      state.mosaic.root = Mosaic.renameTab(state.mosaic.root, tabKey, title);
    },
    setActiveTheme: (state, { payload: key }: SetActiveTheme) => {
      state.activeTheme = key;
    },
    toggleActiveTheme: (state) => {
      const keys = Object.keys(state.themes);
      const index = keys.indexOf(state.activeTheme);
      const next = keys[(index + 1) % keys.length];
      state.activeTheme = next;
    },
    setNavdrawerEntryState: (
      state,
      { payload: { location, state: entryState } }: SetNavdrawerEntryState
    ) => {
      state.nav.drawer[location] = {
        ...state.nav.drawer[location],
        ...entryState,
      };
    },
  },
});
