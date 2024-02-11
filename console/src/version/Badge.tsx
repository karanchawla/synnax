// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.
import { type ReactElement, useState} from "react";
import React from 'react';
import { Text } from "@synnaxlabs/pluto";
import { type Optional } from "@synnaxlabs/x";
import { selectUpdate, useSelect, useselectUpdate as useSelectUpdate } from "@/version/selectors";
import { checkUpdate, installUpdate } from '@tauri-apps/api/updater';
import { FaExclamationCircle } from "react-icons/fa";
import { Button} from "@synnaxlabs/pluto"
import { Layout } from "@/layout";
import { versionWindowLayout } from "./updateDialog";

type BadgeProps<L extends Text.Level> = Optional<Text.TextProps<L>, "level">;

export const Badge = <L extends Text.Level>({
  level = "p",
  ...props
}: BadgeProps<L>): ReactElement => {
  const v = useSelect();
  const openWindow = Layout.usePlacer();
  const updateAvailable = useSelectUpdate();
  if (updateAvailable){
    return (<Button.Button variant="text" tooltip ={<h1>Update Available!</h1>}level={level} color={"white"} onClick={() => openWindow(versionWindowLayout)} 
    {...props}> {"v" + v}
    <FaExclamationCircle style={{ marginLeft: '4px', color: 'red', fontSize: '1.8em' }} />
    </Button.Button>) 
  }
  else{
    return((<Text.Text level={level} color={"white"} {...props} style={{ display: 'flex', alignItems: 'center' }}>
    {"v" + v}
  </Text.Text>));
  }  

  };

