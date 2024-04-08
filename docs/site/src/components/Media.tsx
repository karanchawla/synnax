// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { useEffect, type DetailedHTMLProps, type ReactElement } from "react";

import { Video as Core } from "@synnaxlabs/pluto/video";

export interface VideoProps
  extends DetailedHTMLProps<
    React.VideoHTMLAttributes<HTMLVideoElement>,
    HTMLVideoElement
  > {
  id: string;
  themed?: boolean;
}

const CDN_ROOT = "https://synnax.nyc3.cdn.digitaloceanspaces.com/docs";

export const Video = ({ id, ...props }: VideoProps): ReactElement => {
  const theme = localStorage.getItem("theme") ?? "light";
  const modifier = theme?.toLowerCase().includes("dark") ? "dark" : "light";
  return (
    <Core.Video
      href={`${CDN_ROOT}/${id}-${modifier}.mp4`}
      loop
      autoPlay
      muted
      {...props}
    />
  );
};

export const Image = ({ id, themed = true }: VideoProps): ReactElement => {
  let url = `${CDN_ROOT}/${id}`;
  if (themed) {
    const theme = localStorage.getItem("theme") ?? "light";
    const modifier = theme?.toLowerCase().includes("dark") ? "dark" : "light";
    url += `-${modifier}`;
  }
  url += ".png";
  return <img src={url} className="image" />;
};
