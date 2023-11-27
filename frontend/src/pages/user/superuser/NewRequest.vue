<script setup lang="ts">
import {useUserStore} from "@/stores/user-store";
import {useRouter} from "vue-router";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";
import {onMounted, ref} from "vue";
import type {NewRequest, User, Vehicle} from "@/lib/models";
// @ts-ignore
import Bus from "vue-material-design-icons/Bus.vue";
// @ts-ignore
import Tram from "vue-material-design-icons/Tram.vue";
// @ts-ignore
import Tank from "vue-material-design-icons/Tank.vue";
// @ts-ignore
import Pencil from "vue-material-design-icons/Pencil.vue";
// @ts-ignore
import Delete from "vue-material-design-icons/Delete.vue";
// @ts-ignore
import Check from "vue-material-design-icons/Check.vue";
// @ts-ignore
import Close from 'vue-material-design-icons/Close.vue';
import type {Malfunction, MalfunctionReport} from "@/lib/models";
import {formatDate} from "@/lib/utils";
// @ts-ignore
import Hammer from "vue-material-design-icons/HammerSickle.vue";


const user = useUserStore()
const router = useRouter();
let notifications = useNotificationStore();
const loading = ref<boolean>(false)

const malfunctionId = router.currentRoute.value.params.id.toString() || ""

const malfunction = ref<Malfunction>()

const vehicles = ref<Vehicle[]>([])
const technicians = ref<User[]>([])

const newRequests = ref<NewRequest>({
  Deadline: "",
  ResolvedByRef: {
    value: null,
    label: "",
  },
})

const submitRequest = async () => {
  try {
    const response = await axios.post(Endpoints.createRequest, {
      Deadline: newRequests.value.Deadline,
      ResolvedByRef: newRequests.value.ResolvedByRef?.value || null,
      MalfuncRepRef: parseInt(malfunctionId),
    },{withCredentials: true})
    if (response.status === 200) {
      notifications.addNotification("Maintenance request created", 'success')
      await router.push('/profile/superuser/requests');
    }
  } catch (error) {
    notifications.addNotification("Failed to create maintenance requests: " + error, 'error')
  }
}

const loadMalfunction = async () => {
  loading.value = true
  try {
    const response = await axios.get<Malfunction>(Endpoints.retrieveMalfunction(malfunctionId), {withCredentials: true})
    malfunction.value = response.data
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

const loadUsers = async () => {
  try {
    loading.value = true
    const response = await axios.get<User[]>(Endpoints.listUsersByRole("technician"), {withCredentials: true})
    technicians.value = response.data
  } catch (error) {
    notifications.addNotification("Failed to get users: " + error, "error")
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadVehicles()
  loadUsers()
  loadMalfunction()
})


</script>

<template>
  <div>
    <div class="header">
      <h2>New maintenance request</h2>
    </div>

    <Loader v-if="loading"/>
    <div v-else>
      <div v-if="malfunction">
        <div class="header">
          <h2>{{malfunction.Title}}</h2>
        </div>

        <div class="details-item">
          <p>Status:</p>
          <p v-if="malfunction.Acknowledged" class="green">Acknowledged</p>
          <p v-else class="red">Unacknowledged</p>
        </div>

        <br/>
        <div class="details-item">
          <p>Created by:</p>
          <p>{{ malfunction.CreatedBy.FirstName + " " + malfunction.CreatedBy.LastName}} <br/> {{malfunction.CreatedBy.Email}}</p>
        </div>
        <br/>
        <div class="details-item">
          <p>Created at:</p>
          <p>{{ formatDate(malfunction.CreatedAt) }}</p>
        </div>
        <br/>
        <div class="details-item">
          <p>Vehicle:</p>
          <p class="connection-title">
            <Bus v-if="malfunction.Vehicle.VehicleTypeName === 'bus'" class="connection-icon"/>
            <Tram v-if="malfunction.Vehicle.VehicleTypeName === 'tram'" class="connection-icon"/>
            <Tank v-if="malfunction.Vehicle.VehicleTypeName === 'obrnena_dodavka'" class="connection-icon"/>
            {{ malfunction.Vehicle.Registration }}
          </p>
        </div>
        <br/>
        <p>Description:</p>
        <p>{{ malfunction.Description }}</p>

        <div class="hr"></div>

        <form @submit.prevent="submitRequest" class="form">
          <input
            type="date"
            name="deadline"
            v-model="newRequests.Deadline"
            required
          />
          <v-select v-model="newRequests.ResolvedByRef" :options="technicians.map(({ ID, LastName, FirstName }) => ({ value: ID, label: FirstName + ' ' + LastName }))" placeholder="Select driver"></v-select>
          <button
              type="submit"
          >Create request</button>
        </form>
      </div>
    </div>
  </div>
</template>

<style>

</style>