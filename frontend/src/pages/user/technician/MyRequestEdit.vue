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

const malfunctionId = ref("")

const malfunction = ref<Malfunction>()

const vehicles = ref<Vehicle[]>([])
const technicians = ref<User[]>([])

const newRequests = ref<NewRequest>({
  Status: null,
  Deadline: "",
  ResolvedByRef: {
    value: null,
    label: "",
  },
})

const submitRequest = async () => {
  try {
    const response = await axios.post(Endpoints.createRequest, {
      Status: newRequests.value.Status,
      Deadline: newRequests.value.Deadline,
      ResolvedByRef: newRequests.value.ResolvedByRef?.value || null,
      MalfuncRepRef: parseInt(malfunctionId.value),
    },{withCredentials: true})
    if (response.status === 200) {
      notifications.addNotification("Maintenance request created", 'success')
      await router.push('/profile/superuser/requests');
    }
  } catch (error) {
    notifications.addNotification("Failed to create maintenance requests: " + error, 'error')
  }
}

const loadRequest = async () => {
  try {
    loading.value = true
    const response = await axios.get(Endpoints.retrieveRequest(router.currentRoute.value.params.id.toString()), {withCredentials: true})
    newRequests.value.Status = response.data.Status
    newRequests.value.Deadline = response.data.Deadline.split('T')[0]
    newRequests.value.ResolvedByRef = {
      value: response.data.ResolvedBy?.ID || null,
      label: response.data.ResolvedBy?.FirstName + " " + response.data.ResolvedBy?.LastName || "",
    }
    malfunctionId.value = response.data.MalfuncRep.ID.toString()
    if (response.status === 200) {
      await loadMalfunction()
    }
  } catch (error: any) {
    notifications.addNotification("Failed to load request: " + error, "error")
  } finally {
    loading.value = false
  }

}

const loadMalfunction = async () => {
  loading.value = true
  try {
    const response = await axios.get<Malfunction>(Endpoints.retrieveMalfunction(malfunctionId.value), {withCredentials: true})
    malfunction.value = response.data
  } catch (error) {
    notifications.addNotification("Failed to load malfunction: " + error, 'error')
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
  loadUsers()
  loadRequest()
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
          <select v-model="newRequests.Status">
            <option value="pending">Pending</option>
            <option value="progress">In progress</option>
            <option value="done">Done</option>
          </select>
          <button
              type="submit"
          >Update request</button>
        </form>
      </div>
    </div>
  </div>
</template>

<style>

</style>