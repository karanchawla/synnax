// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { useState, type ReactElement, useCallback } from "react";

import { box, xy, location, type UnknownRecord, direction } from "@synnaxlabs/x";

import { Aether } from "@/aether";
import { Align } from "@/align";
import { type Color } from "@/color";
import { CSS } from "@/css";
import { useResize } from "@/hooks";
import { Control } from "@/telem/control";
import { Text } from "@/text";
import { Theming } from "@/theming";
import { Tooltip } from "@/tooltip";
import { Button as CoreButton } from "@/vis/button";
import { useInitialViewport } from "@/vis/diagram/aether/Diagram";
import { Labeled, type LabelExtensionProps } from "@/vis/schematic/Labeled";
import { Primitives } from "@/vis/schematic/primitives";
import { Toggle } from "@/vis/toggle";
import { Value as CoreValue } from "@/vis/value";

import "@/vis/schematic/Symbols.css";

export interface ControlStateProps extends Omit<Align.SpaceProps, "direction"> {
  showChip?: boolean;
  showIndicator?: boolean;
  chip?: Control.ChipProps;
  indicator?: Control.IndicatorProps;
  orientation?: location.Outer;
}

const swapXLocation = (l: location.Outer): location.Outer =>
  direction.construct(l) === "x" ? (location.swap(l) as location.Outer) : l;
const swapYLocation = (l: location.Outer): location.Outer =>
  direction.construct(l) === "y" ? (location.swap(l) as location.Outer) : l;

export const ControlState = ({
  showChip = true,
  showIndicator = true,
  indicator,
  orientation = "left",
  chip,
  children,
  ...props
}: ControlStateProps): ReactElement => (
  <Align.Space
    direction={location.rotate90(orientation)}
    align="center"
    justify="center"
    empty
    {...props}
  >
    <Align.Space
      direction={direction.construct(orientation)}
      align="center"
      className={CSS(CSS.B("control-state"))}
      size="small"
    >
      {showChip && <Control.Chip size="small" {...chip} />}
      {showIndicator && <Control.Indicator {...indicator} />}
    </Align.Space>
    {children}
  </Align.Space>
);

export type SymbolProps<P extends object = UnknownRecord> = P & {
  symbolKey: string;
  position: xy.XY;
  selected: boolean;
  onChange: (value: P) => void;
};

export interface ThreeWayValveProps
  extends Primitives.ThreeWayValveProps,
    Omit<Toggle.UseProps, "aetherKey"> {
  label?: LabelExtensionProps;
  control?: ControlStateProps;
}

export const ThreeWayValve = Aether.wrap<SymbolProps<ThreeWayValveProps>>(
  "ThreeWayValve",
  ({
    aetherKey,
    label,
    onChange,
    control,
    source,
    sink,
    orientation = "left",
    color,
  }): ReactElement => {
    const { enabled, triggered, toggle } = Toggle.use({
      aetherKey,
      source,
      sink,
    });
    return (
      <Labeled {...label} onChange={onChange}>
        <ControlState {...control} orientation={swapXLocation(orientation)}>
          <Primitives.ThreeWayValve
            enabled={enabled}
            triggered={triggered}
            onClick={toggle}
            orientation={orientation}
            color={color}
          />
        </ControlState>
      </Labeled>
    );
  },
);

export const ThreeWayValvePreview = (props: ThreeWayValveProps): ReactElement => (
  <Primitives.ThreeWayValve {...props} />
);

export interface ValveProps
  extends Primitives.ValveProps,
    Omit<Toggle.UseProps, "aetherKey"> {
  label?: LabelExtensionProps;
  control?: ControlStateProps;
}

export const Valve = Aether.wrap<SymbolProps<ValveProps>>(
  "Valve",
  ({
    control,
    aetherKey,
    label,
    onChange,
    source,
    sink,
    orientation,
    color,
  }): ReactElement => {
    const { enabled, triggered, toggle } = Toggle.use({ aetherKey, source, sink });
    return (
      <Labeled {...label} onChange={onChange}>
        <ControlState {...control} orientation={orientation}>
          <Primitives.Valve
            enabled={enabled}
            triggered={triggered}
            onClick={toggle}
            orientation={orientation}
            color={color}
          />
        </ControlState>
      </Labeled>
    );
  },
);

