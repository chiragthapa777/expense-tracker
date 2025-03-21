import { Button as NavigationButton, Text } from "@react-navigation/elements";
import { SafeAreaView, StyleSheet, TextInput, View } from "react-native";
import { Button } from "@/components/ui/Button";
import { userAuthStore } from "@/store/auth";
import { removeData } from "@/utils/asyncStore";
import Toast from "react-native-toast-message";

export function Home() {
  const { setUser } = userAuthStore();
  const logout = async () => {
    setUser(null);
    await removeData("accessToken");
    Toast.show({
      type: "customInfo",
      text1: "Logged Out",
    });
  };
  return (
    <SafeAreaView style={styles.container}>
      <Text>Home Screen</Text>
      <Text>Open up 'src/App.tsx' to start working on your app!</Text>
      <NavigationButton screen="Profile" params={{ user: "jane" }}>
        Go to Profile
      </NavigationButton>
      <NavigationButton screen="Settings">Go to Settings</NavigationButton>
      <Button onPress={logout}>Logout</Button>
      <TextInput></TextInput>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    // justifyContent: 'center',
    // alignItems: 'center',
    gap: 10,
  },
});
