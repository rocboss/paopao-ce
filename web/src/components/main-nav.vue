<template>
    <div v-if="store.state.drawerModelShow">
        <n-drawer
            v-model:show="activeDrawerRef"
            :width=212
            :placement="placementRef"
            resizable
        >
            <n-drawer-content>
                <sidebar />
            </n-drawer-content>
        </n-drawer>
    </div>
    <n-card size="small" :bordered="true" class="nav-title-card">
        <template #header>
            <div class="navbar">
                <n-button
                    class="drawer-btn"
                    v-if="store.state.drawerModelShow && !back"
                    @click="activeDrawer"
                    quaternary
                    circle
                    size="medium"
                >
                    <template #icon>
                        <n-icon><dehaze-round /></n-icon>
                    </template>
                </n-button>
                <n-button
                    class="back-btn"
                    v-if="back"
                    @click="goBack"
                    quaternary
                    circle
                    size="small"
                >
                    <template #icon>
                        <n-icon><chevron-left-round /></n-icon>
                    </template>
                </n-button>

                {{ props.title }}

                <n-switch
                    v-if="props.theme"
                    :value="store.state.theme === 'dark'"
                    @update:value="switchTheme"
                    size="small"
                    class="theme-switch-wrap"
                >
                    <template #checked-icon>
                        <n-icon :component="LightModeOutlined" />
                    </template>
                    <template #unchecked-icon>
                        <n-icon :component="DarkModeOutlined" />
                    </template>
                </n-switch>
            </div>
        </template>
    </n-card>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { useStore } from 'vuex';
import { useRouter } from 'vue-router';
import { useOsTheme, DrawerPlacement} from 'naive-ui';
import { LightModeOutlined, DarkModeOutlined, ChevronLeftRound, DehazeRound } from '@vicons/material';

const store = useStore();
const router = useRouter();
const activeDrawerRef = ref(false)
const placementRef = ref<DrawerPlacement>('left')

const props = withDefaults(
    defineProps<{
        title: string;
        back?: boolean;
        theme?: boolean;
    }>(),
    {
        title: '',
        back: false,
        theme: true,
    }
);
const switchTheme = (theme: boolean) => {
    if (theme) {
        localStorage.setItem('PAOPAO_THEME', 'dark');
        store.commit('triggerTheme', 'dark');
    } else {
        localStorage.setItem('PAOPAO_THEME', 'light');
        store.commit('triggerTheme', 'light');
    }
};
const goBack = () => {
    if (window.history.length <= 1) {
        router.push({
            path: '/',
        });
    } else {
        router.go(-1);
    }
};
const activeDrawer = () => {
    activeDrawerRef.value = true
};

onMounted(() => {
    if (!localStorage.getItem('PAOPAO_THEME')) {
        switchTheme((useOsTheme() as unknown as string) === 'dark');
    }
});
</script>

<style lang="less">
.nav-title-card {
    z-index: 99;
    width: 100%;
    top: 0;
    position: sticky;
    border-radius: 0;
    border-bottom: 0;
    background-color: rgba(255, 255, 255, 0.75);
    backdrop-filter: blur(12px);

    .navbar {
        height: 30px;
        position: relative;
        display: flex;
        align-items: center;

        .drawer-btn, 
        .back-btn {
            margin-right: 8px;
        }

        .theme-switch-wrap {
            position: absolute;
            right: 0;
            top: calc(50% - 9px);
        }
    }
}
.dark {
    .nav-title-card {
        background-color: rgba(16, 16, 20, 0.75);
    }
}
</style>