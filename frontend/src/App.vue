<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import Navigation from './components/Navigation.vue'
import Notification from "@/components/Notification.vue";
import axios from "axios";
import type {User} from "@/lib/models";
import {Endpoints} from "@/lib/variables";
import {onMounted, ref} from "vue";
import {useUserStore} from "@/stores/user-store";
import {useNotificationStore} from "@/stores/notification-store";

const loading = ref<boolean>(false)
const user = useUserStore()
const notifications = useNotificationStore()


const getUser = async () => {
  try {
    loading.value = true
    const response = await axios.get<User>(Endpoints.retrieveCurrentUser, {withCredentials: true})
    user.setUserData(response.data)
  } catch (error: any) {
    if (error.response.status != 401) {
      notifications.addNotification("Failed to get user: " + error, "error")
    }
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  getUser()
})
</script>

<template>
  <div>
    <Navigation />
  </div>

  <RouterView />

  <Notification />
</template>

<style>
.router-link-active {
  color: #00bd7e;
}
.red {
  color: #e74c3c;
}
@media (min-width: 1024px) {

}
</style>
