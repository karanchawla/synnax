// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { type ReactElement, useEffect, useState } from "react";

import { Icon } from "@synnaxlabs/media";
import { Button } from "@synnaxlabs/pluto/button";
import { Dropdown } from "@synnaxlabs/pluto/dropdown";
import { Tabs } from "@synnaxlabs/pluto/tabs";
import { Tree } from "@synnaxlabs/pluto/tree";

import { componentsPages, guidesPages } from "@/pages/nav";

export type PageNavNode = Tree.Node;

export interface TOCProps {
  currentPage: string;
}

export const useDocumentSize = (): number | null => {
  const [width, setWidth] = useState<number | null>(null);
  useEffect(() => {
    const handleResize = (): void => setWidth(document.documentElement.clientWidth);
    window.addEventListener("resize", handleResize);
    handleResize();
    return () => window.removeEventListener("resize", handleResize);
  }, []);
  return width;
};

interface ReferenceTreeProps {
  currentPage: string;
}

const Reference = ({ currentPage }: ReferenceTreeProps): ReactElement => {
  let parts = currentPage.split("/").filter((part) => part !== "");
  if (parts.length === 1) parts = componentsPages.map((p) => p.key);
  const treeProps = Tree.use({
    nodes: componentsPages,
    initialExpanded: parts,
    sort: false,
  });
  return (
    <Tree.Tree
      {...treeProps}
      className="tree reference-tree styled-scrollbar"
      itemHeight={35}
      virtual={false}
      selected={[currentPage]}
      useMargin
    />
  );
};

const Guides = ({ currentPage }: TOCProps): ReactElement => {
  let parts = currentPage.split("/").filter((part) => part !== "");
  if (parts.length === 1) parts = guidesPages.map((p) => p.key);
  const treeProps = Tree.use({
    nodes: guidesPages,
    initialExpanded: parts,
    sort: false,
  });
  return (
    <Tree.Tree
      {...treeProps}
      className="tree role-tree styled-scrollbar"
      itemHeight={35}
      virtual={false}
      selected={[currentPage]}
      useMargin
    />
  );
};

export const PageNav = ({ currentPage }: TOCProps): ReactElement | null => {
  const width = useDocumentSize();

  // Split the current page by slashes and remove and get the first part
  const selectedTab = currentPage.split("/").filter((part) => part !== "")[0];

  const { visible, toggle } = Dropdown.use({ initialVisible: false });

  const content: Tabs.TabsProps["content"] = ({ tabKey }) => {
    switch (tabKey) {
      case "guides":
        return <Guides currentPage={currentPage} />;
      default:
        return <Reference currentPage={currentPage} />;
    }
  };

  const tabsProps = Tabs.useStatic({
    selected: selectedTab,
    tabs: [
      { tabKey: "guides", name: "Guides" },
      { tabKey: "reference", name: "Reference" },
    ],
    content,
  });

  const tree = <Tabs.Tabs {...tabsProps} />;

  if (width == null) return null;
  if (width > 700) return tree;
  return (
    <Dropdown.Dialog visible={visible} bordered={false} location="top">
      <Button.Button
        justify="spaceBetween"
        endIcon={<Icon.Copy />}
        variant="text"
        onClick={() => toggle(!visible)}
        size="large"
        style={{
          border: "none",
        }}
      >
        Menu
      </Button.Button>
      {tree}
    </Dropdown.Dialog>
  );
};