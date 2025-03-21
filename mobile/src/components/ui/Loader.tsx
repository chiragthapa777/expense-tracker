import React, { useEffect, useRef } from "react";
import { Animated, ColorValue, Easing } from "react-native";
import { AntDesign } from "@expo/vector-icons";
type Props = {
  size: number;
  color?: ColorValue;
};
const SpinningLoader = ({ size = 16, color = "black" }: Props) => {
  const spinValue = useRef(new Animated.Value(0)).current;

  useEffect(() => {
    const spinAnimation = Animated.loop(
      Animated.timing(spinValue, {
        toValue: 1,
        duration: 1000, // 1 second per full rotation
        easing: Easing.linear,
        useNativeDriver: true, // Improves performance
      })
    );
    spinAnimation.start();
    return () => spinAnimation.stop(); // Cleanup on unmount
  }, []);

  const spin = spinValue.interpolate({
    inputRange: [0, 1],
    outputRange: ["0deg", "360deg"],
  });

  return (
    <Animated.View style={{ transform: [{ rotate: spin }] }}>
      <AntDesign name="loading1" size={size} color={color} />
    </Animated.View>
  );
};

export default SpinningLoader;
