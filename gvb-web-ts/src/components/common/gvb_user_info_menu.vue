<template>
  <div class="gvb_user_info_menu">
    <a-dropdown @select="select" class="gvb_dropdown">
      <div class="gvb_user_info_menu_dropdown">
        <img :src="store.userInfo.avatar">
        <span class="gvb_user_info_menu_dropdown_span">{{ store.userInfo.nick_name }}</span>
        <IconMenu></IconMenu>
      </div>
      <template #content>
        <template v-for="(item,index) in dopList" :key="index">
          <a-dgroup v-if="item.type === 'line'"></a-dgroup>
          <a-doption v-else :value="item.name">{{ item.title }}</a-doption>
        </template>

      </template>
    </a-dropdown>
  </div>
</template>


<script setup lang="ts">
import {IconMenu} from "@arco-design/web-vue/es/icon";
import {useStore} from "@/stores/counter";
import type {tabType} from "@/types";
import router from "@/router";

interface dopType extends tabType {
  type?: string
}

const store = useStore()
const dopList: dopType[] = [
  {name: "user_info", title: "个人信息"},
  {name: "", title: "", type: "line"},
  {name: "article_list", title: "文章列表"},
  {name: "user_list", title: "用户列表"},
  {name: "log_list", title: "系统日志"},
  {name: "", title: "", type: "line"},
  {name: "logout", title: "注销退出"},
]

function select(value: any) {
  let val = value as string
  if (val === "logout") {
    return
  }
  router.push({name: val})
}

</script>


<style lang="scss">
.gvb_dropdown {
  .arco-dropdown-group-title {
    margin-top: 0;

    &::after {
      content: "";
      display: block;
      width: 100%;
      height: 1px;
      background-color: var(--bg);
      position: absolute;
      left: 0;
    }
  }
}
</style>