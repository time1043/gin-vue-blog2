<template>
  <div class="gvb_bread_crumbs">
    <a-breadcrumb>
      <a-breadcrumb-item v-for="item in route.matched">{{ (item.meta as MetaType).title }}</a-breadcrumb-item>
    </a-breadcrumb>
  </div>

</template>


<script setup lang="ts">
import type {RouteMeta} from "vue-router";
import {useRoute} from "vue-router";

const route = useRoute()

interface MetaType extends RouteMeta {
  title: string
}
</script>


<style scoped lang="scss">

</style>