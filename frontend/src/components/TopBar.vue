<script setup>
import {RouterLink, useRouter} from 'vue-router'
import {ref} from "vue";

const router = useRouter()

const open = ref(true) // determines whether the topbar is open on mobile
const query = ref('')
const search = () => router.push({'name': 'home', query: {'query': query.value}})
</script>

<template>
  <nav class="fixed top-0 left-0 right-0 min-h-[6rem] p-6 z-10 flex items-center justify-between flex-wrap bg-gray-800">
    <!-- Logo -->
    <router-link :to="{'name': 'home'}" class="flex items-center flex-shrink-0 text-white mr-6">
      <span class="self-center text-xl font-semibold">Bios</span>
    </router-link>

    <!-- Hamburger menu on mobile -->
    <div class="block md:hidden" @click="open = !open">
      <button class="flex items-center px-3 py-2 border rounded text-white border-white hover:text-white hover:border-white">
        <svg class="fill-current h-3 w-3" viewBox="0 0 20 20"><path d="M0 3h20v2H0V3zm0 6h20v2H0V9zm0 6h20v2H0v-2z"/></svg>
      </button>
    </div>

    <!-- Navigation items -->
    <div :class="{'hidden': open}" class="w-full block flex-grow md:flex md:items-center md:w-auto md:block">
      <div class="text-md md:flex-grow">
        <router-link :to="{'name': 'home', query: {'status': 'now'}}" class="block mt-4 md:inline-block md:mt-0 text-gray-200 hover:text-white mr-4">
          Now
        </router-link>
        <router-link :to="{'name': 'home', query: {'status': 'soon'}}" class="block mt-4 md:inline-block md:mt-0 text-gray-200 hover:text-white mr-4">
          Soon
        </router-link>
        <router-link :to="{'name': 'home', query: {'status': 'ended'}}" class="block mt-4 md:inline-block md:mt-0 text-gray-200 hover:text-white">
          Ended
        </router-link>
      </div>

      <!-- search bar -->
      <div class="mt-5 md:mt-0">
          <div class="flex border-2 rounded bg-white">
            <button @click="search" class="px-3"><svg class="fill-current w-6 h-6" viewBox="0 0 24 24"><path d="M16.32 14.9l5.39 5.4a1 1 0 0 1-1.42 1.4l-5.38-5.38a8 8 0 1 1 1.41-1.41zM10 16a6 6 0 1 0 0-12 6 6 0 0 0 0 12z" /></svg></button>
            <input type="text" v-model="query" @keyup="search" class="px-4 py-2 w-full outline-0" placeholder="Search movie">
          </div>
        </div>
    </div>
  </nav>
</template>