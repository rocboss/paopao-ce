<template>
    <div class="detail-item" @click="goPostDetail(post.id)">
        <n-thing>
            <template #avatar>
                <n-avatar round :size="30" :src="post.user.avatar" />
            </template>
            <template #header>
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
                <span class="username-wrap"> @{{ post.user.username }} </span>
            </template>
            <template #header-extra>
                <div
                    class="options"
                    v-if="
                        store.state.userInfo.is_admin ||
                        store.state.userInfo.id === post.user.id
                    "
                >
                    <n-dropdown
                        placement="bottom-end"
                        trigger="click"
                        size="small"
                        :options="adminOptions"
                        @select="handlePostAction"
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

                <!-- 删除确认 -->
                <n-modal
                    v-model:show="showDelModal"
                    :mask-closable="false"
                    preset="dialog"
                    title="提示"
                    content="确定删除该泡泡动态吗？"
                    positive-text="确认"
                    negative-text="取消"
                    @positive-click="execDelAction"
                />
                <!-- 锁定确认 -->
                <n-modal
                    v-model:show="showLockModal"
                    :mask-closable="false"
                    preset="dialog"
                    title="提示"
                    content="确定锁定该泡泡动态吗？"
                    positive-text="确认"
                    negative-text="取消"
                    @positive-click="execLockAction"
                />
            </template>
            <div v-if="post.texts.length > 0">
                <span
                    v-for="content in post.texts"
                    :key="content.id"
                    class="post-text"
                    @click.stop="doClickText($event, post.id)"
                    v-html="parsePostTag(content.content).content"
                >
                </span>
            </div>

            <template #footer>
                <post-attachment :attachments="post.attachments" />
                <post-attachment
                    :attachments="post.charge_attachments"
                    :price="post.attachment_price"
                />
                <post-image :imgs="post.imgs" />
                <post-video :videos="post.videos" :full="true" />
                <post-link :links="post.links" />
                <div class="timestamp">
                    发布于 {{ formatRelativeTime(post.created_on) }}
                    <span v-if="post.ip_loc">
                        <n-divider vertical />
                        {{ post.ip_loc }}
                    </span>
                    <span v-if="post.created_on != post.latest_replied_on">
                        <n-divider vertical /> 最后回复
                        {{ formatRelativeTime(post.latest_replied_on) }}
                    </span>
                </div>
            </template>
            <template #action>
                <div class="opts-wrap">
                    <n-space justify="space-between">
                        <div
                            class="opt-item hover"
                            @click.stop="handlePostStar"
                        >
                            <n-icon size="20" class="opt-item-icon">
                                <heart-outline v-if="!hasStarred" />
                                <heart v-if="hasStarred" color="red" />
                            </n-icon>
                            {{ post.upvote_count }}
                        </div>
                        <div class="opt-item">
                            <n-icon size="20" class="opt-item-icon">
                                <chatbox-outline />
                            </n-icon>
                            {{ post.comment_count }}
                        </div>
                        <div
                            class="opt-item hover"
                            @click.stop="handlePostCollection"
                        >
                            <n-icon size="20" class="opt-item-icon">
                                <bookmark-outline v-if="!hasCollected" />
                                <bookmark v-if="hasCollected" color="#ff7600" />
                            </n-icon>
                            {{ post.collection_count }}
                        </div>
                    </n-space>
                </div>
            </template>
        </n-thing>
    </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import { useStore } from 'vuex';
import { useRouter } from 'vue-router';
import { formatRelativeTime } from '@/utils/formatTime';
import { parsePostTag } from '@/utils/content';
import {
    Heart,
    HeartOutline,
    Bookmark,
    BookmarkOutline,
    ChatboxOutline,
} from '@vicons/ionicons5';
import { MoreHorizFilled } from '@vicons/material';
import {
    getPostStar,
    postStar,
    getPostCollection,
    postCollection,
    deletePost,
    lockPost,
} from '@/api/post';

const store = useStore();
const router = useRouter();
const hasStarred = ref(false);
const hasCollected = ref(false);
const props = defineProps({
    post: {
        type: Object,
        default: () => {},
    },
});
const showDelModal = ref(false);
const showLockModal = ref(false);

