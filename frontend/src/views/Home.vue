<script setup lang="ts">
import MovieTile from "@/components/MovieTile.vue";
import {useMovieStore} from "../store/movie";
import {watch} from "vue";
import {useRoute} from "vue-router";

const store = useMovieStore()
const route = useRoute()

// Load data
if (route.query.query) {
  await store.fetchMoviesByQuery(route.query.query as string)
} else {
  await store.fetchMoviesByStatus((route.query.status ?? 'now') as string)
}

// Change data when route params change
watch(() => route.query.status, async () => {
  await store.fetchMoviesByStatus((route.query.status ?? 'now') as string)
})

watch(() => route.query.query, async () => {
  if(route.query.query) await store.fetchMoviesByQuery((route.query.query ?? '') as string)
  else await store.fetchMoviesByStatus((route.query.status ?? 'now') as string)
})
</script>

<template>
  <div class="flex justify-center flex-wrap">
    <!-- movie list -->
    <template v-if="store.movies.length">
      <div v-for="movie in store.movies" :key="movie.uid">
        <router-link :to="{path: `/movie/${movie.uid}`}">
          <movie-tile :movie="movie" />
        </router-link>
      </div>
    </template>

    <!-- not found placeholder -->
    <template v-else>
      <div class="bg-white rounded-3xl m-5 p-10">
        No movies found
      </div>
    </template>
  </div>
</template>

