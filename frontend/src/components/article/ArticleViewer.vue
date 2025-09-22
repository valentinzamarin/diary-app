<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const entry = ref(null)
const error = ref(null)

const loadEntry = async () => {
  try {
    const id = parseInt(route.params.id, 10)
    const result = await window.go.main.App.GetEntry(id)
    entry.value = result
  } catch (err) {
    console.error(err)
    error.value = 'Ошибка при загрузке записи'
  }
}

onMounted(() => {
  loadEntry()
})
</script>

<template>
  <div class="flex-1 p-6">
    <div class="max-w-4xl mx-auto">
      <div class="flex justify-between items-center mb-6">
        <h2 class="text-2xl font-bold">
          {{ entry ? entry.Title : 'Загрузка...' }}
        </h2>
      </div>
      <div class="prose prose-invert max-w-none">
        <div class="whitespace-pre-wrap font-sans text-gray-200" v-html="entry ? entry.Text : ''"></div>
      </div>
      <div v-if="error" class="text-red-500 mt-4">
        {{ error }}
      </div>
    </div>
  </div>
</template>