const emit = defineEmits(['reload']);

const post = computed({
    get: () => {
        let post = Object.assign(
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
    },
    set: (newVal) => {
        props.post.upvote_count = newVal.upvote_count;
        props.post.comment_count = newVal.comment_count;
        props.post.collection_count = newVal.collection_count;
    },
});

const adminOptions = computed(() => {
    let options = [
        {
            label: '删除',
            key: 'delete',
        },
    ];
    if (post.value.is_lock === 0) {
        options.push({
            label: '锁定',
            key: 'lock',
        });
    } else {
        options.push({
            label: '解锁',
            key: 'unlock',
        });
    }
    if (store.state.userInfo.is_admin) {
        options.push(
            {
                label: '置顶',
                key: 'top',
            },
            {
                label: '精华',
                key: 'essence',
            }
        );
    }
    return options;
});

const goPostDetail = (id) => {
    router.push({
        name: 'post',
        query: {
            id,
        },
    });
};
const doClickText = (e, id) => {
    if (e.target.dataset.detail) {
        const d = e.target.dataset.detail.split(':');
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
const handlePostAction = (item) => {
    if (item === 'delete') {
        showDelModal.value = true;
    }
    if (item === 'lock' || item === 'unlock') {
        showLockModal.value = true;
    }
};
const execDelAction = () => {
    deletePost({
        id: post.value.id,
    })
        .then((res) => {
            window.$message.success('删除成功');
            router.replace('/');

            setTimeout(() => {
                store.commit('refresh');
            }, 50);
        })
        .catch((err) => {
            loading.value = false;
        });
};
const execLockAction = () => {
    lockPost({
        id: post.value.id,
    })
        .then((res) => {
            emit('reload');
            if (res.lock_status === 1) {
                window.$message.success('锁定成功');
            } else {
                window.$message.success('解锁成功');
            }
        })
        .catch((err) => {
            loading.value = false;
        });
};
const handlePostStar = () => {
    postStar({
        id: post.value.id,
    })
        .then((res) => {
            hasStarred.value = res.status;
            if (res.status) {
                post.value = {
                    ...post.value,
                    upvote_count: post.value.upvote_count + 1,
                };
            } else {
                post.value = {
                    ...post.value,
                    upvote_count: post.value.upvote_count - 1,
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
            hasCollected.value = res.status;
            if (res.status) {
                post.value = {
                    ...post.value,
                    collection_count: post.value.collection_count + 1,
                };
            } else {
                post.value = {
                    ...post.value,
                    collection_count: post.value.collection_count - 1,
                };
            }
        })
        .catch((err) => {
            console.log(err);
        });
};

onMounted(() => {
    if (store.state.userInfo.id > 0) {
        getPostStar({
            id: post.value.id,
        })
            .then((res) => {
                hasStarred.value = res.status;
            })
            .catch((err) => {
                console.log(err);
            });

        getPostCollection({
            id: post.value.id,
        })
            .then((res) => {
                hasCollected.value = res.status;
            })
            .catch((err) => {
                console.log(err);
            });
    }
});
</script>

<style lang="less">
.detail-item {
    width: 100%;
    padding: 16px;
    box-sizing: border-box;

    background: #f7f9f9;
    .nickname-wrap {
        font-size: 14px;
    }
    .username-wrap {
        font-size: 14px;
        opacity: 0.75;
    }

    .options {
        opacity: 0.75;
    }
    .post-text {
        font-size: 16px;
        text-align: justify;
        overflow: hidden;
        white-space: pre-wrap;
        word-break: break-all;
    }
    .opts-wrap {
        margin-top: 20px;
        .opt-item {
            display: flex;
            align-items: center;
            opacity: 0.7;
            .opt-item-icon {
                margin-right: 10px;
            }
            &.hover {
                cursor: pointer;
            }
        }
    }
    .n-thing {
        .n-thing-avatar-header-wrapper {
            align-items: center;
        }
    }
    .timestamp {
        opacity: 0.75;
        font-size: 12px;
        margin-top: 10px;
    }
}
.dark {
    .detail-item {
        background: #18181c;
    }
}
</style>