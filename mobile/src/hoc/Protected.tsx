import { useIsAuthorized } from "@/hooks/useIsAuthorized";
import { removeData, storeData } from "@/utils/asyncStore";
import { useNavigation, useRoute } from "@react-navigation/native";
import React, { ComponentType, useEffect } from "react";
import Toast from "react-native-toast-message";

type Props = { children: React.ReactNode };

export function Protected({ children }: Props) {
  const isAuthorized = useIsAuthorized();
  const navigation = useNavigation();
  const route = useRoute();

  useEffect(() => {
    if (!isAuthorized) {
      removeData("accessToken");
      storeData("redirectRoute", route.name);
      Toast.show({
        type: "customError",
        text1: "UnAuthorized",
      });
      navigation.reset({
        index: 0,
        routes: [{ name: "Login" }],
      });
    }
  }, [isAuthorized]);

  return isAuthorized ? <>{children}</> : null;
}

const withProtection = <P extends object>(
  Screen: React.ComponentType<P>
): React.ComponentType<P> => {
  const ProtectedScreen: React.ComponentType<P> = (props: P) => (
    <Protected>
      <Screen {...props} />
    </Protected>
  );

  return ProtectedScreen;
};
export default withProtection;
