import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import "./assets/css/bootstrap.min.css"
import "@/assets/js/bootstrap.min"
import "@/assets/css/styles.css"
import 'vue-neat-modal/dist/vue-neat-modal.css'

createApp(App).use(router).mount('#app')
