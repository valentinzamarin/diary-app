<script setup>
import { ref, onMounted } from 'vue'
import { EventsOn, EventsOff } from '../../../wailsjs/runtime'
import { useRouter } from 'vue-router'

const items = ref([])
const loading = ref(true)
const error = ref(null)
const router = useRouter()


const formattedDate = (time) => {
  const date = new Date(time);
  const options = {
    day: '2-digit',
    month: 'long',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  };

  return date.toLocaleDateString('ru-RU', options);
}


const loadEntries = async () => {
  try {
    loading.value = true
    error.value = null
    const entries = await window.go.main.App.GetEntries()

    entries.reverse()

    items.value = entries.map(entry => ({
      title: entry.Title,
      date: formattedDate(entry.Date),
      id: entry.ID
    }))


  } catch (err) {
    console.error('Error loading entries:', err)
    error.value = 'Не удалось загрузить записи'
  } finally {
    loading.value = false
  }
}


const onEntryCreated = (newEntry) => {
  items.value.unshift({
    title: newEntry.Title,
    date: formattedDate(newEntry.Date),
    id: newEntry.ID
  })
}

const openEntry = (id) => {
  router.push({ name: 'article', params: { id } })
}

onMounted(() => {
  loadEntries()
  EventsOn('entry:created', onEntryCreated)
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
      <li v-for="item in items" :key="item.id" @click="openEntry(item.id)"
        class="py-4 border-b-[1px] last:border-none border-white/10 cursor-pointer duration-300 hover:text-sky-700">
        <p class="truncate text-sm text-gray-500 dark:text-gray-400">{{ item.date }}</p>
        <h3 class="font-semibold">{{ item.title }}</h3>
      </li>
    </ul>
  </div>
</template>