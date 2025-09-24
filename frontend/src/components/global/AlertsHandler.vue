<script setup>
import { useAlert } from '../../composables/useAlert'
import Alert from '../ui/Alert.vue';
const { alerts } = useAlert()
</script>

<template>
  <div class="fixed top-4 right-4 z-50 space-y-2">
    <Alert v-for="alert in alerts" :key="alert.id" :type="alert.type">
      {{ alert.message }}
    </Alert>
  </div>
</template>
