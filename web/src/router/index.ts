import { createRouter, createWebHashHistory } from "vue-router";

const routes = [
    {
        path: '/',
        name: 'home',
        meta: {
            title: '发现',
            keepAlive: true,
        },
        component: () => import('@/views/Home.vue'),
    },
    {
        path: '/post',
        name: 'post',
        meta: {
            title: '专题详情',
        },
        component: () => import('@/views/Post.vue'),
    },
    {
        path: '/topic',
        name: 'topic',
        meta: {
            title: '专题',
        },
        component: () => import('@/views/Topic.vue'),
    },
    {
        path: '/anouncement',
        name: 'anouncement',
        meta: {
            title: '排行榜',
        },
        component: () => import('@/views/Anouncement.vue'),
    },
    {
        path: '/profile',
        name: 'profile',
        meta: {
            title: '我的贴子',
        },
        component: () => import('@/views/Profile.vue'),
    },
    {
        path: '/user',
        name: 'user',
        meta: {
            title: '用户详情',
        },
        component: () => import('@/views/User.vue'),
    },
    {
        path: '/messages',
        name: 'messages',
        meta: {
            title: '消息',
        },
        component: () => import('@/views/Messages.vue'),
    },
    {
        path: '/collection',
        name: 'collection',
        meta: {
            title: '收藏',
        },
        component: () => import('@/views/Collection.vue'),
    },
    {
        path: '/contacts',
        name: 'contacts',
        meta: {
            title: '好友',
        },
        component: () => import('@/views/Contacts.vue'),
    },
    {
        path: "/following",
        name: "following",
        meta: {
          title: "关注",
        },
        component: () => import("@/views/Following.vue"),
    },
    {
        path: '/wallet',
        name: 'wallet',
        meta: {
            title: '订阅',
        },
        component: () => import('@/views/Wallet.vue'),
    },
    {
        path: '/setting',
        name: 'setting',
        meta: {
            title: '个人中心',
        },
        component: () => import('@/views/Setting.vue'),
    },
    {
        path: '/sharekey',
        name: 'sharekey',
        meta: {
            title: '我的分享码',
        },
        component: () => import('@/views/ShareKey.vue'),
    },
    {
        path: '/404',
        name: '404',
        meta: {
            title: '404',
        },
        component: () => import('@/views/404.vue'),
    },
    {
        path: '/community',
        name: 'community',
        meta: {
            title: '社区公约',
        },
        component: () => import('@/views/community.vue'),
    },
    {
        path: '/notification',
        name: 'notification',
        meta: {
            title: 'AiMo协议（草案）',
        },
        component: () => import('@/views/notification.vue'),
    },
    {
        path: '/:pathMatch(.*)',
        redirect: '/404',
    },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
    document.title = `${to.meta.title} | Aimo社区`;
    next();
});

export default router;
