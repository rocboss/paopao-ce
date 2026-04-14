<template>
    <n-modal
        v-model:show="authModalShow"
        class="auth-card"
        preset="card"
        size="small"
        :mask-closable="false"
        :bordered="false"
        :style="{
            width: '360px',
        }"
    >
        <div class="auth-wrap">
            <n-card :bordered="false">
                <div v-if="!profile.allowUserRegister">
                    <n-space justify="center"><n-h3><n-text type="success">账号登录</n-text></n-h3></n-space>
                    <n-form
                            ref="loginRef"
                            :model="loginForm"
                            :rules="{
                                username: {
                                    required: true,
                                    message: '请输入账户名',
                                },
                                password: {
                                    required: true,
                                    message: '请输入密码',
                                },
                            }"
                        >
                            <n-form-item-row label="账户" path="username">
                                <n-input
                                    v-model:value="loginForm.username"
                                    placeholder="请输入用户名"
                                    @keyup.enter.prevent="handleLogin"
                                />
                            </n-form-item-row>
                            <n-form-item-row label="密码" path="password">
                                <n-input
                                    type="password"
                                    show-password-on="mousedown"
                                    v-model:value="loginForm.password"
                                    placeholder="请输入账户密码"
                                    @keyup.enter.prevent="handleLogin"
                                />
                            </n-form-item-row>
                        </n-form>
                        <n-button
                            type="primary"
                            block
                            secondary
                            strong
                            :loading="loading"
                            @click="handleLogin"
                        >
                            登录
                        </n-button>
                </div>
                <n-tabs
                    v-if="profile.allowUserRegister" 
                    :default-value="authModelTab"
                    size="large"
                    justify-content="space-evenly"
                >
                    <n-tab-pane name="signin" tab="登录">
                        <n-form
                            ref="loginRef"
                            :model="loginForm"
                            :rules="{
                                username: {
                                    required: true,
                                    message: '请输入账户名',
                                },
                                password: {
                                    required: true,
                                    message: '请输入密码',
                                },
                            }"
                        >
                            <n-form-item-row label="账户" path="username">
                                <n-input
                                    v-model:value="loginForm.username"
                                    placeholder="请输入用户名"
                                    @keyup.enter.prevent="handleLogin"
                                />
                            </n-form-item-row>
                            <n-form-item-row label="密码" path="password">
                                <n-input
                                    type="password"
                                    show-password-on="mousedown"
                                    v-model:value="loginForm.password"
                                    placeholder="请输入账户密码"
                                    @keyup.enter.prevent="handleLogin"
                                />
                            </n-form-item-row>
                        </n-form>
                        <n-button
                            type="primary"
                            block
                            secondary
                            strong
                            :loading="loading"
                            @click="handleLogin"
                        >
                            登录
                        </n-button>
                    </n-tab-pane>
                    <n-tab-pane name="signup" tab="注册">
                        <n-form
                            ref="registerRef"
                            :model="registerForm"
                            :rules="registerRule"
                        >
                            <n-form-item-row label="用户名" path="username">
                                <n-input
                                    v-model:value="registerForm.username"
                                    placeholder="用户名注册后无法修改"
                                />
                            </n-form-item-row>
                            <n-form-item-row label="密码" path="password">
                                <n-input
                                    type="password"
                                    show-password-on="mousedown"
                                    placeholder="密码不少于6位"
                                    v-model:value="registerForm.password"
                                    @keyup.enter.prevent="handleRegister"
                                />
                            </n-form-item-row>
                            <n-form-item-row label="重复密码" path="repassword">
                                <n-input
                                    type="password"
                                    show-password-on="mousedown"
                                    placeholder="请再次输入密码"
                                    v-model:value="registerForm.repassword"
                                    @keyup.enter.prevent="handleRegister"
                                />
                            </n-form-item-row>
                        </n-form>
                        <n-button
                            type="primary"
                            block
                            secondary
                            strong
                            :loading="loading"
                            @click="handleRegister"
                        >
                            注册
                        </n-button>
                    </n-tab-pane>
                </n-tabs>
            </n-card>
        </div>
    </n-modal>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue';
