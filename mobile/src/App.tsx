import { Assets as NavigationAssets } from "@react-navigation/elements";
import { Asset } from "expo-asset";
import * as SplashScreen from "expo-splash-screen";
import * as React from "react";
import {
  AppState,
  AppStateStatus,
  Platform,
  StatusBar,
  useColorScheme,
  View,
} from "react-native";
import { SafeAreaProvider } from "react-native-safe-area-context";
import { Navigation } from "./navigation";
import {
  focusManager,
  onlineManager,
  QueryClient,
  QueryClientProvider,
} from "@tanstack/react-query";
import * as Network from "expo-network";
import Toast, { BaseToast } from "react-native-toast-message";

import {
  ColorProvider,
  darkColors,
  lightColors,
  MyDarkTheme,
  MyLightTheme,
} from "./theme";
import {
  createNavigationContainerRef,
  NavigationContainer,
} from "@react-navigation/native";
import CustomToast from "./components/Toast";

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

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: 2, // Retry failed requests twice
      staleTime: 30 * 1000, // Data stays fresh for 5 minutes
    },
    mutations: {
      retry: 0, // Don’t retry mutations by default
    },
  },
});

export const navigationRef = createNavigationContainerRef<any>();

// Online status management
onlineManager.setEventListener((setOnline) => {
  const eventSubscription = Network.addNetworkStateListener((state) => {
    setOnline(!!state.isConnected);
  });
  return eventSubscription.remove;
});

// App focus management
function onAppStateChange(status: AppStateStatus) {
  if (Platform.OS !== "web") {
    focusManager.setFocused(status === "active");
  }
}

export function App() {
  React.useEffect(() => {
    const subscription = AppState.addEventListener("change", onAppStateChange);
    return () => subscription.remove();
  }, []);
  const colorScheme = useColorScheme();
  const isDarkMode = React.useMemo(() => colorScheme === "dark", [colorScheme]);

  return (
    <QueryClientProvider client={queryClient}>
      <SafeAreaProvider>
        <View
          style={{
            backgroundColor: isDarkMode
              ? darkColors.background
              : lightColors.background,
            flex: 1,
          }}
        >
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
              ref={navigationRef}
              theme={isDarkMode ? MyDarkTheme : MyLightTheme}
              onReady={async () => {
                // await new Promise((r)=> {setTimeout(r,10000)})
                SplashScreen.hideAsync();
              }}
            />
            <CustomToast />
          </ColorProvider>
        </View>
      </SafeAreaProvider>
    </QueryClientProvider>
  );
}
