<template>
    <div class="follow-item">
        <n-thing content-indented>
            <template #avatar>
                <n-avatar :size="54" :src="contact.avatar" />
            </template>
            <template #header>
                <span class="nickname-wrap">
                    <router-link
                        @click.stop
                        class="username-link"
                        :to="{
                            name: 'user',
                            query: { s: contact.username },
                        }"
                    >
                        {{ contact.nickname }}
                    </router-link>
                </span>
                 <span class="username-wrap"> @{{ contact.username }} </span>      
                <n-tag
                    v-if="contact.is_following"
                    class="top-tag" type="success" size="small" round>
                    已关注
                </n-tag>
                <div class="user-info">
                    <span class="info-item">
                        UID. {{ contact.user_id }}
                    </span>
                     <span class="info-item">
                        {{ formatDate(contact.created_on) }}&nbsp;加入
                    </span>
                </div>
            </template>
            <template #header-extra>
                <div class="item-header-extra">
                    <n-dropdown
                        placement="bottom-end"
                        trigger="click"
                        size="small"
                        :options="actionOpts"
                        @select="handleAction"
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
        </n-thing>
    </div>
</template>

<script setup lang="ts">
import { h, computed } from 'vue';
import { NIcon } from 'naive-ui';
import type { Component } from 'vue';
import { useDialog, DropdownOption } from 'naive-ui';
import { followUser, unfollowUser } from '@/api/user';
import { formatDate } from '@/utils/formatTime';
import { MoreHorizFilled } from '@vicons/material';
import { PaperPlaneOutline, BodyOutline, WalkOutline } from '@vicons/ionicons5';

const dialog = useDialog();

const emit = defineEmits<{
  (e: 'send-whisper', user: Item.UserInfo): void;
}>();

const renderIcon = (icon: Component) => {
  return () => {
    return h(NIcon, null, {
      default: () => h(icon),
    });
  };
};

const handleFollowUser = () => {
  dialog.success({
    title: '提示',
    content:
      '确定' +
      (props.contact.is_following ? '取消关注 @' : '关注 @') +
      props.contact.username +
      ' 吗？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      if (props.contact.is_following) {
        unfollowUser({
          user_id: props.contact.user_id,
        })
          .then((_res) => {
            window.$message.success('取消关注成功');
            props.contact.is_following = false;
          })
          .catch((err) => {
            console.log(err);
          });
      } else {
        followUser({
          user_id: props.contact.user_id,
        })
          .then((_res) => {
            window.$message.success('关注成功');
            props.contact.is_following = true;
          })
          .catch((err) => {
            console.log(err);
          });
      }
    },
  });
};

const props = withDefaults(
  defineProps<{
    contact: Item.ContactItemProps;
  }>(),
  {},
);

const actionOpts = computed(() => {
  let options: DropdownOption[] = [
    {
      label: '私信 @' + props.contact.username,
      key: 'whisper',
      icon: renderIcon(PaperPlaneOutline),
    },
  ];
  if (props.contact.is_following) {
    options.push({
      label: '取消关注 @' + props.contact.username,
      key: 'unfollow',
      icon: renderIcon(WalkOutline),
    });
  } else {
    options.push({
      label: '关注 @' + props.contact.username,
      key: 'follow',
      icon: renderIcon(BodyOutline),
    });
  }
  return options;
});

const handleAction = (item: 'follow' | 'unfollow' | 'whisper') => {
  switch (item) {
    case 'follow':
    case 'unfollow':
      handleFollowUser();
      break;
    case 'whisper':
      const user: Item.UserInfo = {
        id: props.contact.user_id,
        avatar: props.contact.avatar,
        username: props.contact.username,
        nickname: props.contact.nickname,
        is_admin: false,
        is_friend: true,
        is_following: false,
        created_on: 0,
        follows: 0,
        followings: 0,
        status: 1,
      };
      emit('send-whisper', user);
      break;
    default:
      break;
  }
};
</script>

<style lang="less" scoped>
.follow-item {
    display: border-box;
    width: 100%;
    padding-left: 16px;
    padding-right: 16px;
    padding-top: 12px;
    padding-bottom: 12px;

    &:hover {
        background: #f7f9f9;
    }

    .nickname-wrap {
        line-height: 16px;
        font-size: 16px;
    }
    .username-wrap {
        line-height: 16px;
        font-size: 16px;
    }

    .top-tag {
        transform: scale(0.75);
    }
    .user-info {
        .info-item {
            font-size: 14px;
            line-height: 14px;
            margin-right: 8px;
            opacity: 0.75;
        }
    }
    .item-header-extra {
        display: flex;
        align-items: center;
        opacity: 0.75;
    }
}

.dark {
    .follow-item {
        &:hover {
            background: #18181c;
        }
        background-color: rgba(16, 16, 20, 0.75);
    }
}
</style>