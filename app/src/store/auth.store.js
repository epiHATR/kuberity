import { defineStore } from "pinia";
import AuthService from "../services/auth.service";
import  router from "../routes";
import { ref } from "vue";

export const useAuthStore = defineStore({
  id: "auth",
  state: () => ({
    user: JSON.parse(localStorage.getItem("user")),
    message: ref(""),
    loading: ref(false),
    returnUrl: ref('')
  }),
  getters: {
    authMessage: (state) => state.message,
    authUser: (state) => (state.user ? state.user.userinfo : {}),
    isLoggedIn: (state) => (state.user ? state.user.loggedIn : false),
  },
  actions: {
    async login(loginData) {
      this.loading = true;
      await AuthService.login(loginData).then(
        (res) => {
          const itemToSet = {
            userinfo: res.userinfo,
            loggedIn: true,
            accesstoken: res.accesstoken,
          };

          localStorage.setItem("user", JSON.stringify(itemToSet));

          this.user = itemToSet;
          this.loggedIn = true;
          this.loading = false;
          this.message = "";
          router.push(this.returnUrl || '/');
        },
        (error) => {
          this.message = error.message;
          this.loading = false;
        }
      );
    },
    async logout() {
      this.user = null;
      this.loggedIn = false;
      localStorage.removeItem("user");
      router.push("/auth/login");
    }
  },
});
