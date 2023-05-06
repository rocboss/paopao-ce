<template>
    <div>
        <main-nav title="设置" theme />
        <n-card title="基本信息" size="small" class="setting-card">
            <div class="base-line avatar">
                <n-avatar
                    class="avatar-img"
                    :size="80"
                    :src="store.state.userInfo.avatar"
                />
                <n-upload
                    v-if="!allowPhoneBind || (
                        allowPhoneBind &&
                        store.state.userInfo.phone &&
                        store.state.userInfo.phone.length > 0)
                    "
                    ref="avatarRef"
                    :action="uploadGateway"
                    :headers="{
                        Authorization: uploadToken,
                    }"
                    :data="{
                        type: uploadType,
                    }"
                    @before-upload="beforeUpload"
                    @finish="finishUpload"
                >
                    <n-button size="small">更改头像</n-button>
                </n-upload>
            </div>
            <div class="base-line">
                <span class="base-label">昵称</span>
                <div v-if="!showNicknameEdit">
                    {{ store.state.userInfo.nickname }}
                </div>
                <n-input
                    ref="inputInstRef"
                    v-show="showNicknameEdit"
                    class="nickname-input"
                    v-model:value="store.state.userInfo.nickname"
                    type="text"
                    size="small"
                    placeholder="请输入昵称"
                    @blur="handleNicknameChange"
                    :maxlength="16"
                />
                <n-button
                    quaternary
                    round
                    type="success"
                    size="small"
                    v-if="!showNicknameEdit && (!allowPhoneBind || (
                        allowPhoneBind &&
                        store.state.userInfo.phone &&
                        store.state.userInfo.phone.length > 0 &&
                        store.state.userInfo.status == 1)
                    )
                    "
                    @click="handleNicknameShow"
                >
                    <template #icon>
                        <n-icon>
                            <edit />
                        </n-icon>
                    </template>
                </n-button>
            </div>
            <div class="base-line">
                <span class="base-label">用户名</span> @{{
                    store.state.userInfo.username
                }}
            </div>
        </n-card>

        <n-card v-if="allowPhoneBind" title="手机号" size="small" class="setting-card">
            <div
                v-if="
                    store.state.userInfo.phone &&
                    store.state.userInfo.phone.length > 0
                "
            >
                {{ store.state.userInfo.phone }}

                <n-button
                    quaternary
                    round
                    type="success"
                    v-if="!showPhoneBind && store.state.userInfo.status == 1"
                    @click="showPhoneBind = true"
                >
                    换绑手机
                </n-button>
            </div>
            <div v-else>
                <n-alert title="手机绑定提示" type="warning">
                    成功绑定手机后，才能进行换头像、发动态、回复等交互~<br />
                    <a
                        class="hash-link"
                        @click="showPhoneBind = true"
                        v-if="!showPhoneBind"
                    >
                        立即绑定
                    </a>
                </n-alert>
            </div>

            <div class="phone-bind-wrap" v-if="showPhoneBind">
                <n-form
                    ref="phoneFormRef"
                    :model="modelData"
                    :rules="bindRules"
                >
                    <n-form-item path="phone" label="手机号">
                        <n-input
                            :value="modelData.phone"
                            @update:value="(v: string) => (modelData.phone = v.trim())"
                            placeholder="请输入中国大陆手机号"
                            @keydown.enter.prevent
                        />
                    </n-form-item>
                    <n-form-item path="img_captcha" label="图形验证码">
                        <div class="captcha-img-wrap">
                            <n-input
                                v-model:value="modelData.imgCaptcha"
                                placeholder="请输入图形验证码后获取验证码"
                            />
                            <div class="captcha-img">
                                <img
                                    v-if="modelData.b64s"
                                    :src="modelData.b64s"
                                    @click="loadCaptcha"
                                />
                            </div>
                        </div>
                    </n-form-item>
                    <n-form-item path="phone_captcha" label="短信验证码">
                        <n-input-group>
                            <n-input
                                v-model:value="modelData.phone_captcha"
                                placeholder="请输入收到的短信验证码"
                            />
                            <n-button
                                type="primary"
                                ghost
                                :disabled="smsDisabled"
                                :loading="sending"
                                @click="sendPhoneCaptcha"
                            >
                                {{
                                    smsCounter > 0 && smsDisabled
                                        ? smsCounter + 's后重新发送'
                                        : '发送验证码'
                                }}
                            </n-button>
                        </n-input-group>
                    </n-form-item>
                    <n-row :gutter="[0, 24]">
                        <n-col :span="24">
                            <div class="form-submit-wrap">
                                <n-button
                                    quaternary
                                    round
                                    @click="showPhoneBind = false"
                                >
                                    取消
                                </n-button>
                                <n-button
                                    secondary
                                    round
                                    type="primary"
                                    :loading="binding"
                                    @click="handlePhoneBind"
                                >
                                    绑定
                                </n-button>
                            </div>
                        </n-col>
                    </n-row>
                </n-form>
            </div>
        </n-card>

        <n-card v-if="allowActivation" title="激活码" size="small" class="setting-card">
            <div
                v-if="
                    store.state.userInfo.activation &&
                    store.state.userInfo.activation.length > 0
                "
            >
                {{ store.state.userInfo.activation }}

                <n-button
                    quaternary
                    round
                    type="success"
                    v-if="!showActivation"
                    @click="showActivation = true"
                >
                    重新激活
                </n-button>
            </div>
            <div v-else>
                <n-alert title="激活码激活提示" type="warning">
                    成功激活后后，才能发（公开/好友可见）动态、回复~<br />
                    <a
                        class="hash-link"
                        @click="showActivation = true"
                        v-if="!showActivation"
                    >
                    立即激活
                    </a>
                </n-alert>
            </div>

            <div class="phone-bind-wrap" v-if="showActivation">
                <n-form
                    ref="activateFormRef"
                    :model="activateData"
                    :rules="activateRules"
                >
                    <n-form-item path="activate_code" label="激活码">
                        <n-input
                            :value="activateData.activate_code"
                            @update:value="(v: string) => (activateData.activate_code = v.trim())"
                            placeholder="请输入激活码"
                            @keydown.enter.prevent
                        />
                    </n-form-item>
                    <n-form-item path="img_captcha" label="图形验证码">
                        <div class="captcha-img-wrap">
                            <n-input
                                v-model:value="activateData.imgCaptcha"
                                placeholder="请输入图形验证码后获取验证码"
                            />
                            <div class="captcha-img">
                                <img
                                    v-if="activateData.b64s"
                                    :src="activateData.b64s"
                                    @click="loadCaptcha4Activate"
                                />
                            </div>
                        </div>
                    </n-form-item>
                    <n-row :gutter="[0, 24]">
                        <n-col :span="24">
                            <div class="form-submit-wrap">
                                <n-button
                                    quaternary
                                    round
                                    @click="showActivation = false"
                                >
                                    取消
                                </n-button>
                                <n-button
                                    secondary
                                    round
                                    type="primary"
                                    :loading="activating"
                                    @click="handleActivation"
                                >
                                    激活
                                </n-button>
                            </div>
                        </n-col>
                    </n-row>
                </n-form>
            </div>
        </n-card>

        <n-card title="账户安全" size="small" class="setting-card">
            您已设置密码
            <n-button
                quaternary
                round
                type="success"
                v-if="!showPasswordSetting"
                @click="showPasswordSetting = true"
            >
                重置密码
            </n-button>
            <div class="phone-bind-wrap" v-if="showPasswordSetting">
                <n-form ref="formRef" :model="modelData" :rules="passwordRules">
                    <n-form-item path="old_password" label="旧密码">
                        <n-input
                            v-model:value="modelData.old_password"
                            type="password"
                            placeholder="请输入当前密码"
                            @keydown.enter.prevent
                        />
                    </n-form-item>
                    <n-form-item path="password" label="新密码">
                        <n-input
                            v-model:value="modelData.password"
                            type="password"
                            placeholder="请输入新密码"
                            @input="handlePasswordInput"
                            @keydown.enter.prevent
                        />
                    </n-form-item>
                    <n-form-item
                        ref="rPasswordFormItemRef"
                        first
                        path="reenteredPassword"
                        label="重复密码"
                    >
                        <n-input
                            v-model:value="modelData.reenteredPassword"
                            :disabled="!modelData.password"
                            type="password"
                            placeholder="请再次输入密码"
                            @keydown.enter.prevent
                        />
                    </n-form-item>
                    <n-row :gutter="[0, 24]">
                        <n-col :span="24">
                            <div class="form-submit-wrap">
                                <n-button
                                    quaternary
                                    round
                                    @click="showPasswordSetting = false"
                                >
                                    取消
                                </n-button>
                                <n-button
                                    secondary
                                    round
                                    type="primary"
                                    :loading="passwordSetting"
                                    @click="handleValidateButtonClick"
                                >
                                    更新
                                </n-button>
                            </div>
                        </n-col>
                    </n-row>
                </n-form>
            </div>
        </n-card>
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref, reactive } from 'vue';
import { useStore } from 'vuex';
import { Edit } from '@vicons/tabler';
import {
    getCaptcha,
    sendCaptcha,
    bindUserPhone,
    activateUser,
    changePassword,
    changeNickname,
    changeAvatar,
} from '@/api/user';
import type {
    UploadInst,
    FormItemRule,
    FormItemInst,
    FormInst,
    InputInst,
} from 'naive-ui';

const uploadGateway = import.meta.env.VITE_HOST + '/v1/attachment';
const uploadToken = 'Bearer ' + localStorage.getItem('PAOPAO_TOKEN');
const uploadType = ref('public/avatar');

const allowPhoneBind = (import.meta.env.VITE_ALLOW_PHONE_BIND.toLowerCase() === 'true')
const allowActivation = (import.meta.env.VITE_ALLOW_ACTIVATION.toLowerCase() === 'true')

const store = useStore();
const sending = ref(false);
const binding = ref(false);
const activating = ref(false)
const avatarRef = ref<UploadInst>();
const inputInstRef = ref<InputInst>();
const showNicknameEdit = ref(false);
const passwordSetting = ref(false);
const showPasswordSetting = ref(false);
const smsDisabled = ref(false);
const smsCounter = ref(60);
const showPhoneBind = ref(false);
const showActivation = ref(false);
const phoneFormRef = ref<FormInst>();
const activateFormRef = ref<FormInst>();
const formRef = ref<FormInst>();
const rPasswordFormItemRef = ref<FormItemInst>();
const modelData = reactive({
    id: '',
    b64s: '',
    imgCaptcha: '',
    phone: '',
    phone_captcha: '',
    password: '',
    old_password: '',
    reenteredPassword: '',
});

const activateData = reactive({
    id: '',
    b64s: '',
    imgCaptcha: '',
    activate_code: '',
});

const beforeUpload = async (data: any) => {
    // 图片类型校验
    if (
        uploadType.value === 'public/avatar' &&
        !['image/png', 'image/jpg', 'image/jpeg'].includes(data.file.file?.type)
    ) {
        window.$message.warning('头像仅允许 png/jpg 格式');
        return false;
    }

    if (uploadType.value === 'image' && data.file.file?.size > 1048576) {
        window.$message.warning('头像大小不能超过1MB');
        return false;
    }

    return true;
};

const finishUpload = ({ file, event }: any): any => {
    try {
        let data = JSON.parse(event.target?.response);

        if (data.code === 0) {
            if (uploadType.value === 'public/avatar') {
                changeAvatar({
                    avatar: data.data.content,
                })
                    .then((res) => {
                        window.$message.success('头像更新成功');
                        avatarRef.value?.clear();

                        store.commit('updateUserinfo', {
                            ...store.state.userInfo,
                            avatar: data.data.content,
                        });
                    })
                    .catch((err) => {
                        console.log(err);
                    });
            }
        }
    } catch (error) {
        window.$message.error('上传失败');
    }
};

const validatePasswordStartWith = (rule: FormItemRule, value: any) => {
    return (
        !!modelData.password &&
        (modelData.password as any).startsWith(value) &&
        (modelData.password as any).length >= value.length
    );
};

const validatePasswordSame = (rule: FormItemRule, value: any) => {
    return value === modelData.password;
};

const handlePasswordInput = () => {
    if (modelData.reenteredPassword) {
        rPasswordFormItemRef.value?.validate({ trigger: 'password-input' });
    }
};

const handleValidateButtonClick = (e: MouseEvent) => {
    e.preventDefault();
    formRef.value?.validate((errors) => {
        if (!errors) {
            passwordSetting.value = true;
            changePassword({
                password: modelData.password,
                old_password: modelData.old_password,
            })
                .then((res) => {
                    passwordSetting.value = false;
                    showPasswordSetting.value = false;
                    window.$message.success('密码重置成功');

                    // 用户退出登录
                    store.commit('userLogout');
                    store.commit('triggerAuth', true);
                    store.commit('triggerAuthKey', 'signin');
                })
                .catch((err) => {
                    passwordSetting.value = false;
                });
        }
    });
};

const handlePhoneBind = (e: MouseEvent) => {
    e.preventDefault();
    phoneFormRef.value?.validate((errors) => {
        if (!errors) {
            binding.value = true;
            bindUserPhone({
                phone: modelData.phone,
                captcha: modelData.phone_captcha,
            })
                .then((res) => {
                    binding.value = false;
                    showPhoneBind.value = false;
                    window.$message.success('绑定成功');

                    store.commit('updateUserinfo', {
                        ...store.state.userInfo,
                        phone: modelData.phone,
                    });

                    modelData.id = '';
                    modelData.b64s = '';
                    modelData.imgCaptcha = '';
                    modelData.phone = '';
                    modelData.phone_captcha = '';
                })
                .catch((err) => {
                    binding.value = false;
                });
        }
    });
};

const handleActivation = (e: MouseEvent) => {
    e.preventDefault();
    activateFormRef.value?.validate((errors) => {
    if (activateData.imgCaptcha === '') {
        window.$message.warning('请输入图片验证码');
        return;
    }
    sending.value = true;
        if (!errors) {
            activating.value = true;
            activateUser({
                activate_code: activateData.activate_code,
                captcha_id: activateData.id,
                imgCaptcha: activateData.imgCaptcha
            })
                .then((res) => {
                    activating.value = false;
                    showActivation.value = false;
                    window.$message.success('激活成功');

                    store.commit('updateUserinfo', {
                        ...store.state.userInfo,
                        activation: activateData.activate_code,
                    });

                    activateData.id = '';
                    activateData.b64s = '';
                    activateData.imgCaptcha = '';
                    activateData.activate_code = '';
                })
                .catch((err) => {
                    activating.value = false;
                    if (err.code === 20012) {
                        loadCaptcha4Activate();
                    }
                });
        }
    });
};

const loadCaptcha = () => {
    getCaptcha()
        .then((res) => {
            modelData.id = res.id;
            modelData.b64s = res.b64s;
        })
        .catch((err) => {
            console.log(err);
        });
};

const loadCaptcha4Activate = () => {
    getCaptcha()
        .then((res) => {
            activateData.id = res.id;
            activateData.b64s = res.b64s;
        })
        .catch((err) => {
            console.log(err);
        });
};

