// Copyright 2024 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import type { PayloadAction, UnknownAction } from "@reduxjs/toolkit";
import { Drift } from "@synnaxlabs/drift";
import { useSelectWindowKey } from "@synnaxlabs/drift/react";
import {
  type AsyncDestructor,
  type Nav,
  OS,
  Theming,
  useAsyncEffect,
  useDebouncedCallback,
  useMemoCompare,
} from "@synnaxlabs/pluto";
import { compare } from "@synnaxlabs/x";
import { getCurrent } from "@tauri-apps/api/window";
import { type Dispatch, type ReactElement, useCallback, useState } from "react";
import { useDispatch, useStore } from "react-redux";

import { useSyncerDispatch } from "@/hooks/dispatchers";
import { State } from "@/layout/layout";
import { select, useSelectNavDrawer, useSelectTheme } from "@/layout/selectors";
import {
  type NavdrawerLocation,
  place,
  remove,
  resizeNavdrawer,
  setActiveTheme,
  setNavdrawerVisible,
  toggleActiveTheme,
} from "@/layout/slice";
import { type RootState } from "@/store";
import { Workspace } from "@/workspace";

export interface CreatorProps {
  windowKey: string;
  dispatch: Dispatch<PayloadAction<any>>;
}

/** A function that creates a layout given a set of utilities. */
export type Creator = (props: CreatorProps) => Omit<State, "windowKey">;

/** A function that places a layout using the given properties or creation func. */
export type Placer = (layout: Omit<State, "windowKey"> | Creator) => {
  windowKey: string;
  key: string;
};

/** A function that removes a layout. */
export type Remover = (...keys: string[]) => void;

/**
 * useLayoutPlacer is a hook that returns a function that allows the caller to place
 * a layout in the central mosaic or in a window.
 *
 * @returns A layout placer function that allows the caller to open a layout using one
 * of two methods. The first is to pass a layout object with the layout's key, type,
 * title, location, and window properties. The second is to pass a layout creator function
 * that accepts a few utilities and returns a layout object. Prefer the first method
 * when possible, but feel free to use the second method for more dynamic layout creation.
 */
export const usePlacer = (): Placer => {
  const syncer = Workspace.useLayoutSyncer();
  const dispatch = useSyncerDispatch(syncer);
  const os = OS.use();
  const windowKey = useSelectWindowKey();
  if (windowKey == null) throw new Error("windowKey is null");
  return useCallback(
    (base) => {
      const layout = typeof base === "function" ? base({ dispatch, windowKey }) : base;
      const { key, location, window, name: title } = layout;
      dispatch(place({ ...layout, windowKey }));
      if (location === "window")
        dispatch(
          Drift.createWindow({
            ...{ ...window, navTop: undefined, decorations: os !== "Windows" },
            url: "/",
            key,
            title,
          }),
        );
      return { windowKey, key };
    },
    [dispatch, windowKey],
  );
};

/**
 * useLayoutRemover is a hook that returns a function that allows the caller to remove
 * a layout.
 *
 * @param key - The key of the layout to remove.
 * @returns A layout remover function that allows the caller to remove a layout. If
 * the layout is in a window, the window will also be closed.
 */
export const useRemover = (...baseKeys: string[]): Remover => {
  const syncer = Workspace.useLayoutSyncer();
  const syncDispatch = useSyncerDispatch(syncer);
  const store = useStore<RootState>();
  const memoKeys = useMemoCompare(
    () => baseKeys,
    ([a], [b]) => compare.primitiveArrays(a, b) === compare.EQUAL,
    [baseKeys],
  );
  return useCallback(
    (...keys) => {
      keys = [...baseKeys, ...keys];
      const s = store.getState();
      keys.forEach((keys) => {
        const l = select(s, keys);
        // Even if the layout is not present, close the window for good measure.
        if (l == null || l.location === "window") {
          store.dispatch(Drift.closeWindow({ key: keys }));
        }
      });
      syncDispatch(remove({ keys }));
    },
    [memoKeys],
  );
};

