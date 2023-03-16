import { createRouter, createWebHistory } from "vue-router";

import { useAuthStore } from "@/store/auth.store";

import authRoute from "./auth.route";
import workloadRoute from "./workload.route";
import dashboardRoute from "./dashboard.route";

const routes = [
  { path: "", component: () => import("../components/DashboardPage") },
  { ...workloadRoute},
  { ...authRoute},
  { ...dashboardRoute},
  { path: "/:pathMatch(.*)*", redirect: "/auth/login" },
];

const router = createRouter({
  history: createWebHistory(),
  linkActiveClass: "active",
  routes,
});

router.beforeEach(async (to) => {
  const publicPages = ["/auth/login", "about", "help"];

  const authRequired = !publicPages.includes(to.path);
  const authStore = useAuthStore();
  if (authRequired && !authStore.isLoggedIn) {
    authStore.returnUrl = to.fullPath;
    return { name: "login", redirect: to.fullPath };
  }
});

export default router;
