<template>
    <div class="rightbar-wrap" v-if="!store.state.collapsedRight">
        <div class="search-wrap">
            <n-input
                round
                clearable
                placeholder="搜一搜..."
                v-model:value="keyword"
                @keyup.enter.prevent="handleSearch"
            >
                <template #prefix>
                    <n-icon :component="Search" />
                </template>
            </n-input>
        </div>
        <n-card title="热门话题" embedded :bordered="false" size="small">
            <n-spin :show="loading">
                <div class="hot-tag-item" v-for="tag in tags" :key="tag.id">
                    <router-link
                        class="hash-link"
                        :to="{
                            name: 'home',
                            query: {
                                q: tag.tag,
                                t: 'tag',
                            },
                        }"
                    >
                        #{{ tag.tag }}
                    </router-link>

                    <div class="post-num">
                        {{ formatQuoteNum(tag.quote_num) }}
                    </div>
                </div>
            </n-spin>
        </n-card>
        <n-card class="copyright-wrap" embedded :bordered="false" size="small">
            <div class="copyright">&copy; 2022 PaoPao.Info</div>
            <div class="copyright">
                <a
                    href="https://beian.miit.gov.cn/"
                    target="_blank"
                    rel="noopener noreferrer"
                    class="beian-link"
                >
                    沪ICP备2020036525号-5
                </a>
            </div>
            <div>
                <n-space>
                    <a
                        href="https://www.rocs.me"
                        target="_blank"
                        class="hash-link"
                        >Roc's Me</a
                    >
                    <a
                        href="https://www.rocboss.com"
                        target="_blank"
                        class="hash-link"
                    >
                        ROCBOSS开源社区
                    </a>
                </n-space>
            </div>
        </n-card>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useStore } from 'vuex';
import { useRouter } from 'vue-router';
import { getTags } from '@/api/post';
import { Search } from '@vicons/ionicons5';

const tags = ref<Item.TagProps[]>([]);
const loading = ref(false);
const keyword = ref('');
const store = useStore();
const router = useRouter();

const loadTags = () => {
    loading.value = true;
    getTags({
        type: 'hot',
        num: 12,
    })
        .then((res) => {
            tags.value = res;
            loading.value = false;
        })
        .catch((err) => {
            loading.value = false;
        });
};
const formatQuoteNum = (num: number) => {
    if (num >= 1000) {
        return (num / 1000).toFixed(1) + 'k';
    }

    return num;
};
const handleSearch = () => {
    router.push({
        name: 'home',
        query: {
            q: keyword.value,
        },
    });
};
onMounted(() => {
    loadTags();
});
</script>

<style lang="less" scoped>
.rightbar-wrap {
    width: 240px;
    position: fixed;
    left: calc(50% + 260px);
    .search-wrap {
        margin: 12px 0;
    }

    .hot-tag-item {
        line-height: 2;
        position: relative;

        .hash-link {
            width: calc(100% - 60px);
            text-overflow: ellipsis;
            white-space: nowrap;
            overflow: hidden;
            display: block;
        }

        .post-num {
            position: absolute;
            right: 0;
            top: 0;
            width: 60px;
            text-align: right;
            line-height: 2;
            opacity: 0.5;
        }
    }

    .copyright-wrap {
        margin-top: 10px;

        .copyright {
            font-size: 12px;
            opacity: 0.75;
        }

        .beian-link {
            color: #333;
            text-decoration: none;
            &:hover {
                opacity: 0.75;
            }
        }

        .hash-link {
            font-size: 12px;
        }
    }
}
</style>