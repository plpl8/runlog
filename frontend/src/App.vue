<script setup>
import { onMounted, ref } from 'vue'

const runs = ref([])

const form = ref({
  date: '',
  distance: '',
  duration_seconds: '',
  type: '',
  notes: ''
})

async function loadRuns() {
  const response = await fetch('http://localhost:8080/api/runs')
  runs.value = await response.json()
}

async function createRun() {
  const response = await fetch('http://localhost:8080/api/runs', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      date: `${form.value.date}T00:00:00Z`,
      distance: Number(form.value.distance),
      duration_seconds: Number(form.value.duration_seconds),
      type: form.value.type,
      notes: form.value.notes
    })
  })

  if (!response.ok) {
    alert('Не удалось добавить пробежку')
    return
  }

  form.value = {
    date: '',
    distance: '',
    duration_seconds: '',
    type: '',
    notes: ''
  }

  await loadRuns()
}

async function deleteRun(id) {
  await fetch(`http://localhost:8080/api/runs/${id}`, {
    method: 'DELETE'
  })

  await loadRuns()
}

onMounted(loadRuns)
</script>

<template>
  <main>
    <h1>🏃 RunLog</h1>
    <p>Трекер пробежек</p>

    <section>
      <h2>Добавить пробежку</h2>

      <form @submit.prevent="createRun">
        <input
          v-model="form.date"
          type="date"
          required
        />

        <input
          v-model="form.distance"
          type="number"
          step="0.1"
          placeholder="Дистанция, км"
          required
        />

        <input
          v-model="form.duration_seconds"
          type="number"
          placeholder="Время, секунд"
          required
        />

        <input
          v-model="form.type"
          type="text"
          placeholder="Тип"
          required
        />

        <input
          v-model="form.notes"
          type="text"
          placeholder="Заметки"
        />

        <button type="submit">
          Добавить
        </button>
      </form>
    </section>

    <section>
      <h2>Мои пробежки</h2>

      <p v-if="runs.length === 0">
        Пробежек пока нет.
      </p>

      <div v-for="run in runs" :key="run.id">
        <strong>{{ run.date.slice(0, 10) }}</strong>
        — {{ run.distance }} км
        — {{ run.duration_seconds }} сек
        — {{ run.type }}

        <span v-if="run.notes">
          — {{ run.notes }}
        </span>

        <button @click="deleteRun(run.id)">
          Удалить
        </button>
      </div>
    </section>
  </main>
</template>
