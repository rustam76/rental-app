import { createMemoryHistory, createRouter } from "vue-router";
import Home from "@/pages/Home.vue";
import DefaultLayout from "@/components/layout/DefaultLayout.vue";
import BlankLayout from "@/components/layout/BlankLayout.vue";

const routes = [
    {
        path: '/',
        component: DefaultLayout,
        children: [
            {
                path: '',
                name: 'Home',
                component: Home,
            }

        ]
    },
    {
        path: '/login',
        component: BlankLayout,
        children: [
            {
                path: '',
                name: 'Login',
                component: () => import("@/pages/Auth/Login.vue"),
            },
        ]
    },
    {
        path: '/register',
        component: BlankLayout,
        children: [
            {
                path: '',
                name: 'Register',
                component: () => import("@/pages/Auth/Register.vue"),
            },
        ]
    }
];

const router = createRouter({
    history: createMemoryHistory(),
    routes,
});


export default router;
