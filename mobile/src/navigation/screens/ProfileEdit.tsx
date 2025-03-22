import BackButton from "@/components/ui/BackButton";
import KeyboardView from "@/components/ui/KeyboardView";
import Text from "@/components/ui/Text";
import React from "react";
import { SafeAreaView, ScrollView } from "react-native";
type Props = {};

export default function ProfileEdit({}: Props) {
  return (
    <SafeAreaView
      style={{
        flex: 1,
      }}
    >
      <KeyboardView style={{ flex: 1 }}>
        <ScrollView
          contentContainerStyle={{
            flexGrow: 1,
            position:"relative"
          }}
        >
          <BackButton />
          <Text>Edit Profile</Text>
        </ScrollView>
      </KeyboardView>
    </SafeAreaView>
  );
}
