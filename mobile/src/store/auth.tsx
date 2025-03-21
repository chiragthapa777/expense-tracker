import { create } from "zustand";
import { User } from "../types/user";

type AuthState = {
  user: User | null;
  isAuthorized: boolean;
};

interface AuthStateAction {
  setUser: (user: User | null) => void;
}

export const userAuthStore = create<AuthState & AuthStateAction>()((set) => ({
  user: null,
  isAuthorized: false,
  setUser: (user) =>
    set((state) => ({ user, isAuthorized: user === null ? false : true })),
}));
