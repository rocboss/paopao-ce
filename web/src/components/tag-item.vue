<template>
    <div v-if="(!checkFollowing && !checkPin) || (checkFollowing && tag.is_following === 1) || (checkPin && tag.is_following === 1 && tag.is_pin === 1)" class="tag-item">
        <n-thing>
            <template #header>
                <n-tag
                        type="success"
                        size="large"
                        round
                        :key="tag.id"
                    >
                        <router-link
                            class="hash-link"
                            :to="{
                                name: 'home',
                                query: {
                                    q: tag.tag,
                                    t: 'tag',
                                },
                            }"
                        >
                            #{{ tag.tag }}
                        </router-link>
                        <span v-if="!showAction" class="tag-quote">({{ tag.quote_num }})</span>
                        <span v-if="showAction" class="tag-quote tag-follow">({{ tag.quote_num }})</span>
                        <template #avatar>
                            <n-avatar :src="tagUserAvatar" />
                        </template>
                    </n-tag>
            </template>
            <template #header-extra>
                <div 
                    v-if="showAction" 
                    class="options">
                    <n-dropdown
                        placement="bottom-end"
                        trigger="click"
                        size="small"
                        :options="tagOptions"
                        @select="handleTagAction"
                    >
                        <n-button type="success" quaternary circle block>
                            <template #icon>
                                <n-icon>
                                    <more-vert-outlined />
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
import { ref, onMounted, computed } from 'vue';
import { MoreVertOutlined } from '@vicons/material';
import type { DropdownOption } from 'naive-ui';
import { pinTopic, stickTopic, followTopic, unfollowTopic } from '@/api/post';
import defaultUserAvatar from '@/assets/img/logo.png';

const hasFollowing = ref(false);
const props = withDefaults(
  defineProps<{
    tag: Item.TagProps;
    showAction: boolean;
    checkFollowing: boolean;
    checkPin: boolean;
  }>(),
  {},
);

const tagUserAvatar = computed(() => {
  if (props.tag.user) {
    return props.tag.user.avatar;
  } else {
    return defaultUserAvatar;
  }
});

const tagOptions = computed(() => {
  let options: DropdownOption[] = [];
  if (props.tag.is_following === 0) {
    options.push({
      label: '关注',
      key: 'follow',
    });
  } else {
    if (props.tag.is_pin === 0) {
      options.push({
        label: '钉住',
        key: 'pin',
      });
    } else {
      options.push({
        label: '取消钉住',
        key: 'unpin',
      });
    }
    if (props.tag.is_top === 0) {
      options.push({
        label: '置顶',
        key: 'stick',
      });
    } else {
      options.push({
        label: '取消置顶',
        key: 'unstick',
      });
    }
    options.push({
      label: '取消关注',
      key: 'unfollow',
    });
  }
  return options;
});

const handleTagAction = (
  item: 'follow' | 'unfollow' | 'pin' | 'unpin' | 'stick' | 'unstick',
) => {
  switch (item) {
    case 'follow':
      followTopic({
        topic_id: props.tag.id,
      })
        .then((_res) => {
          props.tag.is_following = 1;
          window.$message.success(`关注成功`);
        })
        .catch((err) => {
          console.log(err);
        });
      break;
    case 'unfollow':
      unfollowTopic({
        topic_id: props.tag.id,
      })
        .then((_res) => {
          props.tag.is_following = 0;
          window.$message.success(`取消关注`);
        })
        .catch((err) => {
          console.log(err);
        });
      break;
    case 'pin':
      pinTopic({
        topic_id: props.tag.id,
      })
        .then((_res) => {
          props.tag.is_pin = 1;
          window.$message.success(`钉住成功`);
        })
        .catch((err) => {
          console.log(err);
        });
      break;
    case 'unpin':
      pinTopic({
        topic_id: props.tag.id,
      })
        .then((_res) => {
          props.tag.is_pin = 0;
          window.$message.success(`取消钉住`);
        })
        .catch((err) => {
          console.log(err);
        });
      break;
    case 'stick':
      stickTopic({
        topic_id: props.tag.id,
      })
        .then((res) => {
          props.tag.is_top = res.top_status;
          window.$message.success(`置顶成功`);
        })
        .catch((err) => {
          console.log(err);
        });
      break;
    case 'unstick':
      stickTopic({
        topic_id: props.tag.id,
      })
        .then((res) => {
          props.tag.is_top = res.top_status;
          window.$message.success(`取消置顶`);
        })
        .catch((err) => {
          console.log(err);
        });
      break;
    default:
      break;
  }
};

onMounted(() => {
  hasFollowing.value = false;
});
</script>

<style lang="less">
.tag-item {
    .tag-quote {
        margin-left: 12px;
        font-size: 14px;
        opacity: 0.75;
    }
    .tag-follow {
        margin-right: 22px;
    }
    .options {
        margin-left: -32px;
        margin-bottom: 4px;
        opacity: 0.55;
    }
    .n-thing {
        .n-thing-header {
            margin-bottom: 0px;
        }
        .n-thing-avatar-header-wrapper {
            align-items: center;
        }
    }
}
</style>