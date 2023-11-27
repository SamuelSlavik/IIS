<script setup lang="ts">
import {useUserStore} from "@/stores/user-store";

const user = useUserStore()
</script>

<template>
  <h2>Hi {{user.firstName}}</h2>
</template>
