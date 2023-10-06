import { acceptHMRUpdate, defineStore } from "pinia";
import { useAuthStore } from "./auth";

type Method = "GET" | "POST" | "PUT" | "PATCH" | "DELETE";

export const useApiStore = defineStore("api", () => {
  const apiUrl = useRuntimeConfig().public.apiUrl;

  interface ApiMethods {
    <T>(method: Method, path: string, body?: any): Promise<T>;
    (method: Method, path: string, body?: any): Promise<Response>;
    <T>(method: Method, path: string, body?: any): Promise<T>;
  }

  const fetchWrapper: ApiMethods = async <T extends any | Object>(
    method: Method,
    path: string,
    body: any,
  ) => {
    const token = useAuthStore().token;
    const options = {
      method,
      headers: {
        Authorization: token ? `Bearer ${token}` : "",
      },
      body: JSON.stringify(body),
    };

    return await $fetch<T>(`${apiUrl}${path}`, options);
  };

  const GET = async <T>(path: string) => await fetchWrapper<T>("GET", path, undefined);
  const POST = async <T>(path: string, body: any) => await fetchWrapper<T>("POST", path, body);
  const PUT = async <T>(path: string, body: any) => await fetchWrapper<T>("PUT", path, body);
  const PATCH = async <T>(path: string, body: any) => await fetchWrapper<T>("PATCH", path, body);
  const DELETE = async <T>(path: string) => await fetchWrapper<T>("DELETE", path, undefined);

  return { GET, POST, PUT, PATCH, DELETE };
});

if (import.meta.hot) import.meta.hot.accept(acceptHMRUpdate(useApiStore, import.meta.hot));
