import { createRouter, createWebHistory } from "vue-router";
import LoginView from "../views/LoginView.vue";
import RegisterView from "../views/RegisterView.vue";
import DashboardView from "../views/DashboardView.vue";

const routes = [
  {
    path: "/",
    component: LoginView,
  },
  {
    path: "/register",
    component: RegisterView,
  },
  {
    path: "/dashboard",
    component: DashboardView,
  },
];

// create a arouter
const router = createRouter({
  history: createWebHistory(),
  routes,
});
// export router
export default router;
