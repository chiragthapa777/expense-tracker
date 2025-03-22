import { loginApi, registerApi } from "@/api/authApi";
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
import React, { LegacyRef, useEffect, useRef } from "react";
import { Controller, useForm } from "react-hook-form";
import { TextInput, View } from "react-native";
import Toast from "react-native-toast-message";
import { z } from "zod";

const PASSWORD_REGEX =
  /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*])[A-Za-z\d!@#$%^&*]{8,}$/;

const registerFormSchema = z
  .object({
    firstName: z.string().min(3, { message: "minimum length is 3" }),
    lastName: z.string().optional(),
    email: z.string().email(),
    password: z.string(),
    confirmPassword: z
      .string()
      .regex(
        PASSWORD_REGEX,
        "Password must be 8+ characters with uppercase, lowercase, number, and special character."
      ),
    rememberMe: z.boolean(),
  })
  .refine((data) => data.password === data.confirmPassword, {
    message: "Passwords do not match",
    path: ["confirmPassword"],
  });

type Props = {};

export default function RegisterForm({}: Props) {
  const color = useColor();
  const queryClient = useQueryClient();
  const { setUser } = userAuthStore();
  const firstNameRef = React.createRef<TextInput>();
  const lastNameRef = React.createRef<TextInput>();
  const emailRef = React.createRef<TextInput>();
  const passwordRef = React.createRef<TextInput>();
  const confirmPasswordRef = React.createRef<TextInput>();
  const navigation = useNavigation();

  const {
    control,
    formState: { errors, isSubmitting },
    handleSubmit,
    reset,
    setValue,
  } = useForm<z.infer<typeof registerFormSchema>>({
    resolver: zodResolver(registerFormSchema),
    defaultValues: {
      lastName: undefined,
      firstName: "",
      email: "",
      password: "",
      confirmPassword: "",
      rememberMe: false,
    },
    mode: "all",
  });

  useEffect(() => {
    firstNameRef.current?.focus();
  }, []);

  const mutation = useMutation({
    mutationFn: registerApi,
    onSuccess: async (data) => {
      Toast.show({
        type: "customSuccess",
        text1: "Register Successful",
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

  const onSubmit = async (data: z.infer<typeof registerFormSchema>) => {
    if (data.rememberMe) {
      await storeData("rememberMeEmail", data.email);
    } else {
      await removeData("rememberMeEmail");
    }
    mutation.mutate(data);
  };

  return (
    <View style={{ gap: 20, padding: 20 }}>
      {/* First Name Field */}
      <View style={{ gap: 5 }}>
        <Text size="sm" weight="semibold">
          First Name{" "}
          <Text size="xs" color={color.textSecondary}>
            (required)
          </Text>
        </Text>
        <Controller
          control={control}
          rules={{ required: true }}
          render={({ field: { onChange, onBlur, value } }) => (
            <CustomTextInput
              ref={firstNameRef}
              placeholder="First Name"
              onBlur={onBlur}
              onChangeText={onChange}
              value={value}
              hasError={!!errors.firstName}
              returnKeyType="next" // Shows "Next" on keyboard
              onSubmitEditing={() => lastNameRef.current?.focus()} // Moves to password
            />
          )}
          name="firstName"
        />
        {errors.firstName && (
          <Text color={color.error} size="xs">
            {errors.firstName.message}
          </Text>
        )}
      </View>

      {/* Last Name Field */}
      <View style={{ gap: 5 }}>
        <Text size="sm" weight="semibold">
          Last Name
        </Text>
        <Controller
          control={control}
          rules={{ required: true }}
          render={({ field: { onChange, onBlur, value } }) => (
            <CustomTextInput
              ref={lastNameRef}
              placeholder="Last Name"
              onBlur={onBlur}
              onChangeText={onChange}
              value={value}
              hasError={!!errors.lastName}
              returnKeyType="next" // Shows "Next" on keyboard
              onSubmitEditing={() => emailRef.current?.focus()} // Moves to password
            />
          )}
          name="lastName"
        />
        {errors.lastName && (
          <Text color={color.error} size="xs">
            {errors.lastName.message}
          </Text>
        )}
      </View>

      {/* Email Field */}
      <View style={{ gap: 5 }}>
        <Text size="sm" weight="semibold">
          Email{" "}
          <Text size="xs" color={color.textSecondary}>
            (required)
          </Text>
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
          Password{" "}
          <Text size="xs" color={color.textSecondary}>
            (required)
          </Text>
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
              returnKeyType="next" // Shows "Done" (or "Go") to submit
              onSubmitEditing={() => confirmPasswordRef.current?.focus()} // Submits form
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

      {/* Confirm Password Field */}
      <View style={{ gap: 5 }}>
        <Text size="sm" weight="semibold">
          Confirm Password{" "}
          <Text size="xs" color={color.textSecondary}>
            (required)
          </Text>
        </Text>
        <Controller
          control={control}
          rules={{ required: true }}
          render={({ field: { onChange, onBlur, value } }) => (
            <CustomTextInput
              ref={confirmPasswordRef}
              placeholder="Confirm Password"
              onBlur={onBlur}
              onChangeText={onChange}
              value={value}
              hasError={!!errors.confirmPassword}
              secureTextEntry
              returnKeyType="done"
              onSubmitEditing={handleSubmit(onSubmit)}
            />
          )}
          name="confirmPassword"
        />
        {errors.confirmPassword && (
          <Text color={color.error} size="xs">
            {errors.confirmPassword.message}
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
        Register
      </Button>
    </View>
  );
}
