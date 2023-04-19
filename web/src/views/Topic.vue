<template>
    <div>
        <main-nav title="话题" />

        <n-list class="main-content-wrap tags-wrap" bordered>
            <n-tabs type="line" animated @update:value="changeTab">
                <n-tab-pane name="hot" tab="热门" />
                <n-tab-pane name="new" tab="最新" />
                <n-tab-pane v-if="store.state.userLogined"
                name="follow" tab="关注" />
            </n-tabs>
            <n-spin :show="loading">
                <n-space>
                    <tag-item
                        v-for="tag in tags"
                        :tag="tag"
                        :showAction="store.state.userLogined"
                    >
                    </tag-item>
                </n-space>
            </n-spin>
        </n-list>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted} from 'vue';
import { getTags } from '@/api/post';
import { useStore } from 'vuex';

const store = useStore();
const tags = ref<Item.TagProps[]>([]);
const tagType = ref<"hot" | "new" | "follow">('hot');
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
const changeTab = (tab: "hot" | "new" | "follow") => {
    tagType.value = tab;
    if (tab == "follow") {
        tagType.value = "hot"
    }
    loadTags();
};
onMounted(() => {
    loadTags();
});
</script>

<style lang="less" scoped>
.tags-wrap {
    padding: 20px;
}
.dark {
    .tags-wrap {
        background-color: rgba(16, 16, 20, 0.75);
    }
}
</style>