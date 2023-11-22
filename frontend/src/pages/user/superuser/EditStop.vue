<script setup lang="ts">
import {onMounted, ref} from "vue";
import Loader from "@/components/Loader.vue";
import type {NewStop, Stop, User} from "@/lib/models";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";
import router from "@/router";

const stopId = router.currentRoute.value.params.id.toString() || ""

const loading = ref<boolean>(false)
const notifications = useNotificationStore()

const stop = ref<Stop>()

const newStop = ref<NewStop>({
  Name: stop.value?.Name || "",
})

const loadStop = async () => {
  loading.value = true
  try {
     const response = await axios.get<Stop>(Endpoints.stopDetail(stopId), {withCredentials: true})
     newStop.value.Name = response.data.Name
  } catch (error) {
    notifications.addNotification("Failed to load stop: " + error, 'error')
  } finally {
    loading.value = false
  }
}

const updateStop = async () => {
  try {
    const response = await axios.put(Endpoints.editStop(stopId), newStop.value, {withCredentials: true})
    if (response.status === 200) {
      notifications.addNotification("Stop updated")
      await router.push('/profile/superuser/stops');
    }
  } catch (error) {
    notifications.addNotification("Failed to update stop: " + error, 'error')
  } finally {
  }
}

onMounted(() => {
  loadStop()
})

</script>

<template>
  <div>
    <div class="header"><h2>Edit stop</h2></div>
    <Loader v-if="loading"/>
    <form v-else @submit.prevent="updateStop" class="form">
      <input
          type="text"
          name="name"
          placeholder="Stop name"
          v-model="newStop.Name"
          required
      />
      <button
          type="submit"
      >Update Stop</button>
    </form>
  </div>
</template>

<style>

</style>