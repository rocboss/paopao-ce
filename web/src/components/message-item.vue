<template>
    <div class="message-item" :class="{ unread: message.is_read === 0 }" @click="handleReadMessage(message)">
        <n-thing content-indented>
            <template #avatar>
                <n-avatar round :size="30" :src="
                    message.sender_user.id > 0
                        ? message.sender_user.avatar
                        : defaultavatar
                " />
            </template>
            <template #header>
                <div class="sender-wrap">
                    <span class="nickname" v-if="message.sender_user.id > 0">
                        <router-link @click.stop class="username-link" :to="{
                            name: 'user',
                            query: {
                                username: message.sender_user.username,
                            },
                        }">
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
                <n-alert :show-icon="false" class="brief-wrap" :type="message.is_read > 0 ? 'default' : 'success'">
                    <div v-if="message.type != 4" class="brief-content">
                        {{ message.brief }}
                        <span v-if="message.type === 1 || message.type === 2 || message.type === 3"
                            @click.stop="viewDetail(message)" class="hash-link view-link">
                            <n-icon>
                                <share-outline />
                            </n-icon> 查看详情
                        </span>
                    </div>

                    <div v-if="message.type === 4" class="whisper-content-wrap">
                        {{ message.content }}
                    </div>

                    <div v-if="message.type === 5" class="requesting-friend-wrap">
                        {{ message.content }}
                        <span v-if="message.reply_id === 1" @click.stop="agreeAddFriend(message)"
                            class="hash-link view-link">
                            <n-icon>
                                <checkmark-outline />
                            </n-icon> 同意
                        </span>
                        <span v-if="message.reply_id === 1" @click.stop="rejectAddFriend(message)"
                            class="hash-link view-link">
                            <n-icon>
                                <close-outline />
                            </n-icon> 拒绝
                        </span>
                        <span v-if="message.reply_id === 2" class="status-info">
                            <n-icon>
                                <checkmark-done-outline />
                            </n-icon> 已同意
                        </span>
                        <span v-if="message.reply_id === 3" class="status-info">
                            <n-icon>
                                <close-outline />
                            </n-icon> 已拒绝
                        </span>
                    </div>
                </n-alert>
            </template>
        </n-thing>
    </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { ShareOutline, CheckmarkOutline, CloseOutline, CheckmarkDoneOutline } from '@vicons/ionicons5';
import { readMessage, addFriend, rejectFriend } from '@/api/user';
import { formatRelativeTime } from '@/utils/formatTime';

const defaultavatar = 'https://assets.paopao.info/public/avatar/default/admin.png';

const router = useRouter();
const props = withDefaults(
    defineProps<{
        message: Item.MessageProps;
    }>(),
    {}
);

const viewDetail = (message: Item.MessageProps) => {
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

const agreeAddFriend = (message: Item.MessageProps) => {
    handleReadMessage(message);
    addFriend({
        user_id: message.sender_user_id,
    })
        .then((res) => {
            message.reply_id = 2;
            window.$message.success('已同意添加好友');
        })
        .catch((err) => {
            console.log(err);
        });
}

const rejectAddFriend = (message: Item.MessageProps) => {
    handleReadMessage(message);
    rejectFriend({
        user_id: message.sender_user_id,
    })
        .then((res) => {
            message.reply_id = 3;
            window.$message.success('已拒绝添加好友');
        })
        .catch((err) => {
            console.log(err);
        });
}

const handleReadMessage = (message: Item.MessageProps) => {
    if (message.is_read === 0) {
        readMessage({
            id: message.id,
        })
            .then((res) => {
                message.is_read = 1;
            })
            .catch((err) => {
                console.log(err);
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

    .brief-wrap {
        margin-top: 10px;

        .brief-content {
            display: flex;
            width: 100%;
        }

        .whisper-content-wrap {
            display: flex;
            width: 100%;
        }

        .requesting-friend-wrap {
            display: flex;
            width: 100%;
        }
    }

    .view-link {
        margin-left: 8px;
        display: flex;
        align-items: center;
    }

    .status-info {
        margin-left: 8px;
        align-items: center;
    }
}

.dark {
    .message-item {
        &.unread {
            background: #0f180b;
        }
        .brief-wrap {
            background-color: #18181c;
        }
        background-color: rgba(16, 16, 20, 0.75);
    }
}
</style>