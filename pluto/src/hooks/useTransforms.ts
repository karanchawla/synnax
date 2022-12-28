import { useCallback, useState } from "react";

import { UnknownRecord } from "@/util/record";
import { ArrayTransform } from "@/util/transform";

export interface ArrayTransformEntry<E extends UnknownRecord = UnknownRecord> {
  transform: ArrayTransform<E>;
  key: string;
  priority: number;
}

export interface UseTransformsProps<E extends UnknownRecord = UnknownRecord> {
  transforms?: Array<ArrayTransformEntry<E>>;
}

export interface UseTransformsReturn<E extends UnknownRecord = UnknownRecord> {
  transform: ArrayTransform<E>;
  setTransform: (key: string, t: ArrayTransform<E>, priority?: number) => void;
  deleteTransform: (key: string) => void;
}

export const useTransforms = <E extends UnknownRecord>({
  transforms: initialTransforms = [],
}: UseTransformsProps<E>): UseTransformsReturn<E> => {
  const [transforms, setTransforms] =
    useState<Array<ArrayTransformEntry<E>>>(initialTransforms);

  const setTransform = (key: string, t: ArrayTransform<E>, priority = 0): void =>
    setTransforms((prev) => {
      const next = prev.filter((t) => t.key !== key);
      next.push({ key, transform: t, priority });
      next.sort((a, b) => b.priority - a.priority);
      return next;
    });

  const deleteTransform = (key: string): void =>
    setTransforms((prev) => prev.filter((t) => t.key !== key));

  const transform = useCallback(
    (data: E[]) => transforms.reduce((data, t) => t.transform(data), data),
    [transforms]
  );

  return { transform, setTransform, deleteTransform };
};
