<template>
    <div>
        <main-nav title="话题" />

        <n-list class="main-content-wrap tags-wrap" bordered>
            <n-tabs type="line" animated @update:value="changeTab">
                <n-tab-pane name="hot" tab="热门" />
                <n-tab-pane name="new" tab="最新" />
            </n-tabs>
            <n-spin :show="loading">
                <n-space>
                    <n-tag
                        class="tag-item"
                        type="success"
                        round
                        v-for="tag in tags"
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
                        <span class="tag-hot">({{ tag.quote_num }})</span>
                        <template #avatar>
                            <n-avatar :src="tag.user.avatar" />
                        </template>
                    </n-tag>
                </n-space>
            </n-spin>
        </n-list>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { getTags } from '@/api/post';

const tags = ref<Item.TagProps[]>([]);
const tagType = ref<"hot" | "new">('hot');
const loading = ref(false);

const loadTags = () => {
    loading.value = true;
    getTags({
        type: tagType.value,
        num: 50,
    })
        .then((res) => {
            tags.value = res.topics;
            loading.value = false;
        })
        .catch((err) => {
            loading.value = false;
        });
};
const changeTab = (tab: "hot" | "new") => {
    tagType.value = tab;
    loadTags();
};
onMounted(() => {
    loadTags();
});
</script>

<style lang="less" scoped>
.tags-wrap {
    padding: 20px;
    .tag-item {
        .tag-hot {
            margin-left: 12px;
            font-size: 12px;
            opacity: 0.75;
        }
    }
}
</style>