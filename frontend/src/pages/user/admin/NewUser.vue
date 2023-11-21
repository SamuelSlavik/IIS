<script setup lang="ts">
import {ref} from "vue";
import type {UserRegistration} from "@/lib/models";
import axios from "axios";
import {Endpoints} from "@/lib/variables";
import router from "@/router";
import {useNotificationStore} from "@/stores/notification-store";

const notifications = useNotificationStore()

const newUser = ref<UserRegistration>({
  firstName: "",
  lastName: "",
  birthDate: "",
  email: "",
  password: "",
  passwordRpt: "",
  role: "",
})

const registerNewUser = async () => {
  try {
    const response = await axios.post(Endpoints.signup, newUser.value)
    if (response.status === 200) {
      alert("User created")
    }
    await router.push('/profile/users');
  } catch (e) {
    notifications.addNotification("Failed to create user: " + e, "error")
  }
}
</script>

<template>
  <div>
    <h2>Create new user</h2>
    <br/>
    <form @submit.prevent="registerNewUser" class="form">
      <input
          type="text"
          name="first-name"
          placeholder="First name"
          v-model="newUser.firstName"
          required
      />
      <input
          type="text"
          name="last-name"
          placeholder="Last name"
          v-model="newUser.lastName"

      />
      <input
          type="date"
          name="birth-date"
          placeholder="Date of birth"
          v-model="newUser.birthDate"

      />
      <input
          type="email"
          name="email"
          placeholder="Email"
          v-model="newUser.email"

      />
      <input
          type="password"
          name="password"
          placeholder="Password"
          v-model="newUser.password"

      />
      <input
          type="password"
          name="password-rpt"
          placeholder="Repeat password"
          v-model="newUser.passwordRpt"
      />
      <select v-model="newUser.role">
        <option value="superuser">Superuser</option>
        <option value="technician">Technician</option>
        <option value="dispatcher">Dispatcher</option>
        <option value="driver">Driver</option>
      </select>
      <button
          type="submit"
      >Register Account</button>
    </form>
  </div>
</template>

<style>
.form {
  width: 100%;
  max-width: 400px;
  margin-left: 0;
  display: flex;
  flex-direction: column;
  gap: 2rem;
}
</style>