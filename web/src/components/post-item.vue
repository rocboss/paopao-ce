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
                                query: { username: post.user.username },
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
                    <div v-if="!store.state.desktopModelShow">
                        <span class="timestamp-mobile">
                            {{ formatPrettyDate(post.created_on) }} {{ post.ip_loc }}
                        </span>
                    </div>
            </template>
            <template #header-extra>
                <div class="item-header-extra">
                    <span v-if="store.state.desktopModelShow" class="timestamp">
                        {{ post.ip_loc ? post.ip_loc + ' · ' : post.ip_loc }}
                        {{ formatPrettyDate(post.created_on) }}
                    </span>
                    <n-dropdown v-if="!store.state.desktopModelShow"
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
                    <span
                    v-for="content in post.texts"
                    :key="content.id"
                    class="post-text"
                    @click.stop="doClickText($event, post.id)"
                    v-html="parsePostTag(content.content).content"
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
                    <div class="opt-item">
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
                    <div class="opt-item">
                        <n-icon size="18" class="opt-item-icon">
                            <bookmark-outline />
                        </n-icon>
                        {{ post.collection_count }}
                    </div>
                    <!-- <div class="opt-item">
                        <n-icon size="18" class="opt-item-icon">
                            <share-social-outline />
                        </n-icon>
                        {{ post.share_count }}
                    </div> -->
                </n-space>
            </template>
        </n-thing>
    </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useStore } from 'vuex';
import type { DropdownOption } from 'naive-ui';
import { useRoute, useRouter } from 'vue-router';
import { formatPrettyDate } from '@/utils/formatTime';
import { parsePostTag } from '@/utils/content';
import {
    HeartOutline,
    BookmarkOutline,
    ChatboxOutline,
    // ShareSocialOutline,
} from '@vicons/ionicons5';
import { MoreHorizFilled } from '@vicons/material';
import copy from "copy-to-clipboard";

const route = useRoute();
const router = useRouter();
const store = useStore();
const props = withDefaults(defineProps<{
    post: Item.PostProps,
}>(), {});

const tweetOptions = computed(() => {
    let options: DropdownOption[] = [
        {
            label: '复制链接',
            key: 'copyTweetLink',
        },
    ];
    return options;
});

const handleTweetAction = (
    item: 'copyTweetLink'
) => {
    switch (item) {
        case 'copyTweetLink':
            copy(`${window.location.origin}/#/post?id=${post.value.id}`);
            window.$message.success('链接已复制到剪贴板');
            break;
        default:
            break;
    }
};

const post = computed(() => {
    let post: Item.PostComponentProps = Object.assign(
        {
            texts: [],
            imgs: [],
            videos: [],
            links: [],
            attachments: [],
            charge_attachments: [],
        },
        props.post
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
});
const goPostDetail = (id: number) => {
    router.push({
        name: 'post',
        query: {
            id,
        },
    });
};
const doClickText = (e: MouseEvent, id: number) => {
    if ((e.target as any).dataset.detail) {
        const d = (e.target as any).dataset.detail.split(':');
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
                        username: d[1],
                    },
                });
            }
            return;
        }
    }
    goPostDetail(id);
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