import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import "./assets/css/bootstrap.min.css"
import "@/assets/js/bootstrap.min"

createApp(App).use(router).mount('#app')
