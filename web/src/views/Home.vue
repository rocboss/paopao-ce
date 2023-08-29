<template>
    <div>
        <main-nav :title="title" />

        <n-list class="main-content-wrap" bordered>
            <n-list-item>
                <!-- 发布器 -->
                <compose @post-success="onPostSuccess" />
            </n-list-item>

            <n-list-item v-if="store.state.desktopModelShow && store.state.userInfo.id > 0" >
            <SlideBar v-model="slideBarList" :wheel-blocks="8" :init-blocks="initBlocks" @click="handleBarClick" tag="div" sub-tag="div">
                <template #default="data">
                    <div class="slide-bar-item">
                        <n-badge value="1" :offset="[0, 48]" dot :show="data.slotData.show">
                            <n-avatar
                                round
                                :size="48"
                                :src="data.slotData.avatar"
                                class="slide-bar-item-avatar"
                            />
                        </n-badge>
                        <div class="slide-bar-item-title slide-bar-user-link">
                            <n-ellipsis :line-clamp="2">
                                {{ data.slotData.title }}
                            </n-ellipsis>
                        </div>
                    </div>
                </template>
            </SlideBar>
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
import { getPosts, getContacts } from '@/api/post';
import { getUserPosts } from '@/api/user';
import SlideBar from '@opentiny/vue-slide-bar';
import allTweets from '@/assets/img/all-tweets.png';
import followingTweets from '@/assets/img/following-tweets.jpeg';

const store = useStore();
const route = useRoute();
const router = useRouter();

const initBlocks = ref(9)
const slideBarList = ref<Item.SlideBarItem[]>([
    { title: '全部动态', style: 1, username: '', avatar: allTweets, show: true },
    // { title: '正在关注', style: 2, username: '', avatar: followingTweets, show: true }
]);

const loading = ref(false);
const noMore = ref(false);
const targetStyle = ref<number>(1)
const targetUsername = ref<string>("")
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

const reset = () => {
    loading.value = false;
    noMore.value = false;
    list.value = [];
    page.value = 1;
    totalPage.value = 0;
}

const handleBarClick = (data: Item.SlideBarItem, index: number) => {
    reset();
    targetStyle.value = data.style
    switch (data.style) {
    case 1:
        loadPosts();
        break;
    case 2:
        // todo: add some other logic
        loadPosts();
        break;
    case 21:
        targetUsername.value = data.username;
        loadUserPosts();
        break;
   
    default:
        break;
   }
   slideBarList.value[index].show = false;
};

const loadContacts = () => {
    if (store.state.userInfo.id === 0) {
        loading.value = true;
    }
    getContacts({
        page: 1,
        page_size: 20,
    }).then((res) => {
        var i = 0;
        const list = res.list || []
        for (; i < res.list.length; i++) {
            let item: Item.ContactItemProps = list[i];
            slideBarList.value.push({
                title: item.nickname, 
                style: 21,
                username: item.username,
                avatar: item.avatar, 
                show: false
            });   
        }
    })
    .catch((err) => {
        console.log(err);
    });
};

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

const loadUserPosts = () => {
    loading.value = true;
    getUserPosts({
        username: targetUsername.value,
        style: "post",
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
                list.value = rsp.list || [];
                window.scrollTo(0, 0);
            }
            totalPage.value = Math.ceil(rsp.pager.total_rows / pageSize.value);
        })
        .catch((err) => {
            list.value = [];
            if (page.value > 1) {
                page.value--;
            }
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

const loadMorePosts = () => {
    switch (targetStyle.value) {
    case 1:
        loadPosts();
        break;
    case 2:
        // todo: add some other logic
        loadPosts();
        break;
    case 21:
        loadUserPosts();
        break;
    default:
        break;
   }
};

const nextPage = () => {
    if (page.value < totalPage.value || totalPage.value == 0) {
        noMore.value = false;
        page.value++;
        loadMorePosts();
    } else {
        noMore.value = true;
    }
};

onMounted(() => {
    loadContacts()
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
            reset();
            setTimeout(() => {
                loadMorePosts();
            }, 0);
            return;
        }
        if (from.path !== '/post' && to.path === '/') {
            reset();
            setTimeout(() => {
                loadMorePosts();
            }, 0);
        }
    }
);

</script>

<style lang="less" scoped>

.tiny-slide-bar .tiny-slide-bar__list > 
div.tiny-slide-bar__select .slide-bar-item .slide-bar-item-title {
    color: #18a058;
    opacity: 0.8;
}

.tiny-slide-bar .tiny-slide-bar__list > 
div:hover .slide-bar-item {
    cursor: pointer;
    .slide-bar-item-avatar {
        opacity: 0.8;
    }
    .slide-bar-item-title {
        color: #18a058;
        opacity: 0.8;
    }
}

.tiny-slide-bar {
    margin-top: -30px;
    margin-bottom: -30px;
    .slide-bar-item {
        min-height: 170px;
        width: 64px;
        display: flex;
        flex-direction:column;
        justify-content: center;
        align-items: center;
        margin-top: 8px;
        .slide-bar-item-title {
            justify-content: center;
            font-size: 12px;
            margin-top: 4px;
            height: 40px;
        }
    }
}

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
    .tiny-slide-bar .tiny-slide-bar__list > 
    div.tiny-slide-bar__select .slide-bar-item .slide-bar-item-title {
        color: #63e2b7;
        opacity: 0.8;
    }

    .tiny-slide-bar .tiny-slide-bar__list > 
    div:hover .slide-bar-item .slide-bar-item-title {
        color: #63e2b7;
        opacity: 0.8;
    }
}
</style>