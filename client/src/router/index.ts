import { createRouter, createWebHistory } from 'vue-router'
import BestScoresView from '../views/BestScoresView.vue'
import DashboardView from '../views/DashboardView.vue'
import LoginView from '../views/LoginView.vue'
import RegisterView from '../views/RegisterView.vue'
import TypingView from '../views/TypingView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'typing',
      component: TypingView,
    },
    {
      path: '/best-scores',
      name: 'best-scores',
      component: BestScoresView,
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView,
    },
    {
      path: '/register',
      name: 'register',
      component: RegisterView,
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: DashboardView,
    }
  ]
})

router.beforeEach(async (to, from, next) => {
  next();
});

export default router
