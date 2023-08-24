<template>
    <div>
        <main-nav :title="title" />

        <n-list class="main-content-wrap" bordered>
            <n-list-item>
                <!-- 发布器 -->
                <compose @post-success="onPostSuccess" />
            </n-list-item>

            <div v-if="loading && list.length === 0" class="skeleton-wrap">
                <post-skeleton :num="pageSize" />
            </div>

            <div>
                <div class="empty-wrap" v-if="list.length === 0">
                    <n-empty size="large" description="暂无数据" />
                </div>
                <div v-if="store.state.desktopModelShow">
                    <n-list-item v-for="post in list" :key="post.id">
                        <post-item :post="post" />
                    </n-list-item>
                </div>
                <div v-else>
                    <n-list-item v-for="post in list" :key="post.id">
                        <mobile-post-item :post="post" />
                    </n-list-item>
                </div>
            </div>
        </n-list>


        <n-space v-if="totalPage > 0" justify="center">
            <InfiniteLoading class="load-more" :slots="{ complete: '没有更多泡泡了', error: '加载出错' }" @infinite="nextPage()">
                <template #spinner>
                    <div class="load-more-wrap">
                        <n-spin :size="14" v-if="!noMore" />
                        <span class="load-more-spinner">{{ noMore ? '没有更多泡泡了' : '加载更多' }}</span>
                    </div>
                </template>
            </InfiniteLoading>
        </n-space>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch } from 'vue';
import { useStore } from 'vuex';
import { useRoute, useRouter } from 'vue-router';
import InfiniteLoading from "v3-infinite-loading";
import { getPosts } from '@/api/post';

const store = useStore();
const route = useRoute();
const router = useRouter();

const loading = ref(false);
const noMore = ref(false);
const list = ref<any[]>([]);
const page = ref(1);
const pageSize = ref(20);
const totalPage = ref(0);
const title = computed(() => {
    let t = '泡泡广场';

    if (route.query && route.query.q) {
        if (route.query.t && route.query.t === 'tag') {
            t = '#' + decodeURIComponent(route.query.q as string);
        } else {
            t = '搜索: ' + decodeURIComponent(route.query.q as string);
        }
    }

    return t;
});

const loadPosts = () => {
    loading.value = true;
    getPosts({
        query: route.query.q ? decodeURIComponent(route.query.q as string) : null,
        type: route.query.t as string,
        page: page.value,
        page_size: pageSize.value,
    })
        .then((rsp) => {
            loading.value = false;
            if (rsp.list.length === 0) {
                noMore.value = true
            }

            if (page.value > 1) {
                list.value = list.value.concat(rsp.list);
            } else {
                list.value = rsp.list;
                window.scrollTo(0, 0);
            }

            totalPage.value = Math.ceil(rsp.pager.total_rows / pageSize.value);
        })
        .catch((err) => {
            loading.value = false;
            if (page.value > 1) {
                page.value--
            }
        });
};

const onPostSuccess = (post: Item.PostProps) => {
    // 如果不在第一页，需要跳转到详情页面
    if (page.value != 1) {
        router.push({
            name: 'post',
            query: {
                id: post.id,
            },
        });
        return;
    }

    // 如果实在第一页，就地插入新推文到文章列表中
    let items = [];
    let length = list.value.length;
    if (length == pageSize.value) {
        length--;
    }
    var i = 0;
    for (; i < length; i++) {
        let item: Item.PostProps = list.value[i];
        if (!item.is_top) {
            break;
        }
        items.push(item);
    }
    items.push(post);
    for (; i < length; i++) {
        items.push(list.value[i]);
    }
    list.value = items;
};

const nextPage = () => {
    if (page.value < totalPage.value || totalPage.value == 0) {
        page.value++;
        loadPosts();
    } else {
        noMore.value = true;
    }
};

onMounted(() => {
    loadPosts();
});
watch(
    () => ({
        path: route.path,
        query: route.query,
        refresh: store.state.refresh,
    }),
    (to, from) => {
        if (to.refresh !== from.refresh) {
            noMore.value = false;
            page.value = 1;
            setTimeout(() => {
                loadPosts();
            }, 0);
            return;
        }
        if (from.path !== '/post' && to.path === '/') {
            noMore.value = false;
            page.value = 1;
            setTimeout(() => {
                loadPosts();
            }, 0);
        }
    }
);
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

.dark {

    .main-content-wrap,
    .pagination-wrap,
    .empty-wrap,
    .skeleton-wrap {
        background-color: rgba(16, 16, 20, 0.75);
    }
}
</style>