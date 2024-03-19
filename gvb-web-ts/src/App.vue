<template>
  <RouterView/>
</template>

<script setup lang="ts">
import {RouterView} from 'vue-router'
import {useStore} from "@/stores/counter";

const store = useStore()
store.loadTheme()

</script>
