import { acceptHMRUpdate, defineStore } from "pinia";
import { ref } from "vue";
import { useApiStore } from "./api";
import type { LoginRequest, LoginResponse } from "@/lib/types/auth";

const COOKIE_TOKEN = "token";

export const useAuthStore = defineStore("auth", () => {
  const token = ref<string | null>(null);

  const load = () => {
    const cToken = useCookie(COOKIE_TOKEN);
    if (cToken.value) token.value = cToken.value;
  };

  const login = async (input: LoginRequest) => {
    const res = await useApiStore().POST<LoginResponse>("/auth/login", input);
    token.value = res.token;
    useCookie(COOKIE_TOKEN).value = token.value;
  };

  const isLoggedIn = computed(() => token.value !== null);

  return { load, token, login, isLoggedIn };
});

if (import.meta.hot) import.meta.hot.accept(acceptHMRUpdate(useAuthStore, import.meta.hot));
