<script setup lang="ts">
import {onMounted, ref} from "vue";
import Loader from "@/components/Loader.vue";
import type {MalfunctionReport, User, Vehicle} from "@/lib/models";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";
import router from "@/router";


const loading = ref<boolean>(false)
const notifications = useNotificationStore()

const vehicles = ref<Vehicle[]>([])

// @ts-ignore
const newReport = ref<MalfunctionReport>({
  Title: "",
  Description: "",
  VehicleRef: "",
})

const loadVehicles = async () => {
  loading.value = true
  try {
    const response = await axios.get<Vehicle[]>(Endpoints.listVehicles, {withCredentials: true})
    vehicles.value = response.data
  } catch (error) {
    notifications.addNotification("Failed to load vehicles: " + error, 'error')
  } finally {
    loading.value = false
  }
}

const submitReport = async () => {
  loading.value = true
  try {
    // @ts-ignore
    const response = await axios.post(Endpoints.reportMalfunction, newReport.value, {withCredentials: true})
    if (response.status === 200) {
      notifications.addNotification("Malfunction reported", 'success')
      await router.push('/profile/driver/reports');
    }
  } catch (error) {
    notifications.addNotification("Failed to report malfunction: " + error, 'error')
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadVehicles()
})

</script>

<template>
  <div>
    <div class="header">
      <h2>Report a malfunction on a vehicle</h2>
    </div>

    <Loader v-if="loading"/>
    <div v-else>
      <form @submit.prevent="submitReport" class="form">
        <input
            type="text"
            name="title"
            placeholder="Title"
            v-model="newReport.Title"
            required
        />
        <textarea
            name="description"
            placeholder="Description"
            v-model="newReport.Description"
            required
        />
        <v-select v-model="newReport.VehicleRef" :options="vehicles.map(({Registration}) => {return Registration})" placeholder="Select vehicle"></v-select>

        <button
            type="submit"
        >Submit report</button>
      </form>
    </div>
  </div>
</template>

<style>

</style>