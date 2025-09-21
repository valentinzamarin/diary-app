<script setup>
import { ref, onMounted } from 'vue'

const items = ref([])
const loading = ref(true)
const error = ref(null)

const loadEntries = async () => {
  try {
    loading.value = true
    error.value = null
    const entries = await window.go.main.App.GetEntries()

    items.value = entries.map(entry => ({
      title: entry.Title,
      date: entry.CreatedAt,
      id: entry.ID
    }))
  } catch (err) {
    console.error('Error loading entries:', err)
    error.value = 'Не удалось загрузить записи'
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadEntries()
})
</script>

<template>
  <div>
    <div v-if="loading" class="py-4 text-center">
      <p class="text-gray-500">Загрузка записей...</p>
    </div>

    <div v-if="error" class="py-4 text-center text-red-500">
      {{ error }}
    </div>

    <div v-if="!loading && items.length === 0" class="py-4 text-center">
      <p class="text-gray-500">Записей пока нет</p>
    </div>

    <ul v-if="!loading && items.length > 0" class="">
      <li v-for="item in items" :key="item.id"
        class="py-4 border-b-[1px] last:border-none border-white/10 cursor-pointer duration-300 hover:text-sky-700">
        <p class="truncate text-sm text-gray-500 dark:text-gray-400">{{ item.date }}</p>
        <h3 class="font-semibold">{{ item.title }}</h3>
      </li>
    </ul>
  </div>
</template>