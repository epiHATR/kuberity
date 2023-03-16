export default {
    path: "/dashboard",
    name: "dashboard",
    meta: {
        title: "dashboard | kuberity for kubernetes",
    },
    component: () => import("../components/DashboardPage"),
};