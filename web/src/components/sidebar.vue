<template>
    <div class="sidebar-wrap">
        <div class="logo-wrap">
            <n-image class="logo-img" width="36" :src="LOGO" :preview-disabled="true" @click="goHome" />
        </div>
        <n-menu :accordion="true" :icon-size="24"
            :options="menuOptions" :render-label="renderMenuLabel" :render-icon="renderMenuIcon" :value="selectedPath"
            @update:value="goRouter" />

        <div class="user-wrap" v-if="store.state.userInfo.id > 0">
            <n-avatar class="user-avatar" round :size="34" :src="store.state.userInfo.avatar" />

            <div class="user-info">
                <div class="nickname">
                    <span class="nickname-txt">
                        {{ store.state.userInfo.nickname }}
                    </span>
                    <n-button class="logout" quaternary circle size="tiny" @click="handleLogout">
                        <template #icon>
                            <n-icon>
                                <log-out-outline />
                            </n-icon>
                        </template>
                    </n-button>
                </div>
                <div class="username">@{{ store.state.userInfo.username }}</div>
            </div>

            <div class="user-mini-wrap">
                <n-button class="logout" quaternary circle @click="handleLogout">
                    <template #icon>
                        <n-icon :size="24">
                            <log-out-outline />
                        </n-icon>
                    </template>
                </n-button>
            </div>
        </div>
        <div class="user-wrap" v-else>
            <div v-if="!allowUserRegister" class="login-only-wrap">
                <n-button strong secondary round type="primary" @click="triggerAuth('signin')">
                    登录
                </n-button>
            </div>
            <div v-if="allowUserRegister" class="login-wrap">
                <n-button strong secondary round type="primary" @click="triggerAuth('signin')">
                    登录
                </n-button>
                <n-button strong secondary round type="info" @click="triggerAuth('signup')">
                    注册
                </n-button>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { h, ref, watch, computed, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useStore } from 'vuex';
import { NIcon, NBadge, useMessage } from 'naive-ui';
import {
    HomeOutline,
    BookmarksOutline,
    MegaphoneOutline,
    ChatbubblesOutline,
    LeafOutline,
    PeopleOutline,
    WalletOutline,
    SettingsOutline,
    LogOutOutline,
} from '@vicons/ionicons5';
import { Hash } from '@vicons/tabler';
import { getUnreadMsgCount } from '@/api/user';
import LOGO from '@/assets/img/logo.png';

const store = useStore();
const route = useRoute();
const router = useRouter();
const hasUnreadMsg = ref(false);
const selectedPath = ref<any>(route.name || '');
const msgLoop = ref();
const allowUserRegister = ref(import.meta.env.VITE_ALLOW_USER_REGISTER.toLowerCase() === 'true')
const defMsgLoopInterval = Number(import.meta.env.VITE_DEFAULT_MSG_LOOP_INTERVAL)

watch(route, () => {
    selectedPath.value = route.name;
});
watch(store.state, () => {
    if (store.state.userInfo.id > 0) {
        if (!msgLoop.value) {
            getUnreadMsgCount()
                .then((res) => {
                    hasUnreadMsg.value = res.count > 0;
                })
                .catch((err) => {
                    console.log(err);
                });

            msgLoop.value = setInterval(() => {
                getUnreadMsgCount()
                    .then((res) => {
                        hasUnreadMsg.value = res.count > 0;
                    })
                    .catch((err) => {
                        console.log(err);
                    });
            }, defMsgLoopInterval);
        }
    } else {
        if (msgLoop.value) {
            clearInterval(msgLoop.value);
        }
    }
});
onMounted(() => {
    window.onresize = () => {
        store.commit('triggerCollapsedLeft', document.body.clientWidth <= 821);
        store.commit('triggerCollapsedRight', document.body.clientWidth <= 821);
    };
});
const menuOptions = computed(() => {
    const options = [
        {
            label: '广场',
            key: 'home',
            icon: () => h(HomeOutline),
            href: '/',
        },
        {
            label: '话题',
            key: 'topic',
            icon: () => h(Hash),
            href: '/topic',
        },
    ];
    if (import.meta.env.VITE_ENABLE_ANOUNCEMENT.toLowerCase() === 'true') {
        options.push( {
            label: '公告',
            key: 'anouncement',
            icon: () => h(MegaphoneOutline),
            href: '/anouncement',
        });
    }
    options.push({
        label: '主页',
        key: 'profile',
        icon: () => h(LeafOutline),
        href: '/profile',
    });
    options.push({
        label: '消息',
        key: 'messages',
        icon: () => h(ChatbubblesOutline),
        href: '/messages',
    })
    options.push({
        label: '收藏',
        key: 'collection',
        icon: () => h(BookmarksOutline),
        href: '/collection',
    });
    options.push({
        label: '好友',
        key: 'contacts',
        icon: () => h(PeopleOutline),
        href: '/contacts',
    });
    if (import.meta.env.VITE_ENABLE_WALLET.toLocaleLowerCase() === 'true') {
        options.push({
            label: '钱包',
            key: 'wallet',
            icon: () => h(WalletOutline),
            href: '/wallet',
        });
    }
    options.push({
        label: '设置',
        key: 'setting',
        icon: () => h(SettingsOutline),
        href: '/setting',
    });

    return store.state.userInfo.id > 0
        ? options
        : [
            {
                label: '广场',
                key: 'home',
                icon: () => h(HomeOutline),
                href: '/',
            },
            {
                label: '话题',
                key: 'topic',
                icon: () => h(Hash),
                href: '/topic',
            },
        ];
});

