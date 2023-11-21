<script setup lang="ts">
import {onMounted, ref} from "vue";
import Loader from "@/components/Loader.vue";
import type {User, Vehicle} from "@/lib/models";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import Pencil from "vue-material-design-icons/Pencil.vue";
import Delete from "vue-material-design-icons/Delete.vue";
import {useNotificationStore} from "@/stores/notification-store";

const loading = ref<boolean>(false)
const notifications = useNotificationStore()

const vehicle = ref<Vehicle | undefined>(undefined)

const loadVehicle = async () => {
  try {
    loading.value = true
    // TODO

  } catch (error: any) {
    notifications.addNotification("Failed to get vehicle: " + error, "error")
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadVehicle()
})

</script>

<template>
  <div>
    <div class="header">
      Create new vehicle
    </div>
  </div>
</template>

<style>

</style>