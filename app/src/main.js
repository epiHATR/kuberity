import * as Vue from 'vue'
import axios from 'axios'
import VueAxios from 'vue-axios'
import router from './routes'
import { createPinia } from 'pinia';
import "bootstrap/dist/css/bootstrap.min.css"
import "bootstrap"
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { library } from '@fortawesome/fontawesome-svg-core'
import { fas } from '@fortawesome/free-solid-svg-icons'
import { far } from '@fortawesome/free-regular-svg-icons'
import { fab } from '@fortawesome/free-brands-svg-icons'
import App from './App.vue'

const app = Vue.createApp(App)
library.add(fas, far, fab)
axios.defaults.baseURL = "/api/v1"
app.use(VueAxios, axios)
app.component('font-awesome-icon', FontAwesomeIcon)
app.use(router)
app.use(createPinia())

app.mount('#app')