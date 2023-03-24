<template>
    <div class="reply-item">
        <div class="header-wrap">
            <div class="username">
                <router-link
                    class="user-link"
                    :to="{
                        name: 'user',
                        query: { username: props.reply.user.username },
                    }"
                >
                    {{ props.reply.user.username }}
                </router-link>
                <span class="reply-name">
                    {{ props.reply.at_user_id > 0 ? '回复' : ':' }}
                </span>

                <router-link
                    class="user-link"
                    :to="{
                        name: 'user',
                        query: { username: props.reply.at_user.username },
                    }"
                    v-if="props.reply.at_user_id > 0"
                >
                    {{ props.reply.at_user.username }}
                </router-link>
            </div>
            <div class="timestamp">
                {{
                    props.reply.ip_loc
                        ? props.reply.ip_loc + ' · '
                        : props.reply.ip_loc
                }}
                {{ formatPrettyTime(props.reply.created_on) }}
                <n-popconfirm
                    v-if="
                        store.state.userInfo.is_admin ||
                        store.state.userInfo.id === props.reply.user.id
                    "
                    negative-text="取消"
                    positive-text="确认"
                    @positive-click="execDelAction"
                >
                    <template #trigger>
                        <n-button quaternary circle size="tiny" class="del-btn">
                            <template #icon>
                                <n-icon>
                                    <trash />
                                </n-icon>
                            </template>
                        </n-button>
                    </template>
                    是否确认删除？
                </n-popconfirm>
            </div>
        </div>

        <div class="base-wrap">
            <div class="content">{{ props.reply.content }}</div>
            <div class="reply-switch" v-if="store.state.userInfo.id > 0">
                <span class="show" @click="focusReply"> 回复 </span>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useStore } from 'vuex';
import { Trash } from '@vicons/tabler';
import { formatPrettyTime } from '@/utils/formatTime';
import { deleteCommentReply } from '@/api/post';

const props = withDefaults(defineProps<{
    reply: Item.ReplyProps,
}>(), {});
const store = useStore();
const emit = defineEmits<{
    (e: 'focusReply', reply: Item.ReplyProps): void,
    (e: 'reload'): void
}>();

const focusReply = () => {
    emit('focusReply', props.reply);
};
const execDelAction = () => {
    deleteCommentReply({
        id: props.reply.id,
    })
        .then((res) => {
            window.$message.success('删除成功');

            setTimeout(() => {
                emit('reload');
            }, 50);
        })
        .catch((err) => {
            console.log(err);
        });
};
</script>


<style lang="less" scoped>
.reply-item {
    display: flex;
    flex-direction: column;
    font-size: 12px;
    padding: 8px;
    border-bottom: 1px solid #f3f3f3;

    .header-wrap {
        display: flex;
        align-items: center;
        justify-content: space-between;
        .username {
            max-width: 50%;
            overflow: hidden;
            text-overflow: ellipsis;
            white-space: nowrap;
            .reply-name {
                margin: 0 3px;
                opacity: 0.75;
            }
        }
        .timestamp {
            opacity: 0.75;
            text-align: right;
            display: flex;
            align-items: center;
            max-width: 50%;
            overflow: hidden;
            text-overflow: ellipsis;
            white-space: nowrap;
        }
    }
    .base-wrap {
        display: flex;
        .content {
            width: calc(100% - 40px);
            margin-top: 4px;
            font-size: 12px;
            text-align: justify;
            line-height: 2;
        }
        .reply-switch {
            width: 40px;
            text-align: right;
            font-size: 12px;
            margin: 10px 0 0;

            .show {
                color: #18a058;
                cursor: pointer;
            }
            .hide {
                opacity: 0.75;
                cursor: pointer;
            }
        }
    }
}
.dark {
    .reply-item {
        border-bottom: 1px solid #262628;
    }
}
</style>