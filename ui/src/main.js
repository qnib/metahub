import Vue from 'vue'
Vue.config.productionTip = false

import './plugins/vuetify'
import './plugins/axios'
import './plugins/account'

import './mixins/features'

import router from './router'
import App from './App.vue'
new Vue({
  router,
  render: h => h(App),
}).$mount('#app') 
 