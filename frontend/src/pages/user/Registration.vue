<script setup lang="ts">

import {ref} from "vue";
import type {UserRegistration} from "@/lib/models";
import axios from "axios";
import {Endpoints} from "@/lib/variables";

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
    alert(response.data)
  } catch (e) {
    alert(e)
  } finally {
    alert("finally")
  }

}

</script>

<template>
  <div class="container">
    <form @submit.prevent="registerNewUser" class="login-form">
      <input
          type="text"
          name="first-name"
          placeholder="First name"
          v-model="newUser.firstName"

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
      <input
          type="text"
          name="role"
          placeholder="Role"
          v-model="newUser.role"

      />
      <button
          type="submit"
      >Register Account</button>
    </form>
  </div>
</template>

<style>

</style>