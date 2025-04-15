<template>
    <div class="reply-item">
        <div class="header-wrap">
            <div class="username">
                <router-link class="user-link" :to="{
                    name: 'user',
                    query: { s: props.reply.user.username },
                }">
                    {{ props.reply.user.username }}
                </router-link>
                <span class="reply-name">
                    {{ props.reply.at_user_id > 0 ? '回复' : ':' }}
                </span>

                <router-link class="user-link" :to="{
                    name: 'user',
                    query: { s: props.reply.at_user.username },
                }" v-if="props.reply.at_user_id > 0">
                    {{ props.reply.at_user.username }}
                </router-link>
            </div>
            <div class="timestamp">
                {{ props.reply.ip_loc }}
                <n-popconfirm v-if="
                    store.state.userInfo.is_admin ||
                    store.state.userInfo.id === props.reply.user.id
                " negative-text="取消" positive-text="确认" @positive-click="execDelAction">
                    <template #trigger>
                        <n-button quaternary circle size="tiny" class="del-btn">
                            <template #icon>
                                <n-icon>
                                    <trash />
                                </n-icon>
                            </template>
                        </n-button>
                    </template>
                    是否删除这条回复？
                </n-popconfirm>
            </div>
        </div>

        <div class="base-wrap">
            <div class="content">
                <n-ellipsis expand-trigger="click" line-clamp="5" :tooltip="false">
                    {{ props.reply.content }}
                </n-ellipsis>
            </div>
            <div class="reply-switch">
                <span class="time-item">
                    {{ formatPrettyTime(props.reply.created_on) }}
                </span>

                <div class="actions">
                    <div v-if="!store.state.userLogined" class="action-item" @click.stop="">
                        <n-icon size="medium">
                            <thumb-up-outlined />
                        </n-icon>
                        <span class="upvote-count">{{ thumbsUpCount }}</span>
                    </div>
                    <div v-if="store.state.userLogined" class="action-item hover" @click.stop="handleThumbsUp">
                        <n-icon size="medium">
                            <thumb-up-outlined v-if="!hasThumbsUp" />
                            <thumb-up-twotone v-if="hasThumbsUp" class="show" />
                        </n-icon>
                        <span class="upvote-count">{{ thumbsUpCount>0 ? thumbsUpCount : "赞" }}</span>
                    </div>
                    <div v-if="!store.state.userLogined" class="action-item">
                        <n-icon size="medium">
                            <thumb-down-outlined />
                        </n-icon>
                    </div>
                    <div v-if="store.state.userLogined" class="action-item hover" @click.stop="handleThumbsDown">
                        <n-icon size="medium">
                            <thumb-down-outlined v-if="!hasThumbsDown" />
                            <thumb-down-twotone v-if="hasThumbsDown" class="show" />
                        </n-icon>
                    </div>
                    <span v-if="store.state.userLogined" class="show opacity-item reply-btn" @click="focusReply"> 回复 </span>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useStore } from 'vuex';
import { Trash } from '@vicons/tabler';
import { formatPrettyTime } from '@/utils/formatTime';
import {
  deleteCommentReply,
  thumbsUpTweetReply,
  thumbsDownTweetReply,
} from '@/api/post';
import {
  ThumbUpTwotone,
  ThumbUpOutlined,
  ThumbDownTwotone,
  ThumbDownOutlined,
} from '@vicons/material';
import { YesNoEnum } from '@/utils/IEnum';

const props = withDefaults(
  defineProps<{
    tweetId: number;
    reply: Item.ReplyProps;
  }>(),
  {},
);
const store = useStore();
const emit = defineEmits<{
  (e: 'focusReply', reply: Item.ReplyProps): void;
  (e: 'reload'): void;
}>();

const hasThumbsUp = ref(props.reply.is_thumbs_up == YesNoEnum.YES);
const hasThumbsDown = ref(props.reply.is_thumbs_down == YesNoEnum.YES);
const thumbsUpCount = ref(props.reply.thumbs_up_count);

const handleThumbsUp = () => {
  thumbsUpTweetReply({
    tweet_id: props.tweetId,
    comment_id: props.reply.comment_id,
    reply_id: props.reply.id,
  })
    .then((_res) => {
      hasThumbsUp.value = !hasThumbsUp.value;
      if (hasThumbsUp.value) {
        thumbsUpCount.value++;
        hasThumbsDown.value = false;
      } else {
        thumbsUpCount.value--;
      }
    })
    .catch((err) => {
      console.log(err);
    });
};
const handleThumbsDown = () => {
  thumbsDownTweetReply({
    tweet_id: props.tweetId,
    comment_id: props.reply.comment_id,
    reply_id: props.reply.id,
  })
    .then((_res) => {
      hasThumbsDown.value = !hasThumbsDown.value;
      if (hasThumbsDown.value) {
        if (hasThumbsUp.value) {
          thumbsUpCount.value--;
          hasThumbsUp.value = false;
        }
      }
    })
    .catch((err) => {
      console.log(err);
    });
};
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
            max-width: 50%;
            overflow: hidden;
            text-overflow: ellipsis;
            white-space: nowrap;
        }
    }

    .base-wrap {
        display: block;

        .content {
            width: calc(100%);
            margin-top: 4px;
            font-size: 12px;
            text-align: justify;
            line-height: 2;
        }

        .reply-switch {
            display: flex;
            align-items: center;
            justify-content: space-between;
            font-size: 12px;

            .actions {
                display: flex;
                align-items: center;
                text-align: right;
                font-size: 12px;
                margin: 10px 0;
            }

            .time-item {
                font-size: 12px;
                opacity: 0.75;
                margin-right: 18px;
            }

            .action-item {
                display: flex;
                align-items: center;
                margin-left: 18px;
                opacity: 0.65;

                .upvote-count {
                    margin-left: 4px;
                    font-size: 12px;
                }

                &.hover {
                    cursor: pointer;
                }
            }

            .opacity-item {
                opacity: 0.75;
            }

            .reply-btn {
                margin-left: 18px;
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
}

.dark {
    .reply-item {
        border-bottom: 1px solid #262628;
        background-color: rgba(16, 16, 20, 0.75);

        .base-wrap {
            .reply-switch {
                .show {
                    color: #63e2b7;
                }
            }
        }
    }
}
</style>