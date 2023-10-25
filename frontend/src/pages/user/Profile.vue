<script setup lang="ts">

import Dashboard from "@/components/Dashboard.vue";
</script>

<template>
  <div class="profile-wrapper">
    <Dashboard />
    <div class="profile">
      <div class="profile-content">
        <h1>Welcome!</h1>
      </div>
    </div>
  </div>
</template>

<style>
  .profile-wrapper {
    display: flex;
    min-height: 100vh;
  }
  .profile {
    width: calc((100% - (1280px - 64px)) / 2 + 930px - 32px);
  }
  .profile-content {
    width: 930px;
    margin-right: auto;
    margin-left: 0;
    padding: 2rem 1rem 2rem 2rem;
  }
</style>