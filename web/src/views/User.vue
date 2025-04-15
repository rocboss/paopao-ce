<template>
    <div>
        <main-nav title="用户详情" />

        <n-list class="main-content-wrap profile-wrap" bordered>
            <!-- 基础信息 -->
            <n-spin :show="userLoading">
                <div class="profile-baseinfo" v-if="user.id > 0">
                    <div class="avatar">
                        <n-avatar :size="72" :src="user.avatar" />
                    </div>
                    <div class="base-info">
                        <div class="username">
                            <strong>{{ user.nickname }}</strong>
                            <span> @{{ user.username }} </span>
                            <n-tag
                                v-if="store.state.profile.useFriendship && store.state.userInfo.id > 0 && store.state.userInfo.username != user.username && user.is_friend"
                                class="top-tag" type="info" size="small" round>
                                好友
                            </n-tag>
                            <n-tag
                                v-if="store.state.userInfo.id > 0 && store.state.userInfo.username != user.username && user.is_following"
                                class="top-tag" type="success" size="small" round>
                                已关注
                            </n-tag>
                            <n-tag v-if="user.is_admin" class="top-tag" type="error" size="small" round>
                                管理员
                            </n-tag>
                        </div>
                        <div class="userinfo">
                            <span class="info-item">UID. {{ user.id }} </span>
                            <span class="info-item">{{ formatDate(user.created_on) }}&nbsp;加入</span>
                        </div>
                        <div class="userinfo">
                            <span class="info-item">
                                <router-link
                                    @click.stop
                                    class="following-link"
                                    :to="{
                                        name: 'following',
                                        query: { 
                                            s: user.username, 
                                            n: user.nickname,
                                            t: 'follows',
                                        },
                                    }"
                                >
                                    关注&nbsp;&nbsp;{{ prettyQuoteNum(user.follows)}}
                                </router-link>
                            </span>
                            <span class="info-item">
                                <router-link
                                    @click.stop
                                    class="following-link"
                                    :to="{
                                        name: 'following',
                                        query: { 
                                            s: user.username, 
                                            n: user.nickname,
                                            t: 'followings',
                                        },
                                    }"
                                >
                                    粉丝&nbsp;&nbsp;{{ prettyQuoteNum(user.followings) }}
                                </router-link>
                            </span>
                            <span class="info-item">
                                泡泡&nbsp;&nbsp;{{ prettyQuoteNum(user.tweets_count || 0) }}
                            </span>
                        </div>
                    </div>

                    <div class="user-opts" v-if="store.state.userInfo.id > 0">
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
            
                <n-tabs v-if="!userLoading" class="profile-tabs-wrap" type="line" animated :value="pageType" @update:value="changeTab">
                    <n-tab-pane name="post" tab="泡泡"></n-tab-pane>
                    <n-tab-pane name="comment" tab="评论"></n-tab-pane>
                    <n-tab-pane name="highlight" tab="亮点"></n-tab-pane>
                    <n-tab-pane name="media" tab="图文"></n-tab-pane>
                    <n-tab-pane name="star" tab="喜欢"></n-tab-pane>
                </n-tabs>
            </n-spin>
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
import { h, ref, reactive, watch, onMounted, computed } from 'vue';
import { NIcon } from 'naive-ui';
import type { Component, Ref } from 'vue';
import { useStore } from 'vuex';
import { useRoute, useRouter } from 'vue-router';
import {
  getUserProfile,
  getUserPosts,
  changeUserStatus,
  deleteFriend,
  followUser,
  unfollowUser,
} from '@/api/user';
import { useDialog, DropdownOption } from 'naive-ui';
import WhisperAddFriend from '../components/whisper-add-friend.vue';
import { MoreHorizFilled } from '@vicons/material';
import { formatDate } from '@/utils/formatTime';
import { prettyQuoteNum } from '@/utils/count';
import {
  SettingsOutline,
  PaperPlaneOutline,
  PersonAddOutline,
  PersonRemoveOutline,
  CubeOutline,
  BodyOutline,
  WalkOutline,
} from '@vicons/ionicons5';
import InfiniteLoading from 'v3-infinite-loading';

