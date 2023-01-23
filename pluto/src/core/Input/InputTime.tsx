// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { forwardRef } from "react";

import { TimeStamp } from "@synnaxlabs/x";

import { Input } from "./Input";
import { InputBaseProps } from "./types";

export interface InputTimeProps extends InputBaseProps<number> {}

export const InputTime = forwardRef<HTMLInputElement, InputTimeProps>(
  ({ size = "medium", value, onChange, ...props }: InputTimeProps, ref) => {
    return (
      <Input
        ref={ref}
        type="time"
        step="1"
        value={new TimeStamp(value, "UTC").fString("time", "local")}
        onChange={(value) =>
          value.length > 0 && onChange(new TimeStamp(value, "local").valueOf())
        }
        {...props}
      />
    );
  }
);
InputTime.displayName = "InputTime";