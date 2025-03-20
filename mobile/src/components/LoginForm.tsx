import React from "react";
import KeyboardView from "./ui/KeyboardView";
import { TextInput } from "react-native";

type Props = {};

export default function LoginForm({}: Props) {
  return (
    <KeyboardView
      style={{
        padding: 20,
      }}
    >
      <TextInput />
    </KeyboardView>
  );
}
