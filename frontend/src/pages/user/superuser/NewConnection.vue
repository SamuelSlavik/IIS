<script setup lang="ts">
import {onMounted, ref} from "vue";
import Loader from "@/components/Loader.vue";
import type {NewConnection, User} from "@/lib/models";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";
import router from "@/router";
import type {LineInList} from "@/lib/models";
import {formatTimeForCreate} from "@/lib/utils";

const loading = ref<boolean>(false)
const notifications = useNotificationStore()

const newConnection = ref<NewConnection>({
  LineName: "",
  DepartureTime: "",
  ArrivalTime: "",
  Direction: true,
  NumberOfDays: null
})

const createConnection = async () => {
  loading.value = true
  try {
    const response = await axios.post(Endpoints.createConnection, {
      LineName: newConnection.value.LineName,
      DepartureTime: formatTimeForCreate(newConnection.value.DepartureTime),
      ArrivalTime: formatTimeForCreate(newConnection.value.ArrivalTime),
      Direction: newConnection.value.Direction,
      NumberOfDays: newConnection.value.NumberOfDays
    }, {withCredentials: true})
    if (response.status === 200) {
      notifications.addNotification("Connection created", 'success')
      await router.push('/profile/superuser/connections');
    }
  } catch (error) {
    notifications.addNotification("Failed to create connection: " + error, 'error')
  } finally {
    loading.value = false
  }
}

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

onMounted(() => {
  loadLines()
} )


</script>

<template>
  <div>
    <div class="header">
      <h2>Create new connection</h2>
    </div>

    <Loader v-if="loading"/>
    <div v-else>
      <form @submit.prevent="createConnection" class="form">
        <v-select
            v-model="newConnection.LineName"
            placeholder="Select line"
            :options="lines.map(({Name}) => {return Name})"
            required
        >
        </v-select>
        <div>
          <label>Departure time:</label>
          <input
              type="datetime-local"
              name="departureTime"
              v-model="newConnection.DepartureTime"
              required
          />
        </div>
        <div>
          <label>Arrival time:</label>
          <input
              type="datetime-local"
              name="arrivalTime"
              v-model="newConnection.ArrivalTime"
              required
          />
        </div>
        <div>
          <label for="direction">Direction from initial to final stop</label>
          <input
              type="checkbox"
              v-model="newConnection.Direction"
              name="direction"
          />
        </div>
        <input
            type="number"
            name="numberOfDays"
            v-model="newConnection.NumberOfDays"
            placeholder="Number of days"
            required
        />
        <button
            type="submit"
        >Create connection</button>
      </form>
    </div>
  </div>
</template>

<style>

</style>