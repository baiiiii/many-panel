import { Layout } from '@/routers/constant';

const manyPanelRouter = {
    sort: 20,
    path: '/mp',
    component: Layout,
    redirect: '/mp/host',
    meta: {
        icon: 'p-host',
        title: '多主机管理',
    },
    children: [
        {
            path: '/mp',
            name: 'mp',
            redirect: '/mp/host',
            component: () => import('@/views/many-panel/index.vue'),
            meta: {},
            children: [
                {
                    path: 'host',
                    name: 'host',
                    component: () => import('@/views/many-panel/host/index.vue'),
                    hidden: true,
                    meta: {
                        activeMenu: '/mp',
                        requiresAuth: false,
                    },
                },
                {
                    path: 'setting',
                    name: 'setting',
                    component: () => import('@/views/many-panel/setting/index.vue'),
                    hidden: true,
                    meta: {
                        activeMenu: '/mp',
                        requiresAuth: false,
                    },
                },
            ],
        },
    ],
};

export default manyPanelRouter;
