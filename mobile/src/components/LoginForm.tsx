import { loginApi } from "@/api/authApi";
import { Button } from "@/components/ui/Button";
import { CustomTextInput } from "@/components/ui/CustomTextInput";
import Text from "@/components/ui/Text";
import { userAuthStore } from "@/store/auth";
import { useColor } from "@/theme";
import { getData, removeData, storeData } from "@/utils/asyncStore";
import { handleError } from "@/utils/errorUtils";
import { zodResolver } from "@hookform/resolvers/zod";
import { useNavigation } from "@react-navigation/native";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import Checkbox from "expo-checkbox";
import React, { useEffect, useRef } from "react";
import { Controller, useForm } from "react-hook-form";
import { View } from "react-native";
import { TextInput } from "react-native-gesture-handler";
import Toast from "react-native-toast-message";
import { z } from "zod";

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
  const { setUser } = userAuthStore();
  const emailRef = React.createRef<TextInput>();
  const passwordRef = React.createRef<TextInput>();
  const navigation = useNavigation();

  const {
    control,
    formState: { errors, isSubmitting },
    handleSubmit,
    reset,
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

  useEffect(() => {
    const loadEmail = async () => {
      const email = await getData("rememberMeEmail");
      if (email) {
        setValue("email", email, {
          shouldValidate: true,
          shouldTouch: true,
        });
        setValue("rememberMe", true, {
          shouldValidate: true,
          shouldTouch: true,
        });
      }
    };
    emailRef.current?.focus();
    loadEmail();
  }, []);

  const mutation = useMutation({
    mutationFn: loginApi,
    onSuccess: async (data) => {
      Toast.show({
        type: "customSuccess",
        text1: "Login Successful",
        position: "bottom",
      });
      reset();
      storeData("accessToken", data.data.token);
      queryClient.invalidateQueries({ queryKey: ["currentUser"] });
      setUser(data.data.user);
      navigation.reset({
        index: 0,
        routes: [{ name: "HomeTabs" }],
      });
    },
    onError: async (err) => {
      const { error, code, statusCode } = handleError(err);
      Toast.show({
        type: "customError",
        text1: error,
        position: "bottom",
      });
    },
  });

  const onSubmit = async (data: z.infer<typeof loginFormSchema>) => {
    if (data.rememberMe) {
      await storeData("rememberMeEmail", data.email);
    } else {
      await removeData("rememberMeEmail");
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
      <Button onPress={handleSubmit(onSubmit)} loading={isSubmitting}>
        Login
      </Button>
    </View>
  );
}
