import { defineStore } from 'pinia'
import {API} from "../helpers/api";
import {Movie} from "./models/movie";
import {Showing} from "./models/showing";

// Store used to load all movies and showings information
export const useMovieStore = defineStore('movie', {
    state: () => ({
        movies: [] as Movie[],
        detailedMovie: null as Movie | null,
        detailedMovieShowings: [] as Showing[]
    }),
    actions: {
        // Fetch all movies with a specific status
        async fetchMoviesByStatus(status = 'now') {
            const {response, ok} = await API.Req('GET', `/api/movies?status=${status}&load=files&load=genres`)
            if (ok) this.movies = response
            else this.movies = []
        },
        // Fetch all movies that match a search query
        async fetchMoviesByQuery(query: string) {
            const {response, ok} = await API.Req('GET', `/api/movies?query=${query}&load=files&load=genres`)
            if (ok) this.movies = response
            else this.movies = []
        },
        // Fetch detailed movie by uid
        async fetchDetailedMovieByUID(uid: string) {
            const {response, ok} = (await API.Req('GET', `/api/movies?uid=${uid}&load=files&load=genres&load=classifications`))
            if (ok) this.detailedMovie = response.find(Boolean)

        },
        // fetch showings for the current detailed movie
        async fetchDetailedMovieShowings() {
            const {response, ok} = await API.Req('GET', `/api/showings?movie=${this.detailedMovie?.uid}&load=room`)
            if (ok) this.detailedMovieShowings = response
        }
    }
})