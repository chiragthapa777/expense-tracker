import axios, { AxiosInstance, InternalAxiosRequestConfig, AxiosResponse } from "axios";
import { navigationRef } from "../App"; // Adjust path to App.tsx
import { getData, removeData } from "@/utils/asyncStore";

const BASE_URL = process.env.BASE_API_URL;

const api: AxiosInstance = axios.create({
  baseURL: BASE_URL,
  timeout: 10000,
  headers: {
    "Content-Type": "application/json",
  },
});

// Request Interceptor
api.interceptors.request.use(
  async (config: InternalAxiosRequestConfig): Promise<InternalAxiosRequestConfig> => {
    const token = await getData("accessToken"); // Replace with real token logic (e.g., AsyncStorage)
    if (token && config.headers) {
      config.headers.set("Authorization", `Bearer ${token}`); // Use set() for type safety
    }
    return config;
  },
  (error) => Promise.reject(error)
);

// Response Interceptor
api.interceptors.response.use(
  (response: AxiosResponse) => response.data,
  async (error) => {
    if (error.response) {
      const { status } = error.response;
      if (status === 401) {
        console.log("Unauthorized - Redirecting to Login...");
        if (navigationRef.isReady()) {
          navigationRef.reset({
            index: 0,
            routes: [{ name: "Login" }],
          });
        }
        await removeData("accessToken");
      } else if (status === 500) {
        console.log("Server error occurred");
      }
    } else if (error.request) {
      console.log("No response received from server");
    } else {
      console.log("Error setting up request:", error.message);
    }
    return Promise.reject(error);
  }
);

export default api;