import Vue from 'vue';
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import App from './App.vue';
import router from './router/index'
import axios from './api/index'
Vue.use(ElementUI);
Vue.prototype.$axios = axios
Vue.config.productionTip = false
new Vue({
  el: '#app',
  render: h => h(App),
  router
}).$mount('#app');