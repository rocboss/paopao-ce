<template>
    <div>
        <main-nav title="收藏" />

        <n-list class="main-content-wrap" bordered>
            <div v-if="loading && list.length === 0" class="skeleton-wrap">
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
        <n-space v-if="totalPage > 0" justify="center">
            <InfiniteLoading class="load-more" :slots="{ complete: '没有更多收藏了', error: '加载出错' }" @infinite="nextPage">
                <template #spinner>
                    <div class="load-more-wrap">
                        <n-spin :size="14" v-if="!noMore" />
                        <span class="load-more-spinner">{{ noMore ? '没有更多收藏了' : '加载更多' }}</span>
                    </div>
                </template>
            </InfiniteLoading>
        </n-space>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useStore } from 'vuex';
import { useRoute } from 'vue-router';
import { useDialog } from 'naive-ui';
import InfiniteLoading from 'v3-infinite-loading';
import { getCollections, followUser, unfollowUser } from '@/api/user';

const store = useStore();
const route = useRoute();
const dialog = useDialog();

const loading = ref(false);
const noMore = ref(false);
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
      '确定' + (post.user.is_following ? '取消关注' : '关注') + '该用户吗？',
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

const loadPosts = () => {
  loading.value = true;
  getCollections({
    page: page.value,
    page_size: pageSize.value,
  })
    .then((res) => {
      loading.value = false;
      if (res.list.length === 0) {
        noMore.value = true;
      }
      if (page.value > 1) {
        list.value = list.value.concat(res.list);
      } else {
        list.value = res.list;
        window.scrollTo(0, 0);
      }
      totalPage.value = Math.ceil(res.pager.total_rows / pageSize.value);
    })
    .catch((_err) => {
      loading.value = false;
      if (page.value > 1) {
        page.value--;
      }
    });
};
const nextPage = () => {
  if (page.value < totalPage.value || totalPage.value == 0) {
    noMore.value = false;
    page.value++;
    loadPosts();
  } else {
    noMore.value = true;
  }
};
onMounted(() => {
  loadPosts();
});
</script>

<style lang="less" scoped>
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
    .main-content-wrap, .empty-wrap, .skeleton-wrap {
        background-color: rgba(16, 16, 20, 0.75);
    }
}
</style>