import {
  Button,
  Keyboard,
  KeyboardAvoidingView,
  Platform,
  StyleProp,
  StyleSheet,
  TextInput,
  TouchableWithoutFeedback,
  View,
  ViewStyle,
} from "react-native";
import { useColor } from "../../theme";
import React from "react";

type Props = {
  children: React.ReactNode;
  style?: StyleProp<ViewStyle>;
};

const KeyboardView = ({ children, style }: Props) => {
  const color = useColor();
  return (
    <KeyboardAvoidingView
      behavior={Platform.OS === "ios" ? "padding" : "height"}
      style={[
        {
          //   backgroundColor: color.surface,
        },
        style,
      ]}
    >
      <TouchableWithoutFeedback onPress={Keyboard.dismiss}>
        <>{children}</>
      </TouchableWithoutFeedback>
    </KeyboardAvoidingView>
  );
};

export default KeyboardView;
