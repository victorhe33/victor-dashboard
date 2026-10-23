<script setup>
import { onMounted, ref } from 'vue'

const health = ref(null)
const widgets = ref([])
const error = ref(null)

async function getJSON(path) {
  const res = await fetch(path)
  if (!res.ok) throw new Error(`${path}: ${res.status}`)
  return res.json()
}

onMounted(async () => {
  try {
    ;[health.value, widgets.value] = await Promise.all([
      getJSON('/api/health'),
      getJSON('/api/widgets'),
    ])
  } catch (e) {
    error.value = e.message
  }
})
</script>

<template>
  <main>
    <header>
      <h1>Victor Dashboard</h1>
      <span class="status" :class="{ ok: health?.status === 'ok' }">
        API: {{ error ? 'unreachable' : health?.status ?? '…' }}
      </span>
    </header>

    <p v-if="error" class="error">{{ error }}</p>

    <section class="grid">
      <article v-for="w in widgets" :key="w.id" class="card">
        <h2>{{ w.title }}</h2>
        <p class="value">{{ w.value }}</p>
      </article>
    </section>
  </main>
</template>
