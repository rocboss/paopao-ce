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

                <div v-if="store.state.desktopModelShow">
                    <n-list-item v-for="post in list" :key="post.id">
                        <post-item :post="post" 
                            :isOwner="store.state.userInfo.id == post.user_id" 
                            :addFollowAction="true"
                            @send-whisper="onSendWhisper"
                            @handle-follow-action="onHandleFollowAction" />
                    </n-list-item>
                </div>
                <div v-else>
                    <n-list-item v-for="post in list" :key="post.id">
                        <mobile-post-item :post="post"
                            :isOwner="store.state.userInfo.id == post.user_id" 
                            :addFollowAction="true"
                            @send-whisper="onSendWhisper"
                            @handle-follow-action="onHandleFollowAction" />
                    </n-list-item>
                </div>
            </div>
            <!-- 私信组件 -->
            <whisper :show="showWhisper" :user="whisperReceiver" @success="whisperSuccess" />
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
import { ref, onMounted } from 'vue';
import { useStore } from 'vuex';
import { useRoute } from 'vue-router';
import { useDialog } from 'naive-ui';
import { getCollections, followUser, unfollowUser } from '@/api/user';

const store = useStore();
const route = useRoute();
const dialog = useDialog();

const loading = ref(false);
const list = ref<any[]>([]);
const page = ref(+(route.query.p as any) || 1);
const pageSize = ref(20);
const totalPage = ref(0);
const showWhisper = ref(false);
const whisperReceiver = ref<Item.UserInfo>({
    id: 0,
    avatar: '',
    username: '',
    nickname: '',
    is_admin: false,
    is_friend: true,
    is_following: false,
    created_on: 0,
    follows: 0,
    followings: 0,
    status: 1,
});

const onSendWhisper =  (user: Item.UserInfo) => {
    whisperReceiver.value = user;
    showWhisper.value = true;
};

const whisperSuccess = () => {
    showWhisper.value = false;
};

const onHandleFollowAction = (post: Item.PostProps) => {
    dialog.success({
        title: '提示',
        content:
            '确定' + (post.user.is_following ? '取消关注' : '关注') + '该用户吗？',
        positiveText: '确定',
        negativeText: '取消',
        onPositiveClick: () => {
            if (post.user.is_following) {
                unfollowUser({
                    user_id: post.user.id,
                }).then((_res) => {
                    window.$message.success('操作成功');
                    post.user.is_following = false;
                })
                .catch((_err) => {});
            } else {
                followUser({
                    user_id: post.user.id,
                }).then((_res) => {
                    window.$message.success('关注成功');
                    post.user.is_following = true;
                })
                .catch((_err) => {});
            }
        },
    });
};

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