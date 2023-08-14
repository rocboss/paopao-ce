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
    visibility_status: import("@/utils/IEnum").VisibilityEnum;
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

  type PostCreateCommentReply = Item.ReplyProps;

  interface PostDeleteCommentReply {}

  interface GetContacts {
    /** 评论列表 */
    list: Item.ContactItemProps[];
    /** 页码信息 */
    pager: Item.PagerProps;
  }

  interface PostStickTopic {
    /** 置顶状态：0为未置顶，1为置顶 */
    top_status: 0 | 1;
  }

  interface PostFollowTopic {}

  interface PostUnfollowTopic {}
}
