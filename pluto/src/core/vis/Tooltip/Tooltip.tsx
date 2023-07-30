// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { ReactElement, useCallback, useEffect, useRef } from "react";

import { XY } from "@synnaxlabs/x";
import { z } from "zod";

import { Aether } from "@/core/aether/main";
import { AetherTooltip } from "@/core/vis/Tooltip/aether";

export interface TooltipProps
  extends Omit<z.input<typeof AetherTooltip.stateZ>, "position"> {}

export const Tooltip = Aether.wrap<TooltipProps>(
  "Tooltip",
  ({ aetherKey }): ReactElement | null => {
    const [, , setState] = Aether.use({
      aetherKey,
      type: AetherTooltip.TYPE,
      schema: AetherTooltip.stateZ,
      initialState: {
        position: null,
      },
    });

    const ref = useRef<HTMLSpanElement>(null);

    const handleMove = useCallback(
      (e: MouseEvent): void => setState({ position: new XY(e) }),
      [setState]
    );

    const handleLeave = useCallback(
      (): void => setState({ position: null }),
      [setState]
    );

    useEffect(() => {
      if (ref.current === null) return;
      // Select the parent node of the tooltip
      const parent = ref.current.parentElement;
      if (parent == null) return;
      // Bind a hover listener to the parent node
      parent.addEventListener("mousemove", handleMove);
      parent.addEventListener("mouseleave", handleLeave);
      return () => {
        parent.removeEventListener("mousemove", handleMove);
        parent.removeEventListener("mouseleave", handleLeave);
      };
    }, [handleMove]);

    return <span ref={ref} />;
  }
);
