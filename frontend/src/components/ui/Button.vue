<script setup>
defineProps({
  type: {
    type: String,
    default: 'button'
  }
})

const emit = defineEmits(['click'])
</script>

<template>
  <button :type="type" @click.prevent="emit('click')"
    class="hover:bg-neutral-700 h-[34px] px-[12px] border-[1px] rounded-[18px] border-white/12 flex items-center justify-center cursor-pointer text-base transition-colors duration-200"
    :class="$attrs.class">
    <slot>
      Button Text
    </slot>
  </button>
</template>