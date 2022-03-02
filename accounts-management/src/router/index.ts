import { createRouter, createWebHistory } from 'vue-router';
import HomeView from '../views/HomeView.vue';
import SignInPage from '../views/SignInPage.vue';
import UserOperations from '../views/UserOperations.vue';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      redirect: '/login'
    },
    {
      path: '/accounts-list',
      name: 'accounts-list',
      component: HomeView
    },
    {
      path: '/add-user',
      name: 'add-user',
      component: UserOperations
    },
    {
      path: '/user-detail/:id',
      name: 'user-detail',
      component: UserOperations,
    },
    {
      path: '/login',
      name: 'login',
      component: SignInPage
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue')
    }
  ]
});

export default router;
