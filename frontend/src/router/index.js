import { createWebHistory, createRouter } from 'vue-router';

const routes = [
  {
    path: "/",
    component: () => import("../components/ItemList")
  },
  {
    path: "/login",
    component: () => import("../components/LoginPage")
  },
  {
    path: "/register",
    component: () => import("../components/RegisterPage")
  },
  {
    path: "/item/:id",
    component: () => import("../components/ItemPage")
  },

];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;