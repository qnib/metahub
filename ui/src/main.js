import Vue from 'vue'
import './plugins/vuetify'
import App from './App.vue'
import VueAxios from 'vue-axios'
import VueAuthenticate from 'vue-authenticate'
import axios from 'axios'

Vue.config.productionTip = false

Vue.use(VueAxios, axios)
Vue.use(VueAuthenticate, {
  //baseUrl: 'http://localhost:8081', // Your API domain  
  providers: {
    github: {
      clientId: '65d9c15a3eb4e0afdd01',
      //redirectUri: 'http://localhost:8081/auth/callback' // Your client app URL
    }
  }
})

new Vue({
  render: h => h(App),
}).$mount('#app')
