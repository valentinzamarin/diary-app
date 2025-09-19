<template>
  <input :type="type" :placeholder="placeholder" :class="classes" />
</template>

<script>
export default {
  props: {
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
    }
  }
}
</script>