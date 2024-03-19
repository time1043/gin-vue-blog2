<template>
  <div class="gvb_tabs">
    <swiper :slides-per-view="slidesPerView" class="mySwiper">
      <swiper-slide v-for="(item, index) in tabList" :key="item.name">
        <span
            :class="{gvb_tab:true, active:route.name===item.name}"
            @click="clickTab(item)"
            @click.middle="closeTab(item)">
            {{ item.title }}
            <IconClose @click.stop="closeTab(item)" v-if="item.name!=='home'"></IconClose>
        </span>
      </swiper-slide>
    </swiper>

    <span class="gvb_tab close_all_tab" @click="closeAllTab">全部关闭</span>

  </div>

</template>


<script setup lang="ts">
import {IconClose} from "@arco-design/web-vue/es/icon";
import type {Ref} from "vue";
import {ref, watch, onMounted} from "vue";
import {useRoute, useRouter} from "vue-router";
import {Swiper, SwiperSlide} from "swiper/vue";
import type {tabType} from "@/types";


const route = useRoute()
const router = useRouter();

const tabList: Ref<tabType[]> = ref([
  {name: "home", title: "首页"},
  {name: "user_info", title: "我的信息"},
])

function clickTab(item: tabType) {
  router.push({name: item.name})
}

watch(() => route.name, () => {
  // console.log(route.name)
  if (!inList(route.name as string)) {
    // 当前路由不在列表里面，则添加当前路由到列表里
    tabList.value.push({
      name: route.name as string,
      title: route.meta.title as string
    })
  }
}, {immediate: true})

function inList(name: string): boolean {
  for (const tab of tabList.value) {
    if (tab.name === name) {
      return true
    }
  }
  return false
}

function closeAllTab() {
  tabList.value = [
    {name: "home", title: "首页"},
  ]
  router.push({name: "home"})
}

function closeTab(item: tabType) {
  // home 不关
  if (item.name === "home") {
    return
  }

  // 找当前tab在列表里的索引
  let index = tabList.value.findIndex((tab) => item.name === tab.name)
  // console.log(index)
  tabList.value.splice(index, 1)

  // 当前页面tab
  if (route.name === item.name) {
    let beforeIndex = index - 1 // 首页一定存在
    let beforeItem = tabList.value[beforeIndex]
    clickTab(beforeItem)
  }

  // 其他tab
}

// 持久化存储
function persistence() {
  // console.log(tabList.value)
  localStorage.setItem("tabList", JSON.stringify(tabList.value))
}

watch(() => tabList.value.length, () => {
  persistence()
})

// 加载
function loadTabs() {
  let val = localStorage.getItem("tabList")
  if (val === null) {
    return
  }

  let tabs = []
  try {
    tabs = JSON.parse(val)
  } catch (e) {
    return;
  }
  tabList.value = tabs
}

loadTabs()

const slidesPerView = ref(12);

onMounted(() => {
  // 总宽度 实际宽度
  let mySwiperWith = document.querySelector(".mySwiper")!.clientWidth
  let actualWidth = document.querySelector(".swiper-wrapper")!.scrollWidth

  if (actualWidth < mySwiperWith) {
    return
  }

  let swiperSlideList = document.querySelectorAll(".swiper-wrapper .swiper-slide")
  let sum = 0
  let count = 0
  for (const slide of swiperSlideList) {
    sum += slide.clientWidth
    if (sum > mySwiperWith) {
      break
    }
    count++
  }
  slidesPerView.value = count
})

</script>


<style lang="scss">
.gvb_tabs {
  height: 30px;
  width: 100%;
  border-bottom: 1px solid var(--bg);
  padding: 0 20px;
  display: flex;
  align-items: center;
  position: relative;
  background-color: var(--color-bg-1);

  .mySwiper {
    width: calc(100% - 94px);
    overflow: hidden;
    white-space: nowrap;
    height: 100%;
    display: flex;
    align-items: center;

    .swiper-wrapper {
      display: flex;
      justify-content: start;
      width: 100%;

      .swiper-slide {
        width: auto !important;
      }
    }
  }

  .gvb_tab {
    border-radius: 5px;
    border: 1px solid var(--bg);
    padding: 2px 8px;
    margin-right: 10px;
    cursor: pointer;

    &.active {
      background-color: var(--active);
      color: white;
      border: none;

      svg:hover {
        background-color: rgb(var(--arcoblue-4));
      }
    }

    svg {
      font-size: 12px;
      border-radius: 50%;

      &:hover {
        background-color: var(--bg)
      }
    }
  }

  .close_all_tab {
    position: absolute;
    right: 20px;
    margin-right: 0;
  }
}

</style>