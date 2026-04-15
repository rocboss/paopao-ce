<template>
    <div>
        <main-nav title="好友" />

        <n-list class="main-content-wrap" bordered>
            <div v-if="loading && list.length === 0" class="skeleton-wrap">
                <post-skeleton :num="pageSize" />
            </div>
            <div v-else>
                <div class="empty-wrap" v-if="list.length === 0">
                    <n-empty size="large" description="暂无数据" />
                </div>

                <n-list-item class="list-item" v-for="contact in list" :key="contact.user_id">
                     <user-card type="contact" :contact="contact" @send-whisper="onSendWhisper" />
                </n-list-item>
            </div>
            <!-- 私信组件 -->
            <whisper :show="showWhisper" :user="whisperReceiver" @success="whisperSuccess" />
        </n-list>

        <infinite-load-more
            :total-page="totalPage"
            :no-more="noMore"
            complete-text="没有更多好友了"
            @load-more="nextPage"
        />
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { Api } from '@/utils/request';
import { usePagination } from '@/composables/usePagination';
import InfiniteLoadMore from '@/components/infinite-load-more.vue';
import UserCard from '@/components/user-card.vue';

const route = useRoute();
const { loading, noMore, page, pageSize, totalPage } = usePagination(20);
const list = ref<Item.ContactItemProps[]>([]);
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

// 初始化页码
page.value = +(route.query.p as string) || 1;

const onSendWhisper = (user: Item.UserInfo) => {
  whisperReceiver.value = user;
  showWhisper.value = true;
};

const whisperSuccess = () => {
  showWhisper.value = false;
};

const nextPage = () => {
  if (page.value < totalPage.value || totalPage.value == 0) {
    noMore.value = false;
    page.value++;
    loadContacts();
  } else {
    noMore.value = true;
  }
};

onMounted(() => {
  loadContacts();
});

const loadContacts = (scrollToBottom: boolean = false) => {
  if (list.value.length === 0) {
    loading.value = true;
  }
  Api.v1.user.get
    .contacts({
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
</script>

<style lang="less" scoped>
.dark {
    .main-content-wrap, .empty-wrap, .skeleton-wrap {
        background-color: rgba(16, 16, 20, 0.75);
    }
}
</style>