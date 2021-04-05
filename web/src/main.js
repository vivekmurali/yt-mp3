import { createApp } from 'vue'
import App from './App.vue'
import './index.css'
import router from "./router";
import 'animate.css'
import Toaster from '@meforma/vue-toaster';

createApp(App).use(router).use(Toaster).mount('#app')
