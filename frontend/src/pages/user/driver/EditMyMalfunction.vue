<script setup lang="ts">
import {onMounted, ref} from "vue";
import Loader from "@/components/Loader.vue";
import type {User} from "@/lib/models";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";

const loading = ref<boolean>(false)
const notifications = useNotificationStore()

const newReport = ref<MalfunctionReport>({
  Title: "",
  Description: "",
  VehicleRef: "",
})

</script>

<template>
  <div>
    Working..
  </div>
</template>

<style>

</style>