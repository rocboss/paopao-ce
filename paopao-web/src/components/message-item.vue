<template>
    <div
        class="message-item"
        :class="{ unread: message.is_read === 0 }"
        @click="handleReadMessage(message)"
    >
        <n-thing content-indented>
            <template #header>
                <div class="sender-wrap">
                    <n-avatar
                        :size="22"
                        :src="
                            message.sender_user.id > 0
                                ? message.sender_user.avatar
                                : defaultavatar
                        "
                    />
                    <span class="nickname" v-if="message.sender_user.id > 0">
                        <router-link
                            @click.stop
                            class="username-link"
                            :to="{
                                name: 'user',
                                query: {
                                    username: message.sender_user.username,
                                },
                            }"
                        >
                            {{ message.sender_user.nickname }}
                        </router-link>
                        <span class="username">
                            @{{ message.sender_user.username }}
                        </span>
                    </span>
                    <span class="nickname" v-else> 系统 </span>
                </div>
            </template>
            <template #header-extra>
                <span class="timestamp">
                    <n-badge v-if="message.is_read === 0" dot processing />
                    <span class="timestamp-txt">
                        {{ formatRelativeTime(message.created_on) }}
                    </span>
                </span>
            </template>
            <template #description>
                <n-alert
                    :show-icon="false"
                    class="breif-wrap"
                    :type="message.is_read > 0 ? 'default' : 'success'"
                >
                    <div class="breif-content">
                        {{ message.breif }}

                        <span
                            @click.stop="viewDetail(message)"
                            class="hash-link view-link"
                        >
                            <n-icon><share-outline /></n-icon> 查看详情
                        </span>
                    </div>
                </n-alert>
            </template>
        </n-thing>
    </div>
</template>

<script setup>
import { useRouter } from 'vue-router';
import { ShareOutline } from '@vicons/ionicons5';
import { readMessage } from '@/api/user';
import { formatRelativeTime } from '@/utils/formatTime';

const defaultavatar =
    'https://assets.paopao.info/public/avatar/default/admin.png';

const router = useRouter();
const props = defineProps({
    message: {
        type: Object,
        default: () => {},
    },
});
const viewDetail = (message) => {
    handleReadMessage(message);
    if (message.type === 1 || message.type === 2 || message.type === 3) {
        if (message.post && message.post.id > 0) {
            router.push({
                name: 'post',
                query: {
                    id: message.post_id,
                },
            });
        } else {
            window.$message.error('该动态已被删除');
        }
    }
};
const handleReadMessage = (message) => {
    if (message.is_read === 0) {
        readMessage({
            id: message.id,
        })
            .then((res) => {
                message.is_read = 1;
            })
            .catch((err) => {
                console.log(res);
            });
    }
};
</script>

<style lang="less" scoped>
.message-item {
    padding: 16px;

    &.unread {
        background: #fcfffc;
    }

    .sender-wrap {
        display: flex;
        align-items: center;

        .nickname {
            margin-left: 10px;
        }
        .username {
            opacity: 0.75;
            font-size: 14px;
        }
    }

    .timestamp {
        opacity: 0.75;
        font-size: 12px;
        display: flex;
        align-items: center;
        .timestamp-txt {
            margin-left: 6px;
        }
    }

    .breif-wrap {
        margin-top: 10px;
        .breif-content {
            display: flex;
            width: 100%;
        }
    }

    .view-link {
        margin-left: 8px;
        display: flex;
        align-items: center;
    }
}
.dark {
    .message-item {
        &.unread {
            background: #0f180b;
        }
    }
}
</style>