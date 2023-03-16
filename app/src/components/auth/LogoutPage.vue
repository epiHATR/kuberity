<template>
  <p>Logout Page</p>
</template>

<script setup>
import { useAuthStore } from "../../store/auth.store"
const { logout } = useAuthStore();
logout()
</script>
