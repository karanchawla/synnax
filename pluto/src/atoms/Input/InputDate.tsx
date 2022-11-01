import { ComponentSize } from "@/util";

export interface InputDateProps {
  size?: ComponentSize;
  onChange?: (value: string) => void;
}

export const InputDate = ({ size, onChange }: InputDateProps) => {
  return (
    <input
      type="date"
      onChange={(e) => onChange && onChange(e.target.value)}
      className={`pluto-input__input pluto-input__input--${size}`}
    />
  );
};
