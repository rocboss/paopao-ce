<template>
    <div>
        <main-nav :title="title" />

        <n-list class="main-content-wrap" bordered>
            <n-list-item>
                <!-- 发布器 -->
                <compose @post-success="onPostSuccess" />
            </n-list-item>

            <n-list-item v-if="showTrendsBar" >
            <SlideBar v-model="slideBarList" :wheel-blocks="wheelBlocks" :init-blocks="initBlocks" @click="handleBarClick" tag="div" sub-tag="div">
                <template #default="data">
                    <div class="slide-bar-item">
                        <n-badge value="1" :offset="[-4, 48]" dot :show="data.slotData.show">
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
            <div  class="style-wrap" v-else-if="showTrendsTag">
            <n-space >
                <n-button v-if="newestTweetsStyle !== 'newest'" size="small" :bordered="false" @click="onNewestTweets" class="style-item" secondary round>
                    全部
                </n-button>
                <n-button v-if="newestTweetsStyle === 'newest'" size="small" type="success"  :bordered="false" @click="onNewestTweets"  class="style-item" secondary round>
                    全部
                </n-button>
                <n-button v-if="newestTweetsStyle !== 'hots'" size="small" :bordered="false" @click="onHotTweets" class="style-item" secondary round>
                    热门推荐
                </n-button>
                <n-button v-if="newestTweetsStyle === 'hots'" size="small" type="success" :bordered="false" @click="onHotTweets" class="style-item" secondary round>
                    热门推荐
                </n-button>
                <n-button v-if="newestTweetsStyle !== 'following'" size="small" :bordered="false" @click="onFollowingTweets" class="style-item" secondary round>
                    正在关注
                </n-button>
                <n-button v-if="newestTweetsStyle === 'following'" size="small" type="success" :bordered="false" @click="onFollowingTweets" class="style-item" secondary round>
                    正在关注
                </n-button>
            </n-space>
            </div>
            <div v-if="loading && list.length === 0" class="skeleton-wrap">
                <post-skeleton :num="pageSize" />
            </div>

            <div>
                <div class="empty-wrap" v-if="list.length === 0">
                    <n-empty size="large" description="暂无数据" />
                </div>
                <div v-if="store.state.desktopModelShow">
                    <n-list-item v-for="post in list" :key="post.id">
                        <post-item :post="post" 
                            :isOwner="store.state.userInfo.id == post.user_id" 
                            :addFollowAction="true"
                            @send-whisper="onSendWhisper"
                            @handle-follow-action="onHandleFollowAction"
                            @handle-friend-action="onHandleFriendAction" />
                    </n-list-item>
                </div>
                <div v-else>
                    <n-list-item v-for="post in list" :key="post.id">
                        <mobile-post-item :post="post"
                            :isOwner="store.state.userInfo.id == post.user_id" 
                            :addFollowAction="true"
                            @send-whisper="onSendWhisper"
                            @handle-follow-action="onHandleFollowAction"
                            @handle-friend-action="onHandleFriendAction" />
                    </n-list-item>
                </div>
            </div>
            <!-- 私信组件 -->
            <whisper :show="showWhisper" :user="whisperReceiver" @success="whisperSuccess" />
            <!-- 加好友组件 -->
            <whisper-add-friend :show="showAddFriendWhisper" :user="user" @success="addFriendWhisperSuccess" />
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
import { ref, onMounted, reactive, computed, watch } from 'vue';
import { useStore } from 'vuex';
import { useRoute, useRouter } from 'vue-router';
import { useDialog } from 'naive-ui';
import InfiniteLoading from 'v3-infinite-loading';
import { getPosts, getIndexTrends } from '@/api/post';
import {
  getUserPosts,
  deleteFriend,
  followUser,
  unfollowUser,
} from '@/api/user';
import SlideBar from '@opentiny/vue-slide-bar';
import allTweets from '@/assets/img/fresh-tweets.png';
import discoverTweets from '@/assets/img/discover-tweets.jpeg';
import followingTweets from '@/assets/img/following-tweets.jpeg';

const store = useStore();
const route = useRoute();
const router = useRouter();
const dialog = useDialog();

const newestTweetsStyle = ref<'newest' | 'hots' | 'following'>('newest');
const onNewestTweets = () => {
  newestTweetsStyle.value = 'newest';
  handleBarClick(slideBarList.value[0], 0);
};
const onHotTweets = () => {
  newestTweetsStyle.value = 'hots';
  handleBarClick(slideBarList.value[1], 1);
};
const onFollowingTweets = () => {
  newestTweetsStyle.value = 'following';
  handleBarClick(slideBarList.value[2], 2);
};

