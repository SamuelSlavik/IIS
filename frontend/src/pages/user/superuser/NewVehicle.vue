<script setup lang="ts">
import {useUserStore} from "@/stores/user-store";
import {useRouter} from "vue-router";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";

const user = useUserStore()
const router = useRouter();
let notifications = useNotificationStore();

</script>

<template>
  <div>
    <h2>Create new vehicle</h2>
  </div>
</template>

<style>

</style>