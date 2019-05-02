import Vue from 'vue'
import VueAuthenticate from 'vue-authenticate'

Vue.use(VueAuthenticate, {
    providers: {
        github: {
            clientId: '65d9c15a3eb4e0afdd01',
        },
        google: {
            clientId: '936040293434-i3m9p4km8it5np2bs253a7rvedchofs6.apps.googleusercontent.com',
        }
    }
})

import Account from "./Account";

Vue.use({
    install(vue) {
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