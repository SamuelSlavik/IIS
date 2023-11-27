<script setup lang="ts">
import {onMounted, ref} from "vue";
import Loader from "@/components/Loader.vue";
import type {Malfunction, NewReport, RequestType, User} from "@/lib/models";
import {Endpoints} from "@/lib/variables";
import axios from "axios";
import {useNotificationStore} from "@/stores/notification-store";
import router from "@/router";
import {formatDate} from "../../../lib/utils";
// @ts-ignore
import Bus from "vue-material-design-icons/Bus.vue";
// @ts-ignore
import Tram from "vue-material-design-icons/Tram.vue";
// @ts-ignore
import Tank from "vue-material-design-icons/Tank.vue";
// @ts-ignore
import Delete from "vue-material-design-icons/Delete.vue";
// @ts-ignore
import Pencil from "vue-material-design-icons/Pencil.vue";
// @ts-ignore
import Hammer from "vue-material-design-icons/HammerSickle.vue";

import {useUserStore} from "@/stores/user-store";

const loading = ref<boolean>(false)
const notifications = useNotificationStore()

const request = ref<RequestType>()

const malfunction = ref<Malfunction>()

const user = useUserStore()

const newReport = ref<NewReport>({
  Title: "",
  Description: "",
  Cost: null,
})

const submitReport = async () => {
  try {
    const response = await axios.post(Endpoints.createReport, {
      Title: newReport.value.Title,
      Description: newReport.value.Description,
      Cost: newReport.value.Cost,
      MaintenReqRef: request.value?.ID,
    }, {withCredentials: true})
    if (response.status === 200) {
      notifications.addNotification("Report created", 'success')
      await router.push('/profile/technician/requests');
    }
  } catch (error) {
    notifications.addNotification("Failed to create report: " + error, 'error')
  } finally {
  }
}

const loadMalfunction = async () => {
  loading.value = true
  try {
    const response = await axios.get<Malfunction>(Endpoints.retrieveMalfunction(request.value?.MalfuncRep.ID || ""), {withCredentials: true})
    malfunction.value = response.data
  } catch (error) {
    notifications.addNotification("Failed to load malfunction: " + error, 'error')
  } finally {
    loading.value = false
  }
}

const loadReport = async () => {
  loading.value = true
  try {
    const response = await axios.get<Malfunction>(Endpoints.retrieveMalfunction(request.value?.MalfuncRep.ID || ""), {withCredentials: true})
    malfunction.value = response.data
  } catch (error) {
    notifications.addNotification("Failed to load malfunction: " + error, 'error')
  } finally {
    loading.value = false
  }
}

const loadRequest = async () => {
  try {
    loading.value = true
    const response = await axios.get(Endpoints.retrieveRequest(router.currentRoute.value.params.id.toString()), {withCredentials: true})
    request.value = response.data
    if (response.status === 200) {
      await loadMalfunction()
    }
  } catch (error: any) {
    notifications.addNotification("Failed to load request: " + error, "error")
  } finally {
    loading.value = false
  }
}

const deleteRequest = async (id: string) => {
  if (!window.confirm("Are you sure you want to delete this request?")) {
    return;
  }

  try {
    const response = await axios.delete(Endpoints.deleteRequest(id), {withCredentials: true})
    if (response.status === 200) {
      notifications.addNotification("Request deleted", "success")
      await router.push("/profile/superuser/requests")
    }
  } catch (error) {
    notifications.addNotification("Failed to delete request: " + error, "error")
  } finally {
  }
}

onMounted(() => {loadRequest()})

</script>

<template>
  <div>
    <div class="header">
      <h2>Create report</h2>
    </div>
    <Loader v-if="loading"/>
    <div v-else v-if="request && malfunction">
      <div class="header">
        <h2>{{request.MalfuncRep.Title}}</h2>
      </div>

      <div class="details">
        <div class="details-item">
          <p>Created at:</p>
          <p>{{ formatDate(request.CreatedAt) }}</p>
        </div>
        <div class="details-item">
          <p>Deadline:</p>
          <p>{{ formatDate(request.Deadline) }}</p>
        </div>
        <br/>
        <div class="details-item" v-if="request.ResolvedByRef">
          <p>Assigned to:</p>
          <p>{{request.ResolvedByRef.FirstName }} {{request.ResolvedByRef.LastName }}<br/>{{request.ResolvedByRef.Email }}</p>
        </div>
        <br/>
        <div class="details-item">
          <p>Status:</p>
          <p v-if="request.Status === 'pending'" class="yellow">Pending</p>
          <p v-if="request.Status === 'progress'" class="yellow">In progress</p>
          <p v-if="request.Status === 'done'" class="green">Done</p>
        </div>
        <br/>
        <div class="details-item">
          <p>Deadline:</p>
          <p>{{ request.Deadline }}</p>
        </div>
      </div>
      <div class="hr"></div>
      <div class="details-item">
        <p>Created by:</p>
        <p>{{ malfunction.CreatedBy.FirstName + " " + malfunction.CreatedBy.LastName}} <br/> {{malfunction.CreatedBy.Email}}</p>
      </div>
      <br/>
      <div class="details-item">
        <p>Created at:</p>
        <p>{{ formatDate(malfunction.CreatedAt) }}</p>
      </div>
      <br/>
      <div class="details-item">
        <p>Vehicle:</p>
        <p class="connection-title">
          <Bus v-if="malfunction.Vehicle.VehicleTypeName === 'bus'" class="connection-icon"/>
          <Tram v-if="malfunction.Vehicle.VehicleTypeName === 'tram'" class="connection-icon"/>
          <Tank v-if="malfunction.Vehicle.VehicleTypeName === 'obrnena_dodavka'" class="connection-icon"/>
          {{ malfunction.Vehicle.Registration }}
        </p>
      </div>
      <br/>
      <p>Description:</p>
      <p>{{malfunction.Description}}</p>

      <div class="hr"></div>

      <div>
        <form @submit.prevent="submitReport" class="form">
          <input
            type="text"
            name="title"
            placeholder="Title"
            v-model="newReport.Title"
            required
          />
          <textarea
            name="description"
            placeholder="Description"
            v-model="newReport.Description"
            required
          ></textarea>
          <input
            type="number"
            name="cost"
            placeholder="Cost"
            v-model="newReport.Cost"
            required
          />
          <button
              type="submit"
          >Submit report</button>
        </form>
      </div>

    </div>
  </div>
</template>

<style>

</style>