export const ValvePreview = (props: ValveProps): ReactElement => (
  <Primitives.Valve {...props} />
);

export interface SolenoidValveProps
  extends Primitives.SolenoidValveProps,
    Omit<Toggle.UseProps, "aetherKey"> {
  label?: LabelExtensionProps;
  control?: ControlStateProps;
}

export const SolenoidValve = Aether.wrap<SymbolProps<SolenoidValveProps>>(
  "SolenoidValve",
  ({
    aetherKey,
    label,
    onChange,
    orientation = "left",
    normallyOpen,
    color,
    source,
    sink,
    control,
  }): ReactElement => {
    const { enabled, triggered, toggle } = Toggle.use({ aetherKey, source, sink });
    return (
      <Labeled {...label} onChange={onChange}>
        <ControlState {...control} orientation={swapYLocation(orientation)}>
          <Primitives.SolenoidValve
            enabled={enabled}
            triggered={triggered}
            onClick={toggle}
            orientation={orientation}
            color={color}
            normallyOpen={normallyOpen}
          />
        </ControlState>
      </Labeled>
    );
  },
);

export const SolenoidValvePreview = (props: SolenoidValveProps): ReactElement => (
  <Primitives.SolenoidValve {...props} />
);

export interface FourWayValveProps
  extends Primitives.FourWayValveProps,
    Omit<Toggle.UseProps, "aetherKey"> {
  label?: LabelExtensionProps;
  control?: ControlStateProps;
}

export const FourWayValve = Aether.wrap<SymbolProps<FourWayValveProps>>(
  "FourWayValve",
  ({
    aetherKey,
    control,
    label,
    onChange,
    orientation,
    color,
    source,
    sink,
  }): ReactElement => {
    const { enabled, triggered, toggle } = Toggle.use({ aetherKey, source, sink });
    return (
      <Labeled {...label} onChange={onChange}>
        <ControlState {...control} orientation={orientation}>
          <Primitives.FourWayValve
            enabled={enabled}
            triggered={triggered}
            onClick={toggle}
            orientation={orientation}
            color={color}
          />
        </ControlState>
      </Labeled>
    );
  },
);

export const FourWayValvePreview = (props: FourWayValveProps): ReactElement => (
  <Primitives.FourWayValve {...props} />
);

export interface AngledValveProps
  extends Primitives.AngledValveProps,
    Omit<Toggle.UseProps, "aetherKey"> {
  label?: LabelExtensionProps;
  control?: ControlStateProps;
}

export const AngledValve = Aether.wrap<SymbolProps<AngledValveProps>>(
  "AngleValve",
  ({
    aetherKey,
    label,
    control,
    onChange,
    orientation = "left",
    color,
    source,
    sink,
  }): ReactElement => {
    const { enabled, triggered, toggle } = Toggle.use({ aetherKey, source, sink });
    return (
      <Labeled {...label} onChange={onChange}>
        <ControlState {...control} orientation={swapXLocation(orientation)}>
          <Primitives.AngledValve
            enabled={enabled}
            triggered={triggered}
            onClick={toggle}
            orientation={orientation}
            color={color}
          />
        </ControlState>
      </Labeled>
    );
  },
);

export const AngledValvePreview = (props: AngledValveProps): ReactElement => (
  <Primitives.AngledValve {...props} />
);

export interface PumpProps
  extends Primitives.PumpProps,
    Omit<Toggle.UseProps, "aetherKey"> {
  label?: LabelExtensionProps;
  control?: ControlStateProps;
}

export const Pump = Aether.wrap<SymbolProps<PumpProps>>(
  "Pump",
  ({
    aetherKey,
    label,
    control,
    onChange,
    orientation,
    color,
    source,
    sink,
  }): ReactElement => {
    const { enabled, triggered, toggle } = Toggle.use({ aetherKey, source, sink });
    return (
      <Labeled {...label} onChange={onChange}>
        <ControlState {...control} orientation={orientation}>
          <Primitives.Pump
            enabled={enabled}
            triggered={triggered}
            onClick={toggle}
            orientation={orientation}
            color={color}
          />
        </ControlState>
      </Labeled>
    );
  },
);

