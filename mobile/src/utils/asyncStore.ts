import AsyncStorage from "@react-native-async-storage/async-storage";

export const storeData = async (key: string, value: string) => {
  try {
    await AsyncStorage.setItem(key, value);
  } catch (e) {
    console.log(e);
  }
};

export const getData = async (key: string): Promise<string | null> => {
  try {
    const value = await AsyncStorage.getItem("my-key");
    return value;
  } catch (e) {
    console.log(e);
    return null;
  }
};
