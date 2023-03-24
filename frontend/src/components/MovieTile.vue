<script setup lang="ts">
import {PropType, defineProps} from "vue";
import {Movie} from "../store/models/movie";

const props = defineProps({
  movie: {
    type: Object as PropType<Movie>,
    required: true,
  }
})
</script>

<template>
  <!-- animation -->
  <div class="transition ease-in-out hover:scale-110 duration-300">
    <!-- card -->
    <div class="bg-white rounded-3xl overflow-hidden w-[25rem] m-5">
      <img class="w-full" :src="props.movie.files?.find(f=>f.type === 'image')?.path" alt="movie picture">
      <!-- info -->
      <div class="px-6 py-4">
        <div class="font-bold text-xl mb-2">{{movie.name}}</div>
        <p class="text-gray-700 text-base max-h-40 overflow-y-auto">
          {{movie.description}}
        </p>
      </div>
      <!-- genres -->
      <div class="px-6 pt-4 pb-2">
        <span v-for="genre in movie.genres" :key="genre.uid" class="inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2 mb-2">{{genre.name}}</span>
      </div>
    </div>
  </div>
</template>