import React from "react";
import { StyleProp, View, ViewStyle } from "react-native";

type Props = {
  children: React.ReactNode;
  style: StyleProp<ViewStyle>;
};

export const ViewY = ({ children, style }: Props) => {
  return (
    <View
      style={[
        {
          flex: 1,
        },
        style,
      ]}
    >
      {children}
    </View>
  );
};
