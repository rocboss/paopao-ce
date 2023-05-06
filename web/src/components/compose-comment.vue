<template>
    <div>
        <div class="compose-wrap" v-if="store.state.userInfo.id > 0">
            <div class="compose-line">
                <div class="compose-user">
                    <n-avatar
                        round
                        :size="30"
                        :src="store.state.userInfo.avatar"
                    />
                </div>
                <n-mention
                    type="textarea"
                    size="large"
                    autosize
                    :bordered="false"
                    :options="optionsRef"
                    :prefix="['@']"
                    :loading="loading"
                    :value="content"
                    :disabled="props.lock === 1"
                    @update:value="changeContent"
                    @search="handleSearch"
                    @focus="focusComment"
                    :placeholder="
                        props.lock === 1
                            ? '泡泡已被锁定，回复功能已关闭'
                            : '快来评论两句吧...'
                    "
                />
            </div>

            <n-upload
                v-if="showBtn"
                ref="uploadRef"
                abstract
                list-type="image"
                :multiple="true"
                :max="9"
                :action="uploadGateway"
                :headers="{
                    Authorization: uploadToken,
                }"
                :data="{
                    type: uploadType,
                }"
                @before-upload="beforeUpload"
                @finish="finishUpload"
                @error="failUpload"
                @remove="removeUpload"
                @update:file-list="updateUpload"
            >
                <div class="compose-line compose-options">
                    <div class="attachment">
                        <n-upload-trigger #="{ handleClick }" abstract>
                            <n-button
                                :disabled="
                                    (fileQueue.length > 0 &&
                                        uploadType === 'public/video') ||
                                    fileQueue.length === 9
                                "
                                @click="
                                    () => {
                                        setUploadType('public/image');
                                        handleClick();
                                    }
                                "
                                quaternary
                                circle
                                type="primary"
                            >
                                <template #icon>
                                    <n-icon
                                        size="20"
                                        color="var(--primary-color)"
                                    >
                                        <image-outline />
                                    </n-icon>
                                </template>
                            </n-button>
                        </n-upload-trigger>

                        <n-tooltip trigger="hover" placement="bottom">
                            <template #trigger>
                                <n-progress
                                    class="text-statistic"
                                    type="circle"
                                    :show-indicator="false"
                                    status="success"
                                    :stroke-width="10"
                                    :percentage="(content.length / 200) * 100"
                                />
                            </template>
                            {{ content.length }} / 200
                        </n-tooltip>
                    </div>

                    <div class="submit-wrap">
                        <n-button
                            quaternary
                            round
                            type="tertiary"
                            class="cancel-btn"
                            size="small"
                            @click="cancelComment"
                        >
                            取消
                        </n-button>
                        <n-button
                            :loading="submitting"
                            @click="submitPost"
                            type="primary"
                            secondary
                            size="small"
                            round
                        >
                            发布
                        </n-button>
                    </div>
                </div>

                <div class="attachment-list-wrap">
                    <n-upload-file-list />
                </div>
            </n-upload>
        </div>

        <div class="compose-wrap" v-else>
            <div class="login-wrap">
                <span class="login-banner"> 登录后，精彩更多</span>
            </div>
            <div v-if="!allowUserRegister" class="login-only-wrap">
                <n-button
                    strong
                    secondary
                    round
                    type="primary"
                    @click="triggerAuth('signin')"
                >
                    登录
                </n-button>
            </div>
            <div v-if="allowUserRegister" class="login-wrap">
                <n-button
                    strong
                    secondary
                    round
                    type="primary"
                    @click="triggerAuth('signin')"
                >
                    登录
                </n-button>
                <n-button
                    strong
                    secondary
                    round
                    type="info"
                    @click="triggerAuth('signup')"
                >
                    注册
                </n-button>
            </div>
        </div>
    </div>
</template>


<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { useStore } from 'vuex';
import { debounce } from 'lodash';
import {
    ImageOutline,
} from '@vicons/ionicons5';
import { createComment } from '@/api/post';
import { getSuggestUsers } from '@/api/user';
import { parsePostTag } from '@/utils/content';
import type { MentionOption, UploadFileInfo, UploadInst } from 'naive-ui';

const emit = defineEmits<{
    (e: 'post-success'): void;
}>();
const props = withDefaults(
    defineProps<{
        lock: number;
        postId: number;
    }>(),
    {
        lock: 0,
        postId: 0,
    }
);

const store = useStore();

const optionsRef = ref<MentionOption[]>([]);
const showBtn = ref(false);
const loading = ref(false);
const submitting = ref(false);
const content = ref('');
const uploadRef = ref<UploadInst>();
const uploadType = ref('public/image');
const fileQueue = ref<UploadFileInfo[]>([]);
const imageContents = ref<Item.CommentItemProps[]>([]);
const allowUserRegister = ref(import.meta.env.VITE_ALLOW_USER_REGISTER.toLowerCase() === 'true')

const uploadGateway = import.meta.env.VITE_HOST + '/v1/attachment';
const uploadToken = ref();

