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
          v-model="inputEmail"
          type="email"
          placeholder="Email"
          class="input input-bordered w-full"
        />
      </div>
      <div>
        <label for="password" class="label font-medium">Password:</label>
        <input
          v-model="inputPassword"
          type="password"
          placeholder="Password"
          class="input input-bordered w-full"
        />
      </div>
      <div>
        <p v-for="error in validationErrors" :key="error" class="text-sm text-red-500">
          {{ error }}
        </p>
      </div>
      <button type="submit" class="btn w-full btn-neutral">
        <span v-if="isLoading">
          <Icon name="tdesign:loading" class="w-6 h-6 animate-spin" />
        </span>
        <span v-else>Sign In</span>
      </button>
    </form>
  </div>
</template>

<script lang="ts" setup>
import { safeParse } from "valibot";
import type { LoginRequest } from "~/lib/types/auth";
import { schemaLoginRequest } from "~/lib/types/auth";

definePageMeta({
  layout: "landing",
});

const router = useRouter();
const authStore = useAuthStore();

const inputEmail = ref("admin@diploid.dev");
const inputPassword = ref("admin1234");
const isLoading = ref(false);
const validationErrors = ref<string[]>([]);

function handleSignIn() {
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
  authStore
    .login(data)
    .then(() => {
      router.push("/");
    })
    .finally(() => {
      isLoading.value = false;
    });
}
</script>
