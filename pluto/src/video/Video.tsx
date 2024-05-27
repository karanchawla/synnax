// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { forwardRef, type ComponentPropsWithoutRef, type ReactElement } from "react";

import { CSS } from "@/css";

import "@/video/Video.css";

export interface VideoProps extends ComponentPropsWithoutRef<"video"> {
  href: string;
}

export const Video = forwardRef<HTMLVideoElement, VideoProps>(
  ({ href, className, ...props }: VideoProps, ref): ReactElement => (
    <video ref={ref} className={CSS(CSS.B("video"), className)} {...props}>
      <source id={href} src={href} type="video/mp4" />
    </video>
  ),
);
Video.displayName = "Video";
