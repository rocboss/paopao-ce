<template>
    <div class="link-wrap">
        <div class="link-item" v-for="link in props.links" :key="link.id">
            <n-icon class="hash-link"><link-outline /></n-icon>
            <div class="link-txt-wrap">
                <a :href="link.content" class="hash-link" target="_blank" @click.stop>
                    <span class="link-txt">{{ link.content }}</span>  
                </a>
            </div>
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
    position: relative;
    .link-item {
        height: 22px;
        display: flex;
        align-items: center;
        position: relative;
        .link-txt-wrap {
            left: calc(1em + 4px);
            width: calc(100% - 1em);
            overflow: hidden;
            white-space: nowrap;
            text-overflow: ellipsis;
            position: absolute;
            .hash-link {
                .link-txt {
                    word-break: break-all;
                }
            }
        }
    }
}
</style>