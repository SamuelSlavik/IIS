<script setup lang="ts">
import {ref} from "vue";
import axios from "axios";
import {Endpoints} from "@/lib/variables";
import { useNotificationStore } from "@/stores/notification-store";
import { useUserStore } from "@/stores/user-store";
import Loader from "@/components/Loader.vue";
import { useRouter } from 'vue-router';

const notifications = useNotificationStore();
const router = useRouter();
const user = useUserStore()

const email =ref<string>("")
const password = ref<string>("")

const loading = ref<boolean>(false)

const login = async () => {
  loading.value = true
  try {
    const response = await axios.post(Endpoints.login,
        {Email: email.value, Password: password.value},
        { withCredentials: true }
    )
    if (response.status === 200) {
      await router.push('/profile');
    }
  } catch (e) {
    notifications.addNotification("Logging in failed: " + e, "error")
  } finally {
    loading.value = false
  }
}

</script>

<template>
<div class="container">
  <Loader v-if="loading"/>
  <form @submit.prevent="login" class="login-form">
    <input
      type="email"
      name="email"
      placeholder="Email"
      v-model="email"
      required
    />
    <input
      type="password"
      name="password"
      placeholder="Password"
      v-model="password"
      required
    />
    <button
      type="submit"
    >Log in</button>
  </form>
</div>
</template>

<style>
  .login-form {
    width: 100%;
    max-width: 400px;
    margin: auto;
    display: flex;
    flex-direction: column;
    gap: 2rem;
  }

  input {
    width: 100%;
    padding: 1rem;
    outline: none;
    border-radius: 4px;
    border: 1px solid rgba(60, 60, 67, .12);
  }
  input:invalid {
    border: 1px solid #ff453a;
  }
  input:hover, input:focus {
    border: 1px solid #00bd7e;
  }

  button {
    width: auto;
    outline: none;
    border:none;
    padding: 1rem;
    transition: all ease 0.3s;
    border-radius: 4px;
    border: 1px solid rgba(60, 60, 67, .12);
  }
  button:hover {
    background-color: #00bd7e;
    cursor: pointer;
  }
  button.small-button {
    padding: 0.5rem;
  }
</style>