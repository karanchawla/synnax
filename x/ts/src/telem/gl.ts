// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

export interface GLBufferController {
  createBuffer: () => WebGLBuffer | null;
  bufferData: ((target: number, data: ArrayBufferLike, usage: number) => void) &
    ((target: number, size: number, usage: number) => void);
  bufferSubData: (target: number, offset: number, data: ArrayBufferLike) => void;
  bindBuffer: (target: number, buffer: WebGLBuffer | null) => void;
  deleteBuffer: (buffer: WebGLBuffer | null) => void;
  ARRAY_BUFFER: number;
  STATIC_DRAW: number;
  DYNAMIC_DRAW: number;
}

export type GLBufferUsage = "static" | "dynamic";
