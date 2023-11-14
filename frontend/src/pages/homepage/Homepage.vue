<script setup lang="ts">
import ConnectionCard from "@/components/ConnectionCard.vue"
import {onMounted, ref} from "vue";
import axios from "axios";
import type {Connection} from "@/lib/models";
import {Endpoints} from "@/lib/variables";
import { useNotificationStore} from "@/stores/notification-store";
import Loader from "@/components/Loader.vue";

const notifications = useNotificationStore()
const searchInProgress = ref<boolean>(false)

const connections = ref<Connection[]>([])

const getConnections = async () => {
  try {
    searchInProgress.value = true
    const response = await axios.get<Connection[]>(Endpoints.connections)
    connections.value = response.data
  } catch (error) {
    notifications.addNotification("Failed to get connections: " + error, "error")
  } finally {
    searchInProgress.value = false
  }

  console.log(connections.value)
}

onMounted(() => {
  getConnections()
})
</script>

<template>
  <div class="container">
    <h1>Connections</h1>
    <Loader v-if="searchInProgress"/>
    <ConnectionCard
      v-for="connection in connections"
      :key="connection.ID"
      :id="connection.ID"
      :lineName="connection.LineName"
      :type="connection.Type"
      :stops="connection.ListStops"
    />
  </div>
</template>
