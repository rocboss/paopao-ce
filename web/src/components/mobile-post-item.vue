<template>
    <div class="post-item">
        <n-thing content-indented>
            <template #avatar>
                <n-avatar round :size="30" :src="post.user.avatar" />
            </template>
            <template #header>
                    <span class="nickname-wrap">
                        <router-link
                            @click.stop
                            class="username-link"
                            :to="{
                                name: 'user',
                                query: { s: post.user.username },
                            }"
                        >
                            {{ post.user.nickname }}
                        </router-link>
                    </span>
                    <span class="username-wrap"> @{{ post.user.username }} </span>
                    <n-tag
                        v-if="post.is_top"
                        class="top-tag"
                        type="warning"
                        size="small"
                        round
                    >
                        置顶
                    </n-tag>
                    <n-tag
                        v-if="post.visibility == 1"
                        class="top-tag"
                        type="error"
                        size="small"
                        round
                    >
                        私密
                    </n-tag>
                    <n-tag
                        v-if="post.visibility == 2"
                        class="top-tag"
                        type="info"
                        size="small"
                        round
                    >
                        好友可见
                    </n-tag>
                    <div>
                        <span class="timestamp-mobile">
                            {{ formatPrettyDate(post.created_on) }} {{ post.ip_loc }}
                        </span>
                    </div>
            </template>
            <template #header-extra>
                <div class="item-header-extra">
                    <n-dropdown
                        placement="bottom-end"
                        trigger="click"
                        size="small"
                        :options="tweetOptions"
                        @select="handleTweetAction"
                    >
                        <n-button quaternary circle>
                            <template #icon>
                                <n-icon>
                                    <more-horiz-filled />
                                </n-icon>
                            </template>
                        </n-button>
                    </n-dropdown>
                </div>
            </template>
            <template #description v-if="post.texts.length > 0">
                <div @click="goPostDetail(post.id)">
                    <span v-for="content in post.texts"
                        :key="content.id"
                        class="post-text"
                        @click.stop="doClickText($event, post.id)"
                        v-html="preparePost(content.content, '展开', '收起', store.state.profile.tweetMobileEllipsisSize, inFoldStyle)"
                    ></span>
                </div>
            </template>

            <template #footer>
                <post-attachment 
                    v-if="post.attachments.length > 0"
                    :attachments="post.attachments" />
                <post-attachment
                    v-if="post.charge_attachments.length > 0"
                    :attachments="post.charge_attachments"
                    :price="post.attachment_price"
                />
                <post-image
                    v-if="post.imgs.length > 0"
                    :imgs="post.imgs" />
                <post-video
                    v-if="post.videos.length > 0"
                    :videos="post.videos" />
                <post-link
                    v-if="post.links.length > 0"
                    :links="post.links" />
            </template>
            <template #action>
                <n-space justify="space-between">
                    <div class="opt-item" @click.stop="handlePostStar">
                        <n-icon size="18" class="opt-item-icon">
                            <heart-outline />
                        </n-icon>
                        {{ post.upvote_count }}
                    </div>
                    <div class="opt-item" @click.stop="goPostDetail(post.id)">
                        <n-icon size="18" class="opt-item-icon">
                            <chatbox-outline />
                        </n-icon>
                        {{ post.comment_count }}
                    </div>
                    <div class="opt-item" @click.stop="handlePostCollection">
                        <n-icon size="18" class="opt-item-icon">
                            <bookmark-outline />
                        </n-icon>
                        {{ post.collection_count }}
                    </div>
                </n-space>
            </template>
        </n-thing>
    </div>
</template>

<script setup lang="ts">
import { h, ref, computed } from 'vue';
import type { Component } from 'vue';
import { NIcon } from 'naive-ui';
import { useStore } from 'vuex';
import type { DropdownOption } from 'naive-ui';
import { useRouter } from 'vue-router';
import { formatPrettyDate } from '@/utils/formatTime';
import { preparePost } from '@/utils/content';
import { postStar, postCollection } from '@/api/post';
import {
  PaperPlaneOutline,
  HeartOutline,
  BookmarkOutline,
  ChatboxOutline,
  ShareSocialOutline,
  PersonAddOutline,
  PersonRemoveOutline,
  BodyOutline,
  WalkOutline,
} from '@vicons/ionicons5';
import { MoreHorizFilled } from '@vicons/material';
import copy from 'copy-to-clipboard';

const router = useRouter();
const store = useStore();
const inFoldStyle = ref<boolean>(true);
const props = withDefaults(
  defineProps<{
    post: Item.PostProps;
    isOwner: boolean;
    addFriendAction: boolean;
    addFollowAction: boolean;
  }>(),
  {},
);

const emit = defineEmits<{
  (e: 'send-whisper', user: Item.UserInfo): void;
  (e: 'handle-follow-action', user: Item.PostProps): void;
  (e: 'handle-friend-action', user: Item.PostProps): void;
}>();

const renderIcon = (icon: Component) => {
  return () => {
    return h(NIcon, null, {
      default: () => h(icon),
    });
  };
};

