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
            children: [
                {
                    path: "",  // http://localhost:5173/admin
                    name: "home",
                    component: () => import('../views/admin/home/index.vue'),
                },
                {
                    path: "user_center",  // http://localhost:5173/admin/user_center/user_info
                    name: "user_center",
                    children: [
                        {
                            path: "user_info",
                            name: "user_info",
                            component: () => import('../views/admin/user_center/user_info.vue'),
                        }
                    ]
                },
                {
                    path: "article",  // http://localhost:5173/admin/article/article_list
                    name: "article",
                    children: [
                        {
                            path: "article_list",
                            name: "article_list",
                            component: () => import('../views/admin/article/article_list.vue'),
                        }
                    ]
                },
                {
                    path: "users",  // http://localhost:5173/admin/users/user_list
                    name: "users",
                    children: [
                        {
                            path: "user_list",
                            name: "user_list",
                            component: () => import('../views/admin/users/user_list.vue'),
                        }
                    ]
                },
                {
                    path: "chat_group",  // http://localhost:5173/admin/chat_group/char_list
                    name: "chat_group",
                    children: [
                        {
                            path: "char_list",
                            name: "char_list",
                            component: () => import('../views/admin/chat_group/char_list.vue'),
                        }
                    ]
                },
                {
                    path: "system",  // http://localhost:5173/admin/system/menu_list
                    name: "system",
                    children: [
                        {
                            path: "menu_list",
                            name: "menu_list",
                            component: () => import('../views/admin/system/menu_list.vue'),
                        }
                    ]
                },

            ]
        }
    ]
})

export default router
