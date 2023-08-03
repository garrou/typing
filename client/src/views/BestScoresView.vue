<script setup lang="ts">
import Footer from '../components/Footer.vue'
import Navigation from '../components/Navigation.vue'
import BestScores from '../components/BestScores.vue'
</script>

<template>
    <header>
        <Navigation />
    </header>
    <main>
        <BestScores />
    </main>
    <footer>
        <Footer />
    </footer>
</template>