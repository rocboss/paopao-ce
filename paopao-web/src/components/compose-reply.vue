<template>
    <div class="reply-compose-wrap">
        <div class="reply-switch">
            <span class="show" v-if="!showReply" @click="switchReply(true)">
                回复
            </span>
            <span class="hide" v-if="showReply" @click="switchReply(false)">
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

<script setup>
import { ref } from 'vue';
import { createCommentReply } from '@/api/post';

const props = defineProps({
    commentId: {
        type: Number,
        default: 0,
    },
    atUserid: {
        type: Number,
        default: 0,
    },
    atUsername: {
        type: String,
        default: '',
    },
});
const emit = defineEmits(['reload', 'reset']);
const inputInstRef = ref(null);
const showReply = ref(false);
const replyContent = ref('');
const submitting = ref(false);
const switchReply = (status) => {
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
        text-align: right;
        font-size: 12px;
        margin: 10px 0;

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
</style>