/**
 * useThemeProvider is a hook that returns the props to pass to a ThemeProvider from
 * @synnaxlabs/pluto. This hook allows theme management to be centralized in the layout
 * redux store, and be synchronized across several windows.
 *
 * @returns The props to pass to a ThemeProvider from @synnaxlabs/pluto.
 */
export const useThemeProvider = (): Theming.ProviderProps => {
  const theme = useSelectTheme();
  const dispatch = useDispatch();

  useAsyncEffect(async () => {
    if (getCurrent().label !== Drift.MAIN_WINDOW) return;
    await setInitialTheme(dispatch);
    const cleanup = await synchronizeWithOS(dispatch);
    return cleanup;
  }, []);

  return {
    theme: Theming.themeZ.parse(theme),
    setTheme: (key: string) => dispatch(setActiveTheme(key)),
    toggleTheme: () => dispatch(toggleActiveTheme()),
  };
};

export const useErrorThemeProvider = (): Theming.ProviderProps => {
  const [theme, setTheme] = useState<Theming.ThemeSpec | null>(Theming.SYNNAX_LIGHT);
  useAsyncEffect(async () => {
    const theme = matchThemeChange({ payload: await getCurrent().theme() });
    setTheme(Theming.SYNNAX_THEMES[theme]);
  }, []);
  return {
    theme: Theming.themeZ.parse(theme),
    setTheme: (key: string) =>
      setTheme(Theming.SYNNAX_THEMES[key as keyof typeof Theming.SYNNAX_THEMES]),
    toggleTheme: () =>
      setTheme((t) =>
        t?.key === Theming.SYNNAX_LIGHT.key
          ? Theming.SYNNAX_DARK
          : Theming.SYNNAX_LIGHT,
      ),
  };
};

const matchThemeChange = ({
  payload: theme,
}: {
  payload: string | null;
}): keyof typeof Theming.SYNNAX_THEMES =>
  theme === "dark" ? "synnaxDark" : "synnaxLight";

const synchronizeWithOS = async (
  dispatch: Dispatch<UnknownAction>,
): AsyncDestructor => {
  await getCurrent().onThemeChanged((e) =>
    dispatch(setActiveTheme(matchThemeChange(e))),
  );
};

const setInitialTheme = async (dispatch: Dispatch<UnknownAction>): Promise<void> =>
  dispatch(setActiveTheme(matchThemeChange({ payload: await getCurrent().theme() })));

export interface NavMenuItem {
  key: string;
  icon: ReactElement;
  tooltip: string;
}

export interface NavDrawerItem extends Nav.DrawerItem, NavMenuItem {}

export interface UseNavDrawerReturn {
  activeItem: NavDrawerItem | undefined;
  menuItems: NavMenuItem[];
  onSelect: (item: string) => void;
  onResize: (size: number) => void;
}

export const useNavDrawer = (
  location: NavdrawerLocation,
  items: NavDrawerItem[],
): UseNavDrawerReturn => {
  const windowKey = useSelectWindowKey() as string;
  const state = useSelectNavDrawer(location);
  const dispatch = useDispatch();
  const onResize = useDebouncedCallback(
    (size) => {
      dispatch(resizeNavdrawer({ windowKey, location, size }));
    },
    100,
    [dispatch, windowKey],
  );
  if (state == null) {
    return {
      activeItem: undefined,
      menuItems: [],
      onSelect: () => {},
      onResize: () => {},
    };
  }
  let activeItem: NavDrawerItem | undefined;
  let menuItems: NavMenuItem[] = [];
  if (state.activeItem != null)
    activeItem = items.find((item) => item.key === state.activeItem);
  menuItems = items.filter((item) => state.menuItems.includes(item.key));

  if (activeItem != null) activeItem.initialSize = state.size;

  return {
    activeItem,
    menuItems,
    onSelect: (key: string) => dispatch(setNavdrawerVisible({ windowKey, key })),
    onResize,
  };
};
