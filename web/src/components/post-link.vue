<template>
    <div class="link-wrap">
        <div class="link-item" v-for="link in props.links" :key="link.id">
            <n-icon class="hash-link"><link-outline /></n-icon>
            <a
                :href="link.content"
                class="hash-link"
                target="_blank"
                @click.stop
            >
                <span class="link-txt">{{ link.content }}</span>
            </a>
        </div>
    </div>
</template>

<script setup lang="ts">
import { LinkOutline } from '@vicons/ionicons5';
const props = withDefaults(defineProps<{
    links: Item.PostItemProps[]
}>(), {
    links: () => []
});
</script>

<style lang="less" scoped>
.link-wrap {
    margin-bottom: 10px;
    .link-item {
        display: flex;
        align-items: center;
        .hash-link {
            .link-txt {
                margin-left: 4px;
                word-break: break-all;
            }
        }
    }
}
</style>