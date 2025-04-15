<template>
    <div>
        <main-nav title="主页" />

        <n-list
            class="main-content-wrap profile-wrap"
            bordered
            v-if="store.state.userInfo.id > 0"
        >
            <!-- 基础信息 -->
            <!-- <n-spin :show="false" > -->
            <div class="profile-baseinfo">
                <div class="avatar">
                    <n-avatar :size="72" :src="store.state.userInfo.avatar" />
                </div>
                <div class="base-info">
                    <div class="username">
                        <strong>{{ store.state.userInfo.nickname }}</strong>
                        <span> @{{ store.state.userInfo.username }} </span>
                        <n-tag v-if="store.state.userInfo.is_admin" class="top-tag" type="error" size="small" round>
                            管理员
                        </n-tag>
                    </div>
                    <div class="userinfo">
                        <span class="info-item">UID. {{ store.state.userInfo.id }} </span>
                        <span class="info-item">{{ formatDate(store.state.userInfo.created_on) }}&nbsp;加入</span>
                    </div>
                    <div class="userinfo">
                        <span class="info-item">
                             <router-link
                                @click.stop
                                class="following-link"
                                :to="{
                                    name: 'following',
                                    query: { 
                                        s: store.state.userInfo.username, 
                                        n: store.state.userInfo.nickname,
                                        t: 'follows',
                                    },
                                }"
                            >
                                关注&nbsp;&nbsp;{{ prettyQuoteNum(store.state.userInfo.follows) }}
                            </router-link>
                        </span>
                        <span class="info-item">
                            <router-link
                                @click.stop
                                class="following-link"
                                :to="{
                                    name: 'following',
                                    query: { 
                                        s: store.state.userInfo.username, 
                                        n: store.state.userInfo.nickname,
                                        t: 'followings',
                                    },
                                }"
                            >
                                粉丝&nbsp;&nbsp;{{ prettyQuoteNum(store.state.userInfo.followings) }}
                            </router-link>
                        </span>
                        <span class="info-item">
                            泡泡&nbsp;&nbsp;{{ prettyQuoteNum(store.state.userInfo.tweets_count) }}
                        </span>
                    </div>
                </div>

                <div class="user-opts">
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
            <!-- </n-spin> -->
            <n-tabs class="profile-tabs-wrap" type="line" animated @update:value="changeTab">
                <n-tab-pane name="post" tab="泡泡"></n-tab-pane>
                <n-tab-pane name="comment" tab="评论"></n-tab-pane>
                <n-tab-pane name="highlight" tab="亮点"></n-tab-pane>
                <n-tab-pane name="media" tab="图文"></n-tab-pane>
                <n-tab-pane name="star" tab="喜欢"></n-tab-pane>
            </n-tabs>
            <div v-if="loading && list.length === 0" class="skeleton-wrap">
                <post-skeleton :num="pageSize" />
            </div>
            <div v-else>
                <div class="empty-wrap" v-if="list.length === 0">
                    <n-empty size="large" description="暂无数据" />
                </div>
                <div v-if="store.state.desktopModelShow">
                    <div v-if="pageType === 'post'">
                        <n-list-item v-for="post in postList" :key="post.id">
                            <post-item :post="post" 
                                :isOwner="store.state.userInfo.id == post.user_id" 
                                :addFollowAction="true"
                                @send-whisper="onSendWhisper"
                                @handle-follow-action="onHandleFollowAction" />
                        </n-list-item>
                    </div>
                    <div v-if="pageType === 'comment'">
                        <n-list-item v-for="post in commentList" :key="post.id">
                            <post-item :post="post" 
                                :isOwner="store.state.userInfo.id == post.user_id" 
                                :addFollowAction="true"
                                @send-whisper="onSendWhisper"
                                @handle-follow-action="onHandleFollowAction" />
                        </n-list-item>
                    </div>
                    <div v-if="pageType === 'highlight'">
                        <n-list-item v-for="post in highlightList" :key="post.id">
                            <post-item :post="post" 
                                :isOwner="store.state.userInfo.id == post.user_id" 
                                :addFollowAction="true"
                                @send-whisper="onSendWhisper"
                                @handle-follow-action="onHandleFollowAction" />
                        </n-list-item>
                    </div>
                    <div v-if="pageType === 'media'">
                        <n-list-item v-for="post in mediaList" :key="post.id">
                            <post-item :post="post" 
                                :isOwner="store.state.userInfo.id == post.user_id" 
                                :addFollowAction="true"
                                @send-whisper="onSendWhisper"
                                @handle-follow-action="onHandleFollowAction" />
                        </n-list-item>
                    </div>
                    <div v-if="pageType === 'star'">
                        <n-list-item v-for="post in starList" :key="post.id">
                            <post-item :post="post" 
                                :isOwner="store.state.userInfo.id == post.user_id" 
                                :addFollowAction="true"
                                @send-whisper="onSendWhisper"
                                @handle-follow-action="onHandleFollowAction" />
                        </n-list-item>
                    </div>
                </div>
                <div v-else>
                    <div v-if="pageType === 'post'">
                        <n-list-item v-for="post in postList" :key="post.id">
                            <mobile-post-item :post="post"
                                :isOwner="store.state.userInfo.id == post.user_id" 
                                :addFollowAction="true"
                                @send-whisper="onSendWhisper"
                                @handle-follow-action="onHandleFollowAction" />
                        </n-list-item>
                    </div>
                    <div v-if="pageType === 'comment'">
                        <n-list-item v-for="post in commentList" :key="post.id">
                            <mobile-post-item :post="post"
                                :isOwner="store.state.userInfo.id == post.user_id" 
                                :addFollowAction="true"
                                @send-whisper="onSendWhisper"
                                @handle-follow-action="onHandleFollowAction" />
                        </n-list-item>
                    </div>
                    <div v-if="pageType === 'highlight'">
                        <n-list-item v-for="post in highlightList" :key="post.id">
                            <mobile-post-item :post="post"
                                :isOwner="store.state.userInfo.id == post.user_id" 
                                :addFollowAction="true"
                                @send-whisper="onSendWhisper"
                                @handle-follow-action="onHandleFollowAction" />
                        </n-list-item>
                    </div>
                    <div v-if="pageType === 'media'">
                        <n-list-item v-for="post in mediaList" :key="post.id">
                            <mobile-post-item :post="post"
                                :isOwner="store.state.userInfo.id == post.user_id" 
                                :addFollowAction="true"
                                @send-whisper="onSendWhisper"
                                @handle-follow-action="onHandleFollowAction" />
                        </n-list-item>
                    </div>
                    <div v-if="pageType === 'star'">
                        <n-list-item v-for="post in starList" :key="post.id">
                            <mobile-post-item :post="post"
                                :isOwner="store.state.userInfo.id == post.user_id" 
                                :addFollowAction="true"
                                @send-whisper="onSendWhisper"
                                @handle-follow-action="onHandleFollowAction" />
                        </n-list-item>
                    </div>
                </div>
            </div>
            <!-- 私信组件 -->
            <whisper :show="showWhisper" :user="whisperReceiver" @success="whisperSuccess" />
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
import { h, ref, Component, onMounted, computed, watch } from 'vue';
import { NIcon } from 'naive-ui';
import { useStore } from 'vuex';
import { useRoute, useRouter } from 'vue-router';
import { useDialog, DropdownOption } from 'naive-ui';
import { getUserPosts, followUser, unfollowUser } from '@/api/user';
import { formatDate } from '@/utils/formatTime';
import { prettyQuoteNum } from '@/utils/count';
import InfiniteLoading from 'v3-infinite-loading';
import { SettingsOutline } from '@vicons/ionicons5';
import { MoreHorizFilled } from '@vicons/material';

