// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { useHeaderContext } from "./Header";

import { Button, ButtonProps } from "@/core/Button";
import { Typography } from "@/core/Typography";
import { CSS } from "@/css";

export interface HeaderButtonTitleProps extends Omit<ButtonProps, "variant" | "size"> {}

export const HeaderButtonTitle = ({
  children = "",
  className,
  onClick,
  ...props
}: HeaderButtonTitleProps): JSX.Element => {
  const { level } = useHeaderContext();
  return (
    <Button
      variant="text"
      size={Typography.LevelComponentSizes[level]}
      onClick={onClick}
      className={CSS(CSS.B("header-button-title"), className)}
      sharp
      {...props}
    >
      {children}
    </Button>
  );
};
