<script setup lang="ts">
import {onMounted, ref} from "vue";
import Loader from "@/components/Loader.vue";
import type {LineInList, User} from "@/lib/models";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";

const loading = ref<boolean>(false)
const notifications = useNotificationStore()
const lines = ref<LineInList[]>([])

const loadLines = async () => {
  try {
    loading.value = true
    const response = await axios.get(Endpoints.listLines, {withCredentials: true})
    lines.value = response.data
  } catch (error: any) {
    notifications.addNotification("Failed to load lines: " + error, "error")
  } finally {
    loading.value = false
  }
}

onMounted(() => {loadLines()})

</script>

<template>
  <div class="container">
    <div class="header"><h2>Search Connections</h2></div>

    <Loader v-if="loading"/>
    <div v-else>
      <div class="table">
        <div v-for="(line, index) in lines" :key="line.Name">
          <div class="list-item">
            <router-link :to="'/connections/' + line.Name" class="list-item__name">
              <b>{{ line.Name }}</b>
            </router-link>
            <p class="list-item__role"><b>From:</b> {{ line.InitialStop }}</p>
            <p><b>To:</b> {{line.FinalStop}}</p>
          </div>
          <!-- Display table-hr only if it's not the last user for the current role -->
          <div v-if="index < lines.length - 1" class="table-hr"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<style>

</style>