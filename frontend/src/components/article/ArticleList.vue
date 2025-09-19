<script setup>
import { ref } from 'vue'

const items = ref([
  {
    title: 'SQL Post 1',
    date: '19 сентября 2025 года'
  },
  {
    title: 'SQL Post 2',
    date: '20 сентября 2025 года'
  },
  {
    title: 'SQL Post 3',
    date: '21 сентября 2025 года'
  },
  {
    title: 'SQL Post 4',
    date: '22 сентября 2025 года'
  }
])
</script>

<template>

  <ul class="">
    <li v-for="(item, index) in items" :key="index"
      class="py-4 border-b-[1px] last:border-none border-white/10 cursor-pointer duration-300 hover:text-sky-700">
      <p class="truncate text-sm text-gray-500 dark:text-gray-400">{{ item.date }}</p>
      <h3 class="font-semibold">{{ item.title }}</h3>
    </li>
  </ul>

</template>