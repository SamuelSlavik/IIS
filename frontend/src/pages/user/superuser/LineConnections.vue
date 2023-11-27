<script setup lang="ts">
import {onMounted, ref} from "vue";
import Loader from "@/components/Loader.vue";
import type {ConnectionList, User} from "@/lib/models";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";
import router from "@/router";
import {formatDate, formatDateTime} from "@/lib/utils";
// @ts-ignore
import Magnify from "vue-material-design-icons/Magnify.vue";
// @ts-ignore
import Close from "vue-material-design-icons/Close.vue";
// @ts-ignore
import Check from "vue-material-design-icons/Check.vue";
// @ts-ignore
import Delete from "vue-material-design-icons/Delete.vue";
// @ts-ignore

import Pencil from "vue-material-design-icons/Pencil.vue";

const loading = ref<boolean>(false)
const notifications = useNotificationStore()
const connections = ref<ConnectionList[]>()
const line = router.currentRoute.value.params.line.toString() || ""

const cetTime = new Date()
cetTime.setHours(cetTime.getHours() + 2)
const querydate = ref<string>(cetTime.toISOString().split("T")[0])

const loadConnections = async () => {
  try {
    loading.value = true
    const querydatefmt = querydate.value.split("T")[0]
    const response = await axios.get(Endpoints.listConnectionsDatetime(line, querydatefmt), {withCredentials: true})
    connections.value = response.data
  } catch (error: any) {
    notifications.addNotification("Failed to load lines: " + error, "error")
  } finally {
    loading.value = false
  }
}

const deleteConnection = async (id: string) => {
  const numberOfDaysInput = window.prompt("Enter the number of days to delete:", "1");

  let numberOfDays = parseInt(numberOfDaysInput || "1")

  try {
    const response = await axios.delete(Endpoints.deleteConnection(id, numberOfDays.toString()), {withCredentials: true})
    if (response.status === 200) {
      notifications.addNotification("Connection deleted", "success")
      await loadConnections()
    }
  } catch (error) {
    notifications.addNotification("Failed to delete connection: " + error, "error")
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
        <div v-for="(conn, index) in connections" :key="conn.ConnectionID" v-if="connections">
          <div class="list-item">
            <router-link :to="'/profile/superuser/connection/edit/' + conn.ConnectionID" class="list-item__name">
              <b>{{ formatDateTime(conn.DepartureTime) }}</b>
            </router-link>
            <p class="list-item__role"><b>From:</b> {{ conn.InitialStop }}</p>
            <p class="list-item__role"><b>To:</b> {{conn.FinalStop}}</p>
            <div class="list-item__tools">
              <router-link :to="'/profile/superuser/connections/edit/' + conn.ConnectionID"><Pencil :size="24" /></router-link>
              <a @click="deleteConnection(conn.ConnectionID)"><Delete :size="24" /></a>
            </div>
          </div>
          <!-- Display table-hr only if it's not the last user for the current role -->
          <div v-if="index < connections.length - 1" class="table-hr"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<style>

</style>
