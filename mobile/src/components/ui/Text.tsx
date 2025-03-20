import React from "react";
import { ColorValue, Text as RnText, StyleProp, TextStyle } from "react-native";
import { useColor } from "../../theme"; // Adjust path as needed

type Props = {
  children: React.ReactNode;
  style?: StyleProp<TextStyle>;
  size?: "xs" | "sm" | "md" | "lg" | "xl" | "2xl" | "3xl" | "4xl";
  weight?:
    | "thin"
    | "light"
    | "normal"
    | "medium"
    | "semibold"
    | "bold"
    | "extrabold"
    | "black";
  color?: ColorValue;
};

export default function Text({
  children,
  style,
  size = "md",
  weight = "normal",
  color,
}: Props) {
  const colors = useColor();

  const getFontSize = () => {
    switch (size) {
      case "xs":
        return 12;
      case "sm":
        return 14;
      case "md":
        return 16;
      case "lg":
        return 18;
      case "xl":
        return 24;
      case "2xl":
        return 28;
      case "3xl":
        return 32;
      case "4xl":
        return 36;
      default:
        return 16;
    }
  };

  const getFontWeight = () => {
    switch (weight) {
      case "thin":
        return "100";
      case "light":
        return "300";
      case "normal":
        return "400";
      case "medium":
        return "500";
      case "semibold":
        return "600";
      case "bold":
        return "700";
      case "extrabold":
        return "800";
      case "black":
        return "900";
      default:
        return "400";
    }
  };

  return (
    <RnText
      style={[
        {
          color: color ?? colors.text,
          fontSize: getFontSize(),
          fontWeight: getFontWeight(),
        },
        style,
      ]}
    >
      {children}
    </RnText>
  );
}