const initBlocks = ref(9);
const wheelBlocks = ref(8);
const slideBarList = ref<Item.SlideBarItem[]>([
  { title: '最新动态', style: 1, username: '', avatar: allTweets, show: true },
  {
    title: '热门推荐',
    style: 2,
    username: '',
    avatar: discoverTweets,
    show: false,
  },
  {
    title: '正在关注',
    style: 3,
    username: '',
    avatar: followingTweets,
    show: false,
  },
  // TODO: 不知道SlideBar抽什么疯，如果没有填充下面这些伪数据的话，直接设置initBlocks为9而给的数据又不足，后面动态添加数据后，吖的竟然不能后划了，
  // f*k，不知道哪姿势不对，总之先凑合着用吧，后期再优化。
  { title: '', style: 1, username: '', avatar: '', show: true },
  { title: '', style: 1, username: '', avatar: '', show: true },
  { title: '', style: 1, username: '', avatar: '', show: true },
  { title: '', style: 1, username: '', avatar: '', show: true },
  { title: '', style: 1, username: '', avatar: '', show: true },
  { title: '', style: 1, username: '', avatar: '', show: true },
  { title: '', style: 1, username: '', avatar: '', show: true },
  { title: '', style: 1, username: '', avatar: '', show: true },
  { title: '', style: 1, username: '', avatar: '', show: true },
]);
const user = reactive<Item.UserInfo>({
  id: 0,
  avatar: '',
  username: '',
  nickname: '',
  is_admin: false,
  is_friend: false,
  is_following: false,
  created_on: 0,
  follows: 0,
  followings: 0,
  status: 1,
});
const inActionPost = ref<Item.PostProps | null>(null);

const title = ref<string>('泡泡广场');
const loading = ref(false);
const noMore = ref(false);
const targetStyle = ref<number>(1);
const targetUsername = ref<string>('');
const list = ref<any[]>([]);
const page = ref(1);
const pageSize = ref(20);
const totalPage = ref(0);
const showWhisper = ref(false);
const showAddFriendWhisper = ref(false);
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

const onSendWhisper = (user: Item.UserInfo) => {
  whisperReceiver.value = user;
  showWhisper.value = true;
};

const whisperSuccess = () => {
  showWhisper.value = false;
};

const openAddFriendWhisper = () => {
  showAddFriendWhisper.value = true;
};

const openDeleteFriend = (post: Item.PostProps) => {
  dialog.warning({
    title: '删除好友',
    content:
      '将好友 “' +
      post.user.nickname +
      '” 删除，将同时删除 点赞/收藏 列表中关于该朋友的 “好友可见” 推文',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      deleteFriend({
        user_id: user.id,
      })
        .then((res) => {
          window.$message.success('操作成功');
          post.user.is_friend = false;
        })
        .catch((_err) => {});
    },
  });
};

const addFriendWhisperSuccess = () => {
  showAddFriendWhisper.value = false;
  inActionPost.value = null;
};

const onHandleFriendAction = (post: Item.PostProps) => {
  inActionPost.value = post;
  user.id = post.user.id;
  user.username = post.user.username;
  user.nickname = post.user.nickname;
  if (post.user.is_friend) {
    openDeleteFriend(post);
  } else {
    openAddFriendWhisper();
  }
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
  for (let index in list.value) {
    if (list.value[index].user_id == userId) {
      list.value[index].user.is_following = isFollowing;
    }
  }
}

const updateTitle = () => {
  title.value = '泡泡广场';
  if (route.query && route.query.q) {
    if (route.query.t && route.query.t === 'tag') {
      title.value = '#' + decodeURIComponent(route.query.q as string);
    } else {
      title.value = '搜索: ' + decodeURIComponent(route.query.q as string);
    }
  }
};

const showTrendsTag = computed(() => {
  return (
    store.state.userInfo.id > 0 &&
    !store.state.profile.enableTrendsBar &&
    store.state.desktopModelShow
  );
});
const showTrendsBar = computed(() => {
  return (
    store.state.profile.useFriendship &&
    store.state.profile.enableTrendsBar &&
    store.state.desktopModelShow &&
    store.state.userInfo.id > 0
  );
});

