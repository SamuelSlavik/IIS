<script setup lang="ts">
import {onMounted, ref} from "vue";
import Loader from "@/components/Loader.vue";
import type {NewStop, User} from "@/lib/models";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";
import router from "@/router";

const loading = ref<boolean>(false)
const notifications = useNotificationStore()

const newStop = ref<NewStop>({
  Name: "",
})

const createStop = async () => {
  try {
    const response = await axios.post(Endpoints.createStop, newStop.value, {withCredentials: true})
    if (response.status === 200) {
      notifications.addNotification("Stop created")
      await router.push('/profile/superuser/stops');
    }
  } catch (error) {
    notifications.addNotification("Failed to create stop: " + error, 'error')
  } finally {
  }
}

</script>

<template>
  <div>
    <div class="header"><h2>Create new stop</h2></div>
    <Loader v-if="loading"/>
    <form v-else @submit.prevent="createStop" class="form">
      <input
          type="text"
          name="name"
          placeholder="Stop name"
          v-model="newStop.Name"
          required
      />
      <button
          type="submit"
      >Create new stop</button>
    </form>
  </div>
</template>

<style>

</style>