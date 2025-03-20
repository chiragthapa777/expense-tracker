import { DarkTheme, DefaultTheme } from "@react-navigation/native";
import { createContext, useContext, useMemo } from "react";
import { useColorScheme } from "react-native";

export const lightColors = {
  primary: "#13C2AD",
  primaryOverlayDim: "rgba(0, 0, 0, 0.1)", // Low opacity for Android ripple
  primaryDim: "rgba(19, 194, 173, 0.1)", // Low opacity for Android ripple
  secondary: "#FFA940", // Slightly softer orange for contrast
  background: "#F7F9FA",
  card: "#FFFFFF",
  text: "#1A1A1A",
  textSecondary: "#4D4D4D", // Slightly darker for readability
  border: "#CCCCCC", // Lighter to match the background better
  accent: "#0E9987",
  success: "#52C41A", // Green for success
  error: "#FF4D4F", // Red for errors
  textSuccess: "#389E0D", // Darker green for readable text
  textError: "#D9363E", // Darker red for readable text
};

export const darkColors = {
  primary: "#13C2AD",
  primaryOverlayDim: "rgba(0, 0, 0, 0.1)",
  primaryDim: "rgba(19, 194, 173, 0.1)", // Low opacity for Android ripple
  secondary: "#FFA940",
  background: "#121212", // Darker for deep contrast
  card: "#1E1E1E",
  text: "#FFFFFF",
  textSecondary: "#B0B0B0",
  border: "#4D4D4D", // Softer for dark mode
  accent: "#10B29B", // Slightly more vibrant for contrast
  success: "#3DBA18", // Adjusted for dark mode visibility
  error: "#FF4D4F", // Keeping it the same for consistency
  textSuccess: "#52C41A", // Bright green for readable success text
  textError: "#FF7875", // Softer red for readability on dark backgrounds
};

export const MyLightTheme = {
  ...DefaultTheme,
  colors: {
    ...DefaultTheme.colors,
    primary: lightColors.primary,
    background: lightColors.background,
    card: lightColors.card,
    text: lightColors.text,
    border: lightColors.border,
  },
};

export const MyDarkTheme = {
  ...DarkTheme,
  colors: {
    ...DarkTheme.colors,
    primary: darkColors.primary,
    background: darkColors.background,
    card: darkColors.card,
    text: darkColors.text,
    border: darkColors.border,
  },
};

const ColorContext = createContext(lightColors);

// Custom hook to use the context
export const useColor = () => {
  const context = useContext(ColorContext);
  if (!context) {
    throw new Error("useColor must be used within a ColorProvider");
  }
  return context;
};

// Provider component
export const ColorProvider = ({ children }: { children: any }) => {
  const scheme = useColorScheme(); // 'light' or 'dark' based on system preference

  // Memoize the colors to avoid unnecessary re-renders
  const colors = useMemo(
    () => (scheme === "dark" ? darkColors : lightColors),
    [scheme]
  );

  return (
    <ColorContext.Provider value={colors}>{children}</ColorContext.Provider>
  );
};
