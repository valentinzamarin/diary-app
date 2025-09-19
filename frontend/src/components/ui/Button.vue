<script setup>
defineProps({
  type: {
    type: String,
    default: 'button'
  }
})

defineEmits(['click'])
</script>

<template>
  <button :type="type"
    class="hover:bg-neutral-700 absolute right-4 bottom-4 h-[34px] px-[12px] border-[1px] rounded-[18px] border-white/12 flex items-center justify-center cursor-pointer text-[13px] transition-colors duration-200">
    <slot>Button</slot>
  </button>
</template>
