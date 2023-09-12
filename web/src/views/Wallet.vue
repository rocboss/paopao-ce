<template>
  <div>
    <main-nav title="订阅" />

    <n-list class="main-content-wrap" bordered>
      <div class="balance-wrap">
        <div class="balance-line">
          <!-- 账户余额 (元) -->
          <n-statistic label="订阅有效期 (天)"
            ><n-number-animation
              :from="0.0"
              :to="timecal(store.state.userInfo.balance) || 0"
              :duration="500"
              :precision="2"
            />
          </n-statistic>
          <div class="balance-opts">
            <n-space vertical>
              <n-button
                size="small"
                secondary
                type="primary"
                @click="doRecharge"
              >
                充值
              </n-button>
              <!-- <n-button
                size="small"
                secondary
                type="tertiary"
                @click="doWithdraw"
              >
                客服
              </n-button> -->
            </n-space>
          </div>
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
        <div class="empty-wrap" v-if="list.length === 0">
          <n-empty size="large" description="暂无数据" />
        </div>

        <n-list-item v-for="bill in list" :key="bill.id">
          <div class="bill-line">
            <div>NO.{{ bill.id }}</div>
            <div>{{ bill.reason }}</div>
            <div
              :class="{
                income: bill.change_amount >= 0,
                out: bill.change_amount < 0,
              }"
            >
              {{
                (bill.change_amount > 0 ? "+" : "") +
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
    <n-modal v-model:show="showRecharge">
      <n-card
        :bordered="false"
        title="请选择购买时长-年付优惠"
        role="dialog"
        aria-modal="true"
        style="width: 100%; max-width: 330px"
      >
        <div class="amount-options" v-if="rechargeQrcode.length === 0">
          <n-space align="baseline">
            <n-button
              v-for="amount in openAmounts"
              :key="amount"
              size="small"
              secondary
              :type="selectedRechargeAmount === amount ? 'info' : 'default'"
              @click.stop="selectedRechargeAmount = amount"
            >
              {{ amount / 100 }}天
            </n-button>
          </n-space>
        </div>
        <div
          v-if="selectedRechargeAmount > 0 && rechargeQrcode.length === 0"
          style="margin-top: 10px"
        >
          <n-button
            :loading="recharging"
            strong
            secondary
            type="info"
            style="width: 100%"
            @click="handleRecharge"
          >
            <template #icon>
              <n-icon><logo-alipay /></n-icon>
            </template>
            前往支付
          </n-button>
        </div>
        <div class="qrcode-wrap" v-show="rechargeQrcode.length > 0">
          <canvas id="qrcode-container"></canvas>
          <div class="pay-tips">
            请使用支付宝扫码支付{{
              (selectedRechargeAmount / 100).toFixed(2)
            }}天
          </div>
          <div class="pay-sub-tips">
            <n-badge :value="100" type="info" dot processing />
            <span style="margin-left: 6px"> 支付结果实时同步中... </span>
          </div>
        </div>
      </n-card>
    </n-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useStore } from "vuex";
import { useRoute } from "vue-router";
import QRCode from "qrcode";
import { LogoAlipay } from "@vicons/ionicons5";
import { userInfo } from "@/api/auth";
import { reqRecharge, getRecharge, getBills } from "@/api/user";
import { formatRelativeTime } from "@/utils/formatTime";
import MainNav from "@/components/main-nav.vue";

const store = useStore();
const route = useRoute();
const showRecharge = ref(false);
const selectedRechargeAmount = ref(3000);
const recharging = ref(false);
const rechargeQrcode = ref("");
const loading = ref(false);
const list = ref<Item.BillProps[]>([]);
const page = ref(+(route.query.p as string) || 1);
const pageSize = ref(20);
const totalPage = ref(0);
const openAmounts = ref([3000, 18500, 36500]);
// const timecal = (timestamp: number) => { //传递时间戳
//   const timestampInSeconds = Math.floor(new Date().getTime() / 1000); // 输出当前时间的时间戳（以秒为单位）
// //获取时间差
// const timeDifference = timestampInSeconds - timestamp;
//   const days = Math.floor(timeDifference / (24 * 3600)); // 计算出相差天数
//   //determine the days >0
//   if (days > 0) {
//     return days;
//   }
//   return 0;
// };
const timecal = (timestamp: number) => {
  const timestampInSeconds = Math.floor(Date.now() / 1000);
  const timeDifference = timestamp - timestampInSeconds;
  const days = Math.floor(timeDifference / (24 * 3600)) + 1;
  return Math.max(0, days); // 确保返回值不会是负数
};
////获取时间差
const loadPosts = () => {
  loading.value = true;
  getBills({
    page: page.value,
    page_size: pageSize.value,
  })
    .then((rsp) => {
      loading.value = false;
      list.value = rsp.list;
      totalPage.value = Math.ceil(rsp.pager.total_rows / pageSize.value);

      window.scrollTo(0, 0);
    })
    .catch((err) => {
      loading.value = false;
    });
};
const updatePage = (p: number) => {
  page.value = p;
  loadPosts();
};
const loadWallet = () => {
  // 获取最新
  const token = localStorage.getItem("AIMO_TOKEN") || "";
  if (token) {
    userInfo(token)
      .then((res) => {
        store.commit("updateUserinfo", res);
        store.commit("triggerAuth", false);
        loadPosts();
      })
      .catch((err) => {
        store.commit("triggerAuth", true);
        store.commit("userLogout");
      });
  } else {
    store.commit("triggerAuth", true);
    store.commit("userLogout");
  }
};
const doRecharge = () => {
  showRecharge.value = true;
};
const handleRecharge = (amount: any) => {
  recharging.value = true;
  reqRecharge({
    amount: selectedRechargeAmount.value,
  })
    .then((res) => {
      recharging.value = false;
      rechargeQrcode.value = res.pay;

      // 生成二维码
      QRCode.toCanvas(document.querySelector("#qrcode-container"), res.pay, {
        width: 150,
        margin: 2,
      });

      const s = setInterval(() => {
        getRecharge({
          id: res.id,
        })
          .then((res) => {
            if (res.status === "TRADE_SUCCESS") {
              clearInterval(s);
              window.$message.success("充值成功");

              showRecharge.value = false;
              rechargeQrcode.value = "";

              loadWallet();
            }
          })
          .catch((err) => {
            console.log(err);
          });
      }, 2000);
    })
    .catch((err) => {
      recharging.value = false;
    });
};
const doWithdraw = () => {
  if (store.state.userInfo.balance == 0) {
    window.$message.warning("售前客服"); //这里可以改成售前客服
    window.location.href = "http://pre-sales.zhai-zhai.com";
  } else {
    window.$message.warning("售后客服"); //这里可以改成售后客服
    window.location.href = "http://after-sales.zhai-zhai.com";
  }
};
onMounted(() => {
  loadWallet();
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