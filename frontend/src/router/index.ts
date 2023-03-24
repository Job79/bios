import {createRouter, createWebHistory, RouteRecordRaw} from 'vue-router'

const Home = () => import ('@/views/Home.vue')
const MovieDetails = () => import ('@/views/MovieDetails.vue')

// Admin pages
const AdminLogin = () => import ('@/views/admin/Login.vue')
const AdminMovies = () => import ('@/views/admin/Movies.vue')

const routes: Array<RouteRecordRaw> = [
    {
        path: '/',
        name: 'home',
        component: Home
    },
    {
        path: '/movie/:uid',
        name: 'movie-details',
        component: MovieDetails
    },
    {
        path: '/admin/login',
        alias: '/admin',
        name: 'admin-login',
        component: AdminLogin
    },
    {
        path: '/admin/movies',
        name: 'admin-movies',
        component: AdminMovies
    }

]

export default createRouter({
    history: createWebHistory(),
    routes
})