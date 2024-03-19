<template>
  <div>
    <a-modal width="400px" modal-class="gvb_login_modal" :visible="props.visible"
             title="用户登录" :footer="false" @cancel="close">
      <Gvb_login_form></Gvb_login_form>
    </a-modal>
  </div>
</template>


<script setup lang="ts">
import Gvb_login_form from "@/components/common/gvb_login_form.vue";

const props = defineProps({
  visible: {type: Boolean}
})
const emits = defineEmits(["update:visible"])

function close() {
  emits("update:visible", false)
}

</script>


<style lang="scss">
.gvb_login_modal {
  .arco-modal-header {
    display: none;
  }
}
</style>