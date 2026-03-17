<template>
    <n-space v-if="totalPage > 0" justify="center">
        <InfiniteLoading
            class="load-more"
            :slots="{ complete: completeText, error: '加载出错' }"
            @infinite="handleInfinite"
        >
            <template #spinner>
                <div class="load-more-wrap">
                    <n-spin :size="14" v-if="!noMore" />
                    <span class="load-more-spinner">{{ noMore ? completeText : '加载更多' }}</span>
                </div>
            </template>
        </InfiniteLoading>
    </n-space>
</template>

<script setup lang="ts">
import InfiniteLoading from 'v3-infinite-loading';

withDefaults(defineProps<{
    totalPage: number;
    noMore: boolean;
    completeText?: string;
}>(), {
    completeText: '没有更多了',
});

const emit = defineEmits<{
    (e: 'load-more'): void;
}>();

const handleInfinite = () => {
    emit('load-more');
};
</script>

<style lang="less" scoped>
.load-more {
    margin: 20px;

    .load-more-wrap {
        display: flex;
        flex-direction: row;
        justify-content: center;
        align-items: center;
        gap: 14px;

        .load-more-spinner {
            font-size: 14px;
            opacity: 0.65;
        }
    }
}
</style>
