<script setup lang="ts">

import Dashboard from "@/components/Dashboard.vue";
import {onMounted, ref} from "vue";
import axios from "axios";
import {Endpoints} from "@/lib/variables";
import {useNotificationStore} from "@/stores/notification-store";
import {useUserStore} from "@/stores/user-store";
import type {User} from "@/lib/models";
import Loader from "@/components/Loader.vue";

const notifications = useNotificationStore()
const user = useUserStore()
const loading = ref<boolean>(false)

const getUser = async () => {
  try {
    loading.value = true
    const response = await axios.get<User>(Endpoints.retrieveUser, {withCredentials: true})
    user.setUserData(response.data)
  } catch (error) {
    notifications.addNotification("Failed to get user: " + error, "error")
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  getUser()
})
</script>

<template>
  <div class="profile-wrapper">
    <Loader v-if="loading"/>
    <Dashboard v-else/>
    <div class="profile">
      <Loader v-if="loading"/>
      <div class="profile-content" v-else>
        <h1>Welcome!</h1>
      </div>
    </div>
  </div>
</template>

<style>
  .profile-wrapper {
    display: flex;
    min-height: 100vh;
  }
  .profile {
    width: calc((100% - (1280px - 64px)) / 2 + 930px - 32px);
  }
  .profile-content {
    width: 930px;
    margin-right: auto;
    margin-left: 0;
    padding: 2rem 1rem 2rem 2rem;
  }
</style>