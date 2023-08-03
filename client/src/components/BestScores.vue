<script setup lang="ts">
import { onMounted, ref } from 'vue'
import type { Ref } from 'vue'

interface BestScore {

    username: string

    score: number
}

const scores: Ref<BestScore[]> = ref([])

onMounted(() => {
    getBestScores()
})

const getBestScores = async (): Promise<void> => {
    const res = await fetch("http://localhost:8080/api/best-scores", {
        headers: {
            "Content-Type": "application/json",
        }
    })

    if (res.status === 200) {
        scores.value = await res.json()
    }
}
</script>

<template>
    <div class="scores">
        <h1>Best scores</h1>

        <table>
            <tr>
                <td>Username</td>
                <td>Score</td>
            </tr>
            <tr v-for="score in scores">
                <td>{{ score.username }}</td>
                <td>{{ score.score }}</td>
            </tr>
        </table>
    </div>
</template>

<style scoped>

table {
    width: 100%;
}

.scores {
    text-align: center;
    margin: auto 25%;
    padding: 10px;
}

td {
    border: 1px solid #000;
    padding: 1em;
    font-size: 1em;
}
</style>