<script setup lang="ts">
import {onMounted, ref} from "vue";
import Loader from "@/components/Loader.vue";
import type {NewLine, User} from "@/lib/models";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";
import type {Stop} from "@/lib/models";
import router from "@/router";

const loading = ref<boolean>(false)
const notifications = useNotificationStore()

const stops = ref<Stop[]>([])

const newLine = ref<NewLine>({
  Name: "",
  StopsSequence: [{
    StopName: "",
    Duration: null,
  }],
})

const createLine = async () => {
  console.log(newLine.value)
  loading.value = true
  try {
    const response = await axios.post(Endpoints.createLine, newLine.value, {withCredentials: true})
    if (response.status === 200) {
      notifications.addNotification("Line created", 'success')
      await router.push('/profile/superuser/lines');
    }
  } catch (error) {
    notifications.addNotification("Failed to create line: " + error, 'error')
  } finally {
    loading.value = false
  }
}



const loadStops = async () => {
  loading.value = true
  try {
    const response = await axios.get(Endpoints.listStops(""), {withCredentials: true})
    stops.value = response.data
    console.log(stops.value)
  } catch (error) {
    notifications.addNotification("Failed to load stops: " + error, "error")
  } finally {
    loading.value = false
  }
}

const addStop = () => {
  newLine.value.StopsSequence.push({
    StopName: "",
    Duration: null,
  })
}
const removeStop = (index: number) => {
  newLine.value.StopsSequence.splice(index, 1)
}

onMounted(() => {
  loadStops()
})

</script>

<template>
  <div>
    <div class="header">
      <h2>Create new line</h2>
    </div>

    <Loader v-if="loading"/>
    <div v-else>
      <form @submit.prevent="createLine" class="form">
        <input
          type="text"
          placeholder="Line name"
          v-model="newLine.Name"
          required
        />

        <div v-for="(stop, index) in newLine.StopsSequence" :key="index">
          <label>Stop {{ index + 1 }}:</label>
          <v-select
              v-model="stop.StopName"
              placeholder="Select stop"
              :options="stops.map(({Name}) => {return Name})"
              required
          >
          </v-select>
          <input
              type="number"
              v-model="stop.Duration"
              placeholder="Duration"
              required
          />
          <button type="button" @click="removeStop(index)">Remove Stop</button>
        </div>

        <button type="button" @click="addStop">Add Stop</button>

        <button
            type="submit"
        >Create new line</button>
      </form>
    </div>
  </div>
</template>

<style>

</style>