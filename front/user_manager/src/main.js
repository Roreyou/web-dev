import Vue from 'vue'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import App from './App.vue'
import router from './router'
import store from './store';
//import './api/mock'
import axios from 'axios'

Vue.config.productionTip = false
Vue.prototype.$axios = axios
Vue.use(ElementUI)

//axios.defaults.baseURL = '/api'        //关键代码
Vue.config.productionTip = false

//实例化vue对象
new Vue({
  router,
  axios,
  store,
  render: h => h(App),
}).$mount('#app')
