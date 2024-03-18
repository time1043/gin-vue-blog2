<template>
  <div class="gvb_menu">
    <a-menu
        @menu-item-click="clickMenu"
        v-model:selected-keys="selectedKey"
        v-model:open-keys="openKeys"
    >
      <template v-for="item in menuList" :key="item.key">
        <a-menu-item :key="item.name" v-if="item.child?.length === 0">
          {{ item.title }}
          <template #icon>
            <component :is="item.icon"></component>
          </template>
        </a-menu-item>

        <a-sub-menu :key="item.name" v-if="item.child?.length !== 0">
          <template #icon>
            <component :is="item.icon"></component>
          </template>
          <template #title>{{ item.title }}</template>
          <a-menu-item :key="sub.name" v-for="sub in item.child">{{ sub.title }}</a-menu-item>
        </a-sub-menu>
      </template>

    </a-menu>
  </div>

</template>


<script setup lang="ts">
import {ref, watch} from "vue";
import type {Component} from "vue";
import {useRoute, useRouter} from "vue-router";
import type {RouteMeta} from "vue-router";
import {IconMenu, IconUser} from "@arco-design/web-vue/es/icon";

const router = useRouter()
const route = useRoute()

interface MetaType extends RouteMeta {
  title: string
}

interface MenuType {
  key: string
  title: string
  icon?: Component
  name?: string // 路由名字
  child?: MenuType[]
}

const menuList: MenuType[] = [
  {key: "1", title: "首页", icon: IconMenu, name: "home", child: []},
  {
    key: "2", title: "个人中心", icon: IconUser, name: "user_center", child: [
      {key: "2-1", title: "我的信息", icon: IconUser, name: "user_info"},
    ]
  },
  {
    key: "3", title: "文章管理", icon: IconMenu, name: "article", child: [
      {key: "3-1", title: "文章列表", icon: IconUser, name: "article_list"},
    ]
  },
  {
    key: "4", title: "用户管理", icon: IconMenu, name: "users", child: [
      {key: "4-1", title: "用户列表", icon: IconUser, name: "user_list"},

    ]
  },
  {
    key: "5", title: "群聊管理", icon: IconMenu, name: "chat_group", child: [
      {key: "5-1", title: "聊天记录", icon: IconUser, name: "char_list"},

    ]
  },
  {
    key: "6", title: "系统管理", icon: IconMenu, name: "system", child: [
      {key: "6-1", title: "菜单列表", icon: IconUser, name: "menu_list"},

    ]
  },
]

const selectedKey = ref([route.name])
const openKeys = ref([route.matched[1].name])

function clickMenu(name: string) {
  router.push({name: name})
}

watch(() => route.name, () => {
  selectedKey.value = [route.name]
  openKeys.value = [route.matched[1].name]
})
</script>


<style scoped lang="scss">

</style>