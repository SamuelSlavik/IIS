<script setup lang="ts">
import {onMounted, ref} from "vue";
import Loader from "@/components/Loader.vue";
import type {Malfunction, User} from "@/lib/models";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";
import Delete from "vue-material-design-icons/Delete.vue";
import Pencil from "vue-material-design-icons/Pencil.vue";
import Close from "vue-material-design-icons/Close.vue";
import Check from "vue-material-design-icons/Check.vue";
import Bus from "vue-material-design-icons/Bus.vue";
import Tram from "vue-material-design-icons/Tram.vue";
import Tank from "vue-material-design-icons/Tank.vue";
import router from "@/router";

const loading = ref<boolean>(false)
const notifications = useNotificationStore()

const query = ref<string>("")

const ackMalfunctions = ref<Malfunction[]>([])
const unAckMalfunctions = ref<Malfunction[]>([])

const loadMalfunctions = async () => {
  loading.value = true
  try {
    const response = await axios.get(Endpoints.listMalfunctions("ack"), {withCredentials: true})
    ackMalfunctions.value = response.data

    const response2 = await axios.get(Endpoints.listMalfunctions("unack"), {withCredentials: true})
    unAckMalfunctions.value = response2.data
  } catch (error) {
    notifications.addNotification("Failed to load malfunctions: " + error, 'error')
  } finally {
    loading.value = false
  }
}

const deleteMalfunction = async (id: string) => {
  if (!window.confirm("Are you sure you want to delete this malfunction?")) {
    return;
  }

  try {
    const response = await axios.delete(Endpoints.deleteMalfunction(id), {withCredentials: true})
    if (response.status === 200) {
      notifications.addNotification("Malfunction deleted", "success")
      await loadMalfunctions()
    }
  } catch (error) {
    notifications.addNotification("Failed to delete malfunction: " + error, "error")
  } finally {
  }
}

onMounted(() => {
  loadMalfunctions()
})

</script>

<template>
  <div>
    <div class="header">
      <h2>My reported malfunctions</h2>
    </div>

    <Loader v-if="loading"/>
    <div v-else>
      <div v-if="unAckMalfunctions" class="table">
        <div v-for="(report, index) in unAckMalfunctions" :key="report.ID">
          <div class="list-item">
            <router-link :to="'/profile/malfunctions/detail/' + report.ID" class="list-item__name">
              <b>{{ report.Title }}</b>
            </router-link>
            <p class="list-item__role"><Close fill-color="#e74c3c"/></p>
            <p class="list-item__role connection-title">
              <Bus v-if="report.Vehicle.VehicleTypeName === 'bus'" class="connection-icon"/>
              <Tram v-if="report.Vehicle.VehicleTypeName === 'tram'" class="connection-icon"/>
              <Tank v-if="report.Vehicle.VehicleTypeName === 'obrnena_dodavka'" class="connection-icon"/>
              {{report.Vehicle.Registration}}
            </p>
            <div class="list-item__tools">
              <router-link :to="'/profile/malfunctions/edit/' + report.ID"><Pencil :size="24" /></router-link>
              <a @click="deleteMalfunction(report.ID)"><Delete :size="24" /></a>
            </div>
          </div>
          <!-- Display table-hr only if it's not the last user for the current role -->
          <div v-if="index < unAckMalfunctions.length - 1" class="table-hr"></div>
      </div>
      </div>

      <div v-if="ackMalfunctions" class="table">
        <div v-for="(report, index) in ackMalfunctions" :key="report.ID">
          <div class="list-item">
            <router-link :to="'/profile/malfunctions/detail/' + report.ID" class="list-item__name">
              <b>{{ report.Title }}</b>
            </router-link>
            <p class="list-item__role"><Check fill-color="#2ecc71"/></p>
            <p class="list-item__role connection-title">
              <Bus v-if="report.Vehicle.VehicleTypeName === 'bus'" class="connection-icon"/>
              <Tram v-if="report.Vehicle.VehicleTypeName === 'tram'" class="connection-icon"/>
              <Tank v-if="report.Vehicle.VehicleTypeName === 'obrnena_dodavka'" class="connection-icon"/>
              {{report.Vehicle.Registration}}
            </p>
            <div class="list-item__tools">
              <a @click="deleteMalfunction(report.ID)"><Delete :size="24" /></a>
            </div>
          </div>
          <!-- Display table-hr only if it's not the last user for the current role -->
          <div v-if="index < unAckMalfunctions.length - 1" class="table-hr"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<style>

</style>