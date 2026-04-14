<template>
    <div>
        <main-nav title="主页" />

        <n-list
            class="main-content-wrap profile-wrap"
            bordered
            v-if="userInfo.id > 0"
        >
            <!-- 基础信息 -->
            <!-- <n-spin :show="false" > -->
            <div class="profile-baseinfo">
                <div class="avatar">
                    <n-avatar :size="72" :src="userInfo.avatar" />
                </div>
                <div class="base-info">
                    <div class="username">
                        <strong>{{ userInfo.nickname }}</strong>
                        <span> @{{ userInfo.username }} </span>
                        <n-tag v-if="userInfo.is_admin" class="top-tag" type="error" size="small" round>
                            管理员
                        </n-tag>
                    </div>
                    <div class="userinfo">
                        <span class="info-item">UID. {{ userInfo.id }} </span>
                        <span class="info-item">{{ formatDate(userInfo.created_on) }}&nbsp;加入</span>
                    </div>
                    <div class="userinfo">
                        <span class="info-item">
                             <router-link
                                @click.stop
                                class="following-link"
                                :to="{
                                    name: 'following',
                                    query: { 
                                        s: userInfo.username, 
                                        n: userInfo.nickname,
                                        t: 'follows',
                                    },
                                }"
                            >
                                关注&nbsp;&nbsp;{{ prettyQuoteNum(userInfo.follows) }}
                            </router-link>
                        </span>
                        <span class="info-item">
                            <router-link
                                @click.stop
                                class="following-link"
                                :to="{
                                    name: 'following',
                                    query: { 
                                        s: userInfo.username, 
                                        n: userInfo.nickname,
                                        t: 'followings',
                                    },
                                }"
                            >
                                粉丝&nbsp;&nbsp;{{ prettyQuoteNum(userInfo.followings) }}
                            </router-link>
                        </span>
                        <span class="info-item">
                            泡泡&nbsp;&nbsp;{{ prettyQuoteNum(userInfo.tweets_count) }}
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
                  <n-list-item v-for="post in listData" :key="post.id">
                      <post-item :post="post"
                          :isOwner="userInfo.id == post.user_id"
                          :addFollowAction="true"
                          :isMobile="!desktopModelShow"
                          @send-whisper="onSendWhisper"
                          @post-follow-action="postFollowAction" />
                  </n-list-item>
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
import { useStoreMain } from '@/store/main';
import { useRoute, useRouter } from 'vue-router';
import { useDialog, DropdownOption } from 'naive-ui';
import { formatDate } from '@/utils/formatTime';
import { prettyQuoteNum } from '@/utils/count';
import InfiniteLoading from 'v3-infinite-loading';
import { SettingsOutline } from '@vicons/ionicons5';
import { MoreHorizFilled } from '@vicons/material';
import { useStoreUser } from '@/store/user';
import { storeToRefs } from 'pinia';
import { Api } from '@/utils/request';

type PageType = 'post' | 'comment' | 'highlight' | 'media' | 'star';

const storeMain = useStoreMain();
const storeUser = useStoreUser();
const { refresh, desktopModelShow } = storeToRefs(storeMain);
const { userInfo } = storeToRefs(storeUser);

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
const pageType = ref<PageType>('post');
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

const listData = computed(() => {
	switch (pageType.value) {
		case 'post':
			return postList.value;
		case 'comment':
			return commentList.value;
		case 'highlight':
			return highlightList.value;
		case 'media':
			return mediaList.value;
		case 'star':
			return starList.value;
		default:
			return [];
	}
})

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
  loadPostsByStyle(pageType.value);
};
const styleListMap = {
  post: { list: postList, totalPage: postTotalPage },
  comment: { list: commentList, totalPage: commentTotalPage },
  highlight: { list: highlightList, totalPage: highlightTotalPage },
  media: { list: mediaList, totalPage: mediaTotalPage },
  star: { list: starList, totalPage: starTotalPage },
};
const stylePageMap = {
  post: postPage,
  comment: commentPage,
  highlight: highlightPage,
  media: mediaPage,
  star: starPage,
};
function loadPostsByStyle(style: keyof typeof styleListMap) {
  loading.value = true;
  Api.v1.user.get.posts({
    username: userInfo.value.username,
    style,
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
      styleListMap[style].list.value = list.value;
      styleListMap[style].totalPage.value = totalPage.value;
    })
    .catch((_err) => {
      list.value = [];
      if (page.value > 1) {
        page.value--;
      }
      loading.value = false;
    });
}
const changeTab = (tab: PageType) => {
  pageType.value = tab;
  list.value = styleListMap[tab].list.value;
  page.value = stylePageMap[tab].value;
  totalPage.value = styleListMap[tab].totalPage.value;
  loadPostsByStyle(tab);
};
const updatePage = () => {
  stylePageMap[pageType.value].value = page.value;
  loadPostsByStyle(pageType.value);
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
    refresh: refresh.value,
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