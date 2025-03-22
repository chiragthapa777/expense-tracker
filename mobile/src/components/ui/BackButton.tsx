import React from "react";
import { Button } from "./Button";
import Ionicons from "@expo/vector-icons/Ionicons";
import { useColor } from "@/theme";
import { useNavigation } from "@react-navigation/native";

type Props = {};

export default function BackButton({}: Props) {
  const color = useColor();
  const navigation = useNavigation();
  return (
    <Button
      size="icon"
      variant="ghost"
      style={{
        position: "absolute",
        top: 10,
        left: 10,
        zIndex: 90,
      }}
      onPress={() => {
        navigation.goBack();
      }}
    >
      <Ionicons name="arrow-back-sharp" size={28} color={color.text} />
    </Button>
  );
}
