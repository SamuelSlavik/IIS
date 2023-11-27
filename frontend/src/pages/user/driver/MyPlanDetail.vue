<script setup lang="ts">
import {onMounted, ref} from "vue";
import Loader from "@/components/Loader.vue";
import type {ConnectionList, User} from "@/lib/models";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";
import lines from "@/pages/user/superuser/Lines.vue";
import {useUserStore} from "@/stores/user-store";
import {formatDateTime} from "../../../lib/utils";
// @ts-ignore
import Bus from "vue-material-design-icons/Bus.vue";
// @ts-ignore
import Tram from "vue-material-design-icons/Tram.vue";
// @ts-ignore
import Tank from "vue-material-design-icons/Tank.vue";
import router from "@/router";

const loading = ref<boolean>(false)
const notifications = useNotificationStore()

const connectionId = router.currentRoute.value.params.id.toString() || ""

const user = useUserStore()

const connection = ref<ConnectionList>()

const loadConnection = async () => {

}

onMounted(() => {loadConnection()})

</script>

<template>
  <div>
    <Loader v-if="loading"/>
    <div v-else>
      <div class="header">
        <h2></h2>
      </div>
      <div class="hr"></div>
      <div class="details">

      </div>
    </div>
  </div>
</template>

<style>

</style>