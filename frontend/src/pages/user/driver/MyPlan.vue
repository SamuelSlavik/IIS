<script setup lang="ts">
import {onMounted, ref} from "vue";
import Loader from "@/components/Loader.vue";
import type {ConnectionList, User} from "@/lib/models";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";
import lines from "@/pages/user/superuser/Lines.vue";
import {useUserStore} from "@/stores/user-store";
import {formatDateTime} from "../../../lib/utils";
// @ts-ignore
import Bus from "vue-material-design-icons/Bus.vue";
// @ts-ignore
import Tram from "vue-material-design-icons/Tram.vue";
// @ts-ignore
import Tank from "vue-material-design-icons/Tank.vue";

const loading = ref<boolean>(false)
const notifications = useNotificationStore()

const user = useUserStore()

const connections = ref<ConnectionList[]>([])

const loadPlan = async () => {
  loading.value = true
  try {
    const response = await axios.get(Endpoints.listConnectionsByDriver(user.id), {withCredentials: true})
    connections.value = response.data
  } catch (error) {
    notifications.addNotification("Failed to load lines: " + error, 'error')
  } finally {
    loading.value = false
  }
}

onMounted(() => {loadPlan()})

</script>

<template>
  <div>
    <div class="header">
      <h2>My plan</h2>
    </div>

    <Loader v-if="loading"/>
    <div v-else>
      <div class="table">
        <div v-for="(conn, index) in connections" :key="conn.ConnectionID" v-if="connections">
          <div class="list-item">
            <router-link :to="'/profile/driver/connection/detail/' + conn.ConnectionID" class="list-item__name">
              <b>{{ formatDateTime(conn.DepartureTime) }} - {{ formatDateTime(conn.ArrivalTime) }}</b>
            </router-link>
            <p class="list-item__role">{{conn.LineName}}</p>
            <p class="list-item__role"><b>From:</b> {{ conn.InitialStop }}</p>
            <p class="list-item__role"><b>To:</b> {{conn.FinalStop}}</p>
            <p class="list-item__role connection-title">
              <Bus v-if="conn.VehicleType === 'bus'" class="connection-icon"/>
              <Tram v-if="conn.VehicleType === 'tram'" class="connection-icon"/>
              <Tank v-if="conn.VehicleType === 'obrnena_dodavka'" class="connection-icon"/>
              {{conn.VehicleReg}}
            </p>
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