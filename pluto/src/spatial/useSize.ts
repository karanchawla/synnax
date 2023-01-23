// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { RefObject, useEffect, useRef, useState } from "react";

import { ResizeObserver } from "@juggle/resize-observer";
import { debounce as debounceF } from "@synnaxlabs/x";

import { Box, BoxF, ZERO_BOX } from "./box";
import { Direction, isDirection } from "./core";

import { compareArrayDeps, useMemoCompare } from "@/hooks";

/** A list of events that can trigger a resize. */
export type Trigger = "moveX" | "moveY" | "resizeX" | "resizeY";

export interface UseResizeOpts {
  /** A list of triggers that should cause the callback to be called. */
  triggers?: Array<Trigger | Direction>;
  /**  Debounce the resize event by this many milliseconds.
  Useful for preventing expensive renders until rezizing has stopped. */
  debounce?: number;
}

/**
 * Tracks the dimensions of an element and executes a callback when they change.
 *
 * @param onResize - A callback that receives a box representing the dimensions and
 * position of the element.
 * @param opts -  Options for the hook. See UseResizeOpts.
 *
 * @returns a ref callback to attach to the desire element.
 */
export const useResize = <E extends HTMLElement>(
  onResize: BoxF,
  { triggers: _triggers = [], debounce = 0 }: UseResizeOpts
): RefObject<E> => {
  const prev = useRef<Box>(ZERO_BOX);
  const ref = useRef<E | null>(null);
  const triggers = useMemoCompare(
    () => normalizeTriggers(_triggers),
    compareArrayDeps,
    [_triggers] as const
  );
  useEffect(() => {
    const el = ref.current;
    if (el == null) return;
    prev.current = ZERO_BOX;
    const deb = debounceF(() => {
      const next = new Box(el);
      if (shouldResize(triggers, prev.current, next)) onResize(next);
      prev.current = next;
    }, debounce);
    const obs = new ResizeObserver(deb);
    obs.observe(el);
    return () => obs.disconnect();
  }, [triggers, onResize, debounce]);
  return ref;
};

export type UseSizeOpts = UseResizeOpts;

/**
 * Tracks the size of an element and returns it.
 *
 * @param opts - Options for the hook. See UseSizeOpts.
 *
 * @returns A Box representing the size of the element and a ref callback to attach to
 * the element.
 */
export const useSize = <E extends HTMLElement>(
  opts: UseSizeOpts
): [Box, RefObject<E>] => {
  const [size, onResize] = useState<Box>(ZERO_BOX);
  const ref = useResize<E>(onResize, opts);
  return [size, ref];
};

const normalizeTriggers = (triggers: Array<Direction | Trigger>): Trigger[] =>
  triggers
    .map((t): Trigger | Trigger[] => {
      if (isDirection(t))
        return t === "x" ? ["moveX", "resizeX"] : ["moveY", "resizeY"];
      return t;
    })
    .flat();

const shouldResize = (
  triggers: Array<Trigger | Direction>,
  prev: Box,
  next: Box
): boolean => {
  if (triggers.length === 0)
    return (
      prev.width !== next.width ||
      prev.height !== next.height ||
      prev.left !== next.left ||
      prev.top !== next.top
    );
  if (triggers.includes("resizeX") && prev.width !== next.width) return true;
  if (triggers.includes("resizeY") && prev.height !== next.height) return true;
  if (triggers.includes("moveX") && prev.left !== next.left) return true;
  if (triggers.includes("moveY") && prev.top !== next.top) return true;
  return false;
};