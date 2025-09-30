<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import Button from '../ui/Button.vue'


const route = useRoute()
const router = useRouter()
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

const editEntry = () => {
  router.push({ name: 'edit', params: { id: route.params.id } })
}

const removePermannently = async () => {
  const id = parseInt(route.params.id, 10)
  await window.go.main.App.DeleteEntryApp(id)
  router.push({ name: 'home' })
}


onMounted(() => {
  loadEntry()
})
</script>

<template>
  <div class="flex-1 p-6">
    <div class="max-w-4xl mx-auto h-screen overflow-scroll">
      <div class="flex justify-between items-center mb-6">
        <h2 class="text-3xl font-bold">
          {{ entry ? entry.Title : 'Загрузка...' }}
        </h2>
        <div class="flex gap-6 items-center">
          <Button @click="editEntry">
            Редактировать
          </Button>
          <Button @click="removePermannently()">
            Удалить
          </Button>
        </div>
      </div>
      <div class="prose prose-invert max-w-none">
        <div class="space-y-4 whitespace-pre-wrap text-[20px]/[38px] line-height-[24px] text-gray-200"
          v-html="entry ? entry.Text : ''">
        </div>
      </div>
      <div v-if="error" class="text-red-500 mt-4">
        {{ error }}
      </div>
    </div>
  </div>
</template>