<script setup>
import Input from '../ui/Input.vue';

const searchByTitle = () => {
  alert('hello')
}
</script>

<template>
  <form @submit.prevent="searchByTitle()" class="flex flex-1 align-center pb-8">
    <Input type="text" placeholder="Search.." class="w-full outline-none h-12 px-4 bg-[#43454a] rounded-xl" />
  </form>
</template>