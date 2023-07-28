<script setup lang="ts">
import Footer from '../components/Footer.vue'
import Navigation from '../components/Navigation.vue'
import Typing from '../components/Typing.vue'
</script>

<template>
  <header>
    <Navigation />
  </header>
  <main>
    <Typing />
  </main>
  <footer>
    <Footer />
  </footer>
</template>

<style scoped>
footer {
    position: fixed;
    bottom: 0;
    margin: 0 12.5%;
    width: 75%;
    background-color: #000;
    color: #FFF;
    text-align: center;
}
</style>