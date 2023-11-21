import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../pages/homepage/Homepage.vue'
import Login from '../pages/user/Login.vue'
import Registration from '../pages/user/Registration.vue'
import Profile from '../pages/user/Profile.vue'
import UsersList from "@/pages/user/admin/UsersList.vue";
import NewUser from "@/pages/user/admin/NewUser.vue";
import VehiclesList from "@/pages/user/superuser/VahiclesList.vue";
import NewVehicle from "@/pages/user/superuser/NewVehicle.vue";
import UserDetail from "@/pages/user/admin/UserDetail.vue";
import MaintenanceRequests from "@/pages/user/superuser/Requests.vue";
import NewMaintenanceRequest from "@/pages/user/superuser/NewRequest.vue";

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
      component: Profile,
      children: [
        {
          path: 'users',
          component: UsersList
        },
        {
          path: 'users/new',
          component: NewUser
        },
        {
          path: 'users/detail/:id',
          component: UserDetail
        },
        {
          path: 'superuser/vehicles',
          component: VehiclesList
        },
        {
          path: 'superuser/vehicles/new',
          component: NewVehicle
        },
        {
          path: 'superuser/requests',
          component: MaintenanceRequests
        },
        {
          path: 'superuser/requests/new',
          component: NewMaintenanceRequest
        },

      ]
    }
  ]
})

export default router
/*
router.beforeEach(async (to, from, next) => {
  const requiresAuth = to.matched.some(record => record.meta.requiresAuth);
  const user = useUserStore()

  if (requiresAuth && !user.checkAuthentication() ) {
    // Redirect to the homepage if the route requires authentication and the user is not signed in
    next({ name: 'home' });
  } else {
    // Continue to the next route
    next();
  }
});

 */

/*


router.beforeEach(async (to, from, next) => {
  const user = useUserStore()
  const requiresAuth = to.matched.some(record => record.meta.requiresAuth);

  if (requiresAuth && !user.checkAuthentication()) {
    try {
      const response = await axios.get<User>(Endpoints.retrieveUser, {withCredentials: true})
      user.setUserData(response.data)
      if (response.status == 200) {
        next()
      } else {
        next({ name: 'home' })
      }
    } catch (error: any) {
      notifications.addNotification("Failed to get user: " + error, "error")
    }
  } else {
    next();
  }
});

export default router


 */
