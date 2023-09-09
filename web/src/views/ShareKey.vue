<template>
    <div>
      <n-list class="main-content-wrap" bordered>
        <div class="balance-wrap">
          <n-statistic style="color: black; font-weight: bold">我的分享码</n-statistic>
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
  
          <n-table :bordered="false" :single-line="false" striped>
            <thead>
              <tr>
                <th>分享码</th>
                <th>名称</th>
                <th>描述</th>
                <th>操作</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="sharekey in shareKeys" :key="sharekey.share_key">
                <td @click="copyToClipboard(sharekey.share_key)" class="copy-code">
                  {{ sharekey.share_key }}
                </td>
                <td>{{ sharekey.name }}</td>
                <td><n-ellipsis style="max-width: 240px">{{ sharekey.description }}</n-ellipsis></td>
                <td>
                  <n-button size="small" secondary type="primary" @click="isDelete(sharekey)">删除</n-button>
                </td>
              </tr>
            </tbody>
          </n-table>
        </div>
      </n-list>
      
      <n-modal v-model:show="showRecharge">
        <n-card :bordered="false" title="删除" role="dialog" aria-modal="true" style="width: 100%; max-width: 330px">
          <div class="amount-options" v-if="rechargeQrcode.length === 0 && selectedShareKey">
            <n-space align="baseline">
              <div>删除分享码会影响排行榜分数评定,是否删除:</div>
              <div>{{ selectedShareKey.name }} {{ selectedShareKey.share_key }}</div>
            </n-space>
          </div>
          <div v-if="selectedShareKey" style="margin-top: 20px">
            <n-button style="width: 30%; left: 10%" type="primary" @click="deleteShareKey(selectedShareKey.share_key)">确定</n-button>
            <n-button style="width: 30%; left: 30%" @click="showRecharge = false">取消</n-button>
          </div>
        </n-card>
      </n-modal>
    </div>
  </template>  

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useStore } from "vuex";
import { useRoute } from "vue-router";
import { getShareKeys, deleteThisShareKey } from "@/api/shareKey";

const store = useStore();
const route = useRoute();
const showRecharge = ref(false);
const rechargeQrcode = ref("");
const loading = ref(false);
// 获取 ShareKey
const shareKeys = ref<Item.ShareKeyProps[]>([]);
const selectedShareKey = ref<Item.ShareKeyProps | null>(null);
const showFullContentId = ref<string | null>(null);
const page = ref(+(route.query.p as string) || 1);
const pageSize = ref(10);
const totalPage = ref(0);

const copyToClipboard = (text: any) => {
  const textarea = document.createElement('textarea');
  textarea.value = text;
  document.body.appendChild(textarea);
  textarea.select();
  document.execCommand('copy');
  document.body.removeChild(textarea);
  window.$message.success("复制成功，快去分享吧~");
};

const loadKeys = () => {
  loading.value = true;
  const token = localStorage.getItem("PAOPAO_TOKEN") || "";
  if (token) {
    const params: NetParams.UserGetShareKeys = {
      page: page.value,
      page_size: pageSize.value,
    };
    getShareKeys(params, token)
      .then((rsp) => {
        loading.value = false;
        shareKeys.value = rsp.list;
        totalPage.value = Math.ceil(rsp.pager.total_rows / pageSize.value);
        window.scrollTo(0, 0);
      })
      .catch((err) => {
        loading.value = false;
        store.commit("triggerAuth", true);
        store.commit("userLogout");
      });
  } else {
    loading.value = false;
    store.commit("triggerAuth", true);
    store.commit("userLogout");
  }
};
const isDelete = (sharekey: Item.ShareKeyProps) => {
  selectedShareKey.value = sharekey;
  showRecharge.value = true;
};
const deleteShareKey = (keyToDelete: string) => {
  const params: NetParams.UserDeleteShareKey = {
    share_key: keyToDelete,
  };
  const token = localStorage.getItem("PAOPAO_TOKEN") || "";
  if (token) {
    deleteThisShareKey(params, token)
      .then((rsp) => {
        if (rsp.Status === "DELETE_SUCCESS") {
          window.$message.success("删除成功");
          showRecharge.value = false;
          loadKeys();
        }
        loading.value = false;
        window.scrollTo(0, 0);
      })
      .catch((err) => {
        loading.value = false;
        store.commit("triggerAuth", true);
        store.commit("userLogout");
      });
  } else {
    loading.value = false;
    store.commit("triggerAuth", true);
    store.commit("userLogout");
  }
  showRecharge.value = false;
};

const updatePage = (p: number) => {
  page.value = p;
  loadKeys();
};

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
  }
}

.pagination-wrap {
  padding: 10px;
  display: flex;
  justify-content: center;
}

.copy-code {
  cursor: pointer;
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