import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../pages/homepage/Homepage.vue'
import Login from '../pages/user/Login.vue'
import Registration from '../pages/user/Registration.vue'
import Profile from '../pages/user/Profile.vue'
import UsersList from "@/pages/user/admin/UsersList.vue";
import NewUser from "@/pages/user/admin/NewUser.vue";
import VehiclesList from "@/pages/user/superuser/VehiclesList.vue";
import NewVehicle from "@/pages/user/superuser/NewVehicle.vue";
import UserDetail from "@/pages/user/admin/UserDetail.vue";
import MaintenanceRequests from "@/pages/user/superuser/Requests.vue";
import NewMaintenanceRequest from "@/pages/user/superuser/NewRequest.vue";
import Connections from "@/pages/user/superuser/Connections.vue";
import NewStop from "@/pages/user/superuser/NewStop.vue";
import Stops from "@/pages/user/superuser/Stops.vue";
import Lines from "@/pages/user/superuser/Lines.vue";
import LineDetail from "@/pages/user/superuser/LineDetail.vue";
import NewLine from "@/pages/user/superuser/NewLine.vue";
import EditLine from "@/pages/user/superuser/EditLine.vue";
import EditStop from "@/pages/user/superuser/EditStop.vue";
import EditRequest from "@/pages/user/superuser/EditRequest.vue";
import RequestDetail from "@/pages/user/superuser/RequestDetail.vue";
import EditUser from "@/pages/user/admin/EditUser.vue";
import VehicleDetail from "@/pages/user/superuser/VehicleDetail.vue";
import EditVehicle from "@/pages/user/superuser/EditVehicle.vue";
import MyRequests from "@/pages/user/technician/MyRequests.vue";
import MyRequestDetail from "@/pages/user/technician/MyRequestDetail.vue";
import MyRequestComplete from "@/pages/user/technician/MyRequestComplete.vue";
import MyPlan from "@/pages/user/driver/MyPlan.vue";
import NewMalfunction from "@/pages/user/driver/NewMalfunction.vue";
import EditMyMalfunction from "@/pages/user/driver/EditMyMalfunction.vue";
import MyMalfunctions from "@/pages/user/driver/MyMalfunctions.vue";
import MalfunctionDetail from "@/pages/user/driver/MalfunctionDetail.vue";
import Malfunctions from "@/pages/user/superuser/Malfunctions.vue";
import LinesListDispatcher from "@/pages/user/dispatcher/Lines.vue";
import LineConnectionsListDispatcher from "@/pages/user/dispatcher/Connections.vue";
import NewConnection from "@/pages/user/superuser/NewConnection.vue";
import LineConnections from "@/pages/user/superuser/LineConnections.vue";
import EditConnection from "@/pages/user/superuser/EditConnection.vue";
import LineConnectionDetailDispatcher from "@/pages/user/dispatcher/ConnectionDetail.vue";
import Hello from "@/pages/user/Hello.vue";
import MyPlanDetail from "@/pages/user/driver/MyPlanDetail.vue";
import Drivers from "@/pages/user/admin/Drivers.vue";
import DriversPlan from "@/pages/user/admin/DriversPlan.vue";

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
          path: '',
          component: Hello
        },
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
          path: 'users/edit/:id',
          component: EditUser
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
          path: 'superuser/vehicles/detail/:id',
          component: VehicleDetail
        },
        {
          path: 'superuser/vehicles/edit/:id',
          component: EditVehicle
        },
        {
          path: 'superuser/requests',
          component: MaintenanceRequests
        },
        {
          path: 'superuser/requests/create/:id',
          component: NewMaintenanceRequest
        },
        {
          path: 'superuser/requests/detail/:id',
          component: RequestDetail
        },
        {
          path: 'superuser/requests/edit/:id',
          component: EditRequest
        },
        {
          path: 'superuser/connections',
          component: Connections
        },
        {
          path: 'superuser/connections/:line',
          component: LineConnections
        },
        {
          path: 'superuser/connections/new',
          component: NewConnection
        },
        {
          path: 'superuser/connections/edit/:id',
          component: EditConnection
        },
        {
          path: 'superuser/stops/new',
          component:  NewStop
        },
        {
          path: 'superuser/stops',
          component: Stops
        },
        {
          path: 'superuser/stops/edit/:id',
          component: EditStop
        },
        {
          path: 'superuser/lines',
          component: Lines
        },
        {
          path: 'superuser/lines/detail/:name',
          component: LineDetail
        },
        {
          path: 'superuser/lines/new',
          component: NewLine
        },
        {
          path: 'superuser/lines/edit/:name',
          component: EditLine
        },
        {
          path: 'malfunctions/detail/:id',
          component: MalfunctionDetail
        },
        {
          path: 'superuser/malfunctions',
          component: Malfunctions
        },




        {
          path: 'technician/requests',
          component: MyRequests
        },
        {
          path: 'technician/requests/detail/:id',
          component: MyRequestDetail
        },
        {
          path: 'technician/requests/complete/:id',
          component: MyRequestComplete
        },
        {
          path: 'admin/drivers',
          component: Drivers
        },
        {
          path: 'admin/drivers/detail/:id',
          component: DriversPlan
        },
        {
          path: 'driver/plans',
          component: MyPlan
        },
        {
          path: 'driver/connection/detail/:id',
          component: MyPlanDetail
        },
        {
          path: 'driver/reports/new',
          component: NewMalfunction
        },
        {
          path: 'driver/reports',
          component: MyMalfunctions
        },
        {
          path: 'malfunctions/edit/:id',
          component: EditMyMalfunction
        },

        {
          path: 'dispatcher/lines',
          component: LinesListDispatcher
        },
        {
          path: 'dispatcher/connections/:line',
          component: LineConnectionsListDispatcher
        },
        {
          path: '/profile/dispatcher/connection/detail/:id',
          component: LineConnectionDetailDispatcher
        }
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
