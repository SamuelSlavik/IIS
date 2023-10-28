<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import Navigation from './components/Navigation.vue'
</script>

<template>
  <div>
    <Navigation />
  </div>

  <RouterView />
</template>

<style scoped>

@media (min-width: 1024px) {

}
</style>
