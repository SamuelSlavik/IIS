<script setup lang="ts">
import {onMounted, ref} from "vue";
import Loader from "@/components/Loader.vue";
import type {User} from "@/lib/models";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";

const loading = ref<boolean>(false)
const notifications = useNotificationStore()

const loadConnections = async () => {
  try {
    loading.value = true
    const response = await axios.get<User[]>(Endpoints.listUsers(query.value), {withCredentials: true})
    users.value = response.data
    console.log(users.value)
  } catch (error) {
    notifications.addNotification("Failed to get users: " + error, "error")
  } finally {
    loading.value = false
  }
}



</script>

<template>
  <div>
    <div class="header">
      <h2>Manage connections</h2>
    </div>

    <Loader v-if="loading"/>
    <div v-else>

    </div>
  </div>
</template>

<style>

</style>