<script setup>
import ArticleList from '../article/ArticleList.vue'
import ArticleEditor from '../article/ArticleEditor.vue'

</script>

<template>
  <div class="bg-[#151517] flex flex-col h-screen text-white overflow-hidden">
    <div class="flex">
      <aside class="bg-[#1b1b1c] h-screen overflow-auto w-xs border-r-[1px] border-white/10">
        <div class="p-4">
          <ArticleList />
        </div>
      </aside>
      <router-view :key="$route.fullPath" />
    </div>
  </div>
</template>