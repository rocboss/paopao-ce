declare module NetParams {
  interface AuthUserLogin {
    /** 用户名 */
    username: string;
    /** 密码 */
    password: string;
  }

  interface AuthUserRegister {
    /** 用户名 */
    username: string;
    /** 密码 */
    password: string;
  }

  type AuthUserInfo = string;

  interface AuthUpdateUserPassword {
    /** 新密码 */
    password: string;
    /** 旧密码 */
    old_password: string;
  }

  interface UserGetCollections {
    page: number;
    page_size: number;
  }

  interface UserPrecheckAttachment {
    id: number;
  }

  interface UserGetAttachment {
    id: number;
  }

  interface UserGetUnreadMsgCount {}

  interface UserGetMessages {
    page: number;
    page_size: number;
  }

  interface UserGetUserPosts {
    username: string;
    style: string;
    page: number;
    page_size: number;
  }

  interface UserGetUserProfile {
    username: string;
  }

  interface UserGetBills {
    page: number;
    page_size: number;
  }

  interface UserStatusReq {
    id: number;
    status: number;
  }

  interface FollowUserReq {
    user_id: number;
  }

  interface UnfollowUserReq {
    user_id: number;
  }

  interface UserReqRecharge {
    amount: number;
  }

  interface UserGetRecharge {
    id: number;
  }

  interface UserBindUserPhone {
    phone: string;
    captcha: string;
  }

  interface UserActivation {
    activate_code: string;
    captcha_id: string;
    imgCaptcha: string;
  }

  interface UserGetCaptcha {}

  interface UserWhisper {
    user_id: number;
    content: string;
  }

  interface RequestingFriend {
    user_id: number;
    greetings: string;
  }

  interface AddFriend {
    user_id: number;
  }

  interface RejectFriend {
    user_id: number;
  }

  interface DeleteFriend {
    user_id: number;
  }

  interface GetContacts {
    page: number;
    page_size: number;
  }

  interface GetUserFollows {
    username: string;
    page: number;
    page_size: number;
  }

  interface GetUserFollowings {
    username: string;
    page: number;
    page_size: number;
  }

  interface UserChangePassword {
    /** 新密码 */
    password: string;
    /** 旧密码 */
    old_password: string;
  }

  interface UserChangeNickname {
    /** 昵称 */
    nickname: string;
  }

  interface PostGetPost {
    id: number;
  }

  interface PostGetPosts {
    query: string | null;
    type: string;
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
    visibility: import("@/utils/IEnum").VisibilityEnum;
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
    type: "hot" | "new" | "follow" | "hot_extral";
    num: number;
    extral_num?: number;
  }

  interface PostGetPostComments {
    id: number;
    sort_strategy: "default" | "newest";
    page?: number;
    page_size?: number;
  }

  interface GetContacts {}

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
    visibility: import("@/utils/IEnum").VisibilityEnum;
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

  interface PostFollowTopic {
    topic_id: number;
  }

  interface PostUnfollowTopic {
    topic_id: number;
  }
}
