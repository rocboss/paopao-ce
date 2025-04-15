<template>
    <div class="rightbar-wrap" v-if="!store.state.collapsedRight">
        <div class="search-wrap">
            <n-input
                round
                clearable
                placeholder="搜一搜..."
                v-model:value="keyword"
                @keyup.enter.prevent="handleSearch"
            >
                <template #prefix>
                    <n-icon :component="Search" />
                </template>
            </n-input>
        </div>
        <n-card v-if="showFollowTopics" class="hottopic-wrap" title="关注话题" embedded :bordered="false" size="small">
            <n-spin :show="loading">
                <div class="hot-tag-item" v-for="tag in followTags" :key="tag.id">
                    <router-link
                        class="hash-link"
                        :to="{
                            name: 'home',
                            query: {
                                q: tag.tag,
                                t: 'tag',
                            },
                        }"
                    >
                        #{{ tag.tag }}
                    </router-link>

                    <div class="post-num">
                        {{ formatQuoteNum(tag.quote_num) }}
                    </div>
                </div>
            </n-spin>
        </n-card>
        <n-card class="hottopic-wrap" title="热门话题" embedded :bordered="false" size="small">
            <n-spin :show="loading">
                <div class="hot-tag-item" v-for="tag in hotTags" :key="tag.id">
                    <router-link
                        class="hash-link"
                        :to="{
                            name: 'home',
                            query: {
                                q: tag.tag,
                                t: 'tag',
                            },
                        }"
                    >
                        #{{ tag.tag }}
                    </router-link>

                    <div class="post-num">
                        {{ formatQuoteNum(tag.quote_num) }}
                    </div>
                </div>
            </n-spin>
        </n-card>
        <n-card class="copyright-wrap" embedded :bordered="false" size="small">
            <div class="copyright">&copy; {{ store.state.profile.copyrightTop }}</div>
            <div>
                <n-space>
                    <a
                        :href="store.state.profile.copyrightLeftLink"
                        target="_blank"
                        class="hash-link"
                    >
                        {{ store.state.profile.copyrightLeft }}
                    </a>
                    <a
                        :href="store.state.profile.copyrightRightLink"
                        target="_blank"
                        class="hash-link"
                    >
                        {{ store.state.profile.copyrightRight }}
                    </a>
                </n-space>
            </div>
        </n-card>
        <div class="site-info" v-if="store.state.userInfo.is_admin" ref="userInfoElement">
            <span class="site-info-item">{{ registerUserCount }} 注册用户，{{ onlineUserCount }} 人在线，最高在线 {{ historyMaxOnline }} 人，站点上线于 {{ formatRelativeTime(serverUpTime) }}</span>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch } from 'vue';
import { useStore } from 'vuex';
import { useRouter } from 'vue-router';
import { getTags } from '@/api/post';
import { getSiteInfo } from '@/api/user';
import { Search } from '@vicons/ionicons5';
import { formatRelativeTime } from '@/utils/formatTime';

const hotTags = ref<Item.TagProps[]>([]);
const followTags = ref<Item.TagProps[]>([]);
const loading = ref(false);
const keyword = ref('');
const store = useStore();
const router = useRouter();
const registerUserCount = ref(0);
const onlineUserCount = ref(0);
const historyMaxOnline = ref(0);
const serverUpTime = ref(0);
const userInfoElement = ref<HTMLElement | null>(null);
const rightFollowTopicMaxSize = Number(
  import.meta.env.VITE_RIGHT_FOLLOW_TOPIC_MAX_SIZE,
);
const rightHotTopicMaxSize = Number(
  import.meta.env.VITE_RIGHT_HOT_TOPIC_MAX_SIZE,
);

const loadSiteInfo = () => {
  getSiteInfo()
    .then((res) => {
      registerUserCount.value = res.register_user_count;
      onlineUserCount.value = res.online_user_count;
      historyMaxOnline.value = res.history_max_online;
      serverUpTime.value = res.server_up_time;
    })
    .catch((_err) => {
      // do nothing
    });
  observer.disconnect();
};
const loadHotTags = () => {
  loading.value = true;
  getTags({
    type: 'hot_extral',
    num: rightHotTopicMaxSize,
    extral_num: rightFollowTopicMaxSize,
  })
    .then((res) => {
      hotTags.value = res.topics;
      followTags.value = res.extral_topics ?? [];
      showFollowTopics.value = true;
      loading.value = false;
    })
    .catch((_err) => {
      loading.value = false;
    });
};
const formatQuoteNum = (num: number) => {
  if (num >= 1000) {
    return (num / 1000).toFixed(1) + 'k';
  }
  return num;
};
const handleSearch = () => {
  router.push({
    name: 'home',
    query: {
      q: keyword.value,
    },
  });
};
const showFollowTopics = computed({
  get: () => {
    return store.state.userLogined && followTags.value.length !== 0;
  },
  set: (newVal) => {
    // do nothing
  },
});
watch(
  () => ({
    refreshTopicFollow: store.state.refreshTopicFollow,
    userLogined: store.state.userLogined,
  }),
  (to, from) => {
    if (to.refreshTopicFollow !== from.refreshTopicFollow || to.userLogined) {
      loadHotTags();
    }
    if (store.state.userInfo.is_admin) {
      loadSiteInfo();
    }
  },
);
const observer = new IntersectionObserver(
  (entries: IntersectionObserverEntry[]) => {
    entries.forEach((entry) => {
      if (entry.isIntersecting) {
        loadSiteInfo();
      }
    });
  },
  {
    root: null,
    rootMargin: '0px',
    threshold: 1,
  },
);
onMounted(() => {
  // 不知道为什么 store.state.userInfo.is_admin 在这里就是不起作用f*k，所以才用这么一种蹩脚的法子来凑合
  if (userInfoElement.value) {
    observer.observe(userInfoElement.value);
  }
  loadHotTags();
});
</script>

<style lang="less" scoped>
.rightbar-wrap::-webkit-scrollbar {
  width: 0; /* 隐藏滚动条的宽度 */
  height: 0; /* 隐藏滚动条的高度 */
}
.rightbar-wrap {
    width: 240px;
    position: fixed;
    left: calc(50% + var(--content-main) / 2 + 10px);
    max-height: calc(100vh); /* 调整高度 */
    overflow: auto;
    .search-wrap {
        margin: 12px 0;
    }

    .hot-tag-item {
        line-height: 2;
        position: relative;

        .hash-link {
            width: calc(100% - 60px);
            text-overflow: ellipsis;
            white-space: nowrap;
            overflow: hidden;
            display: block;
        }

        .post-num {
            position: absolute;
            right: 0;
            top: 0;
            width: 60px;
            text-align: right;
            line-height: 2;
            opacity: 0.5;
        }
    }

    .hottopic-wrap {
        margin-bottom: 10px;
    }

    .site-info {
        margin-top: 8px;
        padding-left: 16px;
        padding-right: 16px;
        .site-info-item {
            font-size: 10px;
            opacity: 0.75;
        }
    }

    .copyright-wrap {
        .copyright {
            font-size: 12px;
            opacity: 0.75;
        }

        .hash-link {
            font-size: 12px;
        }
    }
}
.dark {
    .hottopic-wrap {
        background-color: #18181c;
    }
    .copyright-wrap {
        background-color: #18181c;
    }
}
</style>