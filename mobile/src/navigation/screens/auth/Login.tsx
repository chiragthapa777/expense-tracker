import { View, ScrollView } from "react-native";
import Text from "../../../components/ui/Text";
import { useColor } from "../../../theme";
import LoginTab from "../../../components/LoginTab";
import { useState } from "react";
import LoginForm from "../../../components/LoginForm";
import { Button } from "../../../components/ui/Button";
import { SafeAreaView } from "react-native-safe-area-context";
import KeyboardView from "@/components/ui/KeyboardView";

type Props = {};

const Login = ({}: Props) => {
  const [activeTab, setActiveTab] = useState<"login" | "register">("login");
  const color = useColor();

  return (
    <SafeAreaView style={{ flex: 1, backgroundColor: color.background }}>{/* If there is not header in the navigation add, safeAreaView, for iso compulsory */}
      <KeyboardView style={{ flex: 1 }}>
        <ScrollView
          contentContainerStyle={{
            flexGrow: 1, // Ensures content can grow and scroll
            paddingBottom: 20, // Adds space at the bottom for scrolling
          }}
        >
          <View style={{ padding: 20 }}>
            <Text size="3xl" weight="bold" color={color.primary}>
              Expense Tracker
            </Text>
            <Text size="3xl">Welcome</Text>
          </View>
          <LoginTab activeTab={activeTab} setActiveTab={setActiveTab}>
            <View style={{ flexGrow: 1 }}>
              <LoginForm />
            </View>
          </LoginTab>
        </ScrollView>
      </KeyboardView>
    </SafeAreaView>
  );
};

export default Login;