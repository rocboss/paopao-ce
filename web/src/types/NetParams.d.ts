declare namespace NetParams {
  type AuthUserInfo = string;

  interface AuthUpdateUserPassword {
    /** 新密码 */
    password: string;
    /** 旧密码 */
    old_password: string;
  }

  interface SiteInfoReq {}

  interface IndexTrendsReq {
    page: number;
    page_size: number;
  }

  interface PostGetPost {
    id: number;
  }

  interface PostGetPosts {
    query: string | null;
    type: string;
    style: 'newest' | 'hots' | 'following' | 'search';
    page: number;
    page_size: number;
  }

  interface PostLockPost {
    id: number;
  }

  interface PostStickPost {
    id: number;
  }

  interface PostHighlightPost {
    id: number;
  }

  interface PostVisibilityPost {
    id: number;
    /** 可见性：0为公开，1为私密，2为好友可见 */
    visibility: import('@/utils/IEnum').VisibilityEnum;
  }

  interface PostGetPostStar {
    id: number;
  }

  interface PostPostStar {
    id: number;
  }

  interface PostGetPostCollection {
    id: number;
  }

  interface PostPostCollection {
    id: number;
  }

  interface PostGetTags {
    type: 'hot' | 'new' | 'follow' | 'pin' | 'hot_extral';
    num: number;
    extral_num?: number;
  }

  interface PostGetPostComments {
    id: number;
    style: 'default' | 'hots' | 'newest';
    page?: number;
    page_size?: number;
  }

  interface PostCreatePost {
    /** 帖子内容列表 */
    contents: Partial<Item.PostItemProps>[];
    /** 标签列表 */
    tags: string[];
    /** 艾特用户列表 */
    users: string[];
    /** 附件价格 */
    attachment_price: number;
    /** 可见性：0为公开，1为私密，2为好友可见 */
    visibility: import('@/utils/IEnum').VisibilityEnum;
  }

  interface PostDeletePost {
    id: number;
  }

  interface PostTweetCommentThumbs {
    tweet_id: number;
    comment_id: number;
  }

  interface PostTweetReplyThumbs {
    tweet_id: number;
    comment_id: number;
    reply_id: number;
  }

  interface PostCreateComment {
    /** 内容ID */
    post_id: number;
    /** 帖子内容列表 */
    contents: Partial<Item.CommentItemProps>[];
    /** 艾特用户列表 */
    users: string[];
  }

  interface PostDeleteComment {
    id: number;
  }

  interface PostHighlightComment {
    id: number;
  }

  interface PostCreateCommentReply {
    /** 艾特的用户UID */
    at_user_id: number;
    /** 回复的评论ID */
    comment_id: number;
    /** 回复内容 */
    content: string;
  }

  interface PostDeleteCommentReply {
    id: number;
  }

  interface PostStickTopic {
    topic_id: number;
  }

  interface PostPinTopic {
    topic_id: number;
  }

  interface PostFollowTopic {
    topic_id: number;
  }

  interface PostUnfollowTopic {
    topic_id: number;
  }
}
