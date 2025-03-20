import { View } from "react-native";
import Text from "../../../components/ui/Text";
import { useColor } from "../../../theme";
import LoginTab from "../../../components/LoginTab";
import { useState } from "react";
import LoginForm from "../../../components/LoginForm";
import { Button } from "../../../components/ui/Button";

type Props = {};

const Login = ({}: Props) => {
  const [activeTab, setActiveTab] = useState<"email" | "phone">("email");

  const color = useColor();
  return (
    <View
      style={{
        flex: 1,
      }}
    >
      <View
        style={{
          padding: 20,
        }}
      >
        <Text size="3xl" weight="bold" color={color.primary}>
          Expense Tracker
        </Text>
        <Text size="3xl">Welcome</Text>
        <Text>Login</Text>
      </View>
      <LoginTab activeTab={activeTab} setActiveTab={setActiveTab}>
        <>
          <LoginForm />
          <View
            style={{
              borderWidth: 1,
              flex: 1,
            }}
          >
            <View
              style={{
                width: "100%",
                flexDirection:"row"
              }}
            >
              <Button variant="default" size="default" >
                Test
              </Button>
            </View>
          </View>
        </>
      </LoginTab>
    </View>
  );
};

export default Login;
