<template>
    <div class="post-item" @click="goPostDetail(post.id)">
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
            </template>
            <template #header-extra>
                <span class="timestamp">
                    {{ post.ip_loc ? post.ip_loc + ' · ' : post.ip_loc }}
                    {{ formatRelativeTime(post.created_on) }}
                </span>
            </template>
            <template #description v-if="post.texts.length > 0">
                <span
                    v-for="content in post.texts"
                    :key="content.id"
                    class="post-text"
                    @click.stop="doClickText($event, post.id)"
                    v-html="parsePostTag(content.content).content"
                >
                </span>
            </template>

            <template #footer>
                <post-attachment :attachments="post.attachments" />
                <post-attachment
                    :attachments="post.charge_attachments"
                    :price="post.attachment_price"
                />
                <post-image :imgs="post.imgs" />
                <post-video :videos="post.videos" />
                <post-link :links="post.links" />
            </template>
            <template #action>
                <n-space justify="space-between">
                    <div class="opt-item">
                        <n-icon size="18" class="opt-item-icon">
                            <heart-outline />
                        </n-icon>
                        {{ post.upvote_count }}
                    </div>
                    <div class="opt-item">
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
                </n-space>
            </template>
        </n-thing>
    </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useStore } from 'vuex';
import { useRoute, useRouter } from 'vue-router';
import { formatRelativeTime } from '@/utils/formatTime';
import { parsePostTag } from '@/utils/content';
import {
    HeartOutline,
    BookmarkOutline,
    ChatboxOutline,
} from '@vicons/ionicons5';

const route = useRoute();
const router = useRouter();
const store = useStore();
const props = withDefaults(defineProps<{
    post: Item.PostProps,
}>(), {});

const post = computed(() => {
    let post: Required<Item.PostProps> = Object.assign(
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
    .timestamp {
        opacity: 0.75;
        font-size: 12px;
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
    }
}
</style>