<script setup lang="ts">
import {useUserStore} from "@/stores/user-store";
import {useRouter} from "vue-router";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";
import {onMounted, ref} from "vue";
import type {User, Vehicle} from "@/lib/models";
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


const user = useUserStore()
const router = useRouter();
let notifications = useNotificationStore();
const loading = ref<boolean>(false)

const vehicles = ref<Vehicle[]>([])
const technicians = ref<User[]>([])

const newRequests = ref({
  Status: "",
  Deadline: "",
  MalfuncRefRef: "",
  CreatedBy: "",
  ResolvedByRef: "",
})

const submitRequest = async () => {
  loading.value = true
  try {
    const response = await axios.post(Endpoints.createRequest, {withCredentials: true})
    newRequests.value = response.data
  } catch (error) {
    notifications.addNotification("Failed to create maintenance requests: " + error, 'error')
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

/*
const loadUsers = async () => {
  try {
    loading.value = true
    const response = await axios.get<User[]>(Endpoints.listUsers(query.value), {withCredentials: true})
    users.value = response.data
    console.log(users.value)
  } catch (error) {
    notifications.addNotification("Failed to get users: " + error, "error")
  } finally {
    loading.value = false
  }
}
*/

</script>

<template>
  <div>
    <div class="header">
      <h2>New maintenance request</h2>
    </div>

    <Loader v-if="loading"/>
    <div v-else>

    </div>
  </div>
</template>

<style>

</style>