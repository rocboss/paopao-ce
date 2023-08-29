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
      embedded
      :bordered="false"
      size="small"
    >
    <div class="ranking-header">
        <div class="ranking-title">{{ rankingTitles[currentRankingType] }}</div>
        <div class="toggle-button" @click="toggleRankingType">
          {{ rankingTitles[NextRankingType] }} <n-icon :component="ChevronForward" />
        </div>
      </div>
      <n-spin :show="rankloading">
        <div
          class="ranking-item"
          v-for="(item, index) in getCurrentRankingList"
          :key="item.name"
        >
          <div class="ranking-number">{{ index + 1 }}</div>
          <!-- 添加排序编号 -->
          <div class="ranking-avatar">
            <!-- 用户头像 -->
            <img :src="item.avatar"/>
          </div>

          <div class="ranking-info">
            <div style="flex: 1">
              <div class="user-name">
                <router-link
                  class="hash-link"
                  :to="{
                    name: 'user',
                    query: {
                      s: item.name,
                    },
                  }"
                >
                  {{ formatUserName(item.nickname) }}
                </router-link>

                <div class="score">
                  <div class="score-value" v-if="currentRankingType === 'highQuality'">
                    {{ item.comprehensive_score }}
                  </div>
                  <div class="score-value" v-else-if="currentRankingType !== 'highQuality'">
                    {{ item.download }}
                  </div>
                </div>
              </div>

              <div class="name-stats">
                <div class="stats" v-if="currentRankingType === 'highQuality'">
                  <div class="stat-item">
                    <div class="stat-value">📃{{ formatQuoteNumStats(item.post_count || item.download) }}</div>
                  </div>
                  <div class="stat-drop">·</div>
                  <!-- 加粗的点 -->
                  <div class="stat-item">
                    <div class="stat-value">❤️{{ formatQuoteNumStats(item.likes) }}</div>
                  </div>
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
import { getDownloadRank, getHighQuailty, getTags } from "@/api/post";
import { Search } from "@vicons/ionicons5";
import { ChevronForward } from "@vicons/ionicons5";

const hotTags = ref<Item.TagProps[]>([]);
const followTags = ref<Item.TagProps[]>([]);
const loading = ref(false);
const rankloading = ref(false);
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
const rankingList = ref<Item.RankingDataProps[]>([]);
const allDownloadRankingList = ref<Item.RankingDataProps[]>([]);
const DownloadPreWeekRankingList = ref<Item.RankingDataProps[]>([]);
const DownloadPreMonthRankingList = ref<Item.RankingDataProps[]>([]);

//获取排行榜数据
const locadHeighQuailtyRankingList = () => {
  rankloading.value = true;
  getHighQuailty()
    .then((res) => {
      rankingList.value = res.list;
      rankloading.value = false;
    })
    .catch((err) => {
      rankloading.value = false;
    });
};

const loadDowmloadRankingByType = (type: number) => {
  rankloading.value = true;
  getDownloadRank(type)
    .then((res) => {
      if (type ===1 ) allDownloadRankingList.value = res.list;
      if (type === 2) DownloadPreWeekRankingList.value = res.list;
      if (type === 3) DownloadPreMonthRankingList.value = res.list;
      rankloading.value = false;
    })
    .catch((err) => {
      rankloading.value = false;
    });
};

const rankingTitles: { [key: string]: string } = {
  highQuality: "用户榜",
  downloadPreWeek: "下载周榜",
  downloadPreMonth: "下载月榜",
  downloadAll: "下载总榜",
};

const rankingTypes = ['highQuality', 'downloadPreWeek', 'downloadPreMonth', 'downloadAll'];
let currentRankingTypeIndex = 0;
const currentRankingType = ref("highQuality");
const NextRankingType = ref("downloadPreWeek");

const toggleRankingType = () => {
  currentRankingTypeIndex = (currentRankingTypeIndex + 1) % rankingTypes.length;
  currentRankingType.value = rankingTypes[currentRankingTypeIndex];
  NextRankingType.value = rankingTypes[(currentRankingTypeIndex + 1) % rankingTypes.length];
};

//1总 2周 3月
const getCurrentRankingList = computed(() => {
  if (currentRankingType.value === "highQuality") {
    return rankingList.value;
  } else if (currentRankingType.value === "downloadAll") {
    if (allDownloadRankingList.value.length === 0) {
      loadDowmloadRankingByType(1);
    }
    return allDownloadRankingList.value;
  } else if (currentRankingType.value === "downloadPreWeek") {
    if (DownloadPreWeekRankingList.value.length === 0) {
      loadDowmloadRankingByType(2);
    }
    return DownloadPreWeekRankingList.value;
  } else if (currentRankingType.value === "downloadPreMonth") {
    if (DownloadPreMonthRankingList.value.length === 0) {
      loadDowmloadRankingByType(3);
    }
    return DownloadPreMonthRankingList.value;
  }
  return [];
});

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
const formatQuoteNumStats = (num: number) => {
  if (num >= 1000) {
    const formattedNum = (num / 1000).toFixed(1); // Get one decimal place
    return formattedNum + 'k';
  } else {
    return num.toString().padStart(3, '0'); // Ensure 3-digit format
  }
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
  return name.length > 10 ? name.substring(0, 10) + "..." : name;
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
      display: flex;
      align-items: center;
      margin-top: 10px;
    }

    .ranking-info {
      margin-left: auto; /* 将左边距设置为 auto，使排行榜内容靠右 */
      display: flex;
      align-items: center; /* 将 align-items 设置为 center */
      justify-content: flex-end; /* 将 justify-content 设置为 flex-end */
      flex: 1;

      .name-stats {
        display: flex;
        align-items: flex-end; /* 将 align-items 设置为 flex-end */
        justify-content: right;
        flex: 1;

        .stats {
          display: flex;
          align-items: center;
          .stat-drop {
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

      .score {
        margin-left: auto;
        text-align: right;
        align-items: center;

        .score-value {
          font-weight: bold;
          font-size: 15px;
        }
      }
    }
  }
}
.ranking-card {
  margin-bottom: 10px;
}

.ranking-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px;
}

.ranking-title {
  font-size: 16px;
  font-weight: bold;
}

.toggle-button {
  display: flex;
  align-items: center;
  cursor: pointer;
}

.toggle-button n-icon {
  margin-left: 5px;
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