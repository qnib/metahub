import Vue from 'vue'

//import axios from '../axios.js'
import axios from 'axios';

import Account from "./Account";

Vue.use({
    install(vue) {
        axios.interceptors.request.use(function (config) {
            config.headers.Authorization = 'Bearer DUMMY';
            return config;
        }, function (error) {
            return Promise.reject(error);
        });

        const root = new Vue({
            data: {},
            render: createElement => createElement(Account),
            mounted() {
                vue.prototype.$account = this.$root.$children[0];
            },
        })
        root.$mount(document.body.appendChild(document.createElement('div')))
    }
})