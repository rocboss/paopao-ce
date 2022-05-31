<template>
  <div>
    <main-nav title="点赞" />

    <n-list class="main-content-wrap" bordered>
      <template #footer>
        <div v-if="totalPage > 0" class="pagination-wrap">
          <n-pagination
            :page="page"
            :page-slot="!store.state.collapsedRight ? 8 : 5"
            :page-count="totalPage"
            @update:page="updatePage"
          />
        </div>
      </template>

      <div v-if="loading" class="skeleton-wrap">
        <post-skeleton :num="pageSize" />
      </div>
      <div v-else>
        <div v-if="list.length === 0" class="empty-wrap">
          <n-empty size="large" description="暂无数据" />
        </div>

        <n-list-item v-for="post in list" :key="post.id">
          <post-item :post="post" />
        </n-list-item>
      </div>
    </n-list>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch } from 'vue';
import { useStore } from 'vuex';
import { useRoute, useRouter } from 'vue-router';
import { getStars } from '@/api/user';

const store = useStore();
const route = useRoute();
const router = useRouter();

const loading = ref(false);
const list = ref<Item.PostProps[]>([]);
const page = ref(+(route.query.p as string) || 1);
const pageSize = ref(20);
const totalPage = ref(0);

const loadPosts = () => {
  loading.value = true;
  getStars({
    page: page.value,
    page_size: pageSize.value,
  })
    .then((rsp) => {
      loading.value = false;
      list.value = rsp.list;
      totalPage.value = Math.ceil(rsp.pager.total_rows / pageSize.value);

      window.scrollTo(0, 0);
    })
    .catch((err) => {
      loading.value = false;
    });
};
const updatePage = (p: number) => {
  page.value = p;
  loadPosts();
};
onMounted(() => {
  loadPosts();
});
</script>

<style lang="less" scoped>
.pagination-wrap {
  padding: 10px;
  display: flex;
  justify-content: center;
  overflow: hidden;
}
</style>
