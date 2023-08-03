<template>
    <div class="images-wrap">
        <n-image-group v-if="[1].includes(props.imgs.length)">
            <n-grid :x-gap="4" :y-gap="4" :cols="2">
                <template v-for="img in props.imgs" :key="img.id">
                    <n-gi>
                        <n-image
                            @error="() => (img.content = defaultImg)"
                            @click.stop
                            class="post-img x1"
                            object-fit="cover"
                            :src="img.content + thumbnail"
                            :preview-src="img.content"
                        />
                    </n-gi>
                </template>
            </n-grid>
        </n-image-group>

        <n-image-group v-if="[2, 3].includes(props.imgs.length)">
            <n-grid :x-gap="4" :y-gap="4" :cols="3">
                <template v-for="img in props.imgs" :key="img.id">
                    <n-gi>
                        <n-image
                            @error="() => (img.content = defaultImg)"
                            @click.stop
                            class="post-img x2"
                            object-fit="cover"
                            :src="img.content + thumbnail"
                            :preview-src="img.content"
                        />
                    </n-gi>
                </template>
            </n-grid>
        </n-image-group>

        <n-image-group v-if="[4].includes(props.imgs.length)">
            <n-grid :x-gap="4" :y-gap="4" :cols="4">
                <template v-for="img in props.imgs" :key="img.id">
                    <n-gi>
                        <n-image
                            @error="() => (img.content = defaultImg)"
                            @click.stop
                            class="post-img x3"
                            object-fit="cover"
                            :src="img.content + thumbnail"
                            :preview-src="img.content"
                        />
                    </n-gi>
                </template>
            </n-grid>
        </n-image-group>

        <n-image-group v-if="[5].includes(props.imgs.length)">
            <n-grid :x-gap="4" :y-gap="4" :cols="3">
                <template v-for="(img, idx) in props.imgs" :key="img.id">
                    <n-gi v-if="idx < 3">
                        <n-image
                            @error="() => (img.content = defaultImg)"
                            @click.stop
                            class="post-img x2"
                            object-fit="cover"
                            :src="img.content + thumbnail"
                            :preview-src="img.content"
                        />
                    </n-gi>
                </template>
            </n-grid>
            <n-grid :x-gap="4" :y-gap="4" :cols="2" style="margin-top: 4px">
                <template v-for="(img, idx) in props.imgs" :key="img.id">
                    <n-gi v-if="idx >= 3">
                        <n-image
                            @error="() => (img.content = defaultImg)"
                            @click.stop
                            class="post-img x1"
                            object-fit="cover"
                            :src="img.content + thumbnail"
                            :preview-src="img.content"
                        />
                    </n-gi>
                </template>
            </n-grid>
        </n-image-group>

        <n-image-group v-if="[6].includes(props.imgs.length)">
            <n-grid :x-gap="4" :y-gap="4" :cols="3">
                <template v-for="(img, idx) in props.imgs" :key="img.id">
                    <n-gi v-if="idx < 3">
                        <n-image
                            @error="() => (img.content = defaultImg)"
                            @click.stop
                            class="post-img x2"
                            object-fit="cover"
                            :src="img.content + thumbnail"
                            :preview-src="img.content"
                        />
                    </n-gi>
                </template>
            </n-grid>
            <n-grid :x-gap="4" :y-gap="4" :cols="3" style="margin-top: 4px">
                <template v-for="(img, idx) in props.imgs" :key="img.id">
                    <n-gi v-if="idx >= 3">
                        <n-image
                            @error="() => (img.content = defaultImg)"
                            @click.stop
                            class="post-img x2"
                            object-fit="cover"
                            :src="img.content + thumbnail"
                            :preview-src="img.content"
                        />
                    </n-gi>
                </template>
            </n-grid>
        </n-image-group>

        <n-image-group v-if="props.imgs.length === 7">
            <n-grid :x-gap="4" :y-gap="4" :cols="4">
                <template v-for="(img, idx) in props.imgs">
                    <n-gi :key="img.id" v-if="idx < 4">
                        <n-image
                            @error="() => (img.content = defaultImg)"
                            @click.stop
                            class="post-img x3"
                            object-fit="cover"
                            :src="img.content + thumbnail"
                            :preview-src="img.content"
                        />
                    </n-gi>
                </template>
            </n-grid>
            <n-grid :x-gap="4" :y-gap="4" :cols="3" style="margin-top: 4px">
                <template v-for="(img, idx) in props.imgs">
                    <n-gi :key="img.id" v-if="idx >= 4">
                        <n-image
                            @error="() => (img.content = defaultImg)"
                            @click.stop
                            class="post-img x2"
                            object-fit="cover"
                            :src="img.content + thumbnail"
                            :preview-src="img.content"
                        />
                    </n-gi>
                </template>
            </n-grid>
        </n-image-group>

        <n-image-group v-if="props.imgs.length === 8">
            <n-grid :x-gap="4" :y-gap="4" :cols="4">
                <template v-for="(img, idx) in props.imgs">
                    <n-gi :key="img.id" v-if="idx < 4">
                        <n-image
                            @error="() => (img.content = defaultImg)"
                            @click.stop
                            class="post-img x3"
                            object-fit="cover"
                            :src="img.content + thumbnail"
                            :preview-src="img.content"
                        />
                    </n-gi>
                </template>
            </n-grid>
            <n-grid :x-gap="4" :y-gap="4" :cols="4" style="margin-top: 4px">
                <template v-for="(img, idx) in props.imgs">
                    <n-gi :key="img.id" v-if="idx >= 4">
                        <n-image
                            @error="() => (img.content = defaultImg)"
                            @click.stop
                            class="post-img x3"
                            object-fit="cover"
                            :src="img.content + thumbnail"
                            :preview-src="img.content"
                        />
                    </n-gi>
                </template>
            </n-grid>
        </n-image-group>

        <n-image-group v-if="props.imgs.length === 9">
            <n-grid :x-gap="4" :y-gap="4" :cols="3">
                <template v-for="(img, idx) in props.imgs">
                    <n-gi :key="img.id" v-if="idx < 3">
                        <n-image
                            @error="() => (img.content = defaultImg)"
                            @click.stop
                            class="post-img x2"
                            object-fit="cover"
                            :src="img.content + thumbnail"
                            :preview-src="img.content"
                        />
                    </n-gi>
                </template>
            </n-grid>
            <n-grid :x-gap="4" :y-gap="4" :cols="3" style="margin-top: 4px">
                <template v-for="(img, idx) in props.imgs">
                    <n-gi :key="img.id" v-if="idx >= 3 && idx < 6">
                        <n-image
                            @error="() => (img.content = defaultImg)"
                            @click.stop
                            class="post-img x2"
                            object-fit="cover"
                            :src="img.content + thumbnail"
                            :preview-src="img.content"
                        />
                    </n-gi>
                </template>
            </n-grid>
            <n-grid :x-gap="4" :y-gap="4" :cols="3" style="margin-top: 4px">
                <template v-for="(img, idx) in props.imgs">
                    <n-gi :key="img.id" v-if="idx >= 6">
                        <n-image
                            @error="() => (img.content = defaultImg)"
                            @click.stop
                            class="post-img x2"
                            object-fit="cover"
                            :src="img.content + thumbnail"
                            :preview-src="img.content"
                        />
                    </n-gi>
                </template>
            </n-grid>
        </n-image-group>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';

const defaultImg = import.meta.env.VITE_DEFAULT_TWEET_IMAGE_404;
const thumbnail = import.meta.env.VITE_TWEET_IMAGE_THUMBNAIL;
const props = withDefaults(defineProps<{
    imgs: Item.PostItemProps[],
}>(), {
    imgs: () => []
});
</script>

<style lang="less">
.images-wrap {
    margin-top: 10px;
}
.post-img {
    display: flex;
    margin: 0;
    border-radius: 3px;
    overflow: hidden;
    background: rgba(0, 0, 0, 0.1);
    border: 1px solid #eee;
    img {
        width: 100%;
        height: 100%;
    }
}
.x1 {
    height: 168px;
}
.x2 {
    height: 108px;
}
.x3 {
    height: 96px;
}
.dark {
    .post-img {
        border: 1px solid #333;
    }
}
@media screen and (max-width: 821px) {
    .x1 {
        height: 100px;
    }
    .x2 {
        height: 70px;
    }
    .x3 {
        height: 50px;
    }
}
</style>