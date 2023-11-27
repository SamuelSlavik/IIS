import { defineStore } from 'pinia';
import { ref, computed } from 'vue';

export const useNotificationStore = defineStore('notification', () => {
    const notifications = ref([]);

    const addNotification = (message: string = "", type: string = 'info') => {
        // @ts-ignore
        notifications.value.push({ message, type });
        setTimeout(() => {
            removeNotification(0);
        }, 5000); // Automatically remove the notification after 5 seconds
    };

    const removeNotification = (index: any) => {
        notifications.value.splice(index, 1);
    };

    // Computed property to get the current notifications
    const currentNotifications = computed(() => notifications.value);

    return {
        notifications: currentNotifications,
        addNotification,
        removeNotification,
    };
});