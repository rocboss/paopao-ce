<template>
    <div class="attachment-wrap">
        <div
            class="attach-item"
            v-for="attachment in attachments"
            :key="attachment.id"
        >
            <n-button
                @click.stop="download(attachment)"
                type="primary"
                size="tiny"
                dashed
            >
                <template #icon>
                    <n-icon>
                        <cloud-download-outline />
                    </n-icon>
                </template>
                {{ attachment.type === 8 ? '收费' : '免费' }}附件
            </n-button>
        </div>

        <!-- 删除确认 -->
        <n-modal
            v-model:show="showDownloadModal"
            :mask-closable="false"
            preset="dialog"
            title="下载提示"
            :content="downloadTip"
            positive-text="确认下载"
            negative-text="取消"
            icon-placement="top"
            @positive-click="execDownloadAction"
        />
    </div>
</template>

<script setup lang="ts">
import { h, ref } from 'vue';
import { CloudDownloadOutline } from '@vicons/ionicons5';
import { precheckAttachment, getAttachment } from '@/api/user';

const props = withDefaults(
  defineProps<{
    attachments: Item.PostItemProps[];
    price?: number;
  }>(),
  {
    attachments: () => [],
    price: 0,
  },
);
const showDownloadModal = ref(false);
const downloadTip = ref<any>('');
const attachmentID = ref(0);

const download = (attachment: Item.PostItemProps) => {
  showDownloadModal.value = true;
  attachmentID.value = attachment.id;

  downloadTip.value = '这是一个免费附件，您可以直接下载？';
  if (attachment.type === 8) {
    downloadTip.value = () =>
      h('div', {}, [
        h(
          'p',
          {},
          '这是一个收费附件，下载将收取' +
            (props.price / 100).toFixed(2) +
            '元',
        ),
      ]);

    precheckAttachment({
      id: attachmentID.value,
    })
      .then((res) => {
        if (res.paid) {
          downloadTip.value = () =>
            h('div', {}, [
              h('p', {}, '此次下载您已支付或无需付费，请确认下载'),
            ]);
        }
      })
      .catch((err) => {
        showDownloadModal.value = false;
      });
  }
};
const execDownloadAction = () => {
  getAttachment({
    id: attachmentID.value,
  })
    .then((res) => {
      window.open(res.signed_url.replace('http://', 'https://'), '_blank');
    })
    .catch((err) => {
      console.log(err);
    });
};
</script>

<style lang="less" scoped>
.attach-item {
    margin: 10px 0;
}
</style>