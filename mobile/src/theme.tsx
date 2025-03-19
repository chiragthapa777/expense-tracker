import { DarkTheme, DefaultTheme } from "@react-navigation/native";
import { createContext, useContext, useMemo } from "react";
import { useColorScheme } from "react-native";

export const lightColors = {
  primary: '#13C2AD',
  secondary: '#F5A623',
  background: '#F7F9FA',
  surface: '#FFFFFF',
  text: '#1A1A1A',
  textSecondary: '#666666',
  accent: '#0E9987',
};

export const darkColors = {
  primary: '#13C2AD',
  secondary: '#F5A623',
  background: '#1A1A1A',
  surface: '#252525',
  text: '#FFFFFF',
  textSecondary: '#B3B3B3',
  accent: '#3DD9C2',
};

export const MyLightTheme = {
  ...DefaultTheme,
  colors: {
    ...DefaultTheme.colors,
    primary: lightColors.primary,
    background: lightColors.background,
    card: lightColors.surface,
    text: lightColors.text,
    border: lightColors.textSecondary,
  },
};

export const MyDarkTheme = {
  ...DarkTheme,
  colors: {
    ...DarkTheme.colors,
    primary: darkColors.primary,
    background: darkColors.background,
    card: darkColors.surface,
    text: darkColors.text,
    border: darkColors.textSecondary,
  },
};

const ColorContext = createContext(lightColors);

// Custom hook to use the context
export const useColor = () => {
  const context = useContext(ColorContext);
  if (!context) {
    throw new Error('useColor must be used within a ColorProvider');
  }
  return context;
};

// Provider component
export const ColorProvider = ({ children }:{children:any}) => {
  const scheme = useColorScheme(); // 'light' or 'dark' based on system preference

  // Memoize the colors to avoid unnecessary re-renders
  const colors = useMemo(() => (scheme === 'dark' ? darkColors : lightColors), [scheme]);

  return (
    <ColorContext.Provider value={colors}>
      {children}
    </ColorContext.Provider>
  );
};