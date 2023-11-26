<script setup lang="ts">
import {useUserStore} from "@/stores/user-store";
import {useRouter} from "vue-router";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";
import {onMounted, ref} from "vue";
import type {Vehicle} from "@/lib/models";
import Bus from "vue-material-design-icons/Bus.vue";
import Tram from "vue-material-design-icons/Tram.vue";
import Tank from "vue-material-design-icons/Tank.vue";
import Pencil from "vue-material-design-icons/Pencil.vue";
import Delete from "vue-material-design-icons/Delete.vue";
import Check from "vue-material-design-icons/Check.vue";
import Close from 'vue-material-design-icons/Close.vue';


const user = useUserStore()
const router = useRouter();
let notifications = useNotificationStore();
const loading = ref<boolean>(false)

const vehicles = ref<Vehicle[]>([])

const uniqueVehicles = ref(['tram', 'bus', 'obrnena_dodavka'])
const getVehiclesByType = (type: string) => {
  console.log(type)
  console.log(vehicles)
  return vehicles.value.filter((vehicle) => vehicle.VehicleTypeName === type);
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
          <div v-for="(vehicle, index) in getVehiclesByType(type)" :key="vehicle.ID">
            <div class="list-item">
              <router-link :to="'/profile/superuser/vehicles/detail/' + vehicle.ID" class="list-item__name">
                <Bus v-if="vehicle.VehicleTypeName === 'bus'" class="connection-icon"/>
                <Tram v-if="vehicle.VehicleTypeName === 'tram'" class="connection-icon"/>
                <Tank v-if="vehicle.VehicleTypeName === 'obrnena_dodavka'" class="connection-icon"/>
                {{vehicle.Registration}}
              </router-link>
              <p class="list-item__role" v-if="vehicle.Malfunctions"><Close fill-color="#e74c3c"/></p>
              <p class="list-item__role" v-else><Check fill-color="#2ecc71"/></p>
              <p class="list-item__role">{{ vehicle.Capacity }}</p>
              <div class="list-item__tools">
                <router-link to="/profile/superuser/vehicles/edit"><Pencil :size="24" /></router-link>
                <a @click="deleteVehicle(vehicle.ID)"><Delete :size="24" /></a>
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