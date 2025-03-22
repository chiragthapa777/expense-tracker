import { ApiResponse } from "@/types/response";
import api from "./axios";
import { LoginResponse } from "@/types/auth";
import { User } from "@/types/user";

type LoginBody = { email: string; password: string }
export const loginApi = async (body:LoginBody ) => {
  return api.post<LoginBody, ApiResponse<LoginResponse>>("/api/v1/auth/login", body);
};

type RegisterBody = { email: string; password: string, firstName:string, lastName?:string }
export const registerApi = async (body:RegisterBody ) => {
  return api.post<RegisterBody, ApiResponse<LoginResponse>>("/api/v1/auth/register", body);
};

export const getCurrentUserApi = async () => {
  return api.get<LoginBody, ApiResponse<User>>("/api/v1/auth/me");
};


