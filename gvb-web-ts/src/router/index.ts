import {createRouter, createWebHistory} from 'vue-router'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'web',
            component: () => import('../views/web/web.vue'),
            children: [
                {
                    path: "",
                    name: "index",
                    component: () => import('../views/web/index.vue'),
                }
            ]
        },
        {
            path: "/admin",
            name: "admin",
            component: () => import('../views/admin/index.vue'),
            meta: {title: "首页"},
            children: [
                {
                    path: "",  // http://localhost:5173/admin
                    name: "home",
                    meta: {title: "home"},
                    component: () => import('../views/admin/home/index.vue'),
                },
                {
                    path: "user_center",  // http://localhost:5173/admin/user_center/user_info
                    name: "user_center",
                    meta: {title: "个人中心"},
                    children: [
                        {
                            path: "user_info",
                            name: "user_info",
                            meta: {title: "我的信息"},
                            component: () => import('../views/admin/user_center/user_info.vue'),
                        }
                    ]
                },
                {
                    path: "article",  // http://localhost:5173/admin/article/article_list
                    name: "article",
                    meta: {title: "文章管理"},
                    children: [
                        {
                            path: "article_list",
                            name: "article_list",
                            meta: {title: "文章列表"},
                            component: () => import('../views/admin/article/article_list.vue'),
                        }
                    ]
                },
                {
                    path: "users",  // http://localhost:5173/admin/users/user_list
                    name: "users",
                    meta: {title: "用户管理"},
                    children: [
                        {
                            path: "user_list",
                            name: "user_list",
                            meta: {title: "用户列表"},
                            component: () => import('../views/admin/users/user_list.vue'),
                        }
                    ]
                },
                {
                    path: "chat_group",  // http://localhost:5173/admin/chat_group/char_list
                    name: "chat_group",
                    meta: {title: "群聊管理"},
                    children: [
                        {
                            path: "char_list",
                            name: "char_list",
                            meta: {title: "聊天记录"},
                            component: () => import('../views/admin/chat_group/char_list.vue'),
                        }
                    ]
                },
                {
                    path: "system",  // http://localhost:5173/admin/system/menu_list
                    name: "system",
                    meta: {title: "系统管理"},
                    children: [
                        {
                            path: "menu_list",
                            name: "menu_list",
                            meta: {title: "菜单列表"},
                            component: () => import('../views/admin/system/menu_list.vue'),
                        }
                    ]
                },

            ]
        }
    ]
})

export default router
