<script setup lang="ts">
import {onMounted, ref} from "vue";
import Loader from "@/components/Loader.vue";
import type {ConnectionUnauth} from "@/lib/models";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";
import lines from "@/pages/user/superuser/Lines.vue";
import {useUserStore} from "@/stores/user-store";
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

const connection = ref<ConnectionUnauth>()

const loadConnection = async () => {
  try {
    loading.value = true
    const response = await axios.get(Endpoints.ConnectionDetailNotLoggedDatetime(connectionId), {withCredentials: true})
    connection.value = response.data
  } catch (error: any) {
    notifications.addNotification("Failed to load connection: " + error, "error")
  } finally {
    loading.value = false
  }
}

onMounted(() => {loadConnection()})

</script>

<template>
  <div class="container">
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
            <Bus v-if="connection.Type === 'bus'" class="connection-icon"/>
            <Tram v-if="connection.Type === 'tram'" class="connection-icon"/>
            <Tank v-if="connection.Type === 'obrnena_dodavka'" class="connection-icon"/>
          </p>
        </div>
      </div>
      <div class="table">
        <div v-for="(stop, index) in connection.ListStops" :key="stop.StopName">
          <div class="list-item">
            <p class="list-item__name">
              <b>{{ stop.StopName}}</b>
            </p>
            <p class="list-item__role">
              {{stop.DepartureTime}}
            </p>
          </div>
          <div v-if="index < connection.ListStops.length - 1" class="table-hr"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<style>

</style>
