import Feather from "@expo/vector-icons/Feather";
import Ionicons from "@expo/vector-icons/Ionicons";
import { createBottomTabNavigator } from "@react-navigation/bottom-tabs";
import { Home } from "./Home";

import { Updates } from "./Updates";
import withProtection from "@/hoc/Protected";

export const HomeTabs = createBottomTabNavigator({
  screens: {
    Home: {
      screen: withProtection(Home),
      options: {
        headerShown: false,
        title: "Feed",
        tabBarIcon: ({ color, size }) => (
          <Feather name="home" size={size} color={color} />
        ),
      },
    },
    Updates: {
      screen: withProtection(Updates),
      options: {
        tabBarIcon: ({ color, size }) => (
          <Ionicons name="options" size={size} color={color} />
        ),
      },
    },
  },
});
