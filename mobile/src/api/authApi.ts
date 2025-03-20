import { ApiResponse } from "@/types/response";
import api from "./axios";
import { LoginResponse } from "@/types/auth";

type LoginBody = { email: string; password: string }
export const LoginApi = async (body:LoginBody ) => {
  return api.post<LoginBody, ApiResponse<LoginResponse>>("/api/v1/auth/login", body);
};
