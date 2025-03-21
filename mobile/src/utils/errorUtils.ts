import { ApiErrorResponse } from "@/types/response";
import { AxiosError } from "axios";

export const handleError = (
  err: Error
): { error: string; code?: string; statusCode?: number } => {
  if (err instanceof AxiosError) {
    return {
      error: err?.response?.data?.error || "server error",
      code: err?.response?.data?.code,
      statusCode: err?.status,
    };
  }
  return {
    error: "server error",
  };
};
