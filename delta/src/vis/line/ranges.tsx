// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { useMemo } from "react";

import { TimeStamp } from "@synnaxlabs/x";

import { useMemoSelect } from "@/hooks";
import { LayoutStoreState } from "@/layout";
import { XAxisKey, X_AXIS_KEYS } from "@/vis/axis";
import {
  GOOD_STATUS,
  INVALID_VIS_STATUS,
  LineVis,
  RangesState,
  Status,
  StatusProvider,
} from "@/vis/line/core";
import { selectRequiredVis, VisStoreState } from "@/vis/store";
import { Range, selectRanges, WorkspaceStoreState } from "@/workspace";

export class Ranges implements StatusProvider {
  readonly core: RangesState;
  readonly ranges: Record<string, Range>;
  readonly status: Status;

  private constructor(
    core: RangesState,
    ranges: Record<string, Range>,
    status: Status = GOOD_STATUS
  ) {
    this.core = core;
    this.ranges = ranges;
    this.status = status;
  }

  static useSelectCore(key: string): RangesState {
    return useMemoSelect(
      (state: VisStoreState & LayoutStoreState & WorkspaceStoreState) =>
        selectRequiredVis<LineVis>(state, key, "line").ranges,
      [key]
    );
  }

  static use(key: string): Ranges {
    const { core, ranges } = useMemoSelect(
      (state: VisStoreState & LayoutStoreState & WorkspaceStoreState) => {
        const core = selectRequiredVis<LineVis>(state, key, "line").ranges;
        const ranges = Ranges.rangesFromArray(selectRanges(state, Ranges.keys(core)));
        return { core, ranges };
      },
      [key]
    );
    return useMemo(() => {
      let status = GOOD_STATUS;
      if (!Ranges.isValid(core)) status = INVALID_VIS_STATUS;
      return new Ranges(core, ranges, status);
    }, [core, ranges]);
  }

  static isValid(core: RangesState): boolean {
    return Object.values(core).flat().length > 0;
  }

  private static keys(core: RangesState): string[] {
    return Object.values(core).flat();
  }

  private static rangesFromArray(ranges: Range[]): Record<string, Range> {
    return Object.fromEntries(ranges.map((r) => [r.key, r]));
  }

  forEach(callback: (range: Range, axes: XAxisKey[]) => void): void {
    Object.entries(this.ranges).forEach(([key, range]) => {
      const axes = X_AXIS_KEYS.filter((axis) => this.core[axis].includes(key));
      callback(range, axes);
    });
  }

  axis(key: XAxisKey): Range[] {
    return this.core[key].map((key) => this.ranges[key]);
  }

  axisKeys(key: XAxisKey): readonly string[] {
    return this.core[key];
  }

  get array(): Range[] {
    return Object.values(this.ranges);
  }

  get isLive(): boolean {
    const now = TimeStamp.now();
    return this.array.some((r) => new TimeStamp(r.end).after(now));
  }

  get valid(): boolean {
    return this.array.length > 0;
  }
}