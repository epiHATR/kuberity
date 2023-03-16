<template>
  <NavigationBarComp />

  <div class=".container-fluid m-3">
    <router-view />
  </div>
</template>

<script>
import NavigationBarComp from './components/shared/NavigationBarComp.vue';

export default {
  name: "App",
  components: {
    NavigationBarComp
  },
  watch: {
    $route: {
      immediate: true,
      handler() {
        document.title = "kuberity for Kubernetes";
      },
    },
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif !important;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: left;
  color: #2c3e50;
  font-size: 12px;
}

.dropdown-menu .dropdown-item {
  font-size: 12px !important;
}

.router-link-exact-active {
  font-weight: bold !important;
}

.active-item {
  color: white !important;
  background: #2c3e50;
}

.breadcums a {
  color: #2c3e50 !important;
}
</style>
