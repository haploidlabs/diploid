<script lang="ts" setup>
import { is, parse, safeParse } from "valibot";
import { LoginRequest, schemaLoginRequest } from "~/lib/types/auth";
import { useAuthStore } from "~/stores/auth";

definePageMeta({
  layout: "landing",
});

const authStore = useAuthStore();

const inputEmail = ref("");
const inputPassword = ref("");
const isLoading = ref(false);
const validationErrors = ref<string[]>([]);

const handleSignIn = () => {
  isLoading.value = true;
  validationErrors.value = [];
  const data = {
    email: inputEmail.value,
    password: inputPassword.value,
  } satisfies LoginRequest;
  const validation = safeParse(schemaLoginRequest, data);
  if (!validation.success) {
    validationErrors.value = validation.issues.map((issue) => issue.message);
    isLoading.value = false;
    return;
  }
  authStore.login(data).finally(() => {
    isLoading.value = false;
  });
};
</script>

<template>
  <div class="max-w-md mx-auto p-4 space-y-6">
    <div class="text-center space-y-2">
      <h1 class="text-4xl font-bold">Sign In</h1>
      <p class="text-black/80">Sign in to your account.</p>
    </div>
    <form class="form-control flex flex-col gap-4" @submit.prevent="handleSignIn">
      <div>
        <label for="email" class="label font-medium">Email:</label>
        <input
          type="email"
          placeholder="Email"
          class="input input-bordered w-full"
          v-model="inputEmail"
        />
      </div>
      <div>
        <label for="password" class="label font-medium">Password:</label>
        <input
          type="password"
          placeholder="Password"
          class="input input-bordered w-full"
          v-model="inputPassword"
        />
      </div>
      <div>
        <p v-for="error in validationErrors" class="text-sm text-red-500">{{ error }}</p>
      </div>
      <button type="submit" class="btn w-full btn-neutral">
        <span v-if="isLoading"><LucideLoader2 class="w-6 h-6 animate-spin" /></span>
        <span v-else>Sign In</span>
      </button>
    </form>
  </div>
</template>
