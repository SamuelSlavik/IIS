<script setup lang="ts">
import {useUserStore} from "@/stores/user-store";
import {useRouter} from "vue-router";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";
import {ref} from "vue";
import type {NewVehicle} from "@/lib/models";

const user = useUserStore()
const router = useRouter();
let notifications = useNotificationStore();

const newVehicle = ref<NewVehicle>({
  Registration: "",
  Capacity: null,
  Type: "",
  Brand: "",
})

const submitVehicle = async () => {
  try {
    const response = await axios.post(Endpoints.createVehicle, newVehicle.value, {withCredentials: true})
    if (response.status === 200) {
      notifications.addNotification("Vehicle created", 'success')
      await router.push('/profile/superuser/vehicles');
    }
  } catch (error: any) {
    notifications.addNotification("Failed to create vehicle: " + error, "error")
  }
}

</script>

<template>
  <div>
    <div class="header">
      <h2>Create new vehicle</h2>
    </div>

    <div>
      <form @submit.prevent="submitVehicle" class="form">
        <input
          type="text"
          name="registration"
          placeholder="Vehicle registration: XXX0000"
          v-model="newVehicle.Registration"
          required
        />
        <input
            type="number"
            name="capacity"
            placeholder="Capacity"
            v-model="newVehicle.Capacity"
            required
        />
        <input
            type="text"
            name="brand"
            placeholder="Brand"
            v-model="newVehicle.Brand"
        />
        <select v-model="newVehicle.Type" required>
          <option value="" disabled>Vehicle type</option>
          <option value="bus">Bus</option>
          <option value="tram">Tram</option>
          <option value="obrnena_dodavka">Obrnena dodavka</option>
        </select>
        <button
            type="submit"
        >Create vehicle</button>
      </form>
    </div>
  </div>
</template>

<style>

</style>