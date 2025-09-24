<script setup>
import { ref } from 'vue'
import ArticleList from '../article/ArticleList.vue'
import Searchbar from '../article/Searchbar.vue'
import { useRouter } from 'vue-router'

const searchText = ref('')

const router = useRouter()
const addPage = () => {
  router.push({ name: 'home' })
}
</script>

<template>
  <div class="bg-[#151517] flex flex-col h-screen text-white overflow-hidden">
    <div class="flex">
      <aside class="bg-[#1b1b1c] h-screen overflow-auto w-xs border-r-[1px] border-white/10">
        <div class="p-4">
          <Searchbar v-model:searchText="searchText" />
          <button @click="addPage"
            class="w-full outline-none h-12 px-4 bg-[#5686fe] duration-300 hover:opacity-80 rounded-xl cursor-pointer">
            Добавить
          </button>
          <ArticleList :searchText="searchText" />
        </div>
      </aside>
      <router-view :key="$route.fullPath" />
    </div>
  </div>
</template>