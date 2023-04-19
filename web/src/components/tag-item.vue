<template>
    <div class="tag-item">
        <n-thing>
            <template #header>
                <n-tag
                        type="success"
                        size="large"
                        round
                        :key="tag.id"
                    >
                        <router-link
                            class="hash-link"
                            :to="{
                                name: 'home',
                                query: {
                                    q: tag.tag,
                                    t: 'tag',
                                },
                            }"
                        >
                            #{{ tag.tag }}
                        </router-link>
                        <span v-if="!showAction" class="tag-quote">({{ tag.quote_num }})</span>
                        <span v-if="showAction" class="tag-quote tag-follow">({{ tag.quote_num }})</span>
                        <template #avatar>
                            <n-avatar :src="tag.user.avatar" />
                        </template>
                    </n-tag>
            </template>
            <template #header-extra>
                <div 
                    v-if="showAction" 
                    class="options">
                    <n-dropdown
                        placement="bottom-end"
                        trigger="click"
                        size="small"
                        :options="tagOptions"
                        @select="handleTagAction"
                    >
                        <n-button type="success" quaternary circle block>
                            <template #icon>
                                <n-icon>
                                    <more-vert-outlined />
                                </n-icon>
                            </template>
                        </n-button>
                    </n-dropdown>
                </div>
            </template>
        </n-thing>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { MoreVertOutlined } from '@vicons/material';
import type { DropdownOption } from 'naive-ui';

const hasFollowing= ref(false);
const props = withDefaults(
    defineProps<{
        tag: Item.TagProps;
        showAction: boolean;
    }>(),
    {}
);

const tagOptions = computed(() => {
    let options: DropdownOption[] = [];
    if (props.tag.is_following == 0) {
        options.push({
            label: '关注',
            key: 'follow',
        });
    } else {
        options.push({
            label: '取消关注',
            key: 'unfollow',
        });
    }
    return options;
});

const handleTagAction = (
    item: 'follow' | 'unfollow'
) => {
    switch (item) {
        case 'follow':
            window.$message.success(`关注话题 - ${props.tag.tag}`);
            break;
        case 'unfollow':
            window.$message.success(`取消关注话题 - ${props.tag.tag}`);
            break;
        default:
            break;
    }
};
const execFollowAction = () => {
    // TODO
    window.$message.success('修改follow/unfllow');
};

onMounted(() => {
    hasFollowing.value = false
});
</script>

<style lang="less">
.tag-item {
    .tag-quote {
        margin-left: 12px;
        font-size: 14px;
        opacity: 0.75;
    }
    .tag-follow {
        margin-right: 22px;
    }
    .options {
        margin-left: -32px;
        margin-bottom: 4px;
        opacity: 0.55;
    }
    .n-thing {
        .n-thing-header {
            margin-bottom: 0px;
        }
        .n-thing-avatar-header-wrapper {
            align-items: center;
        }
    }
}
</style>