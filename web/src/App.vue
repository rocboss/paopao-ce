<template>
    <n-config-provider :theme="iTheme">
        <n-message-provider>
            <n-dialog-provider>
                <div
                    class="app-container"
                    :class="{ dark: iTheme?.name === 'dark', mobile: !desktopModelShow }"
                >
                    <div has-sider class="main-wrap" position="static" >
                        <!-- 侧边栏 -->
                        <div v-if="desktopModelShow">
                            <sidebar />
                        </div>

                        <div class="content-wrap">
                            <router-view
                                class="app-wrap"
                                v-slot="{ Component }"
                            >
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
            </n-dialog-provider>
        </n-message-provider>
        <n-global-style />
    </n-config-provider>
</template>

<script setup lang="ts">
import { onMounted, computed } from 'vue';
import { useStoreMain } from '@/store/main';
import { darkTheme } from 'naive-ui';
import { getSiteProfile } from '@/api/site';
import { useStoreProfile } from '@/store/profile';
import { storeToRefs } from 'pinia';

const storeMain = useStoreMain();
const storeProfile = useStoreProfile();
const { theme, desktopModelShow } = storeToRefs(storeMain);

const iTheme = computed(() => (theme.value === 'dark' ? darkTheme : null));

function loadSiteProfile() {
    storeProfile.loadDefaultSiteProfile();
    if (import.meta.env.VITE_USE_WEB_PROFILE.toLowerCase() === 'true') {
        getSiteProfile()
            .then((res) => {
                storeProfile.updateSiteProfile(res);
            }).catch((err) => {
                console.log(err);
            });
    }
}

onMounted(() => {
  loadSiteProfile();
});
</script>

<style lang="less">
@import '@/assets/css/main.less';
</style>