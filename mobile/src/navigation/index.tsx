import { HeaderButton, Text } from "@react-navigation/elements";
import {
  createStaticNavigation,
  StaticParamList,
} from "@react-navigation/native";
import { createNativeStackNavigator } from "@react-navigation/native-stack";
import Login from "./screens/auth/Login";
import { NotFound } from "./screens/NotFound";
import { Profile } from "./screens/Profile";
import { Settings } from "./screens/Settings";
import Welcome from "./screens/Welcome";
import { HomeTabs } from "./screens/home-tabs";
import withProtection from "@/hoc/Protected";



const RootStack = createNativeStackNavigator({
  initialRouteName: "Welcome",
  screens: {
    Welcome: {
      screen: Welcome,
      options: {
        headerShown: false,
      },
      config: {},
    },
    Login: {
      screen: Login,
      options: {
        headerShown: false,
      },
      config: {},
    },
    HomeTabs: {
      screen: HomeTabs,
      options: {
        title: "Home",
        headerShown: false,
      },
    },
    Profile: {
      screen: withProtection(Profile),
      linking: {
        path: ":user(@[a-zA-Z0-9-_]+)",
        parse: {
          user: (value) => value.replace(/^@/, ""),
        },
        stringify: {
          user: (value) => `@${value}`,
        },
      },
    },
    Settings: {
      screen: withProtection(Settings),
      options: ({ navigation }) => ({
        presentation: "modal",
        headerRight: () => (
          <HeaderButton onPress={navigation.goBack}>
            <Text>Close</Text>
          </HeaderButton>
        ),
      }),
    },
    NotFound: {
      screen: NotFound,
      options: {
        title: "404",
      },
      linking: {
        path: "*",
      },
    },
  },
});

export const Navigation = createStaticNavigation(RootStack);

type RootStackParamList = StaticParamList<typeof RootStack>;

declare global {
  namespace ReactNavigation {
    interface RootParamList extends RootStackParamList {}
  }
}
