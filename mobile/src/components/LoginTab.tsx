import { Pressable, View } from "react-native";
import React from "react";
import * as Haptics from "expo-haptics";
import Text from "./ui/Text";
import { useColor } from "../theme";

type Props = {
  children: React.ReactNode;
  setActiveTab: React.Dispatch<React.SetStateAction<"login" | "register">>;
  activeTab: "login" | "register";
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
            setActiveTab("login");
          }}
          android_ripple={{
            color: color.primaryDim,
            radius: 80,
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
                borderBottomWidth: 3,
                borderColor: "transparent",
              },
              activeTab === "login" && {
                borderBottomWidth: 3,
                borderColor: color.primary,
                color: color.primary,
              },
            ]}
          >
            Login To Account
          </Text>
        </Pressable>
        <Pressable
          onPress={() => {
            Haptics.selectionAsync();
            setActiveTab("register");
          }}
          android_ripple={{
            color: color.primaryDim,
            radius: 80,
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
                borderBottomWidth: 3,
                borderColor: "transparent",
              },
              activeTab === "register" && {
                borderBottomWidth: 3,
                borderColor: color.primary,
                color: color.primary,
              },
            ]}
          >
            Create New Account
          </Text>
        </Pressable>
      </View>
      {children}
    </View>
  );
};

export default LoginTab;
