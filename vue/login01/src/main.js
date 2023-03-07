import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import axios from 'axios'
import jsCookie from 'js-cookie'
import VueResource from 'vue-resource'

Vue.use(VueResource)
Vue.prototype.$cookie = jsCookie;
Vue.config.productionTip = false
Vue.prototype.$axios=axios//把axios这个包挂载到原型对象上
axios.defaults.baseURL="http://127.0.0.1:8080"
axios.defaults.withCredentials=false


Vue.use(ElementUI);
new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')




