<script setup lang="ts">
import {useUserStore} from "@/stores/user-store";
import {useRouter} from "vue-router";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";

const user = useUserStore()
const router = useRouter();
let notifications = useNotificationStore();

const logOut = async () => {
  try {
    const response = await axios.get(Endpoints.logout, {withCredentials: true})
    if (response.status === 200) {
      user.logOut()
      await router.push('/')
    }
  } catch (error: any) {
    notifications.addNotification("Failed to logout: " + error, "error")
  } finally {

  }
}

</script>

<template>
  <div class="dashboard">
    <div class="dashboard-content">
      <div>
        <h2>{{ user.firstName + " " + user.lastName }}</h2>
        <p>{{user.email}}</p>
        <br/>
        <p>{{user.role}}</p>
      </div>
      <div class="hr"></div>
      <div>
        <p v-if="user.role === 'admin'">
          <router-link to="/profile/users">Manage users</router-link>
        </p>
        <p v-if="user.role === 'admin'">
          <router-link exact to="/profile/users/new">Create new user</router-link>
        </p>
        <br/>
        <p v-if="user.role === 'admin' || user.role === 'superuser'"><router-link to="/profile/superuser/lines">Manage lines</router-link></p>
        <p v-if="user.role === 'admin' || user.role === 'superuser'"><router-link to="/profile/superuser/lines/new">Create new line</router-link></p>
        <p v-if="user.role === 'admin' || user.role === 'superuser'"><router-link to="/profile/superuser/connections">Manage connections</router-link></p>
        <p v-if="user.role === 'admin' || user.role === 'superuser'"><router-link to="/profile/superuser/connections/new">Create new connection</router-link></p>
        <br/>
        <p v-if="user.role === 'admin' || user.role === 'superuser'"><router-link to="/profile/superuser/vehicles">Manage vehicles</router-link></p>
        <p v-if="user.role === 'admin' || user.role === 'superuser'"><router-link to="/profile/superuser/vehicles/new">Create new vehicle</router-link></p>
        <p v-if="user.role === 'admin' || user.role === 'superuser'"><router-link to="/profile/superuser/requests">Manage maintenance requests</router-link></p>
        <p v-if="user.role === 'admin' || user.role === 'superuser'"><router-link to="/profile/superuser/requests/new">Create maintenance request</router-link></p>

        <br/>
        <p v-if="user.role === 'admin' || user.role === 'superuser'"><router-link to="/profile/superuser/stops">Manage stops</router-link></p>
        <p v-if="user.role === 'admin' || user.role === 'superuser'"><router-link to="/profile/superuser/stops/new">Create new stop</router-link></p>
        <br/>
        <p v-if="user.role === 'admin' || user.role === 'technician'"><router-link to="/profile/technician/requests">My maintenance requests</router-link></p>
        <br/>
        <p v-if="user.role === 'admin' || user.role === 'dispatcher'"><router-link to="/profile/dispatcher/lines">Manage lines</router-link></p>
        <br/>
        <p v-if="user.role === 'admin' || user.role === 'driver'"><router-link to="/profile/driver/plans">My Plan</router-link></p>
        <br/>
        <p v-if="user.role === 'admin' || user.role === 'driver'"><router-link to="/profile/driver/reports/new">Report vehicle problem</router-link></p>
        <p v-if="user.role === 'admin' || user.role === 'driver'"><router-link to="/profile/driver/reports">Reported malfunctions</router-link></p>
      </div>
      <div class="hr"></div>
      <div>
        <button :onclick="logOut" class="small-button">Log Out</button>
      </div>
    </div>
  </div>
</template>

<style>
  .dashboard {
    background-color: rgb(246, 246, 247);
    width: calc((100% - (1280px - 64px)) / 2 + 350px - 32px);
    min-width: 350px;
    height: auto;
    min-height: 100vh;
  }
  .dashboard-content {
    width: 350px;
    margin-right: 0;
    margin-left: auto;
    padding: 2rem 2rem 2rem 1rem;
  }
</style>