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

Vue.mixin({
  methods: {
    login: function (cb) {
      if (this.isLoggedIn()) {
        Vue.nextTick(function () {
          cb(true);
        })
        return;
      }
      var self = this;
      const refs = this.$root.$children[0].$refs;
      const loginDialog = refs["login-dialog"];
      const closed = function () {
        loginDialog.$off("close", closed);
        window.console.log("closed")
        if (cb) {
          cb(self.isLoggedIn());
        }
      };
      loginDialog.$on("close", closed);
      loginDialog.show();
    },
    logout: function () {
      this.$auth.logout();
      //this.$router.push("/");
    },
    isLoggedIn: function () {
      return this.$auth.isAuthenticated();
    }
  }
});

import Welcome from "./components/Welcome";
import MachineTypes from "./components/MachineTypes";

const router = new VueRouter({
  mode: 'hash',
  routes: [
    { path: '/', component: Welcome },
    { path: '/machinetypes', component: MachineTypes, meta: { requiresAuth: true } },
    /*    {
          path: '/login', component: Login, beforeEnter: (to, from, next) => {
            if (router.app.$auth.isAuthenticated()) {
              next({
                path: '/',
              })
            }else{
              next()
            }
          }
        },*/
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
    next()
  }
})

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
