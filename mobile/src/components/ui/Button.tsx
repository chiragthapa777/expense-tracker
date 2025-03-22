import React from "react";
import {
  Pressable,
  PressableProps,
  Text,
  StyleSheet,
  ViewStyle,
  TextStyle,
  StyleProp,
} from "react-native";
import { useColor } from "../../theme";
import * as Haptics from "expo-haptics";
import AntDesign from "@expo/vector-icons/AntDesign";
import SpinningLoader from "./Loader";

type ButtonVariants =
  | "default"
  | "secondary"
  | "destructive"
  | "link"
  | "ghost"
  | "outline";

type ButtonSize = "default" | "sm" | "lg" | "icon";

type Props = PressableProps & {
  children: React.ReactNode;
  variant?: ButtonVariants;
  size?: ButtonSize;
  fullWidth?: boolean;
  style?: StyleProp<ViewStyle>;
  loading?: boolean;
};

export const Button = React.forwardRef<
  React.ElementRef<typeof Pressable>,
  Props
>(
  (
    {
      children,
      variant = "default",
      size = "default",
      fullWidth,
      style,
      onPress,
      onLongPress,
      disabled,
      loading,
      ...props
    },
    ref
  ) => {
    const colors = useColor();
    disabled = disabled || loading

    // Variant styles
    const variantStyles: Record<ButtonVariants, ViewStyle> = {
      default: {
        backgroundColor: colors.primary,
        borderColor: "transparent",
      },
      secondary: {
        backgroundColor: colors.secondary,
        borderColor: "transparent",
      },
      destructive: {
        backgroundColor: colors.error,
        borderColor: "transparent",
      },
      outline: {
        backgroundColor: "transparent",
        borderColor: colors.primary,
        borderWidth: 2,
      },
      ghost: {
        backgroundColor: "transparent",
        borderColor: "transparent",
      },
      link: {
        backgroundColor: "transparent",
        borderColor: "transparent",
      },
    };

    // Size styles
    const sizeStyles: Record<ButtonSize, ViewStyle> = {
      default: {
        paddingVertical: 10,
        paddingHorizontal: 12,
      },
      sm: {
        paddingVertical: 6,
        paddingHorizontal: 10,
      },
      lg: {
        paddingVertical: 14,
        paddingHorizontal: 20,
      },
      icon: {
        padding: 5,
      },
    };

    // Text color based on variant
    const textColorStyles: Record<ButtonVariants, TextStyle> = {
      default: { color: "#FFFFFF" },
      secondary: { color: "#FFFFFF" },
      destructive: { color: "#FFFFFF" }, // Changed to white for better contrast
      outline: { color: colors.primary },
      ghost: { color: colors.primary },
      link: { color: colors.primary, textDecorationLine: "underline" },
    };

    return (
      <Pressable
        ref={ref}
        style={[
          styles.baseButton,
          variantStyles[variant],
          sizeStyles[size],
          fullWidth && { width: "100%" },
          (disabled) && { opacity: 0.5 }, // Added for disabled state
          style,
        ]}
        android_ripple={{ color: colors.primaryOverlayDim }} // Consistent ripple for all variants
        onPress={(e) => {
          if (!disabled) {
            Haptics.impactAsync(Haptics.ImpactFeedbackStyle.Medium);
            onPress?.(e);
          }
        }}
        onLongPress={(e) => {
          if (!disabled) {
            Haptics.impactAsync(Haptics.ImpactFeedbackStyle.Heavy);
            onLongPress?.(e);
          }
        }}
        disabled={disabled}
        {...props}
      >
        {loading && (
          <SpinningLoader size={16} color={textColorStyles[variant].color} />
        )}
        {typeof children === "string" ? (
          <Text style={[styles.baseText, textColorStyles[variant]]}>
            {children}
          </Text>
        ) : (
          children
        )}
      </Pressable>
    );
  }
);

Button.displayName = "Button";

const styles = StyleSheet.create({
  baseButton: {
    borderRadius: 8,
    alignItems: "center",
    justifyContent: "center",
    flexDirection:"row",
    gap:4
  },
  baseText: {
    fontSize: 16,
    fontWeight: "600",
  },
});