import { useStoreMain } from '@/store/main';
import { TOKEN_KEY, useStoreUser } from '@/store/user';
import { userInfo } from '@/api/auth';
import type { FormInst, FormItemRule } from 'naive-ui';
import { storeToRefs } from 'pinia';
import { Api } from '@/utils/request';
import { useStoreProfile } from '@/store/profile';

const storeMain = useStoreMain();
const storeUser = useStoreUser();
const storeProfile = useStoreProfile();
const { authModalShow, authModelTab } = storeToRefs(storeMain);
const { profile } = storeToRefs(storeProfile);

const loading = ref(false);
const loginRef = ref<FormInst>();
const loginForm = reactive({
  username: '',
  password: '',
});
const registerRef = ref<FormInst>();
const registerForm = reactive({
  username: '',
  password: '',
  repassword: '',
});
const registerRule = {
  username: {
    required: true,
    message: '请输入账户名',
  },
  password: {
    required: true,
    message: '请输入密码',
  },
  repassword: [
    {
      required: true,
      message: '请输入密码',
    },
    {
      validator: (rule: FormItemRule, value: any) => {
        return (
          !!registerForm.password &&
          registerForm.password.startsWith(value) &&
          registerForm.password.length >= value.length
        );
      },
      message: '两次密码输入不一致',
      trigger: 'input',
    },
  ],
};
const handleLogin = (e: Event) => {
  e.preventDefault();
  e.stopPropagation();

  loginRef.value?.validate((errors) => {
    if (!errors) {
      loading.value = true;

      Api.v1.auth.post.login({
        username: loginForm.username,
        password: loginForm.password,
      })
        .then((res) => {
          const token = res?.token || '';
          // 写入用户信息
          localStorage.setItem(TOKEN_KEY, token);

          return userInfo(token);
        })
        .then((res) => {
          window.$message.success('登录成功');
          loading.value = false;

          storeUser.updateUserinfo(res);
          storeMain.triggerAuth(false);
          storeMain.doRefresh();
          loginForm.username = '';
          loginForm.password = '';
        })
        .catch((err) => {
          loading.value = false;
        });
    }
  });
};

const handleRegister = (e: Event) => {
  e.preventDefault();
  e.stopPropagation();

  registerRef.value?.validate((errors) => {
    if (!errors) {
      loading.value = true;

      Api.v1.auth.post.register({
        username: registerForm.username,
        password: registerForm.password,
      })
        .then((res) => {
          return Api.v1.auth.post.login({
            username: registerForm.username,
            password: registerForm.password,
          });
        })
        .then((res) => {
          const token = res?.token || '';
          // 写入用户信息
          localStorage.setItem(TOKEN_KEY, token);

          return userInfo(token);
        })
        .then((res) => {
          window.$message.success('注册成功');
          loading.value = false;

          storeUser.updateUserinfo(res);
          storeMain.triggerAuth(false);
          registerForm.username = '';
          registerForm.password = '';
          registerForm.repassword = '';
        })
        .catch((err) => {
          loading.value = false;
        });
    }
  });
};

onMounted(() => {
  const token = localStorage.getItem(TOKEN_KEY) || '';
  if (token) {
    userInfo(token)
      .then((res) => {
        storeUser.updateUserinfo(res);
        storeMain.triggerAuth(false);
      })
      .catch((err) => {
        storeUser.userLogout();
      });
  } else {
    storeUser.userLogout();
  }
});
</script>

<style lang="less" scoped>
.auth-wrap {
    margin-top: -30px;
}
.dark {
    .auth-wrap {
        background-color: rgba(16, 16, 20, 0.75);
    }
}
</style>