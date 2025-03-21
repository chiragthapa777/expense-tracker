import { userAuthStore } from "@/store/auth";

export function useIsAuthorized(){
    const {isAuthorized} = userAuthStore()
    return isAuthorized
}