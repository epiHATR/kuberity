<template>
  <ul v-if="isLoggedIn" class="navbar-nav">
    <li class="nav-item dropdown">
      <a
        class="nav-link dropdown-toggle"
        href="#"
        id="navbarDropdown"
        role="button"
        data-toggle="dropdown"
        aria-haspopup="true"
        aria-expanded="false"
      >
        <strong>Hello, {{  authUser.username }}</strong>
      </a>
      <div
        class="dropdown-menu dropdown-menu-right"
        aria-labelledby="navbarDropdown"
      >
        <router-link to="/auth/logout" class="dropdown-item"
          >logout</router-link
        >
      </div>
    </li>
  </ul>
</template>

<script setup>
import { storeToRefs } from "pinia";
import { useAuthStore } from "../../../store/auth.store";
const { authUser, isLoggedIn } = storeToRefs(useAuthStore());
</script>
