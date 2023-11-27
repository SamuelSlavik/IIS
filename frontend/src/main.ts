import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import vSelect from "vue-select";
import "vue-select/dist/vue-select.css";

import App from './App.vue'
import router from './router'
import axios from "axios";
import { Endpoints } from './lib/variables';
import { useUserStore } from './stores/user-store';
import { useNotificationStore } from './stores/notification-store';

//axios.defaults.headers.common['Access-Control-Allow-Origin'] = '*'

const app = createApp(App)

const user = useUserStore()
let notifications = useNotificationStore()

const logOut = async () => {
    try {
      const response = await axios.get(Endpoints.logout, {withCredentials: true})
      if (response.status === 200) {
        user.logOut()
        await router.push('/')
      }
    } catch (error: any) {
      notifications.addNotification("Failed to logout: " + error, "error")
    } finally {
  
    }
}


axios.interceptors.response.use(
    (response) => response,
    (error) => {
        // If the response status is 401 (Unauthorized), navigate to the login page
        if (error.response.status === 401) {
            logOut()
            router.push('/login');
        }
        return Promise.reject(error);
    }
);

app.component("v-select", vSelect);
app.use(createPinia())
app.use(router)

app.mount('#app')
