<script setup lang="ts">
import { ref } from "vue";
import { useNotificationStore} from "@/stores/notification-store";

const notifications = useNotificationStore().notifications;

</script>

<template>
  <div v-if="notifications.length" class="notification-container">
    <div v-for="(notification, index) in notifications" :key="index" class="notification" :class="// @ts-ignore
    notification.type">
      {{ 
        // @ts-ignore
        notification.message 
      }}
    </div>
  </div>
</template>

<style scoped>
.notification-container {
  position: fixed;
  bottom: 10px;
  right: 10px;
  z-index: 1000;
}

.notification {
  margin-bottom: 10px;
  padding: 10px;
  border-radius: 5px;
  cursor: pointer;
}

.info {
  background-color: #3498db;
  color: #fff;
}

.success {
  background-color: #2ecc71;
  color: #fff;
}

.error {
  background-color: #e74c3c;
  color: #fff;
}
</style>