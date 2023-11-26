<script setup lang="ts">
import {onMounted, ref} from "vue";
import Loader from "@/components/Loader.vue";
import type {User} from "@/lib/models";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";
import type {Malfunction, MalfunctionReport, Vehicle} from "@/lib/models";
import router from "@/router";
import {useUserStore} from "@/stores/user-store";

const loading = ref<boolean>(false)
const notifications = useNotificationStore()

const user = useUserStore()

const vehicles = ref<Vehicle[]>([])

const reportId = router.currentRoute.value.params.id.toString() || ""

const newReport = ref<MalfunctionReport>({
  Title: "",
  Description: "",
  VehicleRef: "",
})

const loadMalfunction = async () => {
  loading.value = true
  try {
    const response = await axios.get<Malfunction>(Endpoints.retrieveMalfunction(reportId), {withCredentials: true})
    newReport.value.Title = response.data.Title
    newReport.value.Description = response.data.Description
    newReport.value.VehicleRef = response.data.Vehicle.Registration
  } catch (error) {
    notifications.addNotification("Failed to load malfunction: " + error, 'error')
  } finally {
    loading.value = false
  }
}
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
    const response = await axios.put(Endpoints.editMalfunction(reportId), newReport.value, {withCredentials: true})
    if (response.status === 200) {
      notifications.addNotification("Malfunction updated", 'success')
      user.role === "driver" ? await router.push('/profile/driver/reports') : await router.push('/profile/superuser/malfunctions');
    }
  } catch (error) {
    notifications.addNotification("Failed to update malfunction: " + error, 'error')
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadMalfunction()
  loadVehicles()
})

</script>

<template>
  <div>
    <div class="header">
      <h2>Edit malfunction report</h2>
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