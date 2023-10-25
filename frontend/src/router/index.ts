import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../pages/homepage/Homepage.vue'
import Login from '../pages/user/Login.vue'
import Registration from '../pages/user/Registration.vue'
import Profile from '../pages/user/Profile.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/registration',
      name: 'registration',
      component: Registration
    },
    {
      path: '/profile',
      name: 'profile',
      component: Profile
    }
  ]
})

export default router
