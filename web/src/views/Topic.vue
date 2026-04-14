<template>
    <div>
        <main-nav title="话题" />

        <n-list class="main-content-wrap tags-wrap" bordered>
            <n-tabs type="line" animated @update:value="changeTab">
                <n-tab-pane name="hot" tab="热门" />
                <n-tab-pane name="new" tab="最新" />
                <n-tab-pane name="follow" tab="关注" v-if="userLogined" />
                <n-tab-pane name="pin" tab="钉住" v-if="userLogined" />
                <template v-if="userLogined" #suffix>
                    <n-tag v-model:checked="tagsChecked" checkable>
                        {{tagsEditText}}
                    </n-tag>
                </template>
            </n-tabs>
            <n-spin :show="loading">
                <n-space>
                    <tag-item
                        v-for="tag in tags"
                        :tag="tag"
                        :showAction="userLogined && tagsChecked"
                        :checkFollowing="inFollowTab"
                        :checkPin="inPinTab"
                    >
                    </tag-item>
                </n-space>
                <div class="empty-wrap" v-if="tags.length === 0">
                    <n-empty size="large" description="暂无数据" />
                </div>
            </n-spin>
        </n-list>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch } from 'vue';
import { getTags } from '@/api/post';
import { useStoreMain } from '@/store/main';
import { useStoreUser } from '@/store/user';
import { storeToRefs } from 'pinia';

const storeMain = useStoreMain();
const storeUser = useStoreUser();
const { userLogined } = storeToRefs(storeUser);

const tags = ref<Item.TagProps[]>([]);
const tagType = ref<'hot' | 'new' | 'follow' | 'pin'>('hot');
const loading = ref(false);
const tagsChecked = ref(false);
const inFollowTab = ref(false);
const inPinTab = ref(false);

watch(tagsChecked, () => {
  if (!tagsChecked.value) {
    window.$message.success('保存成功');
    storeMain.doRefreshTopicFollow();
  }
});
const tagsEditText = computed({
  get: () => {
    let text = '编辑';
    if (tagsChecked.value) {
      text = '保存';
    }
    return text;
  },
  set: (newVal) => {
    // do nothing
  },
});
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
      tags.value = [];
      console.log(err);
      loading.value = false;
    });
};
const changeTab = (tab: 'hot' | 'new' | 'follow' | 'pin') => {
  tagType.value = tab;
  inFollowTab.value = tab === 'follow';
  inPinTab.value = tab === 'pin';
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
    .tags-wrap, .empty-wrap {
        background-color: rgba(16, 16, 20, 0.75);
    }
}
</style>