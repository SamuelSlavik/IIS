<script setup lang="ts">
import {useUserStore} from "@/stores/user-store";
import {useRouter} from "vue-router";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";
import {onMounted, ref} from "vue";
import type {User} from "@/lib/models";
import {Roles} from "@/lib/models";
import Loader from "@/components/Loader.vue";
import {formatDate} from "../../../lib/utils";
// @ts-ignore
import Delete from "vue-material-design-icons/Delete.vue";
// @ts-ignore
import Pencil from "vue-material-design-icons/Pencil.vue";

const user = useUserStore()
const router = useRouter();
const notifications = useNotificationStore();
const loading = ref<boolean>(false)

const userId = router.currentRoute.value.params.id.toString() || ""
const userData = ref<User | undefined >(undefined)

const loadUser = async () => {
  try {
    loading.value = true
    const response = await axios.get(Endpoints.retrieveUser(userId), {withCredentials: true})
    userData.value = response.data

  } catch (error) {
    notifications.addNotification("Failed to load user: " + error, "error")
  } finally {
    loading.value = false
  }
}

const deleteUser = async(id: string) => {
  if (!window.confirm("Are you sure you want to delete this user?")) {
    return;
  }

  try {
    const response = await axios.delete(Endpoints.deleteUser(id), {withCredentials: true})
    if (response.status === 200) {
      notifications.addNotification("User deleted")
      router.push("/profile/users")
    }
  } catch (error) {
    notifications.addNotification("Failed to delete user: " + error, "error")
  } finally {
  }
}

onMounted(() => {
  loadUser()
})

</script>

<template>
  <div>
    <Loader v-if="loading" />
    <div v-else v-if="userData">
      <div class="header">
        <h2>{{ userData.FirstName + " " + userData.LastName }}</h2>
      </div>
      <div class="details">
        <div class="details-item">
          <p>Role:</p>
          <p>{{ userData.Role }}</p>
        </div>
        <br/>
        <div class="details-item">
          <p>Email:</p>
          <p>{{ userData.Email }}</p>
        </div>
        <div class="details-item">
          <p>Date of birth:</p>
          <p>{{ formatDate(userData.BirthDate) }}</p>
        </div>
      </div>
      <div class="hr"></div>
      <div class="tools">
        <router-link :to='"/profile/users/edit/" + userData.ID'><Pencil :size="24" /></router-link>
        <a @click="deleteUser(userData.ID)"><Delete :size="24" /></a>
      </div>
    </div>
  </div>
</template>

<style>
.tools {
  display: flex;
  justify-content: left;
  gap: 1rem;
}
.details {
  width: 100%;
}

.details-item {
  display: flex;
  flex-direction: row;
  justify-content: left;
  gap: 1rem;
}

.details-item p:first-child {
  width: 200px;
  margin: 0;
}
</style>
