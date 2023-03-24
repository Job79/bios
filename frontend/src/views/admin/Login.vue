<script setup lang="ts">
import {ref} from "vue";
import {useAdminStore} from "../../store/admin";
import {useRouter} from "vue-router";

const router = useRouter()
const adminStore = useAdminStore()

const name = ref('')
const password = ref('')
const message = ref('')

async function login() {
  const ok = await adminStore.login(name.value, password.value)
  if (ok) await router.push({name: 'admin-movies'})
  else message.value = 'invalid credentials'
}
</script>

<template>
  <div class="flex justify-center items-center absolute inset-0">
    <div class="bg-white border-l-gray-800 border-l-8 rounded-lg p-8 md:px-16 md:py-12">
      <div class="mb-4">
        <label class="block text-gray-700 text-md font-bold mb-2" for="username">
          Username
        </label>
        <input v-model="name" class="shadow border rounded py-2 px-3 text-gray-800 w-80" id="username" type="text"
               placeholder="Username"
               @keydown.enter="login">
      </div>
      <div class="mb-6">
        <label class="block text-gray-700 text-md font-bold mb-2" for="password">
          Password
        </label>
        <input v-model="password" class="shadow border rounded py-2 px-3 text-gray-800 mb-3 w-80" id="password"
               type="password" placeholder="******************"
               @keydown.enter="login">
        <p class="text-red-600">{{ message }}</p>
      </div>
      <button
          class="bg-gray-800 text-white font-bold py-3 px-5 rounded"
          type="button"
          @click="login">
        Sign In
      </button>
    </div>
  </div>
</template>

