<script setup lang="ts">
import {useMovieStore} from "../store/movie";
import {useRoute} from "vue-router";

const store = useMovieStore()
const route = useRoute()

// Load data
await store.fetchDetailedMovieByUID((route.params.uid ?? '') as string)
await store.fetchDetailedMovieShowings()

// toReadableDateTime converts a json date string to a human readable date string
function toReadableDateTime(jsonDate: string, includeYear = false) {
  const date = new Date(jsonDate)
  const months = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec']
  return date.getDate() + ' ' + months[date.getMonth()] + (includeYear ? ' ' + date.getFullYear() : '') + ' ' + date.getHours() + ':' + date.getMinutes()
}
</script>

<template>
  <div class="flex justify-center">
    <template v-if="store.detailedMovie">
      <!-- card -->
      <div class="flex flex-col bg-white overflow-hidden w-full max-w-[80rem] md:flex-row md:rounded-3xl md:m-5 ">
        <img class="w-80 md:min-w-[25rem] self-center justify-self-center rounded"
             :src="store.detailedMovie.files?.find(f=>f.type === 'image')?.path" alt="movie picture">
        <!-- info -->
        <div class="">
          <div class="px-6 py-2">
            <div class="font-bold text-2xl mb-4">{{ store.detailedMovie.name }}</div>
            <p class="text-gray-700 text-se max-h-100 overflow-y-auto">
              {{ store.detailedMovie.description }}
            </p>
          </div>

          <!-- genres -->
          <div v-if="store.detailedMovie.genres" class="px-6 py-2">
            <span class="text-gray-600 text-sm py-2">Genres</span>
            <div class="mt-2">
              <span v-for="genre in store.detailedMovie.genres" :key="genre.uid"
                    class="inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2 mb-2">
                {{ genre.name }}
              </span>
            </div>
          </div>

          <!-- release date -->
          <div v-if="store.detailedMovie.classifications" class="px-6 py-2">
            <span class="text-gray-600 text-sm py-2">Release date</span>
            <div class="mt-2">
              {{ toReadableDateTime(store.detailedMovie.release_date, true) }}
            </div>
          </div>

          <!-- play time -->
          <div v-if="store.detailedMovie.classifications" class="px-6 py-2">
            <span class="text-gray-600 text-sm py-2">Play time</span>
            <div class="mt-2">
              {{ store.detailedMovie.total_time }} minutes
            </div>
          </div>

          <!-- classification -->
          <div v-if="store.detailedMovie.classifications" class="px-6 py-2">
            <span class="text-gray-600 text-sm py-2">Classification</span>
            <div class="mt-2">
              <span v-for="classification in store.detailedMovie.classifications" :key="classification.uid"
                    class="inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2 mb-2">
                  {{ classification.name }}
              </span>
            </div>
          </div>

          <!-- showings -->
          <div v-if="store.detailedMovieShowings.length" class="px-6 py-2">
            <table class="min-w-full shadow-md rounded">
              <thead class="bg-gray-50">
              <tr>
                <th class="p-4 text-left font-bold">Room</th>
                <th class="p-4 text-left font-bold">Time</th>
              </tr>
              </thead>
              <tbody class="divide-y divide-gray-300">
              <tr v-for="showing in store.detailedMovieShowings" :key="showing.uid">
                <td class="p-4">{{ showing.room.code }}</td>
                <td class="p-4">{{ toReadableDateTime(showing.start_time) }}</td>
              </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </template>

    <!-- not found placeholder -->
    <template v-else>
      <div class="bg-white rounded-3xl m-5 p-10">
        Movie not found
      </div>
    </template>
  </div>
</template>

