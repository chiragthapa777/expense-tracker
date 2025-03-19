import { Button, Text } from '@react-navigation/elements';
import { StyleSheet, TextInput, View } from 'react-native';

export function Home() {
  return (
    <View style={styles.container}>
      <Text>Home Screen</Text>
      <Text>Open up 'src/App.tsx' to start working on your app!</Text>
      <Button screen="Profile" params={{ user: 'jane' }}>
        Go to Profile
      </Button>
      <Button screen="Settings">Go to Settings</Button>
      <TextInput></TextInput>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    gap: 10,
  },
});
