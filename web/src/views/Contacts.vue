<template>
    <div>
        <main-nav title="好友" />

        <n-list class="main-content-wrap" bordered>
            <div v-if="loading" class="skeleton-wrap">
                <post-skeleton :num="pageSize" />
            </div>
            <div v-else>
                <div class="empty-wrap" v-if="list.length === 0">
                    <n-empty size="large" description="暂无数据" />
                </div>

                <n-list-item v-for="contact in list" :key="contact.user_id">
                     <contact-item
                        :contact="contact"
                     />
                </n-list-item>
            </div>
        </n-list>
    </div>

    <div class="pagination-wrap" v-if="totalPage > 0">
        <n-pagination 
            :page="page" 
            @update:page="updatePage"
            :page-slot="!store.state.collapsedRight ? 8 : 5" 
            :page-count="totalPage" />
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { getContacts } from '@/api/post';
import { useStore } from 'vuex';
import { useRoute } from 'vue-router';

const store = useStore();
const route = useRoute();

const loading = ref(false);
const list = ref<Item.ContactItemProps[]>([]);
const page = ref(+(route.query.p as string) || 1);
const pageSize = ref(20);
const totalPage = ref(0);

const updatePage = (p: number) => {
    page.value = p;
    loadContacts();
};

onMounted(() => {
    loadContacts()
});

const loadContacts = (scrollToBottom: boolean = false) => {
    if (list.value.length === 0) {
        loading.value = true;
    }
    getContacts({
        page: page.value,
        page_size: pageSize.value,
    })
        .then((res) => {
            loading.value = false;
            list.value = res.list;
            totalPage.value = Math.ceil(res.pager.total_rows / pageSize.value);

            if (scrollToBottom) {
                setTimeout(() => {
                    window.scrollTo(0, 99999);
                }, 50);
            }
        })
        .catch((err) => {
            loading.value = false;
        });
};
</script>

<style lang="less" scoped>
.pagination-wrap {
    padding: 10px;
    display: flex;
    justify-content: center;
    overflow: hidden;
}
.dark {
    .main-content-wrap, .empty-wrap, .skeleton-wrap {
        background-color: rgba(16, 16, 20, 0.75);
    }
}
</style>