<script setup lang="ts">
import {useUserStore} from "@/stores/user-store";
import {useRouter} from "vue-router";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";
import {onMounted, ref} from "vue";
import type {Vehicle, VehicleInList} from "@/lib/models";
import Bus from "vue-material-design-icons/Bus.vue";
import Tram from "vue-material-design-icons/Tram.vue";
import Tank from "vue-material-design-icons/Tank.vue";
import Pencil from "vue-material-design-icons/Pencil.vue";
import Delete from "vue-material-design-icons/Delete.vue";
import Check from "vue-material-design-icons/Check.vue";
import Close from 'vue-material-design-icons/Close.vue';
import {formatDate} from "../../../lib/utils";


const user = useUserStore()
const router = useRouter();
let notifications = useNotificationStore();
const loading = ref<boolean>(false)

const vehicles = ref<VehicleInList[]>([])

const uniqueVehicles = ref(['tram', 'bus', 'obrnena_dodavka'])
const getVehiclesByType = (type: string) => {
  console.log(type)
  console.log(vehicles)
  return vehicles.value.filter((vehicle) => vehicle.Type === type);
};


const loadVehicles = async () => {
  try {
    loading.value = true
    const response = await axios.get(Endpoints.listVehicles, {withCredentials: true})
    vehicles.value = response.data
  } catch (error: any) {
    notifications.addNotification("Failed to load vehicles: " + error, "error")
  } finally {
    loading.value = false
  }
}

const deleteVehicle = async (id: string) => {

}

onMounted(() => {
  loadVehicles()
})

</script>

<template>
  <div>
    <div class="header">
      <h2>Manage vehicles</h2>
    </div>

    <Loading v-if="loading"/>
    <div v-else>
      <div v-for="type in uniqueVehicles" :key="type">
        <div v-if="getVehiclesByType(type).length > 0" class="table">
          <div v-for="(vehicle, index) in getVehiclesByType(type)" :key="vehicle.Registration">
            <div class="list-item">
              <router-link :to="'/profile/superuser/vehicles/detail/' + vehicle.Registration" class="list-item__name">
                <Bus v-if="vehicle.Type === 'bus'" class="connection-icon"/>
                <Tram v-if="vehicle.Type === 'tram'" class="connection-icon"/>
                <Tank v-if="vehicle.Type === 'obrnena_dodavka'" class="connection-icon"/>
                {{vehicle.Registration}}
              </router-link>
              <p v-if="vehicle.LastMaintenance.Date != '-'" class="list-item__role" :class="{ 'yellow': vehicle.LastMaintenance.Status === 'pending' || vehicle.LastMaintenance.Status === 'progress' }">
                Maintenance: {{ formatDate(vehicle.LastMaintenance.Date)  }}
              </p>
              <p class="list-item__role" :class="{ 'yellow': vehicle.LastMaintenance.Status === 'pending' || vehicle.LastMaintenance.Status === 'progress' }">
                Maintenance: {{ vehicle.LastMaintenance.Date  }}
              </p>
              <div class="list-item__tools">
                <router-link :to="'/profile/superuser/vehicles/edit/' + vehicle.Registration"><Pencil :size="24" /></router-link>
                <a @click="deleteVehicle(vehicle.Registration)"><Delete :size="24" /></a>
              </div>
            </div>
            <div v-if="index < getVehiclesByType(type).length - 1" class="table-hr"></div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style>
.table {
  width: 100%;
  background-color: rgb(246, 246, 247);
  border-radius: 5px;
  margin: 2rem 0;
  padding: 1rem;
  position: relative;
}
</style>
<style scoped>
.list-item__name {
  display: flex;
  gap: 1rem;
  justify-content: left;
}
.list-item__name * {
  margin-top: auto;
  margin-bottom: auto;
}
</style>