const reset = () => {
  loading.value = false;
  noMore.value = false;
  list.value = [];
  page.value = 1;
  totalPage.value = 0;
};

const handleBarClick = (data: Item.SlideBarItem, index: number) => {
  reset();
  targetStyle.value = data.style;
  if (route.query.q) {
    route.query.q = null;
    updateTitle();
  }
  switch (data.style) {
    case 1:
      loadPosts('newest');
      break;
    case 2:
      loadPosts('hots');
      break;
    case 3:
      route.query.q = null;
      loadPosts('following');
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
  slideBarList.value = slideBarList.value.slice(0, 3);
  if (
    !store.state.profile.useFriendship ||
    !store.state.profile.enableTrendsBar ||
    store.state.userInfo.id === 0
  ) {
    return;
  }
  getIndexTrends({
    page: 1,
    page_size: 50,
  })
    .then((res) => {
      var i = 0;
      const list = res.list || [];
      let barItems: Item.SlideBarItem[] = [];
      for (; i < list.length; i++) {
        let item: Item.IndexTrendsItem = list[i];
        barItems.push({
          title: item.nickname,
          style: 21,
          username: item.username,
          avatar: item.avatar,
          show: item.is_fresh,
        });
      }
      if (barItems.length > 0) {
        slideBarList.value = slideBarList.value.concat(barItems);
      }
    })
    .catch((err) => {
      console.log(err);
    });
};

const loadPosts = (style: 'newest' | 'hots' | 'following' | 'search') => {
  loading.value = true;
  getPosts({
    query: route.query.q ? decodeURIComponent(route.query.q as string) : null,
    type: route.query.t as string,
    style: style,
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
        list.value = rsp.list;
        window.scrollTo(0, 0);
      }

      totalPage.value = Math.ceil(rsp.pager.total_rows / pageSize.value);
    })
    .catch((err) => {
      loading.value = false;
      if (page.value > 1) {
        page.value--;
      }
    });
};

const loadUserPosts = () => {
  loading.value = true;
  getUserPosts({
    username: targetUsername.value,
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
  // 暂时统统跳到详情页面，后续再精细化分场景优化
  router.push({
    name: 'post',
    query: {
      id: post.id,
    },
  });
  // // 如果不在第一页，需要跳转到详情页面
  // if (targetStyle.value != 1) {
  //     router.push({
  //         name: 'post',
  //         query: {
  //             id: post.id,
  //         },
  //     });
  //     return;
  // }

  // // 如果是在第一页，就地插入新推文到文章列表中
  // let items = [];
  // let length = list.value.length;
  // if (length == pageSize.value) {
  //     length--;
  // }
  // var i = 0;
  // for (; i < length; i++) {
  //     let item: Item.PostProps = list.value[i];
  //     if (!item.is_top) {
  //         break;
  //     }
  //     items.push(item);
  // }
  // items.push(post);
  // for (; i < length; i++) {
  //     items.push(list.value[i]);
  // }
  // list.value = items;
};

const loadMorePosts = () => {
  switch (targetStyle.value) {
    case 1:
      loadPosts('newest');
      break;
    case 2:
      loadPosts('hots');
      break;
    case 3:
      loadPosts('following');
      break;
    case 21:
      if (route.query.q) {
        loadPosts('search');
      } else {
        loadUserPosts();
      }
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
  reset();
  loadContacts();
  loadPosts('newest');
});

watch(
  () => ({
    path: route.path,
    query: route.query,
    refresh: store.state.refresh,
  }),
  (to, from) => {
    updateTitle();
    if (to.refresh !== from.refresh) {
      reset();
      setTimeout(() => {
        loadContacts();
        loadMorePosts();
      }, 0);
      return;
    }
    if (from.path !== '/post' && to.path === '/') {
      reset();
      setTimeout(() => {
        loadContacts();
        loadMorePosts();
      }, 0);
    }
  },
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
        color: #18a058;
        opacity: 0.8;
    }
    .slide-bar-item-title {
        color: #18a058;
        opacity: 0.8;
    }
}
.style-wrap {
    margin-top: 10px;
    margin-left: 16px;
    margin-bottom: 4px;
    opacity: 0.80;  
    .style-item {
        &.hover {
                cursor: pointer;
        }
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

    .tiny-slide-bar {
        --ti-slider-progress-box-arrow-hover-text-color: #f2f2f2;
        --ti-slider-progress-box-arrow-normal-text-color: #808080;
    }
}
</style>
