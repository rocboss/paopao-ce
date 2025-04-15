declare module NetReq {
  interface AuthUserLogin {
    token: string;
  }

  interface AuthUserRegister {
    /** 用户UID */
    id: number;
    /** 用户名 */
    username: string;
  }

  type AuthUserInfo = Item.UserInfo;

  interface AuthUpdateUserPassword {}

  interface UserGetCollections {
    /** 帖子列表 */
    list: Item.PostProps[];
    /** 页码信息 */
    pager: Item.PagerProps;
  }

  interface UserGetSuggestUsers {
    suggest: string[];
  }

  interface UserGetSuggestTags {
    suggest: string[];
  }

  interface UserPrecheckAttachment {
    paid: number;
  }

  interface UserGetAttachment {
    signed_url: string;
  }

  interface UserGetUnreadMsgCount {
    count: number;
  }

  interface ReadMessageResp {}

  interface ReadAllMessageResp {}

  interface UserGetMessages {
    /** 消息列表 */
    list: Item.MessageProps[];
    /** 页码信息 */
    pager: Item.PagerProps;
  }

  interface UserGetUserPosts {
    /** 帖子列表 */
    list: Item.PostProps[];
    /** 页码信息 */
    pager: Item.PagerProps;
  }

  type UserGetUserProfile = Item.UserInfo;

  interface UserGetBills {
    list: Item.BillProps[];
    /** 页码信息 */
    pager: Item.PagerProps;
  }

  interface UserReqRecharge {
    id: number;
    pay: string;
  }

  interface UserGetRecharge {
    status: string;
  }

  interface UserBindUserPhone {}

  interface UserGetCaptcha {
    id: string;
    /** 头像图片 base64 */
    b64s: string;
  }

  interface UserChangeNickname {}

  interface UserChangePassword {}

  interface UserChangeStatus {}

  interface SiteInfoResp {
    register_user_count: number;
    online_user_count: number;
    history_max_online: number;
    server_up_time: number;
  }

  interface FollowUserResp {}

  interface UnfollowUserResp {}

  interface AddFriend {}

  interface DeleteFriend {}

  interface RejectFriend {}

  interface RequestingFriend {}

  type PostGetPost = Item.PostProps;

  interface PostGetPosts {
    /** 帖子列表 */
    list: Item.PostProps[];
    /** 页码信息 */
    pager: Item.PagerProps;
  }

  interface PostLockPost {
    /** 锁定状态：0为未锁定，1为锁定 */
    lock_status: 0 | 1;
  }

  interface PostStickPost {
    /** 置顶状态：0为未置顶，1为置顶 */
    top_status: 0 | 1;
  }

  interface PostHighlightPost {
    /** 置顶状态：0为未亮点，1为亮点 */
    highlight_status: 0 | 1;
  }

  interface PostVisibilityPost {
    /** 可见性：0为公开，1为私密，2为好友可见 */
    visibility_status: import('@/utils/IEnum').VisibilityEnum;
  }

  interface PostGetPostStar {
    status: boolean;
  }

  interface PostPostStar {
    status: boolean;
  }

  interface PostGetPostCollection {
    status: boolean;
  }

  interface PostPostCollection {
    status: boolean;
  }

  interface PostGetTags {
    topics: Item.TagProps[];
    extral_topics?: Item.TagProps[];
  }

  interface PostGetPostComments {
    /** 评论列表 */
    list: Item.CommentProps[];
    /** 页码信息 */
    pager: Item.PagerProps;
  }

  type PostCreatePost = Item.PostProps;

  interface PostDeletePost {}

  interface PostTweetCommentThumbs {}

  interface PostTweetReplyThumbs {}

  type PostCreateComment = Item.CommentProps;

  interface PostDeleteComment {}

  interface PostHighlightComment {
    highlight_status: import('@/utils/IEnum').YesNoEnum;
  }

  type PostCreateCommentReply = Item.ReplyProps;

  interface PostDeleteCommentReply {}

  interface GetContacts {
    /** 评论列表 */
    list: Item.ContactItemProps[];
    /** 页码信息 */
    pager: Item.PagerProps;
  }

  interface IndexTrendsResp {
    list: Item.IndexTrendsItem[];
    pager: Item.PagerProps;
  }

  interface PostStickTopic {
    /** 置顶状态：0为未置顶，1为置顶 */
    top_status: 0 | 1;
  }

  interface PostPinTopic {
    /** 钉住状态：0为未钉住，1为钉住*/
    pin_status: 0 | 1;
  }

  interface PostFollowTopic {}

  interface PostUnfollowTopic {}

  interface SiteProfile {
    use_friendship?: boolean;
    enable_trends_bar?: boolean;
    enable_wallet?: boolean;
    allow_tweet_attachment?: boolean;
    allow_tweet_attachment_price?: boolean;
    allow_tweet_video?: boolean;
    allow_user_register?: boolean;
    allow_phone_bind?: boolean;
    default_tweet_max_length?: number;
    default_tweet_ellipsis_size?: number;
    default_tweet_visibility?: string;
    default_msg_loop_interval?: number;
    copyright_top?: string;
    copyright_left?: string;
    copyright_left_link?: string;
    copyright_right?: string;
    copyright_right_link?: string;
  }
}
