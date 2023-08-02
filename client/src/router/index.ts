import { createRouter, createWebHistory } from 'vue-router'
import DashboardView from '../views/DashboardView.vue'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import RegisterView from '../views/RegisterView.vue'
import SpeedRecordsView from '../views/SpeedRecordsView.vue'

const isLoggedIn = async () => {
  const res = await fetch("http://localhost:8080/api/user", {
    headers: {
      "Content-Type": "application/json",
      "Authorization": `Bearer ${localStorage.getItem('jwt')}`
    }
  })

  return res.status === 200
}

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/speed-records',
      name: 'speed-records',
      component: SpeedRecordsView
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView
    },
    {
      path: '/register',
      name: 'register',
      component: RegisterView
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: DashboardView
    }
  ]
})

router.beforeEach(async (to, from, next) => {

  if (to.name === "dashboard") {

    const isLogged = await isLoggedIn();

    if (!isLogged) {
      return next({ path: "/login" });
    }
  } else if (to.name === "login" || to.name === "register") {

    const isLogged = await isLoggedIn();

    if (isLogged) {
      return next({ path: "/dashboard" });
    }
  }

  next();
});

export default router
