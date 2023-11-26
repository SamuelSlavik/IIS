<script setup lang="ts">
import {onMounted, ref} from "vue";
import Loader from "@/components/Loader.vue";
import type {User, Vehicle} from "@/lib/models";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import Pencil from "vue-material-design-icons/Pencil.vue";
import Delete from "vue-material-design-icons/Delete.vue";
import {useNotificationStore} from "@/stores/notification-store";
import router from "@/router";

const loading = ref<boolean>(false)
const notifications = useNotificationStore()

const vehicle = ref<Vehicle | undefined>(undefined)
const malfunctions = ref<Vehicle | undefined>(undefined)
const maintenanceRequests = ref<undefined>(undefined)

const loadVehicle = async () => {
  try {
    loading.value = true
    const response = await axios.get<Vehicle>(Endpoints.retrieveVehicle(router.currentRoute.value.params.id.toString()), {withCredentials: true})
    vehicle.value = response.data
  } catch (error: any) {
    notifications.addNotification("Failed to get vehicle: " + error, "error")
  } finally {
    loading.value = false
  }
}

const loadVehiclesMalfunctions = async () => {
  try {
    loading.value = true
    // TODO

  } catch (error: any) {
    notifications.addNotification("Failed to get vehicle malfunctions: " + error, "error")
  } finally {
    loading.value = false
  }
}

const deleteVehicle = async (id: string) => {
  if (!window.confirm("Are you sure you want to delete this vehicle?")) {
    return;
  }

  try {
    const response = await axios.delete(Endpoints.deleteVehicle(id), {withCredentials: true})
    if (response.status === 200) {
      notifications.addNotification("Vehicle deleted", "success")
      await router.push("/profile/superuser/vehicles")
    }
  } catch (error) {
    notifications.addNotification("Failed to delete vehicle: " + error, "error")
  } finally {
  }
}

onMounted(() => {
  loadVehicle()
})

</script>

<template>
  <div>
    <Loader v-if="loading"/>

    <div v-else v-if="vehicle">
      <div class="header">
        <h2>{{vehicle.Registration}}</h2>
      </div>
      <div class="details">
        <div class="details-item">
          <p>Type:</p>
          <p>{{ vehicle.Type }}</p>
        </div>
        <div class="details-item">
          <p>Capacity:</p>
          <p>{{ vehicle.Capacity }}</p>
        </div>
        <div class="details-item" v-if="vehicle.Brand">
          <p>Brand:</p>
          <p>{{ vehicle.Brand }}</p>
        </div>
      </div>
      <div class="hr"></div>

      <div class="hr"></div>
      <div class="tools">
        <router-link :to='"/profile/superuser/vehicles/edit/" + vehicle.Registration'><Pencil :size="24" /></router-link>
        <a @click="deleteVehicle(vehicle.Registration)"><Delete :size="24" /></a>
      </div>
    </div>
  </div>
</template>

<style>

</style>