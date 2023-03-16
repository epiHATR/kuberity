import { defineStore } from "pinia";

import  DeploymentsService from "../services/deployment.service"
import  ServicesService from "../services/service.service"
import  PodsService from "../services/pod.service"

export const useAppStore = defineStore({
  id: "app",
  state: () => ({
    deployments: [],
    services: [],
    pods: [],
    isLoading: false
  }),
  getters: {
    getDeployments: (state) => state.deployments
  },
  actions: {
    async fetchDeployments() {
        this.isLoading = true
        await DeploymentsService.getAllDeployments().then( (res) => {
            this.deployments = res.data
            this.isLoading = false
        })
    },
    async fetchServices() {
        this.isLoading = true
        await ServicesService.getAllServices().then( (res) => {
            this.services = res.data
            this.isLoading = false
        })

        await PodsService.getAllPods().then( (res) => {
            this.pods = res.data
            this.isLoading = false
        })
    },
    async fetchPods() {
        this.isLoading = true
        await PodsService.getAllPods().then( (res) => {
            this.pods = res.data
            this.isLoading = false
        })
    }
  },
});
