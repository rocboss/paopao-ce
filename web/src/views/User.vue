<template>
    <div>
        <main-nav title="用户详情" />

        <n-list class="main-content-wrap profile-wrap" bordered>
            <!-- 基础信息 -->
            <n-spin :show="userLoading">
                <div class="profile-baseinfo" v-if="user.id > 0">
                    <div class="avatar">
                        <n-avatar size="large" :src="user.avatar" />
                    </div>
                    <div class="base-info">
                        <div class="username">
                            <strong>{{ user.nickname }}</strong>
                            <span> @{{ user.username }} </span>
                        </div>
                        <div class="uid">UID. {{ user.id }}</div>
                    </div>

                    <div class="user-opts">
                        <n-space vertical>
                            <n-button
                                size="small"
                                secondary
                                type="primary"
                                @click="openWhisper"
                            >
                                私信
                            </n-button>
                        </n-space>
                    </div>
                </div>

                <!-- 私信组件 -->
                <whisper
                    :show="showWhisper"
                    :user="user"
                    @success="whisperSuccess"
                />
            </n-spin>
            <template #footer>
                <div class="pagination-wrap" v-if="totalPage > 0">
                    <n-pagination
                        :page="page"
                        @update:page="updatePage"
                        :page-slot="!store.state.collapsedRight ? 8 : 5"
                        :page-count="totalPage"
                    />
                </div>
            </template>
            <n-tabs class="profile-tabs-wrap" animated>
                <n-tab-pane name="post" tab="泡泡" />
                <!-- <n-tab-pane name="comment" tab="评论"> </n-tab-pane> -->
            </n-tabs>
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
    </div>
</template>

<script setup lang="ts">
import { ref, reactive, watch, onMounted } from 'vue';
import { useStore } from 'vuex';
import { useRoute } from 'vue-router';
import { getUserProfile, getUserPosts } from '@/api/user';

const store = useStore();
const route = useRoute();

const loading = ref(false);
const user = reactive({
    id: 0,
    avatar: '',
    username: '',
    nickname: '',
});
const userLoading = ref(false);
const showWhisper = ref(false);
const list = ref<Item.PostProps[]>([]);
const username = ref(route.query.username || '');
const page = ref(+(route.query.p as string) || 1);
const pageSize = ref(20);
const totalPage = ref(0);

const loadPosts = () => {
    loading.value = true;
    getUserPosts({
        username: username.value as string,
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
const loadUser = () => {
    userLoading.value = true;
    getUserProfile({
        username: username.value as string,
    })
        .then((res) => {
            userLoading.value = false;
            user.id = res.id;
            user.avatar = res.avatar;
            user.username = res.username;
            user.nickname = res.nickname;
            loadPosts();
        })
        .catch((err) => {
            userLoading.value = false;
            console.log(err);
        });
};

const updatePage = (p: number) => {
    page.value = p;
    loadPosts();
};

const openWhisper = () => {
    // window.$message.warning('您尚未获得私信权限');
    showWhisper.value = true;
};
const whisperSuccess = () => {
    showWhisper.value = false;
};
watch(
    () => ({
        path: route.path,
        query: route.query,
    }),
    (to, from) => {
        if (from.path === '/user' && to.path === '/user') {
            username.value = route.query.username || '';
            loadUser();
        }
    }
);
onMounted(() => {
    loadUser();
});
</script>

<style lang="less" scoped>
.profile-tabs-wrap {
    padding: 0 16px;
}
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

.pagination-wrap {
    padding: 10px;
    display: flex;
    justify-content: center;
    overflow: hidden;
}
</style>