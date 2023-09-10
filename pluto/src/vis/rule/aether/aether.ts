// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { type Box, type Location, type Scale } from "@synnaxlabs/x";
import { z } from "zod";

import { aether } from "@/aether/aether";
import { color } from "@/color/core";
import { theming } from "@/theming/aether";
import { Draw2D } from "@/vis/draw2d";
import { render } from "@/vis/render";

export const ruleStateZ = z.object({
  position: z.number().optional(),
  pixelPosition: z.number().optional(),
  dragging: z.boolean(),
  lineWidth: z.number().optional().default(1),
  lineDash: z.number().optional().default(20),
  color: color.Color.z,
});

export interface RuleProps {
  location: Location;
  decimalToDataScale: Scale;
  plot: Box;
  container: Box;
}

interface InternalState {
  renderCtx: render.Context;
  draw: Draw2D;
}

export class Rule extends aether.Leaf<typeof ruleStateZ, InternalState> {
  static readonly TYPE = "Rule";

  schema = ruleStateZ;

  afterUpdate(): void {
    this.internal.renderCtx = render.Context.use(this.ctx);
    const theme = theming.use(this.ctx);
    this.internal.draw = new Draw2D(this.internal.renderCtx.upper2d, theme);
    render.Controller.requestRender(this.ctx, render.REASON_TOOL);
  }

  afterDelete(): void {
    render.Controller.requestRender(this.ctx, render.REASON_TOOL);
  }

  updatePositions({ decimalToDataScale: scale, plot, container }: RuleProps): number {
    if (this.state.dragging && this.state.pixelPosition != null) {
      const pos = scale.pos(
        (this.state.pixelPosition - plot.top + container.top) / plot.height,
      );
      this.setState((p) => ({ ...p, position: pos }));
      return this.state.pixelPosition;
    }
    if (this.state.position == null) {
      // Calculate the position of the rule at the middle of the plot
      const pos = scale.pos(0.5);
      this.setState((p) => ({ ...p, position: pos }));
    }
    const pixelPos =
      scale.reverse().pos(this.state.position as number) * plot.height +
      plot.top -
      container.top;
    if (!isNaN(pixelPos)) this.setState((p) => ({ ...p, pixelPosition: pixelPos }));
    return pixelPos;
  }

  async render(props: RuleProps): Promise<void> {
    if (this.deleted) return;
    const { renderCtx } = this.internal;
    const { location, plot: plottingRegion } = props;
    const direction = location.direction;
    const { upper2d: canvas } = renderCtx;
    const draw = this.internal.draw;

    let pixelPos = this.updatePositions(props);
    pixelPos += props.container.top;

    draw.rule({
      stroke: this.state.color,
      lineWidth: this.state.lineWidth,
      lineDash: this.state.lineDash,
      direction,
      region: plottingRegion,
      position: pixelPos,
    });

    canvas.fillStyle = this.state.color.hex;
    canvas.beginPath();
    if (location.isLeft) {
      canvas.moveTo(plottingRegion.left, pixelPos);
      canvas.lineTo(plottingRegion.left - 5, pixelPos - 5);
      canvas.lineTo(plottingRegion.left - 5, pixelPos + 5);
    } else if (location.isRight) {
      canvas.moveTo(plottingRegion.right, pixelPos);
      canvas.lineTo(plottingRegion.right + 5, pixelPos - 5);
      canvas.lineTo(plottingRegion.right + 5, pixelPos + 5);
    }
    canvas.closePath();
    canvas.fill();
  }
}

export const REGISTRY: aether.ComponentRegistry = {
  [Rule.TYPE]: Rule,
};
