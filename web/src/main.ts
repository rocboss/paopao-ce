import { createApp } from 'vue';
import { createPinia } from 'pinia';
import router from './router';
import App from './App.vue';
import '@/assets/css/main.less';

import type { MessageApiInjection } from 'naive-ui/lib/message/src/MessageProvider';

// 通用字体
import 'vfonts/Lato.css';
// 等宽字体
import 'vfonts/FiraCode.css';

const pinia = createPinia();

createApp(App).use(router).use(pinia).mount('#app');

declare global {
  interface Window {
    $message: MessageApiInjection;
  }
}
