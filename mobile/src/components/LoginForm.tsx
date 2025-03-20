import React from "react";
import { zodResolver } from "@hookform/resolvers/zod";
import { Controller, useForm } from "react-hook-form";
import { z } from "zod";
import { Button } from "@/components/ui/Button";
import Text from "@/components/ui/Text";
import KeyboardView from "@/components/ui/KeyboardView";
import { CustomTextInput } from "@/components/ui/CustomTextInput";
import { ScrollView, View } from "react-native";
import { useColor } from "@/theme";
import Checkbox from "expo-checkbox";

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
  const {
    control,
    formState: { errors },
    handleSubmit,
    getValues,
    setValue,
  } = useForm<z.infer<typeof loginFormSchema>>({
    resolver: zodResolver(loginFormSchema),
    defaultValues: {
      email: "",
      password: "",
      rememberMe: false,
    },
    mode: "all",
  });
  const onSubmit = (data: any) => console.log(data);

  return (
    <KeyboardView
      style={{
        padding: 20,
        gap: 40,
      }}
    >
      <ScrollView
        contentContainerStyle={{
          rowGap: 20,
        }}
      >
        <View style={{ gap: 5 }}>
          <Text size="sm" weight="semibold">
            Email
          </Text>
          <Controller
            control={control}
            rules={{
              required: true,
            }}
            render={({ field: { onChange, onBlur, value } }) => (
              <CustomTextInput
                placeholder="Email"
                onBlur={onBlur}
                onChangeText={onChange}
                value={value}
                hasError={!!errors.email}
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
        <View style={{ gap: 5 }}>
          <Text size="sm" weight="semibold">
            Password
          </Text>
          <Controller
            control={control}
            rules={{
              required: true,
            }}
            render={({ field: { onChange, onBlur, value } }) => (
              <CustomTextInput
                placeholder="Password"
                onBlur={onBlur}
                onChangeText={onChange}
                value={value}
                hasError={!!errors.password}
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
        <View
          style={{
            flexDirection: "row",
            gap: 10,
          }}
        >
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
        <Button onPress={handleSubmit(onSubmit)}>Login</Button>
      </ScrollView>
    </KeyboardView>
  );
}
