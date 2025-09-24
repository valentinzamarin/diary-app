<script setup>

import { defineEmits, defineProps } from 'vue'

const props = defineProps({
  searchText: String
})

const emits = defineEmits(['update:searchText'])

const onInput = (e) => {
  emits('update:searchText', e.target.value)
}
</script>

<template>
  <form @submit.prevent class="flex flex-1 align-center pb-4">
    <input type="text" :value="searchText" @input="onInput" placeholder="Search.."
      class="w-full outline-none h-12 px-4 bg-[#43454a] rounded-xl" />
  </form>
</template>