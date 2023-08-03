<script setup lang="ts">
import router from '@/router';
import { onMounted, ref } from 'vue'
import type { Ref } from 'vue'
import { RouterLink } from 'vue-router'

const isAuth: Ref<boolean> = ref(false)

onMounted(() => {
    isLoggedIn()
})

const isLoggedIn = async (): Promise<void> => {
    const res = await fetch("http://localhost:8080/api/user", {
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${localStorage.getItem('jwt') ?? ''}`
        }
    })

    isAuth.value = res.status === 200
}

const logout = (): void => {
    localStorage.removeItem('jwt')
    router.push('/')
}
</script>

<template>
    <nav>
        <ul>
            <li>
                <RouterLink to="/">TYPING</RouterLink>
            </li>
            <li>
                <RouterLink to="/best-scores">Best scores</RouterLink>
            </li>
            <li v-if="isAuth">
                <RouterLink to="/dashboard">Dashboard</RouterLink>
            </li>
            <li style="float:right" v-if="!isAuth">
                <RouterLink to="/login">Login</RouterLink>
            </li>
            <li style="float:right" v-if="isAuth">
                <button @click="logout">Logout</button>
            </li>
        </ul>
    </nav>
</template>

<style scoped>
ul {
    width: 75%;
    list-style-type: none;
    margin: 0 12.5%;
    padding: 0;
    overflow: hidden;
    background-color: #000;
}

li {
    float: left;
    border-right: 1px solid #FFF;
}

li a {
    display: block;
    color: white;
    text-align: center;
    padding: 14px 16px;
    text-decoration: none;
}

li a:hover {
    background-color: #05be4c;
}

li button {
    font-size: 16px;
    border: none;
    outline: none;
    color: white;
    padding: 14px 16px;
    background-color: inherit;
    font-family: inherit;
    margin: 0;
    cursor: pointer;
}
</style>