const store = useStore();
const route = useRoute();
const router = useRouter();
const dialog = useDialog();
const loading = ref(false);
const noMore = ref(false);
const list = ref<Item.PostProps[]>([]);
const postList = ref<Item.PostProps[]>([]);
const commentList = ref<Item.PostProps[]>([]);
const highlightList = ref<Item.PostProps[]>([]);
const mediaList = ref<Item.PostProps[]>([]);
const starList = ref<Item.PostProps[]>([]);
const pageType = ref<'post' | 'comment' | 'highlight' | 'media' | 'star'>(
  'post',
);
const postPage = ref(+(route.query.p as string) || 1);
const commentPage = ref(1);
const highlightPage = ref(1);
const mediaPage = ref(1);
const starPage = ref(1);
const page = ref(+(route.query.p as string) || 1);
const pageSize = ref(20);
const totalPage = ref(0);
const postTotalPage = ref(0);
const commentTotalPage = ref(0);
const highlightTotalPage = ref(0);
const mediaTotalPage = ref(0);
const starTotalPage = ref(0);
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

const renderIcon = (icon: Component) => {
  return () => {
    return h(NIcon, null, {
      default: () => h(icon),
    });
  };
};

const userOptions = computed(() => {
  let options: DropdownOption[] = [
    {
      label: '设置',
      key: 'setting',
      icon: renderIcon(SettingsOutline),
    },
  ];
  return options;
});

