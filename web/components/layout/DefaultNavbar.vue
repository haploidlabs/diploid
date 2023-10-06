<template>
  <nav class="navbar">
    <div class="navbar-start">
      <NuxtLink to="/" class="btn btn-ghost text-xl">Diploid</NuxtLink>
    </div>
    <div class="navbar-end">
      <NuxtLink v-if="isLoggedIn" to="/" class="btn btn-sm btn-ghost">
        <!--<Icon name="heroicons:home-20-solid" class="w-5 h-5" />-->
        <span>Dashboard</span>
      </NuxtLink>
      <NuxtLink v-else to="/auth/signin" class="btn btn-neutral btn-sm"> Sign In </NuxtLink>
    </div>
  </nav>
</template>

<script lang="ts" setup>
import { storeToRefs } from "pinia";

const authStore = useAuthStore();
const { isLoggedIn } = storeToRefs(authStore);
</script>
