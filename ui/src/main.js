import Vue from 'vue'
Vue.config.productionTip = false

import './plugins/vuetify'
import './plugins/axios'
// if ENV without-accounts==false
// per default is accounts included
//import './plugins/account'
import './plugins/fakeAccount'
// endif
import './mixins/features'

import router from './router'
import App from './App.vue'
new Vue({
  router,
  render: h => h(App),
}).$mount('#app') 
 