<script lang="ts" setup>
const authStore = useAuthStore();

authStore.load();
</script>

<template>
  <NuxtLayout>
    <NuxtPage />
  </NuxtLayout>
</template>
