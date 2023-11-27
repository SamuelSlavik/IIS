<script setup lang="ts">
import {onMounted, ref} from "vue";
import Loader from "@/components/Loader.vue";
import type {ConnectionList, User} from "@/lib/models";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";
import lines from "@/pages/user/superuser/Lines.vue";
import {useUserStore} from "@/stores/user-store";
import {formatDateTime} from "../../../lib/utils";
// @ts-ignore
import Bus from "vue-material-design-icons/Bus.vue";
// @ts-ignore
import Tram from "vue-material-design-icons/Tram.vue";
// @ts-ignore
import Tank from "vue-material-design-icons/Tank.vue";


const loading = ref<boolean>(false)
const notifications = useNotificationStore()

const drivers = ref<User[]>([])

const loadDrivers = async () => {
  loading.value = true
  try {
    const response = await axios.get<User[]>(Endpoints.listUsersByRole("driver"), {withCredentials: true})
    drivers.value = response.data
  } catch (error) {
    notifications.addNotification("Failed to load drivers: " + error, 'error')
  } finally {
    loading.value = false
  }
}

onMounted(() => {loadDrivers()})
</script>

<template>
  <div>
    <div class="header">
      <h2>Manage drivers plans</h2>
    </div>

    <Loader v-if="loading"/>
    <div v-else>
      <div class="table">
        <div v-for="(driver, index) in drivers" :key="driver.ID" v-if="drivers">
          <div class="list-item">
            <router-link :to="'/profile/admin/drivers/detail/' + driver.ID" class="list-item__name">
              <b>{{ driver.FirstName }} {{ driver.LastName }}</b>
            </router-link>
          </div>
          <div v-if="index < drivers.length - 1" class="table-hr"></div>
        </div>
      </div>
    </div>

  </div>
</template>