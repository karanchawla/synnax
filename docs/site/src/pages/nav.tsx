// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { conceptsNav } from "@/pages/concepts/nav";
import { consoleNav } from "@/pages/console/nav";
import { deploymentNav } from "@/pages/deployment/nav";
import { guidesNav } from "@/pages/guides/nav";
import { pythonClientNav } from "@/pages/python-client/nav";
import { typescriptClientNav } from "@/pages/typescript-client/nav";

export const pages = [
  {
    name: "Get Started",
    key: "/",
    href: "/",
  },
  guidesNav,
  conceptsNav,
  deploymentNav,
  pythonClientNav,
  typescriptClientNav,
  consoleNav,
];
