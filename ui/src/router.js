import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

import Welcome from "./views/Welcome";
import MachineTypes from "./views/MachineTypes";
import MachineType from "./views/MachineType"

const router = new Router({
    mode: 'hash',
    routes: [
        {
            name: "welcome", path: '/', component: Welcome, meta: {
                title: "Welcome",
            }
        },
        {
            name: "machine-types",
            path: '/machinetypes',
            components: {
                default: MachineTypes,
            },
            meta: {
                title: "Machine Types",
                requiresAuth: true,
            }
        },
        {
            name: "edit-machine-type",
            path: '/machinetypes/edit/:id',
            component: MachineType,
            meta: {
                title: "Machine Type details",
                requiresAuth: true,
            }
        },
        {
            name: "new-machine-type",
            path: '/machinetypes/new',
            component: MachineType,
            meta: {
                title: "Machine Type details",
                requiresAuth: true,
            }
        },
    ]
})

router.beforeEach((to, from, next) => {
    if (to.matched.some(record => record.meta.requiresAuth)) {
        const account = router.app.$account;
        if (account.isLoggedIn()) {
            next();
        } else {
            account.login(function () {
                if (!account.isLoggedIn()) {
                    next(false);
                    return;
                }
                next();
            });
        }
    } else {
        next();
    }
})

export default router;