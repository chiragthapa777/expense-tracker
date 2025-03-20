import {
  Keyboard,
  KeyboardAvoidingView,
  Platform,
  StyleProp,
  TouchableWithoutFeedback,
  ViewStyle,
} from "react-native";
import React from "react";
import { useColor } from "../../theme";

type Props = {
  children: React.ReactNode;
  style?: StyleProp<ViewStyle>;
};

const KeyboardView = ({ children, style }: Props) => {
  return (
    <KeyboardAvoidingView
      behavior={Platform.OS === "ios" ? "padding" : "height"}
      style={[{ flex: 1 }, style]} // Ensure it takes full space
      // keyboardVerticalOffset={Platform.OS === "ios" ? 0 : 10} // Adjust for Android status bar if needed
    >
      <TouchableWithoutFeedback onPress={Keyboard.dismiss}>
        <>{children}</>
      </TouchableWithoutFeedback>
    </KeyboardAvoidingView>
  );
};

export default KeyboardView;
