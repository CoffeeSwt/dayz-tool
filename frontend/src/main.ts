import { createApp } from 'vue'
import App from './App.vue'
import './assets/index.css'
import 'virtual:uno.css'
import { router } from '@/router'

const app = createApp(App)
app.use(router)

app.mount('#app')