export const PumpPreview = (props: PumpProps): ReactElement => (
  <Primitives.Pump {...props} />
);

export interface TankProps extends Primitives.TankProps {
  label?: LabelExtensionProps;
}

export const Tank = Aether.wrap<SymbolProps<TankProps>>(
  "Tank",
  ({ label, onChange, orientation, color, dimensions, borderRadius }): ReactElement => {
    return (
      <Labeled {...label} onChange={onChange}>
        <Primitives.Tank
          onResize={(dims) => onChange({ dimensions: dims })}
          orientation={orientation}
          color={color}
          dimensions={dimensions}
          borderRadius={borderRadius}
        />
      </Labeled>
    );
  },
);

export const TankPreview = (props: TankProps): ReactElement => (
  <Primitives.Tank {...props} dimensions={{ width: 25, height: 50 }} />
);

export interface ReliefValveProps extends Primitives.ReliefValveProps {
  label?: LabelExtensionProps;
}

export const ReliefValve = ({
  label,
  onChange,
  orientation,
  color,
}: SymbolProps<ReliefValveProps>): ReactElement => {
  return (
    <Labeled {...label} onChange={onChange}>
      <Primitives.ReliefValve orientation={orientation} color={color} />
    </Labeled>
  );
};

export const ReliefValvePreview = (props: ReliefValveProps): ReactElement => (
  <Primitives.ReliefValve {...props} />
);

export interface RegulatorProps extends Primitives.RegulatorProps {
  label?: LabelExtensionProps;
}

export const Regulator = ({
  label,
  onChange,
  orientation,
  color,
}: SymbolProps<RegulatorProps>): ReactElement => {
  return (
    <Labeled {...label} onChange={onChange}>
      <Primitives.Regulator orientation={orientation} color={color} />
    </Labeled>
  );
};

export const RegulatorPreview = (props: RegulatorProps): ReactElement => (
  <Primitives.Regulator {...props} />
);

export interface BurstDiscProps extends Primitives.BurstDiscProps {
  label?: LabelExtensionProps;
}

export const BurstDisc = ({
  label,
  onChange,
  orientation,
  color,
}: SymbolProps<BurstDiscProps>): ReactElement => {
  return (
    <Labeled {...label} onChange={onChange}>
      <Primitives.BurstDisc orientation={orientation} color={color} />
    </Labeled>
  );
};

export const BurstDiscPreview = (props: BurstDiscProps): ReactElement => (
  <Primitives.BurstDisc {...props} />
);

export interface CapProps extends Primitives.CapProps {
  label?: LabelExtensionProps;
}

export const Cap = ({
  label,
  onChange,
  orientation,
  color,
}: SymbolProps<CapProps>): ReactElement => {
  return (
    <Labeled {...label} onChange={onChange}>
      <Primitives.Cap orientation={orientation} color={color} />
    </Labeled>
  );
};

export const CapPreview = (props: CapProps): ReactElement => (
  <Primitives.Cap {...props} />
);

export interface ManualValveProps extends Primitives.ManualValveProps {
  label?: LabelExtensionProps;
}

export const ManualValve = ({
  label,
  onChange,
  orientation,
  color,
}: SymbolProps<ManualValveProps>): ReactElement => {
  return (
    <Labeled {...label} onChange={onChange}>
      <Primitives.ManualValve orientation={orientation} color={color} />
    </Labeled>
  );
};

export const ManualValvePreview = (props: ManualValveProps): ReactElement => (
  <Primitives.ManualValve {...props} />
);

export interface FilterProps extends Primitives.FilterProps {
  label?: LabelExtensionProps;
}

export const Filter = ({
  label,
  onChange,
  orientation,
  color,
}: SymbolProps<FilterProps>): ReactElement => {
  return (
    <Labeled {...label} onChange={onChange}>
      <Primitives.Filter orientation={orientation} color={color} />
    </Labeled>
  );
};

