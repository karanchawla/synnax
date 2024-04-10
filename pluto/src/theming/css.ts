// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { CSS, createHexOpacityVariants, unitProperty } from "@/css";
import { type Theme } from "@/theming/core/theme";

const OPACITIES: readonly number[] = [90, 80, 70, 60, 50, 40, 30, 20, 10];

export const toCSSVars = (theme: Theme): Record<string, number | string | undefined> =>
  Object.entries({
    "theme-name": theme.name,
    "theme-key": theme.key,
    "primary-m1": theme.colors.primary.m1.hex,
    "primary-z": theme.colors.primary.z.hex,
    "primary-p1": theme.colors.primary.p1.hex,
    ...createHexOpacityVariants("primary-z", theme.colors.primary.z, OPACITIES),
    "secondary-m1": theme.colors.secondary.m1.hex,
    "secondary-z": theme.colors.secondary.z.hex,
    "secondary-p1": theme.colors.secondary.p1.hex,
    ...createHexOpacityVariants("secondary-z", theme.colors.secondary.z, OPACITIES),
    "gray-l0": theme.colors.gray.l0.hex,
    "gray-l0-rgb": theme.colors.gray.l0.rgbString,
    ...createHexOpacityVariants("gray-l0", theme.colors.gray.l0, OPACITIES),
    "gray-l1": theme.colors.gray.l1.hex,
    ...createHexOpacityVariants("gray-l1", theme.colors.gray.l1, OPACITIES),
    "gray-l2": theme.colors.gray.l2.hex,
    ...createHexOpacityVariants("gray-l2", theme.colors.gray.l2, OPACITIES),
    "gray-l3": theme.colors.gray.l3.hex,
    ...createHexOpacityVariants("gray-l3", theme.colors.gray.l3, OPACITIES),
    "gray-l4": theme.colors.gray.l4.hex,
    "gray-l5": theme.colors.gray.l5.hex,
    ...createHexOpacityVariants("gray-l5", theme.colors.gray.l5, OPACITIES),
    "gray-l6": theme.colors.gray.l6.hex,
    ...createHexOpacityVariants("gray-l6", theme.colors.gray.l6, OPACITIES),
    "gray-l7": theme.colors.gray.l7.hex,
    ...createHexOpacityVariants("gray-l7", theme.colors.gray.l7, OPACITIES),
    "gray-l8": theme.colors.gray.l8.hex,
    "gray-l9": theme.colors.gray.l9.hex,
    ...createHexOpacityVariants("gray-l9", theme.colors.gray.l9, OPACITIES),
    "gray-l10": theme.colors.gray.l10.hex,
    "logo-color": theme.colors.logo,
    "error-m1": theme.colors.error.m1.hex,
    "error-z": theme.colors.error.z.hex,
    ...createHexOpacityVariants("error-z", theme.colors.error.z, OPACITIES),
    "error-p1": theme.colors.error.p1.hex,
    white: theme.colors.white.hex,
    "white-rgb": theme.colors.white.rgbString,
    black: theme.colors.black.hex,
    "black-rgb": theme.colors.black.rgbString,
    "text-color": theme.colors.text.hex,
    "text-color-rgb": theme.colors.text.rgbString,
    "border-color": theme.colors.border.hex,
    "base-size": unitProperty(theme.sizes.base, "px"),
    "border-radius": unitProperty(theme.sizes.border.radius, "px"),
    "border-width": unitProperty(theme.sizes.border.width, "px"),
    "schematic-element-stroke-width": unitProperty(theme.sizes.schematic.elementStrokeWidth, "px"),
    "font-family": theme.typography.family,
    "h1-size": unitProperty(theme.typography.h1.size, "rem"),
    "h1-weight": theme.typography.h1.weight,
    "h1-line-height": unitProperty(theme.typography.h1.lineHeight, "rem"),
    "h2-size": unitProperty(theme.typography.h2.size, "rem"),
    "h2-weight": theme.typography.h2.weight,
    "h2-line-height": unitProperty(theme.typography.h2.lineHeight, "rem"),
    "h3-size": unitProperty(theme.typography.h3.size, "rem"),
    "h3-weight": theme.typography.h3.weight,
    "h3-line-height": unitProperty(theme.typography.h3.lineHeight, "rem"),
    "h4-size": unitProperty(theme.typography.h4.size, "rem"),
    "h4-weight": theme.typography.h4.weight,
    "h4-line-height": unitProperty(theme.typography.h4.lineHeight, "rem"),
    "h5-size": unitProperty(theme.typography.h5.size, "rem"),
    "h5-weight": theme.typography.h5.weight,
    "h5-line-height": unitProperty(theme.typography.h5.lineHeight, "rem"),
    "h5-text-transform": theme.typography.h5.textTransform,
    "p-size": unitProperty(theme.typography.p.size, "rem"),
    "p-weight": theme.typography.p.weight,
    "p-line-height": unitProperty(theme.typography.p.lineHeight, "rem"),
    "small-size": unitProperty(theme.typography.small.size, "rem"),
    "small-weight": theme.typography.small.weight,
    "small-line-height": unitProperty(theme.typography.small.lineHeight, "rem"),
  }).reduce<Record<string, number | string | undefined>>(
    (acc, [key, value]) => ({
      ...acc,
      [CSS.var(key)]: value,
    }),
    {},
  );
