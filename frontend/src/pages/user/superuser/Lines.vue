<script setup lang="ts">
import {onMounted, ref} from "vue";
import Loader from "@/components/Loader.vue";
import type {LineInList, User} from "@/lib/models";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";
import Delete from "vue-material-design-icons/Delete.vue";
import Pencil from "vue-material-design-icons/Pencil.vue";

const loading = ref<boolean>(false)
const notifications = useNotificationStore()

const lines = ref<LineInList[]>([])

const loadLines = async () => {
  try {
    loading.value = true
    const response = await axios.get(Endpoints.listLines, {withCredentials: true})
    lines.value = response.data
  } catch (error: any) {
    notifications.addNotification("Failed to load lines: " + error, "error")
  } finally {
    loading.value = false
  }
}

const deleteLine = async (name: string) => {
  if (!window.confirm("Are you sure you want to delete this line?")) {
    return;
  }
  try {
    const response = await axios.delete(Endpoints.deleteLine(name), {withCredentials: true})
    if (response.status === 200) {
      notifications.addNotification("Line deleted", "success")
      loadLines()
    }
  } catch (error: any) {
    notifications.addNotification("Failed to delete line: " + error, "error")
  }
}

onMounted(() => {
  loadLines()
})

</script>

<template>
  <div>
    <div class="header"><h2>Manage lines</h2></div>

    <Loader v-if="loading"/>
    <div v-else>
      <div class="table">
        <div v-for="(line, index) in lines" :key="line.Name">
          <div class="list-item">
            <router-link :to="'/profile/superuser/lines/detail/' + line.Name" class="list-item__name">
              <b>{{ line.Name }}</b>
            </router-link>
            <p class="list-item__role">{{ line.InitialStop }}</p>
            <p class="list-item__role">{{line.FinalStop}}</p>
            <div class="list-item__tools">
              <router-link :to="'/profile/superuser/lines/edit/' + line.Name"><Pencil :size="24" /></router-link>
              <a @click="deleteLine(line.Name)"><Delete :size="24" /></a>
            </div>
          </div>
          <!-- Display table-hr only if it's not the last user for the current role -->
          <div v-if="index < lines.length - 1" class="table-hr"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<style>

</style>