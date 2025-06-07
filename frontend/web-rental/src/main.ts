import { createApp } from 'vue'
import './index.css'
// import './style.css'
import App from './App.vue'
import router from '@/routes/routes.ts'

// createApp(App).mount('#app')

const app = createApp(App)
app.use(router)
app.mount('#app')