import { View } from "react-native";
import Text from "../../components/ui/Text";
import { useColor } from "../../theme";
import { useWaitingDots } from "../../hooks/useWaitingDots";
import { useEffect } from "react";
import { getData } from "../../utils/asyncStore";
import { userAuthStore } from "../../store/auth";
import { useNavigation } from "@react-navigation/native";

const Welcome = () => {
  const color = useColor();
  const { waitingDots } = useWaitingDots();
  const user = userAuthStore((s) => s.user);
  const navigation = useNavigation();

  useEffect(() => {
    const loadUser = async () => {
      const token = await getData("accessToken");
      console.log(token);
      if (!token) {
        navigation.reset({
          index: 0,
          routes: [{ name: "Login" }],
        });
      }
    };
    loadUser();
  }, []);

  return (
    <View
      style={{
        flex: 1,
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <Text size="lg" weight="bold">
        Welcome to{" "}
        <Text size="lg" weight="extrabold" style={{ color: color.primary }}>
          Expense Tracker
        </Text>
      </Text>
      <Text size="xs">Preparing your content. Please wait{waitingDots}</Text>
    </View>
  );
};

export default Welcome;
