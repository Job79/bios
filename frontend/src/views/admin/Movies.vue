<script setup lang="ts">
import {useAdminStore} from "../../store/admin";
import {useRouter} from "vue-router";

const router = useRouter()
const adminStore = useAdminStore()

if (!adminStore.token) router.push({name: 'admin-login'})
</script>

<template>
    <div>TODO movies page</div>
</template>

