<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>

<template>
    <nav>
        <ul>
            <li>
                <RouterLink to="/">TYPING</RouterLink>
            </li>
            <li>
                <RouterLink to="/speed-records">Speed records</RouterLink>
            </li>
            <li style="float:right">
                <RouterLink to="/login">Connexion</RouterLink>
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
</style>