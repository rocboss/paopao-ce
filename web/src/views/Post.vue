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
                <compose-comment :lock="post.is_lock" :post-id="post.id" @post-success="loadComments(true)" />
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
                        <comment-item :comment="comment" @reload="loadComments" />
                    </n-list-item>
                </div>
            </div>

            <div class="load-more-ele" v-if="!noMore" ref="bottomElement">
                加载更多...
            </div>
        </n-list>
    </div>
</template>

<script setup lang="ts">
import { ref, watch, onMounted, onUnmounted, computed, Ref } from 'vue';
import { useRoute } from 'vue-router';
import { getPost, getPostComments } from '@/api/post';

const route = useRoute();
const post = ref<Item.PostProps>({} as Item.PostProps);
const loading = ref(false);
const commentLoading = ref(false);
const comments = ref<Item.CommentProps[]>([]);
const postId = computed(() => +(route.query.id as string));
const sortStrategy = ref<"default" | "newest">('default');
const bottomElement = ref<HTMLElement | null>(null);
const page = ref<number>(1);
const noMore = ref(false);

const commentTab = (tab: "default" | "newest") => {
    sortStrategy.value = tab;
    loadComments();
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
            loadComments();
        })
        .catch((err) => {
            loading.value = false;
        });
};
const loadComments = (scrollToBottom: boolean = false) => {
    if (comments.value.length === 0) {
        commentLoading.value = true;
    }
    getPostComments({
        id: post.value.id as number,
        sort_strategy: sortStrategy.value,
        page: page.value,
        page_size: 20
    })
        .then((res) => {
            if (res.list.length === 0) {
                noMore.value = true
            }

            if (page.value === 1) {
                comments.value = res.list;
            } else {
                comments.value = comments.value.concat(res.list);
            }
            commentLoading.value = false;

            if (scrollToBottom) {
                setTimeout(() => {
                    window.scrollTo(0, 99999);
                }, 50);
            }
        })
        .catch((err) => {
            commentLoading.value = false;
        });
};

const loadMoreComments = () => {
    if (!commentLoading.value && comments.value.length > 0) {
        page.value = page.value + 1;
        loadComments();
    }
};

const observer = new IntersectionObserver((entries: IntersectionObserverEntry[]) => {
    entries.forEach(entry => {
        if (entry.isIntersecting) {
            loadMoreComments();
        }
    });
}, {
    root: null,
    rootMargin: '0px',
    threshold: 1
});

onMounted(() => {
    if (bottomElement.value) {
        observer.observe(bottomElement.value);
    }

    loadPost();
});

onUnmounted(() => {
    observer.disconnect()
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

.load-more-ele {
    font-size: 12px;
    height: 40px;
    display: flex;
    align-items: center;
    justify-content: center;
}

.dark {

    .main-content-wrap,
    .skeleton-wrap {
        background-color: rgba(16, 16, 20, 0.75);
    }
}
</style>