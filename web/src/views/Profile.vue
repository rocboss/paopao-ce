<template>
    <div>
        <main-nav title="主页" />

        <n-list
            class="main-content-wrap profile-wrap"
            bordered
            v-if="store.state.userInfo.id > 0"
        >
            <!-- 基础信息 -->
            <div class="profile-baseinfo">
                <div class="avatar">
                    <n-avatar size="large" :src="store.state.userInfo.avatar" />
                </div>
                <div class="base-info">
                    <div class="username">
                        <strong>{{ store.state.userInfo.nickname }}</strong>
                        <span> @{{ store.state.userInfo.username }} </span>
                    </div>
                    <div class="uid">UID. {{ store.state.userInfo.id }}</div>
                </div>
            </div>
            <n-tabs class="profile-tabs-wrap" type="line" animated @update:value="changeTab">
                <n-tab-pane name="post" tab="泡泡"> </n-tab-pane>
                <n-tab-pane name="comment" tab="评论"> </n-tab-pane>
                <n-tab-pane name="media" tab="图文"> </n-tab-pane>
                <n-tab-pane name="star" tab="喜欢"> </n-tab-pane>
            </n-tabs>
            <div v-if="loading" class="skeleton-wrap">
                <post-skeleton :num="pageSize" />
            </div>
            <div v-else>
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
import { ref, onMounted, watch } from 'vue';
import { useStore } from 'vuex';
import { useRoute } from 'vue-router';
import { getUserPosts } from '@/api/user';

const store = useStore();
const route = useRoute();

const loading = ref(false);
const list = ref<Item.PostProps[]>([]);
const pageType = ref<"post" | "comment" | "media" | "star">('post');
const postPage = ref(+(route.query.p as string) || 1);
const commentPage = ref(1)
const mediaPage = ref(1)
const starPage = ref(1);
const page = ref(+(route.query.p as string) || 1);
const pageSize = ref(20);
const totalPage = ref(0);

const loadPage = () => {
    switch(pageType.value) {
        case "post":
            loadPosts();
            break;
        case "comment":
            loadCommentPosts();
            break;
        case "media":
            loadMediaPosts();
            break;
        case "star":
            loadStarPosts();
            break;
    } 
};
const loadPosts = () => {
    loading.value = true;
    getUserPosts({
        username: store.state.userInfo.username,
        style: "post",
        page: page.value,
        page_size: pageSize.value,
    })
        .then((rsp) => {
            loading.value = false;
            list.value = rsp.list || [];
            totalPage.value = Math.ceil(rsp.pager.total_rows / pageSize.value);
            window.scrollTo(0, 0);
        })
        .catch((err) => {
            list.value = []
            loading.value = false;
        });
};
const loadCommentPosts = () => {
    loading.value = true;
    getUserPosts({
        username: store.state.userInfo.username,
        style: "comment",
        page: page.value,
        page_size: pageSize.value,
    })
        .then((rsp) => {
            loading.value = false;
            list.value = rsp.list || [];
            totalPage.value = Math.ceil(rsp.pager.total_rows / pageSize.value);
            window.scrollTo(0, 0);
        })
        .catch((err) => {
            list.value = []
            loading.value = false;
        });
};
const loadMediaPosts = () => {
    loading.value = true;
    getUserPosts({
        username: store.state.userInfo.username,
        style: "media",
        page: page.value,
        page_size: pageSize.value,
    })
        .then((rsp) => {
            loading.value = false;
            list.value = rsp.list || [];
            totalPage.value = Math.ceil(rsp.pager.total_rows / pageSize.value);
            window.scrollTo(0, 0);
        })
        .catch((err) => {
            list.value = []
            loading.value = false;
        });
};
const loadStarPosts = () => {
    loading.value = true;
    getUserPosts({
        username: store.state.userInfo.username,
        style: "star",
        page: page.value,
        page_size: pageSize.value,
    })
        .then((rsp) => {
            loading.value = false;
            list.value = rsp.list || [];
            totalPage.value = Math.ceil(rsp.pager.total_rows / pageSize.value);
            window.scrollTo(0, 0);
        })
        .catch((err) => {
            list.value = []
            loading.value = false;
        });
};
const changeTab = (tab: "post" | "comment" | "media" | "star") => {
    pageType.value = tab;
    switch(pageType.value) {
        case "post":
            page.value = postPage.value
            loadPosts();
            break;
        case "comment":
            page.value = commentPage.value
            loadCommentPosts();
            break;
        case "media":
            page.value = mediaPage.value
            loadMediaPosts();
            break;
        case "star":
            page.value = starPage.value
            loadStarPosts();
            break;
    }
};
const updatePage = (p: number) => {
    page.value = p;
    switch(pageType.value) {
        case "post":
            postPage.value = p
            loadPosts();
            break;
        case "comment":
            commentPage.value = page.value
            loadCommentPosts();
            break;
        case "media":
            mediaPage.value = page.value
            loadMediaPosts();
            break;
        case "star":
            starPage.value = page.value
            loadStarPosts();
            break;
    }
};
onMounted(() => {
    loadPage();
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
                loadPage();
            }, 0);
            return;
        }
        if (from.path !== '/post' && to.path === '/profile') {
            page.value = +(route.query.p as string) || 1;
            setTimeout(() => {
                loadPage();
            }, 0);
        }
    }
);
</script>

<style lang="less" scoped>
.profile-baseinfo {
    display: flex;
    padding: 16px;
    .avatar {
        width: 55px;
    }

    .base-info {
        position: relative;
        width: calc(100% - 55px);

        .username {
            line-height: 16px;
            font-size: 16px;
        }

        .uid {
            font-size: 14px;
            line-height: 14px;
            margin-top: 10px;
            opacity: 0.75;
        }
    }
}

.profile-tabs-wrap {
    padding: 0 16px;
}

.pagination-wrap {
    padding: 10px;
    display: flex;
    justify-content: center;
    overflow: hidden;
}
.dark {
    .profile-baseinfo {
        background-color: #18181c;
    }
    .profile-wrap, .pagination-wrap {
        background-color: rgba(16, 16, 20, 0.75);
    }
}
</style>