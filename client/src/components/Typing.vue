<script setup lang="ts">
import { ref } from 'vue'
import type { Ref } from 'vue'

const generate = (): string[] => {
  return [
    'un', 'deux', 'trois', 'quatre', 'cinq', 'six', 'sept', 'huit', 'neuf', 'dix',
    'un', 'deux', 'trois', 'quatre', 'cinq', 'six', 'sept', 'huit', 'neuf', 'dix'
  ]
}

let time: Ref<number> = ref(60)
const words: Ref<string[]> = ref(generate())
const score: Ref<number> = ref(0)

const check = (e: any): void => {

  const keyCode = e.keyCode
  const word = e.target.value.trim()

  if (![32, 13].includes(keyCode)) {
    return
  }

  if (word === words.value[0]) {
    removeWord()
    updateScore()
    resetField(e)
  } else {
    errorField(e)
  }
}

const startTimer = () => {
  setInterval(() => {
    time.value--
  }, 1000)
}

const removeWord = (): void => {
  words.value.shift()
}

const updateScore = (): void => {
  score.value++
}

const resetField = (e: any): void => {
  e.target.value = ''
  e.target.style.borderColor = '#000'
}

const errorField = (e: any): void => {
  e.target.style.borderColor = '#f51911'
}

const clear = (e: any): void => {

  if (e.target.value.trim().length === 0) {
    e.target.value = ''
  }
}
</script>

<template>
  <div class="typing">

    <div class="words">
      <span class="word" v-for="word in words">{{ word }}</span>
    </div>

    <input type="text" @keypress="startTimer" @keydown="check" @keyup="clear" />

    <div class="infos">
      <div>Time : {{ time }}</div>
      <div>Score : {{ score }}</div>
    </div>
  </div>
</template>

<style scoped>
p {
  font-size: 1.5em;
}

.typing {
  margin: 10em auto;
  width: 50%;
  padding: 10px;
}

.words {
  padding: 1em;
  border: 3px solid;
  overflow-x: hidden;
  border-radius: 10px;
}

.words span:nth-child(1) {
  color: #05be4c;
}

.word {
  font-size: 1.5em;
  margin-right: 1em;
}

input {
  width: 100%;
  padding: 1em 0.5em;
  margin: 8px 0;
  font-size: 2em;
  border-radius: 10px;
  outline: none !important;
  border: 0.1em solid #000;
}

.infos {
  font-size: 1.5em;
  text-align: center;
  justify-content: space-between;
}
</style>