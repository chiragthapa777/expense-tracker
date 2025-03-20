import React, { useState } from "react";
import { StyleProp, TextInput, TextInputProps } from "react-native";
import { useColor } from "../../theme";

type Props = TextInputProps & {
  style?: StyleProp<TextInput>;
  hasError?: boolean;
};

export const CustomTextInput = React.forwardRef<
  React.ElementRef<typeof TextInput>,
  Props
>(({ style, onFocus, onBlur, hasError, ...props }: Props, ref) => {
  const color = useColor();
  const [active, setActive] = useState(false);
  return (
    <TextInput
      ref={ref}
      {...props}
      onFocus={(e) => {
        setActive(true);
        onFocus?.(e);
      }}
      onBlur={(e) => {
        setActive(false);
        onBlur?.(e);
      }}
      selectionColor={color.primary}
      placeholderTextColor={color.border}
      style={[
        {
          borderWidth: 1,
          borderColor: hasError
            ? color.error
            : active
            ? color.primary
            : color.border,
          borderRadius: 8,
          padding: 15,
          color: color.text,
        },
        style,
      ]}
    />
  );
});
