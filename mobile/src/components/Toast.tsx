import { useColor } from "@/theme";
import React from "react";
import { StyleProp, View, ViewStyle } from "react-native";
import Toast, { BaseToastProps } from "react-native-toast-message";
import Text from "./ui/Text";
import AntDesign from "@expo/vector-icons/AntDesign";
import MaterialIcons from '@expo/vector-icons/MaterialIcons';
import { capitalize } from "@/utils/stringUtils";


export default function CustomToast() {
  const color = useColor();

  const toastStyles: StyleProp<ViewStyle> = {
    maxWidth: "90%",
    backgroundColor: color.card,
    borderColor: color.border,
    borderWidth: 0.1,
    borderRadius: 20,
    padding: 12,
    marginBottom: 50,
    // iOS Shadow
    shadowColor: "#000",
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.2,
    shadowRadius: 3,
    // Android Elevation
    elevation: 2,
    flexDirection: "row",
    gap: 5,
    justifyContent:"flex-start",
    alignItems:"center"
  };

  const renderToast = ({ text1, text2 }: BaseToastProps, textColor: string, icon:React.ReactNode) => (
    <View style={toastStyles}>
      <View>
        {icon}
       
      </View>
      <View
        style={{
        //   flex: 1,
        padding:2
        }}
      >
        {text1 && (
          <Text color={textColor} size="md" weight="semibold" style={{
            
          }}>
            {capitalize(text1)}
          </Text>
        )}
        {text2 && (
          <Text color={color.textSecondary} size="xs" weight="light">
            {capitalize(text2)}
          </Text>
        )}
      </View>
    </View>
  );

  return (
    <Toast
      config={{
        customSuccess: (props) => renderToast(props, color.primary, <AntDesign name="checkcircle" size={24} color={color.primary} />),
        customInfo: (props) => renderToast(props, color.secondary,<AntDesign name="infocirlce" size={24} color={color.secondary} />),
        customError: (props) => renderToast(props, color.error,<MaterialIcons name="error" size={24} color={color.error} />),
      }}
    />
  );
}
