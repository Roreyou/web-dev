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

router.beforeEach((to,from,next)=>{
  const if_token=sessionStorage.getItem("token")
  if(!if_token&&to.name!=='Login')
  {
    next({name:'Login'})
  }
  else if(if_token&&to.name==='Login'){
    next('/manage')//用户已登录时，使其不再访问登录页面
  }
  else if(if_token!=='1'&&(to.name==='Contmanage'||to.name==='GPUmanage'||to.name==='Usermanage'))
  {
    next('/manage')
  }
  else if(if_token==='1'&&(to.name!=='Contmanage'&&to.name!=='GPUmanage'&&to.name!=='Usermanage'))
  {
    next('/gpumanage')
  }
  else{
    next();
  }
})

//axios.defaults.baseURL = '/api'        //关键代码
Vue.config.productionTip = false

//实例化vue对象
new Vue({
  router,
  axios,
  store,
  render: h => h(App),
}).$mount('#app')
