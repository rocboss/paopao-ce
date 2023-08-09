<template>
    <div>
        <main-nav title="泡泡详情" :back="true" />

        <n-list class="main-content-wrap" bordered>
            <n-list-item>
                <n-spin :show="loading">
                    <div class="detail-wrap" v-if="post.id > 1">
                        <post-detail :post="post" @reload="loadPost" />
                    </div>
                    <div class="empty-wrap" v-else>
                        <n-empty size="large" description="暂无数据" />
                    </div>
                </n-spin>
            </n-list-item>
            <div class="comment-opts-wrap" v-if="post.id > 0">
                <n-tabs type="bar" justify-content="end" size="small" animated @update:value="commentTab">
                    <template #prefix>
                        <span class="comment-title-item">评论</span>
                    </template>
                    <n-tab-pane name="default" tab="默认" />
                    <n-tab-pane name="newest" tab="最新" />
                </n-tabs>
            </div>
            <n-list-item v-if="post.id > 0">
                <compose-comment :lock="post.is_lock" :post-id="post.id" @post-success="reloadComments" />
            </n-list-item>

            <div v-if="post.id > 0">
                <div v-if="commentLoading" class="skeleton-wrap">
                    <post-skeleton :num="5" />
                </div>
                <div v-else>
                    <div class="empty-wrap" v-if="comments.length === 0">
                        <n-empty size="large" description="暂无评论，快来抢沙发" />
                    </div>

                    <n-list-item v-for="comment in comments" :key="comment.id">
                        <comment-item :comment="comment" @reload="reloadComments" />
                    </n-list-item>
                </div>
            </div>
            <n-space v-if="comments.length >= pageSize" justify="center">
                <InfiniteLoading class="load-more" :slots="{complete: '没有更多数据了', error: '加载出错'}" @infinite="loadComments">
                    <template #spinner>
                        <span v-if="defaultCommentsSort && defaultNoMore" class="load-more-spinner" ><!-- 注意一定要保留这里 --></span>
                        <span v-if="!defaultCommentsSort && newestNoMore" class="load-more-spinner" ><!-- 注意一定要保留这里 --></span>
                        <span v-if="defaultCommentsSort && !defaultNoMore" class="load-more-spinner" >加载评论</span>
                        <span v-if="!defaultCommentsSort && !newestNoMore" class="load-more-spinner" >加载评论</span>
                    </template>
                </InfiniteLoading>
            </n-space>
        </n-list>
    </div>
</template>

<script setup lang="ts">
import { ref, watch, onMounted, computed } from 'vue';
import { useRoute } from 'vue-router';
import { getPost, getPostComments } from '@/api/post';
import InfiniteLoading from "v3-infinite-loading";
import "v3-infinite-loading/lib/style.css";

const route = useRoute();
const post = ref<Item.PostProps>({} as Item.PostProps);
const loading = ref(false);
const commentLoading = ref(false);
const comments = ref<Item.CommentProps[]>([]);
const postId = computed(() => +(route.query.id as string));
const sortStrategy = ref<"default" | "newest">('default');
const defaultCommentsSort = ref<boolean>(true)
const pageSize = 20

let stateHandler = ({
  loading() {
    //nothing
  },
  loaded() {
    // nothing
  },
  complete() {
    // nothing
  },
  error() {
    // nothing
  },
});

const commentTab = (tab: "default" | "newest") => {
    sortStrategy.value = tab;
    if (tab === "default") {
        defaultCommentsSort.value = true
    }
    loadComments(stateHandler);
};

const loadPost = () => {
    post.value = {
        id: 0,
    } as Item.PostProps;
    loading.value = true;
    getPost({
        id: postId.value,
    })
        .then((res) => {
            loading.value = false;
            post.value = res;

            // 加载评论
            loadComments(stateHandler);
        })
        .catch((err) => {
            loading.value = false;
        });
};

let defaultCommmentsPage = 1;
const defaultNoMore = ref<boolean>(false)
const defaultComments = ref<Item.CommentProps[]>([]);
const loadDefaultComments = ($state: any) => {
    if (defaultNoMore.value) {
        return
    }

    getPostComments({
        id: post.value.id as number,
        sort_strategy: 'default',
        page: defaultCommmentsPage,
        page_size: pageSize,
    })
    .then((res) => {
        if ($state !== null) {
            stateHandler = $state
        }
        if (res.list.length < pageSize) {
            defaultNoMore.value = true
        } else {
            defaultCommmentsPage++
        }
        if (res.list.length > 0)  {
            if (defaultCommmentsPage === 1) {
                defaultComments.value = res.list;
            } else {
                defaultComments.value.push(...res.list);
            }
            comments.value = defaultComments.value
        }
        stateHandler.loaded();
        commentLoading.value = false;
    })
    .catch((err) => {
        commentLoading.value = false;
        stateHandler.error();
    });
};

let newestCommmentsPage = 1;
let newestNoMore = ref<boolean>(false)
const newestComments=ref<Item.CommentProps[]>([]);
const loadNewestComments = ($state: any) => {
    if (newestNoMore.value) {
        return
    }

    getPostComments({
        id: post.value.id as number,
        sort_strategy: 'newest',
        page: newestCommmentsPage,
        page_size: pageSize,
    })
    .then((res) => {
        if ($state !== null) {
            stateHandler = $state
        }
        if (res.list.length < pageSize) {
            newestNoMore.value = true
        } else {
            newestCommmentsPage++
        }
        if (res.list.length > 0) {
            if (newestCommmentsPage === 1) {
                newestComments.value = res.list;
            } else {
                newestComments.value.push(...res.list);
            }
            comments.value = newestComments.value
        }
        stateHandler.loaded();
        commentLoading.value = false;
    })
    .catch((err) => {
        commentLoading.value = false;
        stateHandler.error();
    });
};

const loadComments = ($state: any) => {
    if (postId.value < 1) {
        return
    }
    if (comments.value.length === 0) {
        commentLoading.value = true;
    }
    if (sortStrategy.value === 'default') {
        comments.value = defaultComments.value
        loadDefaultComments($state)
    } else {
        comments.value = newestComments.value
        loadNewestComments($state)
    }
    commentLoading.value = false;
};

const reloadComments = () => {
    // 这里需要做特殊处理,目前暴力处理，一切都重新加载
    // TODO：后续持续优化， 这里有大bug！！！
    defaultCommmentsPage = 1;
    defaultNoMore.value = false
    defaultComments.value = []

    newestCommmentsPage = 1;
    newestNoMore.value = false
    newestComments.value = []

    loadComments(stateHandler)
}

onMounted(() => {
    loadPost();
});

watch(postId, () => {
    if (postId.value > 0 && route.name === 'post') {
        loadPost();
    }
});
</script>

<style lang="less" scoped>
.detail-wrap {
    min-height: 100px;
}

.comment-opts-wrap {
    padding-top: 6px;
    padding-left: 16px;
    padding-right: 16px;
    opacity: 0.75;

    .comment-title-item {
        padding-top: 4px;
        font-size: 16px;
        text-align: center;
    }
}

.main-content-wrap {
    .load-more {
        margin-bottom: 8px;
        .load-more-spinner {
            font-size: 14px;
            opacity: 0.65;
        }
    }
}


.dark {
    .main-content-wrap,
    .skeleton-wrap {
        background-color: rgba(16, 16, 20, 0.75);
    }
}
</style>