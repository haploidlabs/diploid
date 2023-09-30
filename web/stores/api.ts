import { defineStore } from "pinia";
import { useAuthStore } from "./auth";

export const useApiStore = defineStore("api", () => {
    const apiUrl = useRuntimeConfig().public.apiUrl;

    const GET = async <T>(path: string) => {
        const token = useAuthStore().token;

        const response = await fetch(`${apiUrl}${path}`, {
            headers: {
                Authorization: token ? `Bearer ${token}` : "",
            },
        });
        return (await response.json()) as T;
    };

    const POST = async <T>(path: string, body: any) => {
        const token = useAuthStore().token;
        console.log("JDJDJDJDJDJ", body);

        const response = await fetch(`${apiUrl}${path}`, {
            method: "POST",
            headers: {
                Authorization: token ? `Bearer ${token}` : "",
            },
            body: JSON.stringify(body),
        });

        return (await response.json()) as T;
    };

    const PUT = async <T>(path: string, body: any) => {
        const token = useAuthStore().token;

        const response = await fetch(`${apiUrl}${path}`, {
            method: "PUT",
            headers: {
                Authorization: token ? `Bearer ${token}` : "",
            },
            body: JSON.stringify(body),
        });

        return (await response.json()) as T;
    }

    const PATCH = async <T>(path: string, body: any) => {
        const token = useAuthStore().token;

        const response = await fetch(`${apiUrl}${path}`, {
            method: "PATCH",
            headers: {
                Authorization: token ? `Bearer ${token}` : "",
            },
            body: JSON.stringify(body),
        });

        return (await response.json()) as T;
    }

    const DELETE = async <T>(path: string) => {
        const token = useAuthStore().token;

        const response = await fetch(`${apiUrl}${path}`, {
            method: "DELETE",
            headers: {
                Authorization: token ? `Bearer ${token}` : "",
            },
        });

        return (await response.json()) as T;
    }

    return { GET, POST, PUT, PATCH, DELETE };
});