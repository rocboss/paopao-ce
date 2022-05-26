import { createApp } from 'vue'
import router from './router'
import store from './store'
import App from './App.vue'

// 通用字体
import 'vfonts/Lato.css'
// 等宽字体
import 'vfonts/FiraCode.css'

createApp(App).use(router).use(store).mount('#app')
