import { Assets as NavigationAssets } from "@react-navigation/elements";
import { Asset } from "expo-asset";
import * as SplashScreen from "expo-splash-screen";
import * as React from "react";
import { Navigation } from "./navigation";
import { Appearance, StatusBar, useColorScheme, View } from "react-native";
import {
  ColorProvider,
  darkColors,
  lightColors,
  MyDarkTheme,
  MyLightTheme,
} from "./theme";

Asset.loadAsync([
  ...NavigationAssets,
  require("./assets/newspaper.png"),
  require("./assets/bell.png"),
]);

SplashScreen.preventAutoHideAsync();
// SplashScreen.setOptions({
//   duration: 3000,
//   fade: true,
// });

export function App() {
  const colorScheme = useColorScheme();
  const isDarkMode = React.useMemo(() => colorScheme === "dark", [colorScheme]);

  return (
    <View style={{
      backgroundColor : isDarkMode?darkColors.background:lightColors.background,
      flex: 1,
    }}>
      <ColorProvider>
        <StatusBar
          backgroundColor={
            isDarkMode ? darkColors.background : lightColors.background
          }
          barStyle={isDarkMode ? "light-content" : "dark-content"}
        />
        <Navigation
          linking={{
            enabled: "auto",
            prefixes: [
              // Change the scheme to match your app's scheme defined in app.json
              "helloworld://",
            ],
          }}
          theme={isDarkMode ? MyDarkTheme : MyLightTheme}
          onReady={async () => {
            // await new Promise((r)=> {setTimeout(r,10000)})
            SplashScreen.hideAsync();
          }}
        />
      </ColorProvider>
    </View>
  );
}
