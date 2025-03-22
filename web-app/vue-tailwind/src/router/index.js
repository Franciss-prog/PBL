
import { createRouter, createWebHistory } from "vue-router";
import LoginView from "../views/LoginView.vue";
import RegisterView from "../views/RegisterView.vue";

// define a routes (add later royutes)
const routes = [
  // default path
  { path: "/", component: LoginView },
  { path: "/register", component: RegisterView },
];

// initalized a router based on vue-router
const router = createRouter({
  history: createWebHistory(),
  routes,
});

// export it
export default router;
