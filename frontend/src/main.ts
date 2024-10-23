import { createApp } from 'vue'
import App from './App.vue'
import './assets/index.css'
import 'virtual:uno.css'
import { router } from '@/router'
import { createPinia } from 'pinia'

const app = createApp(App)
const pinia = createPinia()

app.use(pinia)
app.use(router)

app.mount('#app')
