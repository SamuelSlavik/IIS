<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import Navigation from './components/Navigation.vue'
import Notification from "@/components/Notification.vue";
</script>

<template>
  <div>
    <Navigation />
  </div>

  <RouterView />

  <Notification />
</template>

<style scoped>

@media (min-width: 1024px) {

}
</style>
