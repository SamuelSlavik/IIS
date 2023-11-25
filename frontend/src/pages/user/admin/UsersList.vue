<script setup lang="ts">
import {onMounted, ref} from "vue";
import Loader from "@/components/Loader.vue";
import type {User} from "@/lib/models";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import Pencil from "vue-material-design-icons/Pencil.vue";
import Delete from "vue-material-design-icons/Delete.vue";
import Tank from "vue-material-design-icons/Tank.vue";
import {useNotificationStore} from "@/stores/notification-store";
import Magnify from "vue-material-design-icons/Magnify.vue";


const loading = ref<boolean>(false)
const notifications = useNotificationStore()

const query = ref<string>("")

const users = ref<User[]>([])

const uniqueRoles = ref(['admin', 'superuser', 'technician', 'dispatcher', 'driver'])
const getUsersByRole = (role: string) => {
  return users.value.filter((user) => user.Role == role);
};

const loadUsers = async () => {
  loading.value = true
  try {
     const response = await axios.get<User[]>(Endpoints.listUsers(query.value), {withCredentials: true})
     users.value = response.data
  } catch (error) {
    notifications.addNotification("Failed to get users: " + error)
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
      notifications.addNotification("User deleted", "success")
      loadUsers()
    }
  } catch (error) {
    notifications.addNotification("Failed to delete user: " + error, "error")
  } finally {
  }
}

onMounted(() => {
  loadUsers()
})

</script>

<template>
  <div>
    <div class="header">
      <h2>Manage users</h2>
    </div>

    <div class="toolbar">
      <form @submit.prevent="loadUsers" class="search-form">
        <input
            type="text"
            name="query"
            placeholder="Search users"
            v-model="query"
        />
        <button type="submit" class="small-button">
          <Magnify size="24px"/>
        </button>
      </form>
    </div>

    <Loader v-if="loading" />
    <div v-else>
      <!-- Iterate over each role -->
      <div v-for="role in uniqueRoles" :key="role">
        <div v-if="getUsersByRole(role).length > 0" class="table">
          <div v-for="(user, index) in getUsersByRole(role)" :key="user.ID">
            <div class="list-item">
              <router-link :to="'/profile/users/detail/' + user.ID" class="list-item__name">
                <b>{{ user.FirstName + " " + user.LastName }}</b>
              </router-link>
              <p class="list-item__role">{{ user.Role }}</p>
              <div class="list-item__tools">
                <router-link :to="'/profile/users/edit/' + user.ID"><Pencil :size="24" /></router-link>
                <a @click="deleteUser(user.ID)"><Delete :size="24" /></a>
              </div>
            </div>
            <!-- Display table-hr only if it's not the last user for the current role -->
            <div v-if="index < getUsersByRole(role).length - 1" class="table-hr"></div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style>
.list-item {
  width: 100%;
  display: flex;
  justify-content: left;
}
.list-item *{
  margin-top: auto;
  margin-bottom: auto;
}

.list-item__name {
  flex: 2;
}
.list-item__role {
  flex: 1;
}
.list-item__tools {
  text-align: right;
  display: flex;
  justify-content: right;
  gap: 1rem;
  flex: 1;
}
.table-hr {
  background-color: rgba(60, 60, 67, 0.12);
  height: 1px;
  width: 100%;
  margin: 1rem 0;
}
.header {
  width: 100%;
  margin-bottom: 2rem;
}
</style>