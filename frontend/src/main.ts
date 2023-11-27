import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import vSelect from "vue-select";
import "vue-select/dist/vue-select.css";

import App from './App.vue'
import router from './router'
import axios from "axios";

//axios.defaults.headers.common['Access-Control-Allow-Origin'] = '*'

const app = createApp(App)


axios.interceptors.response.use(
    (response) => response,
    (error) => {
        // If the response status is 401 (Unauthorized), navigate to the login page
        if (error.response.status === 401) {
            router.push('/login');
        }
        return Promise.reject(error);
    }
);

app.component("v-select", vSelect);
app.use(createPinia())
app.use(router)

app.mount('#app')
