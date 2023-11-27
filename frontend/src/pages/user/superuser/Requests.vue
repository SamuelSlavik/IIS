<script setup lang="ts">
import {useUserStore} from "@/stores/user-store";
import {useRouter} from "vue-router";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";
import {onMounted, ref} from "vue";
import type {RequestType, Vehicle} from "@/lib/models";
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
// @ts-ignore
import Hammer from "vue-material-design-icons/HammerSickle.vue";
import {formatDate} from "../../../lib/utils";
import Magnify from "vue-material-design-icons/Magnify.vue";


const user = useUserStore()
const router = useRouter();
let notifications = useNotificationStore();
const loading = ref<boolean>(false)

const statusQuery = ref<string>("")

const requests = ref<RequestType[]>()

const loadRequests = async () => {
  loading.value = true
  try {
    const response = await axios.get(Endpoints.listRequests(statusQuery.value), {withCredentials: true})
    requests.value = response.data
  } catch (error) {
    notifications.addNotification("Failed to load maintenance requests: " + error, 'error')
  } finally {
    loading.value = false
  }
}

const deleteRequest = async (id: string) => {
  if (!window.confirm("Are you sure you want to delete this maintenance request?")) {
    return;
  }

  try {
    const response = await axios.delete(Endpoints.deleteRequest(id), {withCredentials: true})
    if (response.status === 200) {
      notifications.addNotification("Maintenance request deleted", "success")
      await loadRequests()
    }
  } catch (error) {
    notifications.addNotification("Failed to delete maintenance request: " + error, "error")
  } finally {
  }
}

onMounted(() => {
  loadRequests()
})

</script>

<template>
<div>
  <div class="header">
    <h2>Manage maintenance requests</h2>
  </div>

  <Loader v-if="loading"/>
  <div v-else>
    <div class="toolbar">
      <form @submit.prevent="loadRequests" class="search-form">
        <select v-model="statusQuery">
          <option value="">All</option>
          <option value="pending">Pending</option>
          <option value="progress">In progress</option>
          <option value="done">Done</option>
        </select>
        <button type="submit" class="small-button">
          <Magnify size="24px"/>
        </button>
      </form>
    </div>

    <div class="table" v-if="requests">
      <div v-for="(request, index) in requests" :key="request.ID">
        <div class="list-item">
          <router-link :to="'/profile/superuser/requests/detail/' + request.ID" class="list-item__name">
            <b>{{ request.MalfuncRep.Title }}</b>
          </router-link>
          <p class="list-item__role">{{ formatDate(request.Deadline) }}</p>
          <p class="list-item__role yellow" v-if="request.Status === 'pending'">Pending</p>
          <p class="list-item__role yellow" v-if="request.Status === 'progress'">In progress</p>
          <p class="list-item__role green" v-if="request.Status === 'done'">Done</p>
          <p class="list-item__role connection-title">
            {{request.MalfuncRep.VehicleRef}}
          </p>
          <div class="list-item__tools">
            <router-link :to="'/profile/superuser/requests/edit/' + request.ID"><Pencil :size="24" /></router-link>
            <a @click="deleteRequest(request.ID)"><Delete :size="24" /></a>
          </div>
        </div>
        <!-- Display table-hr only if it's not the last user for the current role -->
        <div v-if="index < requests.length - 1" class="table-hr"></div>
      </div>
    </div>
  </div>
</div>
</template>

<style>

</style>