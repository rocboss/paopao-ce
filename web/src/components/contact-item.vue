<template>
    <div class="contact-item" @click="goUserProfile(contact.username)">
        <div class="avatar">
            <n-avatar :size="54" :src="contact.avatar" />
        </div>
        <div class="base-info">
            <div class="username">
                <strong>{{ contact.nickname }}</strong>
                <span> @{{ contact.username }} </span>
            </div>
            <div class="user-info">
                <span class="info-item">UID. {{ contact.user_id }}</span>
                <span class="info-item">{{ formatDate(contact.created_on) }}&nbsp;加入</span>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
const router = useRouter();
import { formatDate } from '@/utils/formatTime';

const props = withDefaults(defineProps<{
    contact: Item.ContactItemProps
}>(), {})
const goUserProfile = (username: string) => {
    router.push({
        name: 'user',
        query: { s: username },
    });
};
</script>

<style lang="less" scoped>
.contact-item {
    display: flex;
    width: 100%;
    padding-left: 16px;
    padding-right: 16px;
    padding-top: 12px;
    padding-bottom: 12px;

    &:hover {
        background: #f7f9f9;
        cursor: pointer;
    }

    .avatar {
        width: 54px;
    }

    .base-info {
        position: relative;
        margin-left: 12px;
        padding-top: 2px;
        width: calc(100% - 66px);

        .username {
            line-height: 16px;
            font-size: 16px;
        }

        .user-info {
            margin-top: 6px;
            .info-item {
                font-size: 14px;
                line-height: 14px;
                margin-right: 8px;
                opacity: 0.75;
            }
        }
    }
}

.dark {
    .contact-item {
        &:hover {
            background: #18181c;
        }
        background-color: rgba(16, 16, 20, 0.75);
    }
}
</style>