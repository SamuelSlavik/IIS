<script setup lang="ts">
import {onMounted, ref} from "vue";
import Loader from "@/components/Loader.vue";
import type {ConnectionDetail, ConnectionList, User} from "@/lib/models";
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
import router from "@/router";

const loading = ref<boolean>(false)
const notifications = useNotificationStore()

const connectionId = router.currentRoute.value.params.id.toString() || ""

const user = useUserStore()

const connection = ref<ConnectionDetail>()

const loadConnection = async () => {
  try {
    loading.value = true
    const response = await axios.get(Endpoints.driverDetailConnection(connectionId), {withCredentials: true})
    connection.value = response.data
  } catch (error: any) {
    notifications.addNotification("Failed to load lines: " + error, "error")
  } finally {
    loading.value = false
  }
}

onMounted(() => {loadConnection()})

</script>

<template>
  <div>
    <Loader v-if="loading"/>
    <div v-else v-if="connection">
      <div class="header">
        <h2>Connection's detail</h2>
      </div>
      <div class="hr"></div>
      <div class="details">
        <div class="details-item">
          <p>Line:</p>
          <p>{{ connection.LineName }}</p>
        </div>
        <div class="details-item">
          <p>Vehicle:</p>
          <p class="connection-title">
            <Bus v-if="connection.VehicleType === 'bus'" class="connection-icon"/>
            <Tram v-if="connection.VehicleType === 'tram'" class="connection-icon"/>
            <Tank v-if="connection.VehicleType === 'obrnena_dodavka'" class="connection-icon"/>
            {{connection.VehicleReg}}
          </p>
        </div>
      </div>
      <div class="table">
        <div v-for="(stop, index) in connection.StopInConnection" :key="stop.StopName">
          <div class="list-item">
            <p class="list-item__name">
              <b>{{ stop.StopName}}</b>
            </p>
            <p class="list-item__role">
              {{stop.DepartureTime}}
            </p>
          </div>
          <div v-if="index < connection.StopInConnection.length - 1" class="table-hr"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<style>

</style>
