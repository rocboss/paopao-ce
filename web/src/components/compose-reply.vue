<template>
    <div class="reply-compose-wrap">
        <div class="reply-switch">
            <span class="time-item">
                {{ formatPrettyTime(timestamp) }}
            </span>
            <div
                v-if="!store.state.userLogined"
                class="action-item"
            >
                <n-icon size="medium">
                    <thumb-up-outlined/>
                </n-icon>
                <span class="upvote-count">{{ thumbsUpCount }}</span>
            </div>
            <div
                v-if="store.state.userLogined"
                class="action-item hover"
                @click.stop="handleThumbsUp"
            >
                <n-icon size="medium">
                    <thumb-up-outlined v-if="!hasThumbsUp" />
                    <thumb-up-twotone v-if="hasThumbsUp" color="#18a058" />
                </n-icon>
                <span class="upvote-count">{{ thumbsUpCount }}</span>
            </div>
            <div
                v-if="!store.state.userLogined"
                class="action-item"
            >
                <n-icon size="medium">
                    <thumb-down-outlined />
                </n-icon>
            </div>
            <div
                v-if="store.state.userLogined"
                class="action-item hover"
                @click.stop="handleThumbsDown"
            >
                <n-icon size="medium">
                    <thumb-down-outlined v-if="!hasThumbsDown" />
                    <thumb-down-twotone v-if="hasThumbsDown" color="#18a058" />
                </n-icon>
            </div>
            <span class="show" v-if="store.state.userLogined && !showReply" @click="switchReply(true)">
                回复
            </span>
            <span class="hide" v-if="store.state.userLogined && showReply" @click="switchReply(false)">
                取消
            </span>
        </div>

        <div class="reply-input-wrap" v-if="showReply">
            <n-input-group>
                <n-input
                    ref="inputInstRef"
                    size="small"
                    :placeholder="
                        props.atUsername
                            ? '@' + props.atUsername
                            : '请输入回复内容..'
                    "
                    maxlength="100"
                    v-model:value="replyContent"
                    show-count
                    clearable
                />
                <n-button
                    type="primary"
                    size="small"
                    ghost
                    :loading="submitting"
                    @click="submitReply"
                >
                    回复
                </n-button>
            </n-input-group>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useStore } from 'vuex';
import { formatPrettyTime } from '@/utils/formatTime';
import { createCommentReply, thumbsUpTweetComment, thumbsDownTweetComment } from '@/api/post';
import { InputInst } from 'naive-ui';
import {
    ThumbUpTwotone,
    ThumbUpOutlined,
    ThumbDownTwotone,
    ThumbDownOutlined,
CommentFilled,
} from '@vicons/material';

const hasThumbsUp = ref(false)
const hasThumbsDown = ref(false)
const thumbsUpCount = ref(0)

const props = withDefaults(defineProps<{
    timestamp: number,
    tweetId: number,
    commentId: number,
    atUserid: number,
    atUsername: string,
}>(), {
    timestamp: 0,
    tweet_id: 0,
    commentId: 0,
    atUserid: 0,
    atUsername: ''
});
const store = useStore();
const emit = defineEmits<{
    (e: 'reload'): void,
    (e: 'reset'): void
}>();
const inputInstRef = ref<InputInst>();
const showReply = ref(false);
const replyContent = ref('');
const submitting = ref(false);

const handleThumbsUp = () => {
    thumbsUpTweetComment({
        tweet_id: props.tweetId,
        comment_id: props.commentId,
    })
    .then((_res) => {
        hasThumbsUp.value = !hasThumbsUp.value
        if (hasThumbsUp.value) {
            thumbsUpCount.value++
            hasThumbsDown.value = false
        } else {
            thumbsUpCount.value--
        }
    })
    .catch((err) => {
        console.log(err);
    }); 
};
const handleThumbsDown = () => {
    thumbsDownTweetComment({
        tweet_id: props.tweetId,
        comment_id: props.commentId,
    })
    .then((_res) => {
        hasThumbsDown.value = !hasThumbsDown.value
        if (hasThumbsDown.value) {
            if (hasThumbsUp.value) {
                thumbsUpCount.value--
                hasThumbsUp.value = false
            }
        }
    })
    .catch((err) => {
        console.log(err);
    }); 
};
const switchReply = (status: boolean) => {
    showReply.value = status;

    if (status) {
        setTimeout(() => {
            inputInstRef.value?.focus();
        }, 10);
    } else {
        submitting.value = false;
        replyContent.value = '';
        emit('reset');
    }
};
const submitReply = () => {
    submitting.value = true;
    createCommentReply({
        comment_id: props.commentId,
        at_user_id: props.atUserid,
        content: replyContent.value,
    })
        .then((res) => {
            switchReply(false);
            window.$message.success('评论成功');
            emit('reload');
        })
        .catch((err) => {
            submitting.value = false;
        });
};
defineExpose({ switchReply });
</script>

<style lang="less" scoped>
.reply-compose-wrap {
    .reply-switch {
        display: flex;
        align-items: center;
        text-align: right;
        font-size: 12px;
        margin: 10px 0;
        .time-item {
            font-size: 12px;
            opacity: 0.75;
            margin-right: 18px;
        }
        .action-item {
            display: flex;
            align-items: center;
            margin-right: 18px;
            opacity: 0.75;
            .upvote-count {
                margin-left: 4px;
                font-size: 12px;
            }
            &.hover {
                cursor: pointer;
            }
        }
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
.dark {
    .reply-compose-wrap {
        background-color: rgba(16, 16, 20, 0.75);
    }
}
</style>