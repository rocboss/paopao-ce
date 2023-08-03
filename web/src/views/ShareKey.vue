<template>
    <div>
        <!-- <main-nav title="我的分享码" /> -->

        <n-list class="main-content-wrap" bordered>
            <div class="balance-wrap">
                <div class="balance-line">
                    <n-statistic style="color:black; font-weight:bold;">我的分享码</n-statistic>
                </div>
            </div>

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
                        <div>{{ sharekey.description }}</div>
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
                        <div>是否要删除：{{ selectedShareKey.share_key }}</div>
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

const route = useRoute();
const showRecharge = ref(false);
const rechargeQrcode = ref('');
const loading = ref(false);
// 获取 ShareKey
const shareKeys = ref<Item.ShareKeyProps[]>([]);
const selectedShareKey = ref<Item.ShareKeyProps | null>(null);
const page = ref(+(route.query.p as string) || 1);
const pageSize = ref(20);
const totalPage = ref(0);

const loadKeys = () => {
    loading.value = true;
    const token = localStorage.getItem('PAOPAO_TOKEN') || '';
    if(token) {
        getShareKeys(token)
        .then((rsp) => {
            loading.value = false;
            shareKeys.value = rsp.shareKeys;
            window.scrollTo(0, 0);
        })
        .catch((err) => {
            loading.value = false;
        });
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
        });
    }
    showRecharge.value = false;
};
// const updatePage = (p: number) => {
//     page.value = p;
//     loadPosts();
// };
onMounted(() => {
    loadKeys();
});
</script>

<style lang="less" scoped>
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
        color: #18a058;
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
        color: #63e2b7;
    }
    .main-content-wrap {
        background-color: rgba(16, 16, 20, 0.75);
    }
}
</style>