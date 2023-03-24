<script setup lang="ts">
import TopBar from "./components/TopBar.vue";
</script>

<template>
  <div>
    <TopBar />
    <div class="relative mt-[6rem] min-h-[calc(100vh-6rem)] bg-fixed bg-gradient-to-b from-gray-400 to-gray-500">
      <!-- suspense for async viewws -->
      <suspense>
        <router-view />
      </suspense>
    </div>
  </div>
</template>