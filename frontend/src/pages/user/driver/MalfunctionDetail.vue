<script setup lang="ts">
import {onMounted, ref} from "vue";
import Loader from "@/components/Loader.vue";
import type {Malfunction, User} from "@/lib/models";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";
import router from "@/router";
// @ts-ignore
import Delete from "vue-material-design-icons/Delete.vue";
// @ts-ignore
import Pencil from "vue-material-design-icons/Pencil.vue";
import {formatDate} from "@/lib/utils";
// @ts-ignore
import Bus from "vue-material-design-icons/Bus.vue";
// @ts-ignore
import Tram from "vue-material-design-icons/Tram.vue";
// @ts-ignore
import Tank from "vue-material-design-icons/Tank.vue";
import {useUserStore} from "@/stores/user-store";
// @ts-ignore
import Hammer from "vue-material-design-icons/HammerSickle.vue";

const loading = ref<boolean>(false)
const notifications = useNotificationStore()

const user = useUserStore()

const reportId = router.currentRoute.value.params.id.toString() || ""

const report = ref<Malfunction>()

const loadMalfunction = async () => {
  loading.value = true
  try {
    const response = await axios.get<Malfunction>(Endpoints.retrieveMalfunction(reportId), {withCredentials: true})
    report.value = response.data
  } catch (error) {
    notifications.addNotification("Failed to load malfunction: " + error, 'error')
  } finally {
    loading.value = false
  }
}

const deleteReport = async (id: string) => {
  if (!window.confirm("Are you sure you want to delete this malfunction?")) {
    return;
  }

  try {
    const response = await axios.delete(Endpoints.deleteMalfunction(id), {withCredentials: true})
    if (response.status === 200) {
      notifications.addNotification("Malfunction deleted", "success")
      user.role === "driver" ? await router.push('/profile/driver/reports') : await router.push('/profile/superuser/malfunctions');
    }
  } catch (error) {
    notifications.addNotification("Failed to delete malfunction: " + error, "error")
  } finally {
  }
}

onMounted(() => {
  loadMalfunction()
})

</script>

<template>
  <div>
    <Loader v-if="loading"/>
    <div v-else v-if="report">
      <div class="header">
        <h2>{{report.Title}}</h2>
      </div>

      <div class="details-item">
        <p>Status:</p>
        <p v-if="report.Acknowledged" class="green">Acknowledged</p>
        <p v-else class="red">Unacknowledged</p>
      </div>

      <br/>
      <div class="details-item">
        <p>Created by:</p>
        <p>{{ report.CreatedBy.FirstName + " " + report.CreatedBy.LastName}} <br/> {{report.CreatedBy.Email}}</p>
      </div>
      <br/>
      <div class="details-item">
        <p>Created at:</p>
        <p>{{ formatDate(report.CreatedAt) }}</p>
      </div>
      <br/>
      <div class="details-item">
        <p>Vehicle:</p>
        <p class="connection-title">
          <Bus v-if="report.Vehicle.VehicleTypeName === 'bus'" class="connection-icon"/>
          <Tram v-if="report.Vehicle.VehicleTypeName === 'tram'" class="connection-icon"/>
          <Tank v-if="report.Vehicle.VehicleTypeName === 'obrnena_dodavka'" class="connection-icon"/>
          {{ report.Vehicle.Registration }}
        </p>
      </div>
      <br/>
      <p>Description:</p>
      <p>{{report.Description}}</p>

      <div class="hr"></div>
      <div class="tools">
        <router-link v-if="user.Role != 'driver'" :to="'/profile/superuser/requests/create/' + report.ID"><Hammer :size="24" /></router-link>
        <router-link v-if="!report.Acknowledged" :to='"/profile/malfunctions/edit/" + report.ID'><Pencil :size="24" /></router-link>
        <a @click="deleteReport(report.ID)"><Delete :size="24" /></a>
      </div>
    </div>
  </div>
</template>

<style>

</style>