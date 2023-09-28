<script setup lang="ts">
import { useAuthStore } from '@/stores/auth'
import { ref } from 'vue'

const inputEmail = ref('admin@diploid.dev')
const inputPassword = ref('admin')

const authStore = useAuthStore()

const handleLogin = () => {
  authStore.login({
    email: inputEmail.value,
    password: inputPassword.value
  })
}
</script>

<template>
  <main>
    <h1 class="text-5xl">Login</h1>
    <form @submit.prevent="handleLogin">
      <input type="email" placeholder="Email" v-model="inputEmail" />
      <input type="password" placeholder="Password" v-model="inputPassword" />
      <button type="submit" class="bg-red-400">Login</button>
    </form>
  </main>
</template>
