import type { RouteRecordRaw } from "vue-router";

export const routes: RouteRecordRaw[] = [
    {
        path: '/',
        component: () => import('@/pages/Server.vue')
    },
    {
        path: '/map',
        component: () => import('@/pages/Map.vue')
    },
    {
        path: '/wiki',
        component: () => import('@/pages/Wiki.vue')
    }

]