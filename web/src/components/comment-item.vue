<template>
    <div class="comment-item">
        <n-thing content-indented>
            <template #avatar>
                <n-avatar round :size="30" :src="comment.user.avatar" />
            </template>
            <template #header>
                <span class="nickname-wrap">
                    <router-link
                        @click.stop
                        class="username-link"
                        :to="{
                            name: 'user',
                            query: { username: comment.user.username },
                        }"
                    >
                        {{ comment.user.nickname }}
                    </router-link>
                </span>
                <span class="username-wrap">
                    @{{ comment.user.username }}
                </span>
            </template>
            <template #header-extra>
                <div class="opt-wrap">
                    <span class="timestamp">
                        {{  comment.ip_loc}}
                    </span>

                    <n-popconfirm
                        v-if="
                            store.state.userInfo.is_admin ||
                            store.state.userInfo.id === comment.user.id
                        "
                        negative-text="取消"
                        positive-text="确认"
                        @positive-click="execDelAction"
                    >
                        <template #trigger>
                            <n-button
                                quaternary
                                circle
                                size="tiny"
                                class="del-btn"
                            >
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
            </template>
            <template #description v-if="comment.texts.length > 0">
                <span
                    v-for="content in comment.texts"
                    :key="content.id"
                    class="comment-text"
                    @click.stop="doClickText($event, comment.id)"
                    v-html="parsePostTag(content.content).content"
                >
                </span>
            </template>

            <template #footer>
                <post-image
                    v-if="comment.imgs.length > 0"
                    :imgs="comment.imgs" />
                  <!-- 回复编辑器 -->
                  <compose-reply
                    ref="replyComposeRef"
                    :comment="comment"
                    :at-userid="replyAtUserID"
                    :at-username="replyAtUsername"
                    @reload="reload"
                    @reset="resetReply"
                />
                <!-- 回复列表 -->
                <div class="reply-wrap">
                    <reply-item
                        v-for="reply in comment.replies"
                        :key="reply.id"
                        :reply="reply"
                        :tweet-id="comment.post_id"
                        @focusReply="focusReply"
                        @reload="reload"
                    />
                </div>
            </template>
        </n-thing>
    </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';
import { useStore } from 'vuex';
import { useRouter } from 'vue-router';
import { parsePostTag } from '@/utils/content';
import { Trash } from '@vicons/tabler';
import { deleteComment } from '@/api/post';

const store = useStore();
const router = useRouter();
const replyAtUserID = ref(0);
const replyAtUsername = ref('');
const replyComposeRef = ref();

const emit = defineEmits<{
    (e: 'reload'): void
}>();
const props = withDefaults(defineProps<{
    comment: Item.CommentProps
}>(), {})

const comment = computed(() => {
    let comment: Item.CommentComponentProps = Object.assign(
        {
            texts: [],
            imgs: [],
        },
        props.comment
    );
    comment.contents.map((content :any) => {
        if (+content.type === 1 || +content.type === 2) {
            comment.texts.push(content);
        }
        if (+content.type === 3) {
            comment.imgs.push(content);
        }
    });
    return comment;
});

const doClickText = (e: MouseEvent, id: number | string) => {
    let _target = e.target as any;
    if (_target.dataset.detail) {
        const d = _target.dataset.detail.split(':');
        if (d.length === 2) {
            store.commit('refresh');
            if (d[0] === 'tag') {
                window.$message.warning('评论内的无效话题');
            } else {
                router.push({
                    name: 'user',
                    query: {
                        username: d[1],
                    },
                });
            }
        }
    }
};

const focusReply = (reply: Item.ReplyProps) => {
    replyAtUserID.value = reply.user_id;
    replyAtUsername.value = reply.user?.username || '';
    replyComposeRef.value?.switchReply(true);
};
const reload = () => {
    emit('reload');
};
const resetReply = () => {
    replyAtUserID.value = 0;
    replyAtUsername.value = '';
};

const execDelAction = () => {
    deleteComment({
        id: comment.value.id,
    })
        .then((res) => {
            window.$message.success('删除成功');

            setTimeout(() => {
                reload();
            }, 50);
        })
        .catch((err) => {});
};
</script>

<style lang="less" scoped>
.comment-item {
    width: 100%;
    padding: 16px;
    box-sizing: border-box;

    .nickname-wrap {
        font-size: 14px;
    }
    .username-wrap {
        font-size: 14px;
        opacity: 0.75;
    }

    .opt-wrap {
        display: flex;
        align-items: center;
        .timestamp {
            opacity: 0.75;
            font-size: 12px;
        }
        .del-btn {
            margin-left: 4px;
        }
    }

    .comment-text {
        display: block;
        text-align: justify;
        overflow: hidden;
        white-space: pre-wrap;
        word-break: break-all;
    }

    .opt-item {
        display: flex;
        align-items: center;
        opacity: 0.7;
        .opt-item-icon {
            margin-right: 10px;
        }
    }
}

.reply-wrap {
    margin-top: 10px;
    border-radius: 5px;
    background: #fafafc;

    .reply-item {
        &:last-child {
            border-bottom: none;
        }
    }
}
.dark {
    .reply-wrap {
        background: #18181c;
    }
    .comment-item {
        background-color: rgba(16, 16, 20, 0.75);
    }
}
</style>