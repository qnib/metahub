import Vue from 'vue'
import './plugins/vuetify'
import App from './App.vue'
import VueAxios from 'vue-axios'
import VueAuthenticate from 'vue-authenticate'
import axios from 'axios'

import VueRouter from 'vue-router'
Vue.use(VueRouter)

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

import Welcome from "./components/Welcome";
import MachineTypes from "./components/MachineTypes";
import Login from "./components/Login";

const router = new VueRouter({
  mode: 'history',
  routes: [
    { path: '/', component: Welcome },
    { path: '/machinetypes', component: MachineTypes, meta: { requiresAuth: true } },
    {
      path: '/login', component: Login, beforeEnter: (to, from, next) => {
        if (router.app.$auth.isAuthenticated()) {
          next({
            path: '/',
          })
        }else{
          next()
        }
      }
    },
  ]
})
router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (!router.app.$auth.isAuthenticated()) {
      next({
        path: '/login',
        query: { redirect: to.fullPath }
      })
    } else {
      next()
    }
  } else {
    next()
  }
})

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
