<script setup lang="ts">
import {onMounted, ref} from "vue";
import Loader from "@/components/Loader.vue";
import type {NewLine, User} from "@/lib/models";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";
import router from "@/router";
// @ts-ignore
import Delete from "vue-material-design-icons/Delete.vue";
// @ts-ignore
import Pencil from "vue-material-design-icons/Pencil.vue";

const loading = ref<boolean>(false)
const notifications = useNotificationStore()

const lineName = router.currentRoute.value.params.name.toString() || ""

const line = ref<NewLine>()

const loadLine = async () => {
  loading.value = true
  try {
    const response = await axios.get<NewLine>(Endpoints.retrieveLine(lineName), {withCredentials: true})
    line.value = response.data
  } catch (error) {
    notifications.addNotification("Failed to load line: " + error, 'error')
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
      await router.push("/profile/superuser/lines")
    }
  } catch (error: any) {
    notifications.addNotification("Failed to delete line: " + error, "error")
  }
}

onMounted(() => {
  loadLine()
})

</script>

<template>
  <div>
    <Loader v-if="loading"/>
    <div v-else v-if="line">
      <div class="header">
        <h2>{{line.Name}}</h2>
      </div>
      <div class="table">
        <div v-for="(stop, index) in line.StopsSequence" :key="stop.StopName">
          <div class="list-item">
            <p class="list-item__name">
              <b>{{ stop.StopName }}</b>
            </p>
            <p class="list-item__role">
              {{ stop.Duration }} minutes
            </p>
          </div>
          <div v-if="index < line.StopsSequence.length - 1" class="table-hr"></div>
        </div>
      </div>
      <div class="hr"></div>
      <div class="tools">
        <router-link :to="'/profile/superuser/lines/edit/' + line.Name"><Pencil :size="24" /></router-link>
        <a @click="deleteLine(line.Name)"><Delete :size="24" /></a>
      </div>
    </div>
  </div>
</template>

<style>

</style>