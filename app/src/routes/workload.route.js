export default {
  path: "/workloads",
  name: "workloads",
  meta: {
    title: "workloads | kuberity for kubernetes",
  },
  children: [
    {
      path: "/workloads/services",
      name: "services",
      meta: {
        title: "services | kuberity for kubernetes",
      },
      component: () => import("../components/ServicesPage"),
    },
    {
      path: "/workloads/deployments",
      name: "deployments",
      meta: {
        title: "deployments | kuberity for kubernetes",
      },
      component: () => import("../components/DeploymentsPage"),
    },
    {
      path: "/workloads/pods",
      name: "pods",
      meta: {
        title: "pods | kuberity for kubernetes",
      },
      component: () => import("../components/PodsPage"),
    },
  ],
};
