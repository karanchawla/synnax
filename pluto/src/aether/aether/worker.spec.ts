// Copyright 2024 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { alamos } from "@synnaxlabs/alamos";
import { beforeEach, describe, expect, it, vi } from "vitest";
import { z } from "zod";

import { aether } from "@/aether/aether";

export const exampleProps = z.object({
  x: z.number(),
});

interface InternalState {
  contextValue: number;
}

class ExampleLeaf extends aether.Leaf<typeof exampleProps, InternalState> {
  updatef = vi.fn();
  deletef = vi.fn();
  schema = exampleProps;

  async afterUpdate(): Promise<void> {
    this.updatef(this.ctx);
    this.internal.contextValue = this.ctx.getOptional("key") ?? 0;
  }

  async afterDelete(): Promise<void> {
    this.deletef();
  }
}

class ExampleComposite extends aether.Composite<
  typeof exampleProps,
  {},
  ExampleLeaf | ContextSetterComposite
> {
  updatef = vi.fn();
  deletef = vi.fn();

  schema = exampleProps;

  async afterUpdate(): Promise<void> {
    this.updatef(this.ctx);
  }

  async afterDelete(): Promise<void> {
    this.deletef();
  }
}

class ContextSetterComposite extends aether.Composite<
  typeof exampleProps,
  {},
  ExampleLeaf
> {
  updatef = vi.fn();
  deletef = vi.fn();

  schema = exampleProps;

  async afterUpdate(): Promise<void> {
    this.updatef(this.ctx);
    this.ctx.set("key", this.state.x);
  }

  async afterDelete(): Promise<void> {
    this.deletef();
  }
}

const REGISTRY: aether.ComponentRegistry = {
  leaf: ExampleLeaf,
  composite: ExampleComposite,
  context: ContextSetterComposite,
};

const MockSender = {
  send: vi.fn(),
};

const ctx = new aether.Context(MockSender, REGISTRY);

const leafUpdate: aether.Update = {
  ctx,
  variant: "state",
  type: "leaf",
  path: ["test"],
  state: { x: 1 },
  instrumentation: alamos.NOOP,
};

const compositeUpdate: aether.Update = {
  ctx,
  variant: "state",
  type: "composite",
  path: ["test"],
  state: { x: 1 },
  instrumentation: alamos.NOOP,
};

const contextUpdate: aether.Update = {
  ctx,
  variant: "context",
  type: "context",
  path: [],
  state: {},
  instrumentation: alamos.NOOP,
};

