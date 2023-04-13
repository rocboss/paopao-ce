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
                            <n-tag
                                v-if="store.state.userInfo.id > 0 && store.state.userInfo.username != user.username && user.is_friend"
                                class="top-tag" type="info" size="small" round>
                                好友
                            </n-tag>
                            <n-tag v-if="user.is_admin" class="top-tag" type="error" size="small" round>
                                管理员
                            </n-tag>
                        </div>
                        <div class="uid">UID. {{ user.id }}</div>
                    </div>

                    <div class="user-opts"
                        v-if="store.state.userInfo.id > 0 && store.state.userInfo.username != user.username">
                        <n-dropdown placement="bottom-end" trigger="click" size="small" :options="userOptions"
                            @select="handleUserAction">
                            <n-button quaternary circle>
                                <template #icon>
                                    <n-icon>
                                        <more-horiz-filled />
                                    </n-icon>
                                </template>
                            </n-button>
                        </n-dropdown>
                    </div>
                </div>

                <!-- 私信组件 -->
                <whisper :show="showWhisper" :user="user" @success="whisperSuccess" />

                <!-- 加好友组件 -->
                <whisper-add-friend :show="showAddFriendWhisper" :user="user" @success="addFriendWhisperSuccess" />
            </n-spin>
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
import { ref, reactive, watch, onMounted, computed } from 'vue';
import { useStore } from 'vuex';
import { useRoute } from 'vue-router';
import { getUserProfile, getUserPosts, changeUserStatus, deleteFriend } from '@/api/user';
import { useDialog, useMessage, DropdownOption } from 'naive-ui';
import WhisperAddFriend from '../components/whisper-add-friend.vue';
import { MoreHorizFilled } from '@vicons/material';
import { VisibilityEnum } from '@/utils/IEnum';

const message = useMessage();
const dialog = useDialog();
const store = useStore();
const route = useRoute();

const loading = ref(false);
const user = reactive<Item.UserInfo>({
    id: 0,
    avatar: '',
    username: '',
    nickname: '',
    is_admin: false,
    is_friend: true,
    status: 1,
});
const userLoading = ref(false);
const showWhisper = ref(false);
const showAddFriendWhisper = ref(false);
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
            user.is_admin = res.is_admin;
            user.is_friend = res.is_friend;
            user.status = res.status;
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
    showWhisper.value = true;
};
const openAddFriendWhisper = () => {
    showAddFriendWhisper.value = true;
};
const whisperSuccess = () => {
    showWhisper.value = false;
};
const addFriendWhisperSuccess = () => {
    showAddFriendWhisper.value = false;
};
const userOptions = computed(() => {
    let options: DropdownOption[] = [{
        label: '私信',
        key: 'whisper',
    }];
    if (store.state.userInfo.is_admin) {
        if (user.status === 1) {
            options.push({
                label: '禁言',
                key: 'banned',
            });
        } else {
            options.push({
                label: '解封',
                key: 'deblocking',
            });
        }
    }
    if (user.is_friend) {
        options.push({
            label: '删除好友',
            key: 'delete',
        });
    } else {
        options.push({
            label: '添加朋友',
            key: 'requesting',
        });
    }
    return options;
});
const handleUserAction = (
    item: 'whisper' | 'delete' | 'requesting' | 'banned' | 'deblocking'
) => {
    switch (item) {
        case 'whisper':
            openWhisper();
            break;
        case 'delete':
            openDeleteFriend();
            break;
        case 'requesting':
            openAddFriendWhisper();
            break;
        case 'banned':
        case 'deblocking':
            banUser();
            break;
        default:
            break;
    }
};
const openDeleteFriend = () => {
    dialog.warning({
        title: '删除好友',
        content: '将好友 “' + user.nickname + '” 删除，将同时删除 点赞/收藏 列表中关于该朋友的 “好友可见” 推文',
        positiveText: '确定',
        negativeText: '取消',
        onPositiveClick: () => {
            userLoading.value = true;
            deleteFriend({
                user_id: user.id,
            })
                .then((res) => {
                    userLoading.value = false;
                    user.is_friend = false
                    loadPosts();
                })
                .catch((err) => {
                    userLoading.value = false;
                    console.log(err);
                });
        },
    });
};
const banUser = () => {
    dialog.warning({
        title: '警告',
        content:
            '确定对该用户进行' +
            (user.status === 1 ? '禁言' : '解封') +
            '处理吗？',
        positiveText: '确定',
        negativeText: '取消',
        onPositiveClick: () => {
            userLoading.value = true;
            changeUserStatus({
                id: user.id,
                status: user.status === 1 ? 2 : 1,
            })
                .then((res) => {
                    userLoading.value = false;
                    loadUser();
                })
                .catch((err) => {
                    userLoading.value = false;
                    console.log(err);
                });
        },
    });
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

        .top-tag {
            transform: scale(0.75);
        }
    }

    .user-opts {
        position: absolute;
        top: 16px;
        right: 16px;
        opacity: 0.75;
    }
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