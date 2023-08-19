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
    <n-card
      v-if="showFollowTopics"
      class="hottopic-wrap"
      title="关注专题"
      embedded
      :bordered="false"
      size="small"
    >
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
    <n-card
      class="hottopic-wrap"
      title="热门专题"
      embedded
      :bordered="false"
      size="small"
    >
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
      <div class="copyright">&copy; {{ copyrightTop }}</div>
      <div>
        <n-space>
          <a :href="copyrightLeftLink" target="_blank" class="hash-link">
            {{ copyrightLeft }}
          </a>
          <a :href="copyrightRightLink" target="_blank" class="hash-link">
            {{ copyrightRight }}
          </a>
        </n-space>
      </div>
    </n-card>
    <n-card
      class="hottopic-wrap"
      title="排行榜"
      embedded
      :bordered="false"
      size="small"
    >
      <n-spin :show="loading">
        <div
          class="ranking-item"
          v-for="(highQuality, index) in rankingList"
          :key="highQuality.name"
        >
        <div class="ranking-number">{{ index + 1 }}</div> <!-- 添加排序编号 -->
          <div class="ranking-avatar">
            <!-- 用户头像 -->
            <img :src="highQuality.avatar" alt="User Avatar" />
          </div>
          <div class="user-name">
            <router-link
              class="hash-link"
              :to="{
                name: 'user',
                query: {
                  s: highQuality.name
                }
              }"
            >
              {{ formatUserName(highQuality.name) }}
            </router-link>
          </div>
          
          <div class="ranking-info">
            <div class="name-stats">
              <div class="downloads">
                <div class="download-box">
                  <div class="download-value">{{ highQuality.comprehensive_score }}</div>
                </div>
              </div>
              <div class="stats">
                <div class="stat-item">
                  <div class="stat-value">{{ highQuality.post_count }} 贴子</div>
                </div>
                <div class="stat-drop">·</div>
                <!-- 加粗的点 -->
                <div class="stat-item">
                  <div class="stat-value">{{ highQuality.likes }} 点赞</div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </n-spin>
    </n-card>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch } from "vue";
import { useStore } from "vuex";
import { useRouter } from "vue-router";
import { getHighQuailty, getTags } from "@/api/post";
import { Search } from "@vicons/ionicons5";
import { ChevronForward } from "@vicons/ionicons5";

const hotTags = ref<Item.TagProps[]>([]);
const followTags = ref<Item.TagProps[]>([]);
const loading = ref(false);
const keyword = ref("");
const store = useStore();
const router = useRouter();
const copyrightTop = import.meta.env.VITE_COPYRIGHT_TOP;
const copyrightLeft = import.meta.env.VITE_COPYRIGHT_LEFT;
const copyrightLeftLink = import.meta.env.VITE_COPYRIGHT_LEFT_LINK;
const copyrightRight = import.meta.env.VITE_COPYRIGHT_RIGHT;
const copyrightRightLink = import.meta.env.VITE_COPYRIGHT_RIGHT_LINK;
const rightFollowTopicMaxSize = Number(
  import.meta.env.VITE_RIGHT_FOLLOW_TOPIC_MAX_SIZE
);
const rightHotTopicMaxSize = Number(
  import.meta.env.VITE_RIGHT_HOT_TOPIC_MAX_SIZE
);

// 模拟排行榜数据
const rankingList = ref<Item.highqualityProps[]>([]);

//获取排行榜数据
const locadHeighQuailtyRankingList = () => {
  loading.value = true;
  getHighQuailty()
    .then((res) => {
      rankingList.value = res.list;
      loading.value = false;
    })
    .catch((err) => {
      loading.value = false;
    });

}

const loadHotTags = () => {
  loading.value = true;
  getTags({
    type: "hot_extral",
    num: rightHotTopicMaxSize,
    extral_num: rightFollowTopicMaxSize,
  })
    .then((res) => {
      hotTags.value = res.topics;
      followTags.value = res.extral_topics ?? [];
      showFollowTopics.value = true;
      loading.value = false;
    })
    .catch((err) => {
      loading.value = false;
    });
};
const formatQuoteNum = (num: number) => {
  if (num >= 1000) {
    return (num / 1000).toFixed(1) + "k";
  }

  return num;
};
const handleSearch = () => {
  router.push({
    name: "home",
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
const formatUserName = (name: string) => {
  return name.length > 5 ? name.substring(0, 5) + '...' : name;
};
watch(
  () => ({
    refreshTopicFollow: store.state.refreshTopicFollow,
    userLogined: store.state.userLogined,
  }),
  (to, from) => {
    if (to.refreshTopicFollow !== from.refreshTopicFollow || to.userLogined) {
      loadHotTags();
    }
  }
);
onMounted(() => {
  loadHotTags();
  locadHeighQuailtyRankingList();
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
    margin-top: 10px;
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
  .ranking-item {
    display: flex;
    align-items: center;
    padding: 2px;
    height: 60px; /* 调整排行榜每个数据块的高度 */

    border-top: 1px solid #eaeaea; /* 添加分割线 */

    .ranking-number {
        margin-right: 8px; /* 调整排序编号与内容的间距 */
        font-size: 16px; /* 设置排序编号的字体大小 */
        margin-top: -20px;
    }
    .ranking-avatar {
      //头像
      width: 28px;
      height: 28px;
      margin-right: 5px;
      img {
        width: 100%;
        height: 100%;
        border-radius: 50%;
      }
    }
    .user-name {
      //用户名
      font-weight: bold;
      font-size: 15px;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
    }

    .ranking-info {
      margin-left: auto; /* 将左边距设置为 auto，使排行榜内容靠右 */
      flex: 1;
      display: flex;
      align-items: center; /* 将 align-items 设置为 center */
      justify-content: flex-end; /* 将 justify-content 设置为 flex-end */

      .name-stats {
        display: flex;
        flex-direction: column;
        align-items: flex-end; /* 将 align-items 设置为 flex-end */
        justify-content: right;
        flex: 1;

        .stats {
          display: flex;
          align-items: center;
          .stat-drop{
            margin-top: -10px;
          }

          .stat-item {
            .stat-value {
              font-size: 1px;
              margin-top: -10px;
            }
          }
        }
      }

      .downloads {
        margin-left: auto;
        text-align: right;
        margin-top: 13px;

        .download-value {
          font-weight: bold;
          font-size: 15px;
        }
      }
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