describe("Aether Worker", () => {
  describe("AetherLeaf", () => {
    let leaf: ExampleLeaf;
    beforeEach(async () => {
      leaf = await ctx.create(leafUpdate);
    });
    describe("internalUpdate", () => {
      it("should throw an error if the path is empty", async () => {
        await expect(
          leaf.internalUpdate({ ...leafUpdate, path: [] }),
        ).rejects.toThrowError(/empty path/);
      });
      it("should throw an error if the path has a subpath", async () => {
        await expect(
          async () =>
            await leaf.internalUpdate({ ...leafUpdate, path: ["test", "dog"] }),
        ).rejects.toThrowError(/subPath/);
      });
      it("should throw an error if the path does not have the correct key", async () => {
        await expect(
          leaf.internalUpdate({ ...leafUpdate, path: ["dog"] }),
        ).rejects.toThrowError(/key/);
      });
      it("should correctly internalUpdate the state", async () => {
        await leaf.internalUpdate({ ...leafUpdate, state: { x: 2 } });
        expect(leaf.state).toEqual({ x: 2 });
      });
      it("should call the handleUpdate", async () => {
        await leaf.internalUpdate({ ...leafUpdate, state: { x: 2 } });
        expect(leaf.updatef).toHaveBeenCalledTimes(2);
      });
    });
    describe("internalDelete", () => {
      it("should call the bound onDelete handler", async () => {
        await leaf.internalDelete(["test"]);
        expect(leaf.deletef).toHaveBeenCalledTimes(1);
      });
    });
    describe("setState", () => {
      it("should communicate the state call to the main thread Sender", () => {
        leaf.setState((p) => ({ ...p }));
        expect(MockSender.send).toHaveBeenCalledTimes(1);
      });
    });
  });

  describe("AetherComposite", () => {
    let composite: ExampleComposite;
    beforeEach(async () => {
      composite = await ctx.create(compositeUpdate);
    });
    describe("setState", () => {
      it("should set the state of the composite's leaf if the path has one element", async () => {
        await composite.internalUpdate({ ...compositeUpdate, state: { x: 2 } });
      });
      it("should create a new leaf if the path has more than one element and the leaf does not exist", async () => {
        await composite.internalUpdate({
          ...leafUpdate,
          path: ["test", "dog"],
          state: { x: 2 },
        });
        expect(composite.children).toHaveLength(1);
        const c = composite.children[0];
        expect(c.key).toEqual("dog");
        expect(c.state).toEqual({ x: 2 });
      });
      it("should set the state of the composite's leaf if the path has more than one element and the leaf exists", async () => {
        await composite.internalUpdate({
          ...leafUpdate,
          path: ["test", "dog"],
          state: { x: 2 },
        });
        await composite.internalUpdate({
          ...leafUpdate,
          path: ["test", "dog"],
          state: { x: 3 },
        });
        expect(composite.children).toHaveLength(1);
        expect(composite.children[0].state).toEqual({ x: 3 });
      });
      it("should throw an error if the path is too deep and the child does not exist", async () => {
        await expect(
          composite.internalUpdate({
            ...compositeUpdate,
            path: ["test", "dog", "cat"],
          }),
        ).rejects.toThrowError();
      });
    });
    describe("internalDelete", () => {
      it("should remove a child from the list of children", async () => {
        await composite.internalUpdate({ ...compositeUpdate, path: ["test", "dog"] });
        await composite.internalDelete(["test", "dog"]);
        expect(composite.children).toHaveLength(0);
      });
      it("should call the deletion hook on the child of a composite", async () => {
        await composite.internalUpdate({ ...leafUpdate, path: ["test", "dog"] });
        const c = composite.children[0];
        await composite.internalDelete(["test", "dog"]);
        expect(c.deletef).toHaveBeenCalled();
      });
    });

    describe("context propagation", () => {
      it("should properly propagate an existing context change to its children", async () => {
        await composite.internalUpdate({ ...leafUpdate, path: ["test", "dog"] });
        await composite.internalUpdate({ ...leafUpdate, path: ["test", "cat"] });
        expect(composite.children).toHaveLength(2);
        await composite.internalUpdate({ ...contextUpdate });
        expect(composite.updatef).toHaveBeenCalledTimes(2);
        composite.children.forEach((c) => expect(c.updatef).toHaveBeenCalledTimes(2));
      });
      it("should progate a new context change to its children", async () => {
        const c = new ContextSetterComposite({ ...compositeUpdate });
        await c.internalUpdate({ ...leafUpdate, path: ["test", "dog"] });
        await c.internalUpdate({ ...compositeUpdate });
        expect(c.children).toHaveLength(1);
        c.children.forEach((c) => expect(c.updatef).toHaveBeenCalledTimes(2));
      });
      it("should correctly separate individual contexts", async () => {
        await composite.internalUpdate({
          ctx,
          variant: "state",
          type: "context",
          path: ["test", "dog"],
          state: { x: 1 },
          instrumentation: alamos.NOOP,
        });
        await composite.internalUpdate({
          ctx,
          variant: "state",
          type: "context",
          path: ["test", "cat"],
          state: { x: 2 },
          instrumentation: alamos.NOOP,
        });
        expect(composite.children).toHaveLength(2);
        await composite.internalUpdate({
          ctx,
          variant: "state",
          type: "leaf",
          path: ["test", "dog", "dogleaf"],
          state: { x: 3 },
          instrumentation: alamos.NOOP,
        });
        await composite.internalUpdate({
          ctx,
          variant: "state",
          type: "leaf",
          path: ["test", "cat", "catLeaf"],
          state: { x: 4 },
          instrumentation: alamos.NOOP,
        });
        await composite.internalUpdate({
          ctx,
          variant: "state",
          type: "context",
          path: ["test", "dog"],
          state: { x: 5 },
          instrumentation: alamos.NOOP,
        });
      });
    });
  });
});