const dialog = useDialog();
const store = useStore();
const route = useRoute();
const router = useRouter();

const loading = ref(false);
const noMore = ref(false);
const user = reactive<Item.UserInfo>({
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
  tweets_count: 0,
  status: 1,
});
const userLoading = ref(false);
const showWhisper = ref(false);
const showAddFriendWhisper = ref(false);
const list = ref<Item.PostProps[]>([]);
const postList = ref<Item.PostProps[]>([]);
const commentList = ref<Item.PostProps[]>([]);
const highlightList = ref<Item.PostProps[]>([]);
const mediaList = ref<Item.PostProps[]>([]);
const starList = ref<Item.PostProps[]>([]);
const username = ref(route.query.s || '');
const page = ref(+(route.query.p as string) || 1);
const pageType = ref<'post' | 'comment' | 'highlight' | 'media' | 'star'>(
  'post',
);
const postPage = ref(+(route.query.p as string) || 1);
const commentPage = ref(1);
const highlightPage = ref(1);
const mediaPage = ref(1);
const starPage = ref(1);
const pageSize = ref(20);
const totalPage = ref(0);
const postTotalPage = ref(0);
const commentTotalPage = ref(0);
const highlightTotalPage = ref(0);
const mediaTotalPage = ref(0);
const starTotalPage = ref(0);

const onSendWhisper = (receiver: Item.UserInfo) => {
  user.id = receiver.id;
  user.username = receiver.username;
  user.nickname = receiver.nickname;
  user.avatar = receiver.avatar;
  showWhisper.value = true;
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
  updateFolloing(postList, userId, isFollowing);
  updateFolloing(commentList, userId, isFollowing);
  updateFolloing(highlightList, userId, isFollowing);
  updateFolloing(mediaList, userId, isFollowing);
  updateFolloing(starList, userId, isFollowing);
}

function updateFolloing(
  posts: Ref<Item.PostProps[]>,
  userId: number,
  isFollowing: boolean,
) {
  if (posts.value && posts.value.length > 0) {
    for (let index in posts.value) {
      if (posts.value[index].user_id == userId) {
        posts.value[index].user.is_following = isFollowing;
      }
    }
  }
}

