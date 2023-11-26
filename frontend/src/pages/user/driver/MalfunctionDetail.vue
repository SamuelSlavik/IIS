<script setup lang="ts">
import {onMounted, ref} from "vue";
import Loader from "@/components/Loader.vue";
import type {Malfunction, User} from "@/lib/models";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";
import router from "@/router";
import Delete from "vue-material-design-icons/Delete.vue";
import Pencil from "vue-material-design-icons/Pencil.vue";
import {formatDate} from "@/lib/utils";
import Bus from "vue-material-design-icons/Bus.vue";
import Tram from "vue-material-design-icons/Tram.vue";
import Tank from "vue-material-design-icons/Tank.vue";

const loading = ref<boolean>(false)
const notifications = useNotificationStore()

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
        <router-link :to='"/profile/malfunction/edit/" + report.ID'><Pencil :size="24" /></router-link>
        <a @click="deleteReport(report.ID)"><Delete :size="24" /></a>
      </div>
    </div>
  </div>
</template>

<style>

</style>