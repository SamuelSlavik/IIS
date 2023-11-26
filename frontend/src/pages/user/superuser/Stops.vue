<script setup lang="ts">
import {onMounted, ref} from "vue";
import Loader from "@/components/Loader.vue";
import type {Stop, User} from "@/lib/models";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";
import Delete from "vue-material-design-icons/Delete.vue";
import Pencil from "vue-material-design-icons/Pencil.vue";
import Magnify from "vue-material-design-icons/Magnify.vue";

const loading = ref<boolean>(false)
const notifications = useNotificationStore()

const query = ref<string>("")

const stops = ref<Stop[]>([])

const loadStops = async () => {
  loading.value = true
  try {
    const response = await axios.get(Endpoints.listStops(query.value), {withCredentials: true})
    stops.value = response.data
    console.log(stops.value)
  } catch (error) {
    notifications.addNotification("Failed to load stops: " + error, "error")
  } finally {
    loading.value = false
  }
}

const deleteStop = async (id: string) => {
  if (!window.confirm("Are you sure you want to delete this stop?")) {
    return;
  }

  try {
    const response = await axios.delete(Endpoints.deleteStop(id), {withCredentials: true})
    if (response.status === 200) {
      notifications.addNotification("Stop deleted")
      loadStops()
    }
  } catch (error) {
    notifications.addNotification("Failed to delete stop: " + error, "error")
  } finally {
  }
}

onMounted(() => {
  loadStops()
})

</script>

<template>
  <div>
    <div class="header"><h2>Manage stops</h2></div>

    <div class="toolbar">
      <form @submit.prevent="loadStops" class="search-form">
        <input
          type="text"
          name="query"
          placeholder="Search stops"
          v-model="query"
        />
        <button type="submit" class="small-button">
          <Magnify size="24px"/>
        </button>
      </form>
    </div>

    <Loader v-if="loading"/>
    <div v-else>
      <div class="table">
        <div v-for="(stop, index) in stops" :key="stop.ID">
          <div class="list-item">
            <router-link :to="'/profile/superuser/stops/edit/' + stop.ID" class="list-item__name">
              <b>{{ stop.Name }}</b>
            </router-link>
            <p v-if="stop.Active" class="list-item__role green">
              Active
            </p>
            <div class="list-item__tools">
              <router-link :to="'/profile/superuser/stops/edit/' + stop.ID"><Pencil :size="24" /></router-link>
              <a v-if="!stop.Active" @click="deleteStop(stop.ID)"><Delete :size="24" /></a>
            </div>
          </div>
          <div v-if="index < stops.length - 1" class="table-hr"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<style>
.search-form {
  display: flex;
  justify-content: space-between;
  gap: 0;
}
</style>