<template>
    <div>
        <main-nav :title="title" />

        <n-list class="main-content-wrap" bordered>
            <n-list-item>
                <!-- 发布器 -->
                <compose @post-success="onPostSuccess" />
            </n-list-item>

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
                :page-count="totalPage"
            />
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch } from 'vue';
import { useStore } from 'vuex';
import { useRoute, useRouter } from 'vue-router';
import { getPosts } from '@/api/post';

const store = useStore();
const route = useRoute();
const router = useRouter();

const loading = ref(false);
const list = ref<any[]>([]);
const page = ref(+(route.query.p as string) || 1);
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
            list.value = rsp.list;
            totalPage.value = Math.ceil(rsp.pager.total_rows / pageSize.value);

            window.scrollTo(0, 0);
        })
        .catch((err) => {
            loading.value = false;
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
   for (;i < length; i++) {
        let item: Item.PostProps = list.value[i];
        if (!item.is_top) {
            break;
        }
        items.push(item);  
   }
   items.push(post);
   for (;i < length; i++) {
        items.push(list.value[i]);
   }
   list.value = items;
};

const updatePage = (p: number) => {
    router.push({
        name: 'home',
        query: {
            ...route.query,
            p,
        },
    });
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
            page.value = +(route.query.p as string) || 1;
            setTimeout(() => {
                loadPosts();
            }, 0);
            return;
        }
        if (from.path !== '/post' && to.path === '/') {
            page.value = +(route.query.p as string) || 1;
            setTimeout(() => {
                loadPosts();
            }, 0);
        }
    }
);
</script>

<style lang="less" scoped>
.pagination-wrap {
    padding: 10px;
    display: flex;
    justify-content: center;
    overflow: hidden;
}
.dark {
    .main-content-wrap, .pagination-wrap, .empty-wrap, .skeleton-wrap {
        background-color: rgba(16, 16, 20, 0.75);
    }
}
</style>