const handleUserAction = (item: 'setting') => {
  switch (item) {
    case 'setting':
      router.push({
        name: 'setting',
        query: {
          t: new Date().getTime(),
        },
      });
      break;
    default:
      break;
  }
};

const onSendWhisper = (user: Item.UserInfo) => {
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
      '确定' +
      (post.user.is_following ? '取消关注 @' : '关注 @') +
      post.user.username +
      ' 吗？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      if (post.user.is_following) {
        unfollowUser({
          user_id: post.user.id,
        })
          .then((_res) => {
            window.$message.success('操作成功');
            postFollowAction(post.user_id, false);
          })
          .catch((_err) => {});
      } else {
        followUser({
          user_id: post.user.id,
        })
          .then((_res) => {
            window.$message.success('关注成功');
            postFollowAction(post.user_id, true);
          })
          .catch((_err) => {});
      }
    },
  });
};

function postFollowAction(userId: number, isFollowing: boolean) {
  updateFolloing(postList.value, userId, isFollowing);
  updateFolloing(commentList.value, userId, isFollowing);
  updateFolloing(highlightList.value, userId, isFollowing);
  updateFolloing(mediaList.value, userId, isFollowing);
  updateFolloing(starList.value, userId, isFollowing);
}

function updateFolloing(
  posts: Item.PostProps[],
  userId: number,
  isFollowing: boolean,
) {
  if (posts && posts.length > 0) {
    for (let index in posts) {
      if (posts[index].user_id == userId) {
        posts[index].user.is_following = isFollowing;
      }
    }
  }
}

