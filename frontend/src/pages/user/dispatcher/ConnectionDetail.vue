<script setup lang="ts">
import {onMounted, ref} from "vue";
import Loader from "@/components/Loader.vue";
import type {ConnectionList, User} from "@/lib/models";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";
import router from "@/router";
import {formatDate, formatDateTime} from "@/lib/utils";
// @ts-ignore
import Magnify from "vue-material-design-icons/Magnify.vue";
// @ts-ignore
import Close from "vue-material-design-icons/Close.vue";
// @ts-ignore
import Check from "vue-material-design-icons/Check.vue";

const loading = ref<boolean>(false)
const notifications = useNotificationStore()
const connection = ref<ConnectionList[]>()
const connectionID = router.currentRoute.value.params.id.toString() || ""

const cetTime = new Date()
cetTime.setHours(cetTime.getHours() + 2)
const querydate = ref<string>(cetTime.toISOString().split("T")[0])

</script>

<template>
    
</template>

<style>

</style>