const reset = () => {
  noMore.value = false;
  list.value = [];
  postList.value = [];
  commentList.value = [];
  highlightList.value = [];
  mediaList.value = [];
  starList.value = [];
  pageType.value = 'post';
  page.value = 1;
  postPage.value = 1;
  commentPage.value = 1;
  highlightPage.value = 1;
  mediaPage.value = 1;
  starPage.value = 1;
  totalPage.value = 0;
  postTotalPage.value = 0;
  commentTotalPage.value = 0;
  highlightTotalPage.value = 0;
  mediaTotalPage.value = 0;
  starTotalPage.value = 0;
};
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
    username: username.value as string,
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
    username: username.value as string,
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
    username: username.value as string,
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
    username: username.value as string,
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
    username: username.value as string,
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
      user.created_on = res.created_on;
      user.is_following = res.is_following;
      user.follows = res.follows;
      user.followings = res.followings;
      user.status = res.status;
      if (res.tweets_count) {
        user.tweets_count = res.tweets_count;
      }
      loadPage();
    })
    .catch((err) => {
      userLoading.value = false;
      console.log(err);
    });
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
const renderIcon = (icon: Component) => {
  return () => {
    return h(NIcon, null, {
      default: () => h(icon),
    });
  };
};
const userOptions = computed(() => {
  if (store.state.userInfo.username == user.username) {
    return [
      {
        label: '设置',
        key: 'setting',
        icon: renderIcon(SettingsOutline),
      },
    ];
  }
  let options: DropdownOption[] = [
    {
      label: '私信',
      key: 'whisper',
      icon: renderIcon(PaperPlaneOutline),
    },
  ];
  if (store.state.userInfo.is_admin) {
    if (user.status === 1) {
      options.push({
        label: '禁言',
        key: 'banned',
        icon: renderIcon(CubeOutline),
      });
    } else {
      options.push({
        label: '解封',
        key: 'deblocking',
        icon: renderIcon(CubeOutline),
      });
    }
  }
  if (user.is_following) {
    options.push({
      label: '取消关注',
      key: 'unfollow',
      icon: renderIcon(WalkOutline),
    });
  } else {
    options.push({
      label: '关注',
      key: 'follow',
      icon: renderIcon(BodyOutline),
    });
  }
  if (store.state.profile.useFriendship) {
    if (user.is_friend) {
      options.push({
        label: '删除好友',
        key: 'delete',
        icon: renderIcon(PersonRemoveOutline),
      });
    } else {
      options.push({
        label: '添加朋友',
        key: 'requesting',
        icon: renderIcon(PersonAddOutline),
      });
    }
  }
  return options;
});
const handleUserAction = (
  item:
    | 'whisper'
    | 'follow'
    | 'unfollow'
    | 'delete'
    | 'requesting'
    | 'banned'
    | 'deblocking'
    | 'setting',
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
    case 'follow':
    case 'unfollow':
      handleFollowUser();
      break;
    case 'banned':
    case 'deblocking':
      banUser();
      break;
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
const openDeleteFriend = () => {
  dialog.warning({
    title: '删除好友',
    content:
      '将好友 “' +
      user.nickname +
      '” 删除，将同时删除 点赞/收藏 列表中关于该朋友的 “好友可见” 推文',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      userLoading.value = true;
      deleteFriend({
        user_id: user.id,
      })
        .then((res) => {
          userLoading.value = false;
          user.is_friend = false;
          loadPosts();
        })
        .catch((err) => {
          userLoading.value = false;
          console.log(err);
        });
    },
  });
};
const handleFollowUser = () => {
  dialog.success({
    title: '提示',
    content:
      '确定' +
      (user.is_following ? '取消关注 @' : '关注 @') +
      user.username +
      ' 吗？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      userLoading.value = true;
      if (user.is_following) {
        unfollowUser({
          user_id: user.id,
        })
          .then((_res) => {
            userLoading.value = false;
            window.$message.success('操作成功');
            loadUser();
          })
          .catch((err) => {
            userLoading.value = false;
            console.log(err);
          });
      } else {
        followUser({
          user_id: user.id,
        })
          .then((_res) => {
            userLoading.value = false;
            window.$message.success('操作成功');
            loadUser();
          })
          .catch((err) => {
            userLoading.value = false;
            console.log(err);
          });
      }
    },
  });
};
const banUser = () => {
  dialog.warning({
    title: '警告',
    content:
      '确定对该用户进行' + (user.status === 1 ? '禁言' : '解封') + '处理吗？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      userLoading.value = true;
      changeUserStatus({
        id: user.id,
        status: user.status === 1 ? 2 : 1,
      })
        .then((_res) => {
          userLoading.value = false;
          if (user.status === 1) {
            window.$message.success('禁言成功');
          } else {
            window.$message.success('解封成功');
          }
          loadUser();
        })
        .catch((err) => {
          userLoading.value = false;
          console.log(err);
        });
    },
  });
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
  loadUser();
});
watch(
  () => ({
    path: route.path,
    query: route.query,
  }),
  (to, from) => {
    if (from.path === '/u' && to.path === '/u') {
      username.value = route.query.s || '';
      reset();
      loadUser();
    }
  },
);
</script>

<style lang="less" scoped>
.profile-tabs-wrap {
    padding: 0 16px;
}

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
        position: absolute;
        top: 16px;
        right: 16px;
        opacity: 0.75;
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
    .profile-wrap, .pagination-wrap {
        background-color: rgba(16, 16, 20, 0.75);
    }
}
</style>