const renderMenuLabel = (option: AnyObject) => {
    if ('href' in option) {
        return h('div', {}, option.label);
    }
    return option.label;
};
const renderMenuIcon = (option: AnyObject) => {
    if (option.key === 'messages') {
        return h(
            NBadge,
            {
                dot: true,
                show: hasUnreadMsg.value,
                processing: true,
            },
            {
                default: () =>
                    h(
                        NIcon,
                        {
                            color:
                                option.key === selectedPath.value
                                    ? 'var(--n-item-icon-color-active)'
                                    : 'var(--n-item-icon-color)',
                        },
                        { default: option.icon }
                    ),
            }
        );
    }
    return h(NIcon, null, { default: option.icon });
};

const goRouter = (name: string, item: any = {}) => {
    selectedPath.value = name;
    router.push({ name });
};
const goHome = () => {
    if (route.path === '/') {
        store.commit('refresh');
    }
    goRouter('home');
};
const triggerAuth = (key: string) => {
    store.commit('triggerAuth', true);
    store.commit('triggerAuthKey', key);
};
const handleLogout = () => {
    store.commit('userLogout');
    store.commit('refresh')
    goHome()
};
window.$store = store;
window.$message = useMessage();
</script>

<style lang="less">
.sidebar-wrap {
    z-index: 99;
    width: 200px;
    height: 100vh;
    position: fixed;
    right: calc(50% + var(--content-main) / 2 + 10px);
    padding: 12px 0;
    box-sizing: border-box;

    .n-menu .n-menu-item-content::before {
        border-radius: 21px;
    }
}

.logo-wrap {
    display: flex;
    justify-content: flex-start;
    margin-bottom: 12px;

    .logo-img {
        margin-left: 24px;

        &:hover {
            cursor: pointer;
        }
    }
}

.user-wrap {
    display: flex;
    align-items: center;
    position: absolute;
    bottom: 12px;
    left: 12px;
    right: 12px;

    .user-mini-wrap {
        display: none;
    }

    .user-avatar {
        margin-right: 8px;
    }

    .user-info {
        display: flex;
        flex-direction: column;

        .nickname {
            font-size: 16px;
            font-weight: bold;
            line-height: 16px;
            height: 16px;
            margin-bottom: 2px;
            display: flex;
            align-items: center;

            .nickname-txt {
                max-width: 90px;
                text-overflow: ellipsis;
                overflow: hidden;
                white-space: nowrap;
            }

            .logout {
                margin-left: 6px;
            }
        }

        .username {
            font-size: 14px;
            line-height: 16px;
            height: 16px;
            width: 120px;
            text-overflow: ellipsis;
            overflow: hidden;
            white-space: nowrap;
            opacity: 0.75;
        }
    }

    .login-only-wrap {
        display: flex;
        justify-content:center;
        width: 100%;

        button {
            margin: 0 4px;
            width: 80%
        }
    }

    .login-wrap {
        display: flex;
        justify-content: center;
        width: 100%;

        button {
            margin: 0 4px;
        }
    }
}

.auth-card {
    .n-card-header {
        z-index: 999;
    }
}

@media screen and (max-width: 821px) {
    .sidebar-wrap {
        width: 200px;
        right: calc(100% - 200px);
    }

    .logo-wrap {
        .logo-img {
            margin-left: 12px !important;
        }
    }

    .user-wrap {
        .user-avatar,
        .user-info,
        .login-only-wrap,
        .login-wrap {
            margin-bottom: 32px;
        }

    //     .user-mini-wrap {
    //         display: block !important;
    //     }
    }
}
</style>