<template>
  <h1>前台首页</h1>
  <Gvb_login v-model:visible="visible"></Gvb_login>
  <a-button @click="visible=true">登录</a-button>
  <router-link to="/login">登录</router-link>
  <gvb_theme></gvb_theme>
</template>

<script setup lang="ts">
import Gvb_login from "@/components/common/gvb_login.vue";
import {ref} from "vue";
import Gvb_theme from "@/components/common/gvb_theme.vue";

const visible = ref(false)

</script>

<style lang="scss">

</style>