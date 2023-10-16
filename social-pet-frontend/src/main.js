import { createApp } from 'vue';
import App from './App'
import { router } from './router'
import { createPinia } from 'pinia'
import i18n from './locale';
import 'animate.css';
import './layout/index.css';

const app = createApp(App)
app.use(router).use(createPinia()).use(i18n).mount('#app');