export const FilterPreview = (props: FilterProps): ReactElement => (
  <Primitives.Filter {...props} />
);

export interface NeedleValveProps extends Primitives.NeedleValveProps {
  label?: LabelExtensionProps;
}

export const NeedleValve = ({
  label,
  onChange,
  orientation,
  color,
}: SymbolProps<NeedleValveProps>): ReactElement => {
  return (
    <Labeled {...label} onChange={onChange}>
      <Primitives.NeedleValve orientation={orientation} color={color} />
    </Labeled>
  );
};

export const NeedleValvePreview = (props: NeedleValveProps): ReactElement => (
  <Primitives.NeedleValve {...props} />
);

export interface CheckValveProps extends Primitives.CheckValveProps {
  label?: LabelExtensionProps;
}

export const CheckValve = ({
  label,
  onChange,
  orientation,
  color,
}: SymbolProps<CheckValveProps>): ReactElement => {
  return (
    <Labeled {...label} onChange={onChange}>
      <Primitives.CheckValve orientation={orientation} color={color} />
    </Labeled>
  );
};

export const CheckValvePreview = (props: CheckValveProps): ReactElement => (
  <Primitives.CheckValve {...props} />
);

export interface OrificeProps extends Primitives.OrificeProps {
  label?: LabelExtensionProps;
}

export const Orifice = ({
  label,
  onChange,
  orientation,
  color,
}: SymbolProps<OrificeProps>): ReactElement => {
  return (
    <Labeled {...label} onChange={onChange}>
      <Primitives.Orifice orientation={orientation} color={color} />
    </Labeled>
  );
};

export const OrificePreview = (props: OrificeProps): ReactElement => (
  <Primitives.Orifice {...props} />
);

export interface AngledReliefValveProps extends Primitives.AngledReliefValveProps {
  label?: LabelExtensionProps;
}

export const AngledReliefValve = ({
  label,
  onChange,
  orientation,
  color,
}: SymbolProps<AngledReliefValveProps>): ReactElement => {
  return (
    <Labeled {...label} onChange={onChange}>
      <Primitives.AngledReliefValve orientation={orientation} color={color} />
    </Labeled>
  );
};

export const AngledReliefValvePreview = (
  props: Primitives.AngledReliefValveProps,
): ReactElement => <Primitives.AngledReliefValve {...props} />;

export interface ValueProps
  extends Omit<CoreValue.UseProps, "box" | "aetherKey">,
    Primitives.ValueProps {
  position?: xy.XY;
  label?: LabelExtensionProps;
  color?: Color.Crude;
  textColor?: Color.Crude;
  tooltip?: string[];
}

interface ValueDimensionsState {
  outerBox: box.Box;
  labelBox: box.Box;
}

export const Value = Aether.wrap<SymbolProps<ValueProps>>(
  "Value",
  ({
    aetherKey,
    label,
    level = "p",
    position,
    className,
    textColor,
    color,
    telem,
    units,
    onChange,
    tooltip,
  }): ReactElement => {
    const font = Theming.useTypography(level);
    const [dimensions, setDimensions] = useState<ValueDimensionsState>({
      outerBox: box.ZERO,
      labelBox: box.ZERO,
    });

    const valueBoxHeight = (font.lineHeight + 0.5) * font.baseSize + 2;
    const resizeRef = useResize(
      useCallback((b) => {
        // Find the element with the class pluto-symbol__label that is underneath
        // the 'react-flow__node' with the data-id of aetherKey
        const label = document.querySelector(
          `.react-flow__node[data-id="${aetherKey}"] .pluto-symbol__label`,
        );
        let labelBox = { ...box.ZERO };
        if (label != null) {
          labelBox = box.construct(label);
          labelBox = box.resize(labelBox, {
            width: box.width(labelBox),
            height: box.height(labelBox),
          });
        }
        setDimensions({ outerBox: b, labelBox });
      }, []),
      {},
    );

    const { zoom } = useInitialViewport();

    const adjustedBox = adjustBox({
      labelOrientation: label?.orientation ?? "top",
      hasLabel: label?.label != null && label?.label.length > 0,
      valueBoxHeight,
      position,
      zoom,
      ...dimensions,
    });

    const { width: oWidth } = CoreValue.use({
      aetherKey,
      color: textColor,
      level,
      box: adjustedBox,
      telem,
      minWidth: 60,
    });

    return (
      <Tooltip.Dialog
        location={{ y: "top" }}
        hide={tooltip == null || tooltip.length === 0}
      >
        <Align.Space direction="y">
          {tooltip?.map((t, i) => (
            <Text.Text key={i} level="small">
              {t}
            </Text.Text>
          ))}
        </Align.Space>
        <Labeled
          className={CSS(className, CSS.B("value-labeled"))}
          ref={resizeRef}
          onChange={onChange}
          {...label}
        >
          <Primitives.Value
            color={color}
            dimensions={{
              height: valueBoxHeight,
              width: oWidth,
            }}
            units={units}
          />
        </Labeled>
      </Tooltip.Dialog>
    );
  },
);

