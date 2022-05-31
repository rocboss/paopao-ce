<template>
    <div>
        <main-nav title="泡泡详情" :back="true" />

        <n-list class="main-content-wrap" bordered>
            <n-list-item>
                <n-spin :show="loading">
                    <div class="detail-wrap" v-if="post.id > 0">
                        <post-detail :post="post" @reload="loadPost" />
                    </div>
                    <div class="empty-wrap" v-else>
                        <n-empty size="large" description="暂无数据" />
                    </div>
                </n-spin>
            </n-list-item>
            <n-list-item v-if="post.id > 0">
                <compose-comment
                    :lock="post.is_lock"
                    :post-id="post.id"
                    @post-success="loadComments(true)"
                />
            </n-list-item>

            <div v-if="post.id > 0">
                <div v-if="commentLoading" class="skeleton-wrap">
                    <post-skeleton :num="5" />
                </div>
                <div v-else>
                    <div class="empty-wrap" v-if="comments.length === 0">
                        <n-empty
                            size="large"
                            description="暂无评论，快来抢沙发"
                        />
                    </div>

                    <n-list-item v-for="comment in comments" :key="comment.id">
                        <comment-item
                            :comment="comment"
                            @reload="loadComments"
                        />
                    </n-list-item>
                </div>
            </div>
        </n-list>
    </div>
</template>

<script setup lang="ts">
import { ref, watch, onMounted, computed, Ref } from 'vue';
import { useRoute } from 'vue-router';
import { getPost, getPostComments } from '@/api/post';

const route = useRoute();
const post = ref<Item.PostProps>({} as Item.PostProps);
const loading = ref(false);
const commentLoading = ref(false);
const comments = ref<Item.CommentProps[]>([]);
const postId = computed(() => +(route.query.id as string));

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
    })
        .then((res) => {
            comments.value = res.list;
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
</style>