<template>
    <div>
        <!-- <main-nav title="我的分享码" /> -->

        <n-list class="main-content-wrap" bordered>
            <div class="balance-wrap">
                <div class="balance-line">
                    <n-statistic style="color:black; font-weight:bold;">我的分享码</n-statistic>
                </div>
            </div>
            <template #footer>
                <div class="pagination-wrap" v-if="totalPage > 1">
                    <n-pagination
                        :page="page"
                        @update:page="updatePage"
                        :page-slot="!store.state.collapsedRight ? 8 : 5"
                        :page-count="totalPage"
                    />
                </div>
            </template>

            <div v-if="loading" class="skeleton-wrap">
                <post-skeleton :num="pageSize" />
            </div>
            <div v-else>
                <div class="empty-wrap" v-if="shareKeys.length === 0">
                    <n-empty size="large" description="暂无数据" />
                </div>

                <n-list-item v-for="sharekey in shareKeys" :key="sharekey.share_key">
                    <div class="bill-line">
                        <div>{{ sharekey.share_key }}</div>
                        <div>{{ sharekey.name }}</div>
                        <!-- <div>{{ sharekey.description }}</div> -->
                        <div>
                        <span
                            class="truncated-content"
                            @mouseover="showFullContent(sharekey.share_key)"
                            @mouseleave="hideFullContent"
                        >
                            {{ truncatedDescription(sharekey.description) }}
                        </span>
                        <!-- 气泡弹出框 -->
                        <div
                            class="popover-content"
                            v-show="showFullContentId === sharekey.share_key"
                        >
                            {{ sharekey.description }}
                        </div>
                        </div>
                        <n-button
                                size="small"
                                secondary
                                type="primary"
                                @click="isDelete(sharekey)"
                            >
                                删除
                        </n-button>
                    </div>
                </n-list-item>
            </div>
        </n-list>
        <n-modal v-model:show="showRecharge">
            <n-card
                :bordered="false"
                title="删除"
                role="dialog"
                aria-modal="true"
                style="width: 100%; max-width: 330px"
            >
                <div class="amount-options" v-if="rechargeQrcode.length === 0">
                    <div v-if="selectedShareKey !== null" class="amount-options">
                    <n-space align="baseline">
                        <div>  {{selectedShareKey.name}}  {{ selectedShareKey.share_key }}</div>
                    </n-space>
                </div>
                </div>
                <div v-if="selectedShareKey !== null"
                    style="margin-top: 20px"
                >
                    <n-button
                    style="width: 30%;left: 10%;"
                    type="primary"
                    @click="deleteShareKey(selectedShareKey.share_key)"
                    >
                        确定
                    </n-button>
                    <n-button
                    style="width: 30% ; left: 30%"
                    @click="showRecharge = false"
                    >
                        取消
                    </n-button>
                </div>
            </n-card>
        </n-modal>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useStore } from 'vuex';
import { useRoute } from 'vue-router';
import { getShareKeys, deleteThisShareKey } from '@/api/shareKey';

const store = useStore();
const route = useRoute();
const showRecharge = ref(false);
const rechargeQrcode = ref('');
const loading = ref(false);
// 获取 ShareKey
const shareKeys = ref<Item.ShareKeyProps[]>([]);
const selectedShareKey = ref<Item.ShareKeyProps | null>(null);
const showFullContentId = ref<string | null>(null);
const page = ref(+(route.query.p as string) || 1);
const pageSize = ref(10);
const totalPage = ref(0);

const loadKeys = () => {
    loading.value = true;
    const token = localStorage.getItem('PAOPAO_TOKEN') || '';
    if(token) {
        const params: NetParams.UserGetShareKeys = {
            page: page.value,
            page_size: pageSize.value,
        }
        getShareKeys(params, token)
        .then((rsp) => {
            loading.value = false;
            shareKeys.value = rsp.list;
            totalPage.value = Math.ceil(rsp.pager.total_rows / pageSize.value);
            window.scrollTo(0, 0);
        })
        .catch((err) => {
            loading.value = false;
            store.commit('triggerAuth', true);
            store.commit('userLogout');
        });
    } else {
        loading.value = false;
        store.commit('triggerAuth', true);
        store.commit('userLogout');
    }
};
const isDelete = (sharekey: Item.ShareKeyProps) => {
    selectedShareKey.value = sharekey;
    showRecharge.value = true;
};
const deleteShareKey = (keyToDelete: string) => {
    const params: NetParams.UserDeleteShareKey = {
        share_key: keyToDelete,
    }
    const token = localStorage.getItem('PAOPAO_TOKEN') || '';
    if(token) {
        deleteThisShareKey(params, token)
        .then((rsp) => {
            if(rsp.Status === "DELETE_SUCCESS") {
                window.$message.success('删除成功');
                showRecharge.value = false;
                loadKeys();
            }
            loading.value = false;
            window.scrollTo(0, 0);
        })
        .catch((err) => {
            loading.value = false;
            store.commit('triggerAuth', true);
            store.commit('userLogout');
        });
    } else {
        loading.value = false;
        store.commit('triggerAuth', true);
        store.commit('userLogout');
    }
    showRecharge.value = false;
};

const updatePage = (p: number) => {
    page.value = p;
    loadKeys();
};

const showFullContent = (shareKeyId: string) => {
  showFullContentId.value = shareKeyId;
};

const hideFullContent = () => {
  showFullContentId.value = null;
};

const truncatedDescription = (description: string) => {
  const maxLength = 6;
  if (description.length <= maxLength) {
    return description;
  }
  return `${description.slice(0, maxLength)}...`;
};

onMounted(() => {
    loadKeys();
});
</script>

<style lang="less" scoped>
.popover-content {
  position: absolute;
  background-color: #fff; /* 白色背景 */
  color: #000; /* 弹出框文字颜色 */
  padding: 8px; /* 弹出框内边距 */
  border-radius: 4px; /* 弹出框边框圆角 */
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2); /* 弹出框阴影效果 */
  max-width: 200px; /* 弹出框最大宽度 */
  white-space: pre-wrap; /* 弹出框内容换行 */
  z-index: 999; /* 设置弹出框显示在文字上方，值越大显示越靠上 */

  /* 创建气泡三角形 */
  &::before {
    content: '';
    position: absolute;
    bottom: 100%;
    left: 50%;
    margin-left: -8px; /* 三角形的位置在弹出框中心 */
    border-width: 8px;
    border-style: solid;
    border-color: transparent transparent #fff transparent; /* 透明 上方白色 透明 透明 */
  }
}

/* 鼠标悬浮时显示弹出框 */
.truncated-content:hover + .popover-content {
  display: block;
}

/* 隐藏弹出框 */
.popover-content {
  display: none;
}


.balance-wrap {
    padding: 16px;

    .balance-line {
        display: flex;
        justify-content: space-between;

        .balance-opts {
            display: flex;
            flex-direction: column;
        }
    }
}

.bill-line {
    padding: 16px;
    display: flex;
    justify-content: space-between;
    .income,
    .out {
        font-weight: bold;
    }

    .income {
        color: #063c62;
    }
}
.pagination-wrap {
    padding: 10px;
    display: flex;
    justify-content: center;
    overflow: hidden;
}

.qrcode-wrap {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    .pay-tips {
        margin-top: 16px;
    }
    .pay-sub-tips {
        margin-top: 4px;
        font-size: 12px;
        opacity: 0.75;
        display: flex;
        align-items: center;
    }
}

.dark {
    .income {
        color: #2080f0;
    }
    .main-content-wrap {
        background-color: rgba(16, 16, 20, 0.75);
    }
}
</style>