interface AdjustBoxProps {
  labelOrientation: location.Outer;
  zoom: number;
  outerBox: box.Box;
  labelBox: box.Box;
  valueBoxHeight: number;
  position: xy.XY;
  hasLabel: boolean;
}

const LABEL_SCALE = 0.9;

const adjustBox = ({
  labelOrientation,
  outerBox,
  labelBox,
  valueBoxHeight,
  position,
  hasLabel,
  zoom,
}: AdjustBoxProps): box.Box => {
  const labelDims = xy.scale(box.dims(labelBox), 1 / (LABEL_SCALE * zoom));
  const dir = direction.construct(labelOrientation);
  if (dir === "x")
    position = xy.translateY(
      position,
      Math.max((labelDims.y - valueBoxHeight) / 2 - 1, 0),
    );
  if (hasLabel && labelOrientation === "left")
    position = xy.translateX(position, labelDims.x + 4);
  else if (hasLabel && labelOrientation === "top")
    position = xy.translateY(position, labelDims.y + 4);
  return box.construct(position.x, position.y, box.width(outerBox), valueBoxHeight);
};

export const ValuePreview = ({ color }: ValueProps): ReactElement => {
  return (
    <Primitives.Value
      color={color}
      dimensions={{
        width: 60,
        height: 25,
      }}
      units={"psi"}
    >
      <Text.Text level="p">50.00</Text.Text>
    </Primitives.Value>
  );
};

export interface SwitchProps
  extends Primitives.SwitchProps,
    Omit<Toggle.UseProps, "aetherKey"> {
  label?: LabelExtensionProps;
  control?: ControlStateProps;
}

export const Switch = Aether.wrap<SymbolProps<SwitchProps>>(
  "Switch",
  ({
    aetherKey,
    label,
    control,
    onChange,
    orientation,
    color,
    source,
    sink,
  }): ReactElement => {
    const { enabled, triggered, toggle } = Toggle.use({ aetherKey, source, sink });
    return (
      <Labeled {...label} onChange={onChange}>
        <ControlState {...control} orientation={orientation}>
          <Primitives.Switch
            enabled={enabled}
            triggered={triggered}
            onClick={toggle}
            orientation={orientation}
            color={color}
          />
        </ControlState>
      </Labeled>
    );
  },
);

export const SwitchPreview = (props: SwitchProps): ReactElement => (
  <Primitives.Switch {...props} />
);

export interface ButtonProps
  extends Omit<Primitives.ButtonProps, "label" | "onClick">,
    Omit<CoreButton.UseProps, "aetherKey"> {
  label?: LabelExtensionProps;
  control?: ControlStateProps;
}

export const Button = Aether.wrap<SymbolProps<ButtonProps>>(
  "Button",
  ({ aetherKey, label, orientation, sink, control }) => {
    const { click } = CoreButton.use({ aetherKey, sink });
    return (
      <ControlState {...control} className={CSS.B("symbol")} orientation={orientation}>
        <Primitives.Button label={label?.label} onClick={click} />
      </ControlState>
    );
  },
);

export const ButtonPreview = ({ label: _, ...props }: ButtonProps): ReactElement => (
  <Primitives.Button label="Button" {...props} />
);
