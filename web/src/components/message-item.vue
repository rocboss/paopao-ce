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
          <span v-if="message.sender_user.id > 0" class="nickname">
            <router-link
              class="username-link"
              :to="{
                name: 'user',
                query: {
                  username: message.sender_user.username,
                },
              }"
              @click.stop
            >
              {{ message.sender_user.nickname }}
            </router-link>
            <span class="username"> @{{ message.sender_user.username }} </span>
          </span>
          <span v-else class="nickname"> 系统 </span>
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
              v-if="message.type !== 4"
              class="hash-link view-link"
              @click.stop="viewDetail(message)"
            >
              <n-icon><share-outline /></n-icon> 查看详情
            </span>
          </div>

          <div v-if="message.type === 4" class="whisper-content-wrap">
            {{ message.content }}
          </div>
        </n-alert>
      </template>
    </n-thing>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { ShareOutline } from '@vicons/ionicons5';
import { readMessage } from '@/api/user';
import { formatRelativeTime } from '@/utils/formatTime';

const defaultavatar =
  'https://assets.paopao.info/public/avatar/default/admin.png';

const router = useRouter();
const props = withDefaults(
  defineProps<{
    message: Item.MessageProps;
  }>(),
  {},
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
    .whisper-content-wrap {
      margin-top: 12px;
      text-decoration: underline;
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
