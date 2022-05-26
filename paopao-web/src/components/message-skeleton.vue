<template>
    <div class="skeleton-item" v-for="i in new Array(num)" :key="i">
        <div class="content">
            <n-skeleton text :repeat="2" />
            <n-skeleton text style="width: 60%" />
        </div>
    </div>
</template>

<script setup>
const props = defineProps({
    num: {
        type: Number,
        default: 1,
    },
});
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