const tweetOptions = computed(() => {
  let options: DropdownOption[] = [];
  if (!props.isOwner) {
    options.push({
      label: '私信 @' + props.post.user.username,
      key: 'whisper',
      icon: renderIcon(PaperPlaneOutline),
    });
  }
  if (!props.isOwner && props.addFollowAction) {
    if (props.post.user.is_following) {
      options.push({
        label: '取消关注 @' + props.post.user.username,
        key: 'unfollow',
        icon: renderIcon(WalkOutline),
      });
    } else {
      options.push({
        label: '关注 @' + props.post.user.username,
        key: 'follow',
        icon: renderIcon(BodyOutline),
      });
    }
  }
  if (!props.isOwner && props.addFriendAction) {
    if (props.post.user.is_friend) {
      options.push({
        label: '删除好友 @' + props.post.user.username,
        key: 'delete',
        icon: renderIcon(PersonRemoveOutline),
      });
    } else {
      options.push({
        label: '添加朋友 @' + props.post.user.username,
        key: 'requesting',
        icon: renderIcon(PersonAddOutline),
      });
    }
  }
  options.push({
    label: '复制链接',
    key: 'copyTweetLink',
    icon: renderIcon(ShareSocialOutline),
  });
  return options;
});

const handleTweetAction = async (
  item:
    | 'copyTweetLink'
    | 'whisper'
    | 'follow'
    | 'unfollow'
    | 'delete'
    | 'requesting',
) => {
  switch (item) {
    case 'copyTweetLink':
      copy(
        `${window.location.origin}/#/post?id=${post.value.id}&share=copy_link&t=${new Date().getTime()}`,
      );
      window.$message.success('链接已复制到剪贴板');
      break;
    case 'whisper':
      emit('send-whisper', props.post.user);
      break;
    case 'delete':
    case 'requesting':
      emit('handle-friend-action', props.post);
      break;
    case 'follow':
    case 'unfollow':
      emit('handle-follow-action', props.post);
      break;
    default:
      break;
  }
};

const post = computed({
  get: () => {
    let post: Item.PostComponentProps = Object.assign(
      {
        texts: [],
        imgs: [],
        videos: [],
        links: [],
        attachments: [],
        charge_attachments: [],
      },
      props.post,
    );
    post.contents.map((content) => {
      if (+content.type === 1 || +content.type === 2) {
        post.texts.push(content);
      }
      if (+content.type === 3) {
        post.imgs.push(content);
      }
      if (+content.type === 4) {
        post.videos.push(content);
      }
      if (+content.type === 6) {
        post.links.push(content);
      }
      if (+content.type === 7) {
        post.attachments.push(content);
      }
      if (+content.type === 8) {
        post.charge_attachments.push(content);
      }
    });
    return post;
  },
  set: (newVal) => {
    props.post.upvote_count = newVal.upvote_count;
    props.post.collection_count = newVal.collection_count;
  },
});

const handlePostStar = () => {
  postStar({
    id: post.value.id,
  })
    .then((res) => {
      if (res.status) {
        post.value = {
          ...post.value,
          upvote_count: post.value.upvote_count + 1,
        };
      } else {
        post.value = {
          ...post.value,
          upvote_count:
            post.value.upvote_count > 0 ? post.value.upvote_count - 1 : 0,
        };
      }
    })
    .catch((err) => {
      console.log(err);
    });
};

const handlePostCollection = () => {
  postCollection({
    id: post.value.id,
  })
    .then((res) => {
      if (res.status) {
        post.value = {
          ...post.value,
          collection_count: post.value.collection_count + 1,
        };
      } else {
        post.value = {
          ...post.value,
          collection_count:
            post.value.collection_count > 0
              ? post.value.collection_count - 1
              : 0,
        };
      }
    })
    .catch((err) => {
      console.log(err);
    });
};

const goPostDetail = (id: number) => {
  router.push({
    name: 'post',
    query: {
      id,
    },
  });
};
const doClickText = (e: MouseEvent, id: number) => {
  const detail = (e.target as any).dataset.detail;
  if (detail && detail !== 'post') {
    const d = detail.split(':');
    if (d.length === 2) {
      store.commit('refresh');
      if (d[0] === 'tag') {
        router.push({
          name: 'home',
          query: {
            q: d[1],
            t: 'tag',
          },
        });
      } else {
        router.push({
          name: 'user',
          query: {
            s: d[1],
          },
        });
      }
    }
  } else if (detail && detail === 'post') {
    inFoldStyle.value = !inFoldStyle.value;
  } else {
    goPostDetail(id);
  }
};
</script>

<style lang="less">
.post-item {
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

    .top-tag {
        transform: scale(0.75);
    }
    .timestamp-mobile {
        margin-top: 2px;
        opacity: 0.75;
        font-size: 11px;
    }
    .item-header-extra {
        display: flex;
        align-items: center;
        opacity: 0.75;
        .timestamp {
            font-size: 12px;
        }
    }
    .post-text {
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
    &:hover {
        background: #f7f9f9;
        cursor: pointer;
    }

    .n-thing-avatar {
        margin-top: 0;
    }
    .n-thing-header {
        line-height: 16px;
        margin-bottom: 8px !important;
    }
}
.dark {
    .post-item {
        &:hover {
            background: #18181c;
        }
        background-color: rgba(16, 16, 20, 0.75);
    }
}
</style>