<template>
    <div>
        <main-nav title="收藏" />

        <n-list class="main-content-wrap" bordered>
            <div v-if="loading" class="skeleton-wrap">
                <post-skeleton :num="pageSize" />
            </div>
            <div v-else>
                <div class="empty-wrap" v-if="list.length === 0">
                    <n-empty size="large" description="暂无数据" />
                </div>

                <n-list-item v-for="post in list" :key="post.id">
                    <post-item :post="post" />
                </n-list-item>
            </div>
        </n-list>

        <div class="pagination-wrap" v-if="totalPage > 0">
            <n-pagination
            :page="page"
            @update:page="updatePage"
            :page-slot="!store.state.collapsedRight ? 8 : 5"
            :page-count="totalPage" />
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch } from 'vue';
import { useStore } from 'vuex';
import { useRoute, useRouter } from 'vue-router';
import { getCollections } from '@/api/user';

const store = useStore();
const route = useRoute();
const router = useRouter();

const loading = ref(false);
const list = ref<any[]>([]);
const page = ref(+(route.query.p as any) || 1);
const pageSize = ref(20);
const totalPage = ref(0);

const loadPosts = () => {
    loading.value = true;
    getCollections({
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
.dark {
    .main-content-wrap, .empty-wrap, .skeleton-wrap {
        background-color: rgba(16, 16, 20, 0.75);
    }
}
</style>