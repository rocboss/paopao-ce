<template>
  <div v-for="i in new Array(num)" :key="i" class="skeleton-item">
    <div class="content">
      <n-skeleton text :repeat="2" />
      <n-skeleton text style="width: 60%" />
    </div>
  </div>
</template>

<script setup lang="ts">
const props = withDefaults(
  defineProps<{
    num: number;
  }>(),
  {
    num: 1,
  },
);
</script>

<style lang="less" scoped>
.skeleton-item {
  padding: 12px;
  display: flex;

  .content {
    width: 100%;
  }
}
</style>
