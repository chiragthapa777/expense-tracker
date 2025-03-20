import React, { useRef } from "react";
import { zodResolver } from "@hookform/resolvers/zod";
import { Controller, useForm } from "react-hook-form";
import { z } from "zod";
import { Button } from "@/components/ui/Button";
import Text from "@/components/ui/Text";
import { CustomTextInput } from "@/components/ui/CustomTextInput";
import { View } from "react-native";
import { useColor } from "@/theme";
import Checkbox from "expo-checkbox";
import { removeData, storeData } from "@/utils/asyncStore";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { LoginApi } from "@/api/authApi";
import Toast from "react-native-toast-message";
import AntDesign from '@expo/vector-icons/AntDesign';


const PASSWORD_REGEX =
  /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*])[A-Za-z\d!@#$%^&*]{8,}$/;

const loginFormSchema = z.object({
  email: z.string().email(),
  password: z
    .string()
    .regex(
      PASSWORD_REGEX,
      "Password must be 8+ characters with uppercase, lowercase, number, and special character."
    ),
  rememberMe: z.boolean(),
});

type Props = {};

export default function LoginForm({}: Props) {
  const color = useColor();
  const queryClient = useQueryClient();

  const {
    control,
    formState: { errors },
    handleSubmit,
  } = useForm<z.infer<typeof loginFormSchema>>({
    resolver: zodResolver(loginFormSchema),
    defaultValues: {
      email: "",
      password: "",
      rememberMe: false,
    },
    mode: "all",
  });

  const emailRef = useRef<any>(null);
  const passwordRef = useRef<any>(null);

  const mutation = useMutation({
    mutationFn: LoginApi,
    onSuccess: async (data) => {
      queryClient.invalidateQueries({ queryKey: ["todos"] });
    },
    onError: async (err) => {},
  });

  const onSubmit = async (data: z.infer<typeof loginFormSchema>) => {
    if (data.rememberMe) {
      await storeData("rememberMe:email", data.email);
    } else {
      removeData("rememberMe:email");
    }
    mutation.mutate(data);
  };

  return (
    <View style={{ gap: 20, padding: 20 }}>
      {/* Email Field */}
      <View style={{ gap: 5 }}>
        <Text size="sm" weight="semibold">
          Email
        </Text>
        <Controller
          control={control}
          rules={{ required: true }}
          render={({ field: { onChange, onBlur, value } }) => (
            <CustomTextInput
              ref={emailRef}
              placeholder="Email"
              onBlur={onBlur}
              onChangeText={onChange}
              value={value}
              hasError={!!errors.email}
              returnKeyType="next" // Shows "Next" on keyboard
              onSubmitEditing={() => passwordRef.current?.focus()} // Moves to password
            />
          )}
          name="email"
        />
        {errors.email && (
          <Text color={color.error} size="xs">
            {errors.email.message}
          </Text>
        )}
      </View>

      {/* Password Field */}
      <View style={{ gap: 5 }}>
        <Text size="sm" weight="semibold">
          Password
        </Text>
        <Controller
          control={control}
          rules={{ required: true }}
          render={({ field: { onChange, onBlur, value } }) => (
            <CustomTextInput
              ref={passwordRef}
              placeholder="Password"
              onBlur={onBlur}
              onChangeText={onChange}
              value={value}
              hasError={!!errors.password}
              secureTextEntry
              returnKeyType="done" // Shows "Done" (or "Go") to submit
              onSubmitEditing={handleSubmit(onSubmit)} // Submits form
            />
          )}
          name="password"
        />
        {errors.password && (
          <Text color={color.error} size="xs">
            {errors.password.message}
          </Text>
        )}
      </View>

      {/* Remember Me */}
      <View style={{ flexDirection: "row", gap: 10 }}>
        <Controller
          name="rememberMe"
          control={control}
          render={({ field: { onChange, value } }) => (
            <Checkbox
              value={value}
              onValueChange={(bool) => onChange(bool)}
              color={value ? color.primary : color.border}
            />
          )}
        />
        <Text size="sm" weight="semibold">
          Remember Me
        </Text>
      </View>

      {/* Submit Button */}
      <Button onPress={handleSubmit(onSubmit)}>Login</Button>
      <Button
        onPress={() => {
          console.log("test")
          Toast.show({
            type:"customSuccess",
            text1:"This is message",
            position:"bottom"
          })
        }}
      >
        Toast
      </Button>
      <Button
        onPress={() => {
          console.log("test")
          Toast.show({
            type:"customError",
            text1:"This is message",
            text2:"this can be a long message just bear",
            position:"bottom"
          })
        }}
      >
        Toast
      </Button>
      <Button
        onPress={() => {
          console.log("test")
          Toast.show({
            type:"customInfo",
            text1:"This is message",
            text2:"It looks like TypeScript is complaining because the renderToast function is missing explicit types for its parameters. Let's fix that by adding proper TypeScript types. Here's the updated version:",
            position:"bottom"
          })
        }}
      >
        Toast
      </Button>
    </View>
  );
}
