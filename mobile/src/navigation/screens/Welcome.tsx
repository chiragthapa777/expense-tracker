import { getCurrentUserApi } from "@/api/authApi";
import { useNavigation } from "@react-navigation/native";
import { useQuery } from "@tanstack/react-query";
import { useEffect } from "react";
import { View } from "react-native";
import Text from "../../components/ui/Text";
import { useWaitingDots } from "../../hooks/useWaitingDots";
import { userAuthStore } from "../../store/auth";
import { useColor } from "../../theme";
import { getData } from "../../utils/asyncStore";

const Welcome = () => {
  const color = useColor();
  const { waitingDots } = useWaitingDots();
  const navigation = useNavigation<any>();
  const setUser = userAuthStore((s) => s.setUser);

  const { data: currentUser, isLoading, error } = useQuery({
    queryKey: ["currentUser"],
    queryFn: async () => {
      const token = await getData("accessToken");
      if (!token) {
        throw new Error("No token found");
      }
      return getCurrentUserApi(); 
    },
    retry: 0,
    staleTime: Infinity, // Keep data fresh indefinitely for this check
  });

  useEffect(() => {
    if (isLoading) return; 

    if (currentUser) {
      setUser(currentUser.data);
      navigation.reset({
        index: 0,
        routes: [{ name: "HomeTabs" }],
      });
    } else if (error) {
      navigation.reset({
        index: 0,
        routes: [{ name: "Login" }],
      });
    }
  }, [currentUser, isLoading, error, navigation]);

  return (
    <View
      style={{
        flex: 1,
        justifyContent: "center",
        alignItems: "center",
        backgroundColor: color.background, 
      }}
    >
      <Text size="lg" weight="bold">
        Welcome to{" "}
        <Text size="lg" weight="extrabold" style={{ color: color.primary }}>
          Expense Tracker
        </Text>
      </Text>
      <Text size="xs" color={color.textSecondary}>
        Preparing your content. Please wait{waitingDots}
      </Text>
    </View>
  );
};

export default Welcome;