// 加载at用户列表
const loadSuggestionUsers = debounce((k) => {
    getSuggestUsers({
        k,
    })
        .then((res) => {
            let options: MentionOption[] = [];
            res.suggest.map((i) => {
                options.push({
                    label: i,
                    value: i,
                });
            });
            optionsRef.value = options;
            loading.value = false;
        })
        .catch((err) => {
            loading.value = false;
        });
}, 200);

const handleSearch = (k: string, prefix: string) => {
    if (loading.value) {
        return;
    }
    loading.value = true;
    if (prefix === '@') {
        loadSuggestionUsers(k);
    }
};
const changeContent = (v: string) => {
    if (v.length > 200) {
        return;
    }
    content.value = v;
};
const setUploadType = (type: string) => {
    uploadType.value = type;
};
const updateUpload = (list: UploadFileInfo[]) => {
    fileQueue.value = list;
};
const beforeUpload = async (data: any) => {
    // 图片类型校验
    if (
        uploadType.value === 'public/image' &&
        !['image/png', 'image/jpg', 'image/jpeg', 'image/gif'].includes(
            (data.file as any).file?.type
        )
    ) {
        window.$message.warning('图片仅允许 png/jpg/gif 格式');
        return false;
    }

    if (
        uploadType.value === 'image' &&
        (data.file as any).file?.size > 10485760
    ) {
        window.$message.warning('图片大小不能超过10MB');
        return false;
    }

    return true;
};
const finishUpload = ({ file, event }: any): any => {
    try {
        let data = JSON.parse(event.target?.response);

        if (data.code === 0) {
            if (uploadType.value === 'public/image') {
                imageContents.value.push({
                    id: file.id,
                    content: data.data.content,
                } as Item.CommentItemProps);
            }
        }
    } catch (error) {
        window.$message.error('上传失败');
    }
};
const failUpload = ({ file, event }: any): any => {
    try {
        let data = JSON.parse(event.target?.response);

        if (data.code !== 0) {
            let errMsg = data.msg || '上传失败';
            if (data.details && data.details.length > 0) {
                data.details.map((detail: string) => {
                    errMsg += ':' + detail;
                });
            }
            window.$message.error(errMsg);
        }
    } catch (error) {
        window.$message.error('上传失败');
    }
};
const removeUpload = ({ file }: any) => {
    let idx = imageContents.value.findIndex((item) => item.id === file.id);
    if (idx > -1) {
        imageContents.value.splice(idx, 1);
    }
};

const focusComment = () => {
    showBtn.value = true;
};
const cancelComment = () => {
    showBtn.value = false;
    // 置空
    uploadRef.value?.clear();
    fileQueue.value = [];
    content.value = '';
    imageContents.value = [];
};

// 发布动态
const submitPost = () => {
    if (content.value.trim().length === 0) {
        window.$message.warning('请输入内容哦');
        return;
    }

    // 解析用户at
    let { users } = parsePostTag(content.value);

    const contents = [];
    let sort = 100;

    contents.push({
        content: content.value,
        type: 2, // 文字
        sort,
    });
    imageContents.value.map((img) => {
        sort++;
        contents.push({
            content: img.content,
            type: 3, // 图片
            sort,
        });
    });

    submitting.value = true;
    createComment({
        contents,
        post_id: props.postId,
        users: Array.from(new Set(users)),
    })
        .then((res) => {
            window.$message.success('发布成功');
            submitting.value = false;
            emit('post-success');

            // 置空
            cancelComment();
        })
        .catch((err) => {
            submitting.value = false;
        });
};
const triggerAuth = (key: string) => {
    store.commit('triggerAuth', true);
    store.commit('triggerAuthKey', key);
};
onMounted(() => {
    uploadToken.value = 'Bearer ' + localStorage.getItem('PAOPAO_TOKEN');
});
</script>

<style lang="less" scoped>
.compose-wrap {
    width: 100%;
    padding: 16px;
    box-sizing: border-box;

    .compose-line {
        display: flex;
        flex-direction: row;

        .compose-user {
            width: 42px;
            height: 42px;
            display: flex;
            align-items: center;
        }

        &.compose-options {
            margin-top: 6px;
            padding-left: 42px;
            display: flex;
            justify-content: space-between;

            .submit-wrap {
                display: flex;
                align-items: center;

                .cancel-btn {
                    margin-right: 8px;
                }
            }
        }
    }
    .login-only-wrap {
        display: flex;
        justify-content: center;
        width: 100%;
        button {
            margin: 0 4px;
            width: 50%
        }
    }
    .login-wrap {
        display: flex;
        justify-content: center;
        width: 100%;
        .login-banner {
            margin-bottom: 12px;
            opacity: 0.8;
        }
        button {
            margin: 0 4px;
        }
    }
}
.attachment {
    display: flex;
    align-items: center;
    .text-statistic {
        margin-left: 8px;
        width: 18px;
        height: 18px;
        transform: rotate(180deg);
    }
}
.attachment-list-wrap {
    margin-top: 12px;
    margin-left: 42px;
    .n-upload-file-info__thumbnail {
        overflow: hidden;
    }
}
.dark {
    .compose-mention {
        background-color: rgba(16, 16, 20, 0.75);
    }
    .compose-wrap {
        background-color: rgba(16, 16, 20, 0.75);
    }
}
</style>