const loadPage = () => {
  switch (pageType.value) {
    case 'post':
      loadPosts();
      break;
    case 'comment':
      loadCommentPosts();
      break;
    case 'highlight':
      loadHighlightPosts();
      break;
    case 'media':
      loadMediaPosts();
      break;
    case 'star':
      loadStarPosts();
      break;
  }
};
const loadPosts = () => {
  loading.value = true;
  getUserPosts({
    username: store.state.userInfo.username,
    style: 'post',
    page: page.value,
    page_size: pageSize.value,
  })
    .then((rsp) => {
      loading.value = false;
      if (rsp.list.length === 0) {
        noMore.value = true;
      }
      if (page.value > 1) {
        list.value = list.value.concat(rsp.list);
      } else {
        list.value = rsp.list || [];
        window.scrollTo(0, 0);
      }
      totalPage.value = Math.ceil(rsp.pager.total_rows / pageSize.value);
      postList.value = list.value;
      postTotalPage.value = totalPage.value;
    })
    .catch((err) => {
      list.value = [];
      if (page.value > 1) {
        page.value--;
      }
      loading.value = false;
    });
};
const loadCommentPosts = () => {
  loading.value = true;
  getUserPosts({
    username: store.state.userInfo.username,
    style: 'comment',
    page: page.value,
    page_size: pageSize.value,
  })
    .then((rsp) => {
      loading.value = false;
      if (rsp.list.length === 0) {
        noMore.value = true;
      }
      if (page.value > 1) {
        list.value = list.value.concat(rsp.list);
      } else {
        list.value = rsp.list || [];
        window.scrollTo(0, 0);
      }
      totalPage.value = Math.ceil(rsp.pager.total_rows / pageSize.value);
      commentList.value = list.value;
      commentTotalPage.value = totalPage.value;
    })
    .catch((err) => {
      list.value = [];
      if (page.value > 1) {
        page.value--;
      }
      loading.value = false;
    });
};
const loadHighlightPosts = () => {
  loading.value = true;
  getUserPosts({
    username: store.state.userInfo.username,
    style: 'highlight',
    page: page.value,
    page_size: pageSize.value,
  })
    .then((rsp) => {
      loading.value = false;
      if (rsp.list.length === 0) {
        noMore.value = true;
      }
      if (page.value > 1) {
        list.value = list.value.concat(rsp.list);
      } else {
        list.value = rsp.list || [];
        window.scrollTo(0, 0);
      }
      totalPage.value = Math.ceil(rsp.pager.total_rows / pageSize.value);
      highlightList.value = list.value;
      highlightTotalPage.value = totalPage.value;
    })
    .catch((err) => {
      list.value = [];
      if (page.value > 1) {
        page.value--;
      }
      loading.value = false;
    });
};
const loadMediaPosts = () => {
  loading.value = true;
  getUserPosts({
    username: store.state.userInfo.username,
    style: 'media',
    page: page.value,
    page_size: pageSize.value,
  })
    .then((rsp) => {
      loading.value = false;
      if (rsp.list.length === 0) {
        noMore.value = true;
      }
      if (page.value > 1) {
        list.value = list.value.concat(rsp.list);
      } else {
        list.value = rsp.list || [];
        window.scrollTo(0, 0);
      }
      totalPage.value = Math.ceil(rsp.pager.total_rows / pageSize.value);
      mediaList.value = list.value;
      mediaTotalPage.value = totalPage.value;
    })
    .catch((err) => {
      list.value = [];
      if (page.value > 1) {
        page.value--;
      }
      loading.value = false;
    });
};
const loadStarPosts = () => {
  loading.value = true;
  getUserPosts({
    username: store.state.userInfo.username,
    style: 'star',
    page: page.value,
    page_size: pageSize.value,
  })
    .then((rsp) => {
      loading.value = false;
      if (rsp.list.length === 0) {
        noMore.value = true;
      }
      if (page.value > 1) {
        list.value = list.value.concat(rsp.list);
      } else {
        list.value = rsp.list || [];
        window.scrollTo(0, 0);
      }
      totalPage.value = Math.ceil(rsp.pager.total_rows / pageSize.value);
      starList.value = list.value;
      starTotalPage.value = totalPage.value;
    })
    .catch((err) => {
      list.value = [];
      if (page.value > 1) {
        page.value--;
      }
      loading.value = false;
    });
};
const changeTab = (
  tab: 'post' | 'comment' | 'highlight' | 'media' | 'star',
) => {
  pageType.value = tab;
  switch (pageType.value) {
    case 'post':
      list.value = postList.value;
      page.value = postPage.value;
      totalPage.value = postTotalPage.value;
      loadPosts();
      break;
    case 'comment':
      list.value = commentList.value;
      page.value = commentPage.value;
      totalPage.value = commentTotalPage.value;
      loadCommentPosts();
      break;
    case 'highlight':
      list.value = highlightList.value;
      page.value = highlightPage.value;
      totalPage.value = highlightTotalPage.value;
      loadHighlightPosts();
      break;
    case 'media':
      list.value = mediaList.value;
      page.value = mediaPage.value;
      totalPage.value = mediaTotalPage.value;
      loadMediaPosts();
      break;
    case 'star':
      list.value = starList.value;
      page.value = starPage.value;
      totalPage.value = starTotalPage.value;
      loadStarPosts();
      break;
  }
};
const updatePage = () => {
  switch (pageType.value) {
    case 'post':
      postPage.value = page.value;
      loadPosts();
      break;
    case 'comment':
      commentPage.value = page.value;
      loadCommentPosts();
      break;
    case 'highlight':
      highlightPage.value = page.value;
      loadHighlightPosts();
      break;
    case 'media':
      mediaPage.value = page.value;
      loadMediaPosts();
      break;
    case 'star':
      starPage.value = page.value;
      loadStarPosts();
      break;
  }
};
const nextPage = () => {
  if (page.value < totalPage.value || totalPage.value == 0) {
    noMore.value = false;
    page.value++;
    updatePage();
  } else {
    noMore.value = true;
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
  },
);
</script>

<style lang="less" scoped>
.profile-baseinfo {
    display: flex;
    padding: 16px;
    .avatar {
        width: 72px;
    }

    .base-info {
        position: relative;
        margin-left: 12px;
        width: calc(100% - 84px);

        .username {
            line-height: 16px;
            font-size: 16px;
        }

        .userinfo {
            font-size: 14px;
            line-height: 14px;
            margin-top: 10px;
            opacity: 0.75;
            .info-item {
                margin-right: 12px;
            }
        }

        .top-tag {
            transform: scale(0.75);
        }
    }

    .user-opts {
        position: relative;
        opacity: 0.75;
    }
}

.profile-tabs-wrap {
    padding: 0 16px;
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
    .profile-wrap, .pagination-wrap {
        background-color: rgba(16, 16, 20, 0.75);
    }
}
</style>