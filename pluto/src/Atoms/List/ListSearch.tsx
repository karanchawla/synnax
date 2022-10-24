import { ComponentType, useEffect } from "react";
import { useSearch } from "../../Hooks";
import { Input as DefaultInput, InputProps } from "../Input";
import { useListContext } from "./ListContext";
import { ListEntry } from "./Types";
import "./ListSearch.css";

export interface ListSearchProps<E extends ListEntry> {
  Input?: ComponentType<InputProps>;
}

export default function ListSearch<E extends ListEntry>({
  Input = DefaultInput,
}: ListSearchProps<E>) {
  const [query, setQuery, search] = useSearch<E>();
  const { setTransform } = useListContext<E>();
  useEffect(() => setTransform("search", search), [search]);
  return (
    <Input
      value={query}
      onChange={(e) => setQuery(e.target.value)}
      className="pluto-list__search__input"
    />
  );
}
