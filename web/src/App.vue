<template>
    <n-config-provider :theme="theme">
        <n-message-provider>
            <div
                class="app-container"
                :class="{ dark: theme?.name === 'dark' }"
            >
                <div has-sider class="main-wrap" position="static">
                    <!-- 侧边栏 -->
                    <sidebar />

                    <div class="content-wrap">
                        <router-view class="app-wrap" v-slot="{ Component }">
                            <keep-alive>
                                <component
                                    v-if="$route.meta.keepAlive"
                                    :is="Component"
                                />
                            </keep-alive>
                            <component
                                v-if="!$route.meta.keepAlive"
                                :is="Component"
                            />
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
import { version, buildTime } from "../build/info.json";

const store = useStore();
const theme = computed(() => (store.state.theme === 'dark' ? darkTheme : null));

console.log(`%c Release Build Info 
%cVersion			v${version}
BuildTime		${buildTime}`
, "background:#000;color:#FFF;font-weight:bold;"
, "background:#FFF;color:#000;"
)
</script>

<style lang="less">
@import '@/assets/css/main.less';
</style>