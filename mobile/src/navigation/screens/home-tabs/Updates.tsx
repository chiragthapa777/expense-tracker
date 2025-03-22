import { CustomBottomSheet } from "@/components/ui/BottomSheet";
import { Button as MyBtn } from "@/components/ui/Button";
import Text from "@/components/ui/Text";
import BottomSheet from "@gorhom/bottom-sheet";
import { useCallback, useMemo, useRef } from "react";
import { StyleSheet } from "react-native";
import { GestureHandlerRootView } from "react-native-gesture-handler";

const Button = ({ title, onPress }: { title: string; onPress: () => void }) => {
  return <MyBtn onPress={onPress}>{title}</MyBtn>;
};

export function Updates() {
  // hooks
  const sheetRef = useRef<BottomSheet>(null);

  // variables
  const snapPoints = useMemo(() => ["25%", "50%", "90%"], []);

  // callbacks
  const handleSheetChange = useCallback((index: any) => {
    console.log("handleSheetChange", index);
  }, []);
  const handleSnapPress = useCallback((index: any) => {
    sheetRef.current?.snapToPosition("50%");
  }, []);
  const handleClosePress = useCallback(() => {
    sheetRef.current?.close();
  }, []);

  // render
  return (
    <GestureHandlerRootView style={styles.container}>
      <Button title="Snap To 50%" onPress={() => handleSnapPress(1)} />

      <Button title="Close" onPress={() => handleClosePress()} />
      <CustomBottomSheet ref={sheetRef}>
        <Text>tester</Text>
      </CustomBottomSheet>
    </GestureHandlerRootView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
  contentContainer: {
    flex: 1,
    padding: 36,
    alignItems: "center",
    borderWidth: 1,
  },
});
