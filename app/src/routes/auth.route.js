export default {
  path: "/auth",
  name: "auth",
  children: [
    {
      path: "/auth/login",
      name: "login",
      meta: {
        title: "login | kuberity for kubernetes",
      },
      component: () => import("../components/auth/LoginPage"),
    },
    {
      path: "/auth/logout",
      name: "logout",
      meta: {
        title: "logout | kuberity for kubernetes",
      },
      component: () => import("../components/auth/LogoutPage"),
    },
  ],
};
