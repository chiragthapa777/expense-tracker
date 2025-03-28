import RegisterForm from "@/components/RegisterForm";
import Avatar from "@/components/ui/Avatar";
import BackButton from "@/components/ui/BackButton";
import { CustomBottomSheet } from "@/components/ui/BottomSheet";
import Text from "@/components/ui/Text";
import { userAuthStore } from "@/store/auth";
import { useColor } from "@/theme";
import { removeData } from "@/utils/asyncStore";
import FontAwesome5 from "@expo/vector-icons/FontAwesome5";
import MaterialIcons from "@expo/vector-icons/MaterialIcons";
import BottomSheet from "@gorhom/bottom-sheet";
import { StaticScreenProps, useNavigation } from "@react-navigation/native";
import { useCallback, useMemo, useRef } from "react";
import {
  GestureResponderEvent,
  Pressable,
  SafeAreaView,
  ScrollView,
  View
} from "react-native";
import { GestureHandlerRootView } from "react-native-gesture-handler";
import Toast from "react-native-toast-message";

type Props = StaticScreenProps<{
  user: string;
}>;

export function Profile() {
  const color = useColor();
  const user = userAuthStore((state) => state.user);
  const { setUser } = userAuthStore();
  const navigation = useNavigation();

  const fullName = useMemo(
    () => [user?.firstName, user?.lastName].filter((n) => n).join(" "),
    [user]
  );

  // hooks
  const sheetRef = useRef<BottomSheet>(null);

  // variables

  // callbacks
  const handleSheetChange = useCallback((index: any) => {
    console.log("handleSheetChange", index);
    if (index === -1) {
      console.log("Closed");
    }
  }, []);
  const handleSnapPress = useCallback((index: any) => {
    sheetRef.current?.snapToIndex(index);
  }, []);
  const logout = async () => {
    setUser(null);
    await removeData("accessToken");
    Toast.show({
      type: "customInfo",
      text1: "Logged Out",
      position: "bottom",
    });
  };

  return (
    <GestureHandlerRootView
      style={{
        flex: 1,
      }}
    >
      <CustomBottomSheet ref={sheetRef} onChange={handleSheetChange}>
       <RegisterForm/>
      </CustomBottomSheet>
      <SafeAreaView style={{ flex: 1 }}>
        <ScrollView
          style={{
            flexGrow: 1,
            position: "relative",
          }}
        >
          <BackButton />
          <View
            style={{
              justifyContent: "center",
              alignItems: "center",
              padding: 20,
              paddingTop: 50,
              gap: 5,
            }}
          >
            <Avatar
              size={130}
              imageUrl={"https://picsum.photos/seed/696/3000/2000"}
            />
            <Text size="lg" weight="semibold">
              {fullName}
            </Text>
            <Text size="sm">{user?.email}</Text>
          </View>
          {/* Options */}
          <View
            style={{
              backgroundColor: color.card,
              borderRadius: 18,
              borderWidth: 1,
              borderColor: color.border,
              marginHorizontal: 20,
              overflow: "hidden",
              marginTop: 20,
            }}
          >
            <Option
              icon={
                <FontAwesome5 name="user-edit" size={24} color={color.text} />
              }
              onPress={() => {
                navigation.navigate("ProfileEdit");
              }}
              title="Edit Profile"
              description="Update your profile details like name, profile picture etc"
            />
            <Option
              icon={
                <MaterialIcons name="password" size={24} color={color.text} />
              }
              title="Change Password"
              description="Change your password from old to new"
              onPress={() => {
                handleSnapPress(1);
              }}
            />
            <Option
              icon={
                <MaterialIcons name="logout" size={24} color={color.text} />
              }
              title="Log out"
              description={`Log out from ${fullName}`}
              onPress={logout}
            />
          </View>
        </ScrollView>
      </SafeAreaView>
    </GestureHandlerRootView>
  );
}

const Option = ({
  icon,
  title,
  description,
  onPress,
}: {
  icon: React.ReactNode;
  title: string;
  description: string;
  onPress?: (e: GestureResponderEvent) => void;
}) => {
  const color = useColor();
  return (
    <Pressable
      style={({ pressed }) => [
        {
          flexDirection: "row",
          justifyContent: "flex-start",
          gap: 15,
          paddingHorizontal: 20,
          paddingVertical: 15,
          backgroundColor: pressed ? color.background : color.card,
        },
      ]}
      onPress={onPress}
    >
      <View
        style={{
          marginTop: 2,
        }}
      >
        {icon}
      </View>
      <View
        style={{
          flex: 1,
        }}
      >
        <Text weight="semibold" size="sm">
          {title}
        </Text>
        <Text color={color.textSecondary} size="xs">
          {description}
        </Text>
      </View>
    </Pressable>
  );
};
