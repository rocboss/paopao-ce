<template>
    <div>
        <main-nav title="消息" />

        <n-list class="main-content-wrap messages-wrap" bordered>
             <!-- 私信组件 -->
            <whisper :show="showWhisper" :user="whisperReceiver" @success="whisperSuccess" />
            <div v-if="loading" class="skeleton-wrap">
                <message-skeleton :num="pageSize" />
            </div>
            <div v-else>
                <div class="empty-wrap" v-if="list.length === 0">
                    <n-empty size="large" description="暂无数据" />
                </div>

                <n-list-item v-for="m in list" :key="m.id">
                    <message-item :message="m" @send-whisper="onSendWhisper"/>
                </n-list-item>
            </div>
        </n-list>
        <div class="pagination-wrap" v-if="totalPage > 0">
            <n-pagination
                :page="page"
                @update:page="updatePage"
                :page-slot="!store.state.collapsedRight ? 8 : 5"
                :page-count="totalPage"
            />
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useStore } from 'vuex';
import { useRoute } from 'vue-router';
import { getMessages } from '@/api/user';
const route = useRoute();
const store = useStore();
const loading = ref(false);
const page = ref(+(route.query.p as string) || 1);
const pageSize = ref(10);
const totalPage = ref(0);
const list = ref<Item.MessageProps[]>([]);
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

const onSendWhisper =  (user: Item.UserInfo) => {
    whisperReceiver.value = user;
    showWhisper.value = true;
};
const whisperSuccess = () => {
    showWhisper.value = false;
};
const loadMessages = () => {
    loading.value = true;
    getMessages({
        page: page.value,
        page_size: pageSize.value,
    })
        .then((res) => {
            loading.value = false;
            list.value = res.list;
            totalPage.value = Math.ceil(res.pager.total_rows / pageSize.value);
        })
        .catch((err) => {
            loading.value = false;
        });
};
const updatePage = (p: number) => {
    page.value = p;
    loadMessages();
};
onMounted(() => {
    loadMessages();
});
</script>

<style lang="less" scoped>
.pagination-wrap {
    padding: 10px;
    display: flex;
    justify-content: center;
    overflow: hidden;
}
.dark {
    .empty-wrap {
        background-color: rgba(16, 16, 20, 0.75);
    }
    .messages-wrap, .pagination-wrap {
        background-color: rgba(16, 16, 20, 0.75);
    }
}
</style>