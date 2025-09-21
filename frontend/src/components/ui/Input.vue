<script setup>
import { defineProps, defineEmits } from 'vue'

const props = defineProps({
  type: {
    type: String,
    default: 'text'
  },
  placeholder: {
    type: String,
    default: ''
  },
  classes: {
    type: String,
    default: ''
  },
  modelValue: {
    type: String,
    default: ''
  }
})

// Определяем событие на обновление значения
const emit = defineEmits(['update:modelValue'])

function onInput(event) {
  emit('update:modelValue', event.target.value)
}
</script>

<template>
  <input :type="props.type" :placeholder="props.placeholder" :class="props.classes" :value="props.modelValue"
    @input="onInput" />
</template>