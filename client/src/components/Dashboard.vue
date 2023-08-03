<script setup lang="ts">
import { onMounted, ref } from 'vue'
import type { Ref } from 'vue'

interface Score {

    score: number

    createdAt: Date
}

const scores: Ref<Score[]> = ref([])

onMounted(() => {
    getScores()
})

const getScores = async (): Promise<void> => {
    const res = await fetch("http://localhost:8080/api/scores", {
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${localStorage.getItem('jwt')}`
        }
    })

    if (res.status === 200) {
        scores.value = await res.json()
    }
}
</script>

<template>
    <div class="scores">
        <table>
            <tr>
                <td>Date</td>
                <td>Score</td>
            </tr>
            <tr v-for="score in scores">
                <td>{{ score.createdAt }}</td>
                <td>{{ score.score }}</td>
            </tr>
        </table>
    </div>
</template>

<style scoped>
.scores {
    text-align: center;
    margin: auto;
    padding: 10px;
}

td {
    border: 1px solid #000;
    padding: 1em;
    font-size: 1em;
}
</style>