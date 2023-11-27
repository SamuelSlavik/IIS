<script setup lang="ts">
import {onMounted, ref} from "vue";
import Loader from "@/components/Loader.vue";
import type {ConnectionList, User} from "@/lib/models";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";
import router from "@/router";
import {formatDate, formatDateTime} from "@/lib/utils";
import Magnify from "vue-material-design-icons/Magnify.vue";

const loading = ref<boolean>(false)
const notifications = useNotificationStore()
const conncetions = ref<ConnectionList[]>()
const line = router.currentRoute.value.params.line.toString() || ""

const cetTime = new Date()
cetTime.setHours(cetTime.getHours() + 2)
const querydate = ref<string>(cetTime.toISOString().split("T")[0])

const loadConnections = async () => {
  try {
    loading.value = true
    const querydatefmt = querydate.value.split("T")[0]
    const response = await axios.get(Endpoints.listConnectionsDatetime(line, querydatefmt), {withCredentials: true})
    conncetions.value = response.data
  } catch (error: any) {
    notifications.addNotification("Failed to load lines: " + error, "error")
  } finally {
    loading.value = false
  }
}

onMounted(() => {loadConnections()})

</script>

<template>
    <div>
    <div class="header">
      <h2>Manage connections</h2>
      <h3>Line: <i>{{ line }}</i></h3>
    </div>

    <div class="toolbar">
      <form @submit.prevent="loadConnections" class="search-form">
        <input
            type="date"
            name="connection-date"
            placeholder="Date of connections"
            v-model="querydate"
        />
        <button type="submit" class="small-button">
          <Magnify size="24px"/>
        </button>
      </form>
    </div>

    <Loader v-if="loading"/>
    <div v-else>
      <div class="table">
        <div v-for="(conn, index) in conncetions" :key="conn.ConnectionID">
          <div class="list-item">
            <router-link :to="'/profile/dispatcher/connection/detail/' + conn.ConnectionID" class="list-item__name">
              <b>{{ formatDateTime(conn.DepartureTime) }}</b>
            </router-link>
            <p class="list-item__role"><b>From:</b> {{ conn.InitialStop }}</p>
            <p><b>To:</b> {{conn.FinalStop}}</p>
          </div>
          <!-- Display table-hr only if it's not the last user for the current role -->
          <div v-if="index < conncetions.length - 1" class="table-hr"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<style>

</style>
