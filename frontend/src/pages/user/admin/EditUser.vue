<script setup lang="ts">
import {onMounted, ref} from "vue";
import Loader from "@/components/Loader.vue";
import type {User} from "@/lib/models";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";
import router from "@/router";

const loading = ref<boolean>(false)
const notifications = useNotificationStore()

const userId = router.currentRoute.value.params.id.toString() || ""

const newUser = ref<User>()

const loadUser = async () => {
  try {
    loading.value = true
    const response = await axios.get<User>(Endpoints.retrieveUser(userId), {withCredentials: true})
    newUser.value = response.data
    newUser.value.BirthDate = newUser.value.BirthDate.split('T')[0]
    console.log(newUser.value)
  } catch (error: any) {
    if (error.response.status != 401) {
      notifications.addNotification("Failed to get user: " + error, "error")
    }
  } finally {
    loading.value = false
  }
}

const editUser = async () => {
  try {
    loading.value = true
    const response = await axios.patch(Endpoints.updateUser(userId), newUser.value, {withCredentials: true})
    if (response.status === 200) {
      notifications.addNotification("User updated", "success")
      await router.push('/profile/users');
    }
  } catch (error: any) {
    notifications.addNotification("Failed to update user: " + error, "error")
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadUser()
})

</script>

<template>
  <div>
    <div class="header">
      <h2>Edit user</h2>
    </div>

    <Loader v-if="loading"/>
    <div v-else>
      <form @submit.prevent="editUser" class="form" v-if="newUser">
        <input
          type="text"
          placeholder="First name"
          v-model="newUser.FirstName"
          required
        />
        <input
            type="text"
            placeholder="Last name"
            v-model="newUser.LastName"
            required
        />
        <input
            type="email"
            placeholder="Email"
            v-model="newUser.Email"
            required
        />
        <input
            type="date"
            name="birth-date"
            placeholder="Date of birth"
            v-model="newUser.BirthDate"
        />
        <select v-model="newUser.Role" required>
          <option value="superuser">Superuser</option>
          <option value="technician">Technician</option>
          <option value="dispatcher">Dispatcher</option>
          <option value="driver">Driver</option>
        </select>
        <button
            type="submit"
        >Update user</button>
      </form>
    </div>
  </div>
</template>

<style>

</style>