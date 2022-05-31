<template>
  <n-config-provider :theme="theme">
    <n-message-provider>
      <div class="app-container" :class="{ dark: theme?.name === 'dark' }">
        <div has-sider class="main-wrap" position="static">
          <!-- 侧边栏 -->
          <sidebar />

          <div class="content-wrap">
            <router-view v-slot="{ Component }" class="app-wrap">
              <keep-alive>
                <component :is="Component" v-if="$route.meta.keepAlive" />
              </keep-alive>
              <component :is="Component" v-if="!$route.meta.keepAlive" />
            </router-view>
          </div>

          <!-- 右侧 -->
          <rightbar />
        </div>
        <!-- 登录/注册公共组件 -->
        <auth />
      </div>
    </n-message-provider>
    <n-global-style />
  </n-config-provider>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useStore } from 'vuex';
import { darkTheme } from 'naive-ui';

const store = useStore();
const theme = computed(() => (store.state.theme === 'dark' ? darkTheme : null));
</script>

<style lang="less">
@import '@/assets/css/main.less';
</style>
