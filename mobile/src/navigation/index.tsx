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
import ProfileEdit from "./screens/ProfileEdit";



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
      options:{
        headerShown:false,
      },
      linking: {
        path: "profile",
      },
    },
    ProfileEdit: {
      screen: withProtection(ProfileEdit),
      options:{
        headerShown:false,
      },
      linking: {
        path: "profile/edit",
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
