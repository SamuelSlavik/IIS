<script setup lang="ts">
import {onMounted, ref} from "vue";
import Loader from "@/components/Loader.vue";
import type {AssignedConnection, ConnectionList, User} from "@/lib/models";
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
import type {Vehicle} from "@/lib/models";

const loading = ref<boolean>(false)
const notifications = useNotificationStore()
const connection = ref<ConnectionList>()
const connectionID = router.currentRoute.value.params.id.toString() || ""

const cetTime = new Date()
cetTime.setHours(cetTime.getHours() + 2)
const querydate = ref<string>(cetTime.toISOString().split("T")[0])

const drivers = ref<User[]>([])

const selectedDriver = ref('')

const assignedConnection = ref<AssignedConnection>({
  DriverID: {
    value: null,
    label: "",
  },
  VehicleReg: "",
  NumberOfDays: null,
})

const loadConnection = async () => {
  try {
    loading.value = true
    const response = await axios.get(Endpoints.retrieveConnection(connectionID), {withCredentials: true})
    connection.value = response.data
    assignedConnection.value.DriverID.label = response.data.DriverName
    assignedConnection.value.DriverID.value = response.data.DriverID
    assignedConnection.value.VehicleReg = response.data.VehicleReg
  } catch (error: any) {
    notifications.addNotification("Failed to load lines: " + error, "error")
  } finally {
    loading.value = false
  }
}

const loadDrivers = async () => {
  loading.value = true
  try {
    const response = await axios.get<User[]>(Endpoints.listUsersByRole("driver"), {withCredentials: true})
    drivers.value = response.data
  } catch (error) {
    notifications.addNotification("Failed to load drivers: " + error, 'error')
  } finally {
    loading.value = false
  }
}

const vehicles = ref<Vehicle[]>([])

const loadVehicles = async () => {
  loading.value = true
  try {
    const response = await axios.get<Vehicle[]>(Endpoints.listOkVehicles, {withCredentials: true})
    vehicles.value = response.data
  } catch (error) {
    notifications.addNotification("Failed to load vehicles: " + error, 'error')
  } finally {
    loading.value = false
  }
}

const assign = async () => {
  try {
    const response = await axios.patch(Endpoints.assignConnection(connectionID), {
      DriverID: assignedConnection.value.DriverID?.value || null,
      VehicleReg: assignedConnection.value.VehicleReg || null,
      NumberOfDays: assignedConnection.value.NumberOfDays,
    }, {withCredentials: true})
    if (response.status === 200) {
      notifications.addNotification("Connection assigned", 'success')
      await router.push('/profile/dispatcher/connections/' + connection.value?.LineName);
    }
  } catch (error) {
    notifications.addNotification("Failed to assign connection: " + error, 'error')
  } finally {
  }
}

onMounted(() => {
  loadConnection()
  loadDrivers()
  loadVehicles()
})
</script>

<template>
  <div>
    <div class="header">
      <h2>Assign user and vehicle to the connection</h2>
    </div>

    <Loader v-if="loading"/>
    <div v-else>
      <div class="details" v-if="connection">
        <div class="details-item">
          <p>Departure time:</p>
          <p>{{ connection.DepartureTime }}</p>
        </div>
        <div class="details-item">
          <p>Arrival time:</p>
          <p>{{ connection.ArrivalTime }}</p>
        </div>
        <div class="details-item">
          <p>From:</p>
          <p>{{ connection.InitialStop }}</p>
        </div>
        <div class="details-item">
          <p>To:</p>
          <p>{{ connection.FinalStop }}</p>
        </div>
      </div>
      <div class="hr"></div>
      <form @submit.prevent="assign" class="form">
        <v-select v-model="assignedConnection.DriverID" :options="drivers.map(({ ID, LastName, FirstName }) => ({ value: ID, label: FirstName + ' ' + LastName }))" placeholder="Select driver"></v-select>
        <v-select v-model="assignedConnection.VehicleReg" :options="vehicles.map(({Registration}) => {return Registration})" placeholder="Select vehicle"></v-select>
        <input
          type="number"
          name="number-of-days"
          placeholder="Number of days"
          v-model="assignedConnection.NumberOfDays"
          min="1"
          required
        />

        <button
            type="submit"
        >Assign</button>
      </form>
    </div>
  </div>
</template>

<style>
.details-item p:first-child {
  font-weight: bold;
}
</style>
