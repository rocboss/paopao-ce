import { computed } from 'vue';

/**
 * Post 内容解析 composable
 * 将 post.contents 按类型分类到 texts, imgs, videos, links, attachments, charge_attachments
 * 保持与原有代码完全一致的行为
 */
export function usePostContent(
  post: Item.PostProps,
  includeExtraFields: boolean = false,
) {
  return computed({
    get: () => {
      let postData: Item.PostComponentProps = Object.assign(
        {
          texts: [],
          imgs: [],
          videos: [],
          links: [],
          attachments: [],
          charge_attachments: [],
        },
        post,
      );
      postData.contents.map((content) => {
        if (+content.type === 1 || +content.type === 2) {
          postData.texts.push(content);
        }
        if (+content.type === 3) {
          postData.imgs.push(content);
        }
        if (+content.type === 4) {
          postData.videos.push(content);
        }
        if (+content.type === 6) {
          postData.links.push(content);
        }
        if (+content.type === 7) {
          postData.attachments.push(content);
        }
        if (+content.type === 8) {
          postData.charge_attachments.push(content);
        }
      });
      return postData;
    },
    set: (newVal) => {
      post.upvote_count = newVal.upvote_count;
      post.collection_count = newVal.collection_count;
      if (includeExtraFields) {
        post.comment_count = newVal.comment_count;
        post.is_essence = newVal.is_essence;
      }
    },
  });
}
