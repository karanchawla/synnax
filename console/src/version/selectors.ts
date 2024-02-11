// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { useMemoSelect } from "@/hooks";
import { type StoreState } from "@/version/slice";
import { SliceState } from "@synnaxlabs/drift";

export const select = (state: StoreState): string => state.version.version;

export const selectUpdate = (state: StoreState): boolean => state.version.updateAvailable;

export const useSelect = (): string => useMemoSelect(select, []);

export const useselectUpdate = (): boolean => useMemoSelect(selectUpdate, []);

