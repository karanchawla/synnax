// Copyright 2024 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { type runtime } from "@synnaxlabs/x";
import { type ReactElement } from "react";

import { MacOS } from "@/os/Controls/Mac";
import { type InternalControlsProps } from "@/os/Controls/types";
import { Windows } from "@/os/Controls/Windows";
import { use } from "@/os/use";

const Variants: Record<runtime.OS, React.FC<InternalControlsProps>> = {
  MacOS,
  Windows,
  Linux: Windows,
  Docker: Windows,
};

const DEFAULT_OS = "Windows";

export interface ControlsProps extends InternalControlsProps {
  visibleIfOS?: runtime.OS;
}

export const Controls = ({
  forceOS,
  visibleIfOS,
  ...props
}: ControlsProps): ReactElement | null => {
  const os = use({ force: forceOS, default: DEFAULT_OS }) as runtime.OS;
  const C = Variants[os];
  if (visibleIfOS != null && visibleIfOS !== os) return null;
  return <C {...props} />;
};
