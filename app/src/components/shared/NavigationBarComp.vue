<template>
  <nav class="navbar navbar-expand-xl navbar-dark bg-dark">
    <button
      class="navbar-toggler"
      type="button"
      data-toggle="collapse"
      data-target="#navbarSupportedContent"
      aria-controls="navbarSupportedContent"
      aria-expanded="false"
      aria-label="Toggle navigation"
    >
      <span class="navbar-toggler-icon"></span>
    </button>

    <router-link to="/" class="navbar-brand">kuberity</router-link>
    <div v-if="isLoggedIn" class="collapse navbar-collapse" id="navbarSupportedContent">
      <ul class="navbar-nav mr-auto">
        <li class="nav-item dropdown">
          <a
            class="nav-link dropdown-toggle"
            href="#"
            id="navbardrop"
            data-toggle="dropdown"
          >
            Cluster
            <font-awesome-icon icon="fa-solid fa-network-wired" />
          </a>
          <div class="dropdown-menu">
            <router-link to="/workloads/services" class="dropdown-item" exact
              >Namespace</router-link
            >
            <router-link to="/workloads/services" class="dropdown-item" exact
              >Persitence Volumes</router-link
            >
            <router-link to="/workloads/services" class="dropdown-item" exact
              >Nodes</router-link
            >
            <router-link to="/workloads/services" class="dropdown-item" exact
              >Storage Classes</router-link
            >
          </div>
        </li>

        <!-- Dropdown Workfloads-->
        <li class="nav-item dropdown">
          <a
            class="nav-link dropdown-toggle"
            href="#"
            id="navbardrop"
            data-toggle="dropdown"
          >
            Workloads
            <font-awesome-icon icon="fa-solid fa-bars-progress" />
          </a>
          <div class="dropdown-menu">
            <router-link to="/workloads/services" class="dropdown-item" exact
              >Services</router-link
            >
            <router-link to="/workloads/deployments" class="dropdown-item" exact
              >Deployments</router-link
            >
            <router-link to="/workloads/pods" class="dropdown-item" exact
              >Pods</router-link
            >
          </div>
        </li>
      </ul>
      <UserProfileDropdown />
    </div>
  </nav>
</template>

<script setup>
import UserProfileDropdown from "@/components/shared/partials/UserProfileDropdown.vue";
import { storeToRefs } from "pinia";
import { useAuthStore } from "../../store/auth.store";
const { isLoggedIn } = storeToRefs(useAuthStore());

</script>
