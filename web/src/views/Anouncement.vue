<template>
    <div>
        <main-nav title="公告" />

        <n-list class="main-content-wrap" bordered>

            <template #footer>
                <div class="pagination-wrap" v-if="totalPage > 1">
                    <n-pagination :page="page" @update:page="updatePage"
                        :page-slot="!store.state.collapsedRight ? 8 : 5" :page-count="totalPage" />
                </div>
            </template>

            <div v-if="loading" class="skeleton-wrap">
                <post-skeleton :num="pageSize" />
            </div>
            <div v-else>
                <div class="empty-wrap" v-if="list.length === 0">
                    <n-empty size="large" description="暂无数据" />
                </div>

                <n-list-item v-for="bill in list" :key="bill.id">
                    <div class="bill-line">
                        <div>NO.{{ bill.id }}</div>
                        <div>{{ bill.reason }}</div>
                        <div :class="{
                            income: bill.change_amount >= 0,
                            out: bill.change_amount < 0,
                        }">
                            {{
                                    (bill.change_amount > 0 ? '+' : '') +
                                    (bill.change_amount / 100).toFixed(2)
                            }}
                        </div>
                        <div>
                            {{ formatRelativeTime(bill.created_on) }}
                        </div>
                    </div>
                </n-list-item>
            </div>
        </n-list>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useStore } from 'vuex';
import { useRoute } from 'vue-router';
import { formatRelativeTime } from '@/utils/formatTime';

const store = useStore();
const route = useRoute();

const loading = ref(false);
const list = ref<Item.BillProps[]>([]);
const page = ref(+(route.query.p as string) || 1);
const pageSize = ref(20);
const totalPage = ref(0);

const updatePage = (p: number) => {
  page.value = p;
  // TODO
};

onMounted(() => {
  // TODO
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
}
</style>