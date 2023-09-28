import {defineStore} from "pinia";
import {ref} from "vue";
import type {LoginRequest, LoginResponse} from "@/lib/types/auth";
import { useApiStore } from "./api";

export const useAuthStore = defineStore("auth", () => {
    const token = ref<string | null>(null);

    const login = async (input: LoginRequest) => {
        const res = await useApiStore().POST<LoginResponse>("/auth/login", input);
        token.value = res.token;
    }

    return { token, login }
})
