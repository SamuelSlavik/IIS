<script setup lang="ts">
import {useUserStore} from "@/stores/user-store";
import {useRouter} from "vue-router";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";
import {onMounted, ref} from "vue";
import type {Vehicle} from "@/lib/models";
import Bus from "vue-material-design-icons/Bus.vue";
import Tram from "vue-material-design-icons/Tram.vue";
import Tank from "vue-material-design-icons/Tank.vue";
import Pencil from "vue-material-design-icons/Pencil.vue";
import Delete from "vue-material-design-icons/Delete.vue";
import Check from "vue-material-design-icons/Check.vue";
import Close from 'vue-material-design-icons/Close.vue';


const user = useUserStore()
const router = useRouter();
let notifications = useNotificationStore();
const loading = ref<boolean>(false)

const requests = ref()

const loadRequests = async () => {
  loading.value = true
  try {
    const response = await axios.get(Endpoints.listRequests, {withCredentials: true})
    requests.value = response.data
  } catch (error) {
    notifications.addNotification("Failed to load maintenance requests: " + error, 'error')
  } finally {
    loading.value = false
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

  </div>
</div>
</template>

<style>

</style>