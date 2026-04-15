<template>
    <div>
        <main-nav :title="nickname" :back="true" />

        <n-list class="main-content-wrap" bordered>
            <n-tabs type="line" animated :default-value="tabler" @update:value="changeTab">
                <n-tab-pane name="follows" tab="正在关注" />
                <n-tab-pane name="followings" tab="我的粉丝" />
            </n-tabs>
            <div v-if="loading && list.length === 0" class="skeleton-wrap">
                <post-skeleton :num="pageSize" />
            </div>
            <div v-else>
                <div class="empty-wrap" v-if="list.length === 0">
                    <n-empty size="large" description="暂无数据" />
                </div>

                <n-list-item v-for="contact in list" :key="contact.user_id">
                    <user-card type="follow" :contact="contact" @send-whisper="onSendWhisper" @unfollow-success="handleUnfollowSuccess" />
                </n-list-item>
            </div>
            <!-- 私信组件 -->
            <whisper :show="showWhisper" :user="whisperReceiver" @success="whisperSuccess" />
        </n-list>
    </div>
    <n-space v-if="totalPage > 0" justify="center">
		<InfiniteLoading class="load-more" :slots="{ complete: completeStr, error: '加载出错' }" @infinite="handleNextPage">
			<template #spinner>
				<div class="load-more-wrap">
					<n-spin :size="14" v-if="!noMore" />
					<span class="load-more-spinner">{{ noMore ? completeStr : '加载更多' }}</span>
				</div>
			</template>
		</InfiniteLoading>
    </n-space>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import InfiniteLoading from 'v3-infinite-loading';
import { useRoute } from 'vue-router';
import { Api } from '@/utils/request';
import { usePagination } from '@/composables/usePagination';
import UserAction from '@/composables/useUserAction';
import UserCard from '@/components/user-card.vue';

const route = useRoute();

const list = ref<Item.ContactItemProps[]>([]);
const nickname = (route.query.n as string) || '粉丝详情';
const username = (route.query.s as string) || '';
const tabler = ref((route.query.t as string) || 'follows');
const showAddFriendWhisper = ref(false);

// 使用 usePagination composable
const { loading, noMore, page, pageSize, totalPage, reset, nextPage } =
  usePagination(20);

// 使用 UserAction.useWhisper()
const { showWhisper, whisperReceiver, onSendWhisper, whisperSuccess } =
  UserAction.useWhisper();

function resetPage(tab: 'follows' | 'followings') {
  list.value = [];
  reset();
  tabler.value = tab;
}

const completeStr = computed(() => {
  if (tabler.value == 'follows') {
    return '没有更多关注了';
  } else {
    return '没有更多粉丝了';
  }
});

const handleNextPage = () => {
  nextPage(loadPage);
};

const changeTab = (tab: 'follows' | 'followings') => {
  resetPage(tab);
  loadPage();
};

const loadPage = () => {
  if (tabler.value === 'follows') {
    loadFollows(username);
  } else if (tabler.value === 'followings') {
    loadFollowings(username);
  }
};

const handleUnfollowSuccess = () => {
  // 只在【正在关注】页面刷新列表
  if (tabler.value === 'follows') {
    resetPage('follows');
    loadPage();
  }
};

const loadFollows = (username: string, scrollToBottom: boolean = false) => {
  if (list.value.length === 0) {
    loading.value = true;
  }
  Api.v1.user.get
    .follows({
      username: username,
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
        if (scrollToBottom) {
          setTimeout(() => {
            window.scrollTo(0, 99999);
          }, 50);
        }
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

const loadFollowings = (username: string, scrollToBottom: boolean = false) => {
  if (list.value.length === 0) {
    loading.value = true;
  }
  Api.v1.user.get
    .followings({
      username: username,
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
        if (scrollToBottom) {
          setTimeout(() => {
            window.scrollTo(0, 99999);
          }, 50);
        }
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

onMounted(() => {
  loadPage();
});
</script>

<style lang="less" scoped>

.main-content-wrap {
    padding: 20px;
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
    .main-content-wrap, .empty-wrap, .skeleton-wrap {
        background-color: rgba(16, 16, 20, 0.75);
    }
}
</style>