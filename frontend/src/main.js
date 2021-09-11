import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import axios from "axios";

import "./assets/css/bootstrap.min.css"
import "@/assets/js/bootstrap.min"
import "@/assets/css/styles.css"
import 'vue-neat-modal/dist/vue-neat-modal.css'

axios.defaults.baseURL = "http://127.0.0.1:3330/api/"

createApp(App).use(router).mount('#app')
