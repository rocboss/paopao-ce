<template>
    <div v-if="props.videos.length > 0">
        <n-grid :x-gap="4" :y-gap="4" :cols="full ? 1 : 5">
            <n-gi :span="full ? 1 : 3">
                <n-video
                    @click.stop
                    v-for="video in props.videos"
                    :key="video.id"
                    :src="video.content"
                    :colors="['#18a058', '#2aca75']"
                    :hoverable="true"
                    theme="gradient"
                ></n-video>
            </n-gi>
        </n-grid>
    </div>
</template>

<script setup>
import { reactive } from 'vue';
import NVideo from 'nonesir-video';

const props = defineProps({
    videos: {
        type: Array,
        default: () => [],
    },
    full: {
        type: Boolean,
        default: false,
    },
});
</script>