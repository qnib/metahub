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
    },
    google: {
      clientId: '936040293434-i3m9p4km8it5np2bs253a7rvedchofs6.apps.googleusercontent.com',
    }
  }
})

import './mixins/login'

import VueRouter from 'vue-router'
Vue.use(VueRouter)

import Welcome from "./components/Welcome";
import MachineTypes from "./components/MachineTypes";

const router = new VueRouter({
  mode: 'hash',
  routes: [
    { path: '/', component: Welcome },
    { path: '/machinetypes', component: MachineTypes, meta: { requiresAuth: true } },
  ]
})

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    router.app.login(function (loggedIn) {
      if (loggedIn) {
        next();
      } else {
        next(false);
      }
    });
  } else {
    next();
  }
})

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