const handleNicknameChange = () => {
    changeNickname({
        nickname: store.state.userInfo.nickname || '',
    })
        .then((res) => {
            showNicknameEdit.value = false;
            window.$message.success('昵称修改成功');
        })
        .catch((err) => {
            showNicknameEdit.value = true;
        });
};

const sendPhoneCaptcha = () => {
    if (smsCounter.value > 0 && smsDisabled.value) {
        return;
    }
    if (modelData.imgCaptcha === '') {
        window.$message.warning('请输入图片验证码');
        return;
    }
    sending.value = true;
    sendCaptcha({
        phone: modelData.phone,
        img_captcha: modelData.imgCaptcha,
        img_captcha_id: modelData.id,
    })
        .then((res) => {
            smsDisabled.value = true;
            sending.value = false;
            window.$message.success('发送成功');

            let s = setInterval(() => {
                smsCounter.value--;
                if (smsCounter.value === 0) {
                    clearInterval(s);
                    smsCounter.value = 60;
                    smsDisabled.value = false;
                }
            }, 1000);
        })
        .catch((err) => {
            sending.value = false;
            if (err.code === 20012) {
                loadCaptcha();
            }
            console.log(err);
        });
};

const bindRules = {
    phone: [
        {
            required: true,
            message: '请输入手机号',
            trigger: ['input'],
            validator: (rule: FormItemRule, value: any) => {
                return /^[1]+[3-9]{1}\d{9}$/.test(value);
            },
        },
    ],
    phone_captcha: [
        {
            required: true,
            message: '请输入手机验证码',
        },
    ],
};
const activateRules = {
    activate_code: [
        {
            required: true,
            message: '请输入激活码',
            trigger: ['input'],
            validator: (rule: FormItemRule, value: any) => {
                return /\d{6}$/.test(value);
            },
        },
    ],
};
const passwordRules = {
    password: [
        {
            required: true,
            message: '请输入新密码',
        },
    ],
    old_password: [
        {
            required: true,
            message: '请输入旧密码',
        },
    ],
    reenteredPassword: [
        {
            required: true,
            message: '请再次输入密码',
            trigger: ['input', 'blur'],
        },
        {
            validator: validatePasswordStartWith,
            message: '两次密码输入不一致',
            trigger: 'input',
        },
        {
            validator: validatePasswordSame,
            message: '两次密码输入不一致',
            trigger: ['blur', 'password-input'],
        },
    ],
};
const handleNicknameShow = () => {
    showNicknameEdit.value = true;
    setTimeout(() => {
        inputInstRef.value?.focus();
    }, 30);
};
onMounted(() => {
    if (store.state.userInfo.id === 0) {
        store.commit('triggerAuth', true);
        store.commit('triggerAuthKey', 'signin');
    }
    loadCaptcha();
    loadCaptcha4Activate();
});
</script>

<style lang="less" scoped>
.setting-card {
    margin-top: -1px;
    border-radius: 0;
    .form-submit-wrap {
        display: flex;
        justify-content: flex-end;
    }

    .base-line {
        line-height: 2;
        display: flex;
        align-items: center;
        .base-label {
            opacity: 0.75;
            margin-right: 12px;
        }

        .nickname-input {
            margin-right: 10px;
            width: 120px;
        }
    }

    .avatar {
        display: flex;
        flex-direction: column;
        align-items: flex-start;
        margin-bottom: 20px;
        .avatar-img {
            margin-bottom: 10px;
        }
    }

    .hash-link {
        margin-left: 12px;
    }

    .phone-bind-wrap {
        margin-top: 20px;
        .captcha-img-wrap {
            width: 100%;
            display: flex;
            align-items: center;
        }
        .captcha-img {
            width: 125px;
            height: 34px;
            border-radius: 3px;
            margin-left: 10px;
            overflow: hidden;
            cursor: pointer;
            img {
                width: 100%;
                height: 100%;
            }
        }
    }
}
.dark {
    .setting-card {
        background-color: rgba(16, 16, 20, 0.75);
    }
}
</style>