import { Pressable, View } from "react-native";
import React from "react";
import * as Haptics from "expo-haptics";
import Text from "./ui/Text";
import { useColor } from "../theme";

type Props = {
  children: React.ReactNode;
  setActiveTab: React.Dispatch<React.SetStateAction<"email" | "phone">>;
  activeTab: "email" | "phone";
};

const LoginTab = ({ children, activeTab, setActiveTab }: Props) => {
  const color = useColor();
  return (
    <View
      style={{
        flex: 1,
        backgroundColor: color.card,
        gap: 10,
        borderTopEndRadius: 20,
        borderTopStartRadius: 20,
      }}
    >
      <View
        style={{
          paddingHorizontal: 20,
          flexDirection: "row",
          borderBottomWidth: 1,
          borderBlockColor: color.border,
          justifyContent: "center",
          alignItems: "center",
          gap: 10,
        }}
      >
        <Pressable
          onPress={() => {
            Haptics.selectionAsync();
            setActiveTab("email");
          }}
          android_ripple={{
            color: color.primaryDim,
            radius: 70,
          }}
          style={{
            paddingHorizontal: 20,
            borderRadius: 30,
          }}
        >
          <Text
            size="md"
            weight="medium"
            style={[
              {
                paddingVertical: 15,
              },
              activeTab === "email" && {
                borderBottomWidth: 3,
                borderColor: color.primary,
                color: color.primary,
              },
            ]}
          >
            Email Address
          </Text>
        </Pressable>
        <Pressable
          onPress={() => {
            Haptics.selectionAsync();
            setActiveTab("phone");
          }}
          android_ripple={{
            color: color.primaryDim,
            radius: 70,
          }}
          style={{
            paddingHorizontal: 20,
            borderRadius: 30,
          }}
        >
          <Text
            size="md"
            weight="medium"
            style={[
              {
                paddingVertical: 15,
              },
              activeTab === "phone" && {
                borderBottomWidth: 3,
                borderColor: color.primary,
                color: color.primary,
              },
            ]}
          >
            Phone Number
          </Text>
        </Pressable>
      </View>
      {children}
    </View>
  );
};

export default LoginTab;
