declare module NetReq {

    interface AuthUserLogin {
        token: string
    }

    interface AuthUserRegister {
        /** 用户UID */
        id: number,
        /** 用户名 */
        username: string
    }

    type AuthUserInfo = Item.UserInfo

    interface AuthUpdateUserPassword {

    }

    interface UserGetCollections {
        /** 帖子列表 */
        list: Item.PostProps[],
        /** 页码信息 */
        pager: Item.PagerProps
    }

    type UserGetSuggestUsers = string[]
    type UserGetSuggestTags = string[]

    interface UserPrecheckAttachment {
        paid: number
    }

    type UserGetAttachment = string

    interface UserGetUnreadMsgCount {
        count: number
    }

    interface UserGetMessages {
        /** 消息列表 */
        list: Item.MessageProps[],
        /** 页码信息 */
        pager: Item.PagerProps
    }

    interface UserGetUserPosts {
        /** 帖子列表 */
        list: Item.PostProps[],
        /** 页码信息 */
        pager: Item.PagerProps
    }

    interface UserGetStars {
        /** 帖子列表 */
        list: Item.PostProps[],
        /** 页码信息 */
        pager: Item.PagerProps
    }

    type UserGetUserProfile = Item.UserInfo

    interface UserGetBills {
        list: Item.BillProps[],
        /** 页码信息 */
        pager: Item.PagerProps
    }

    interface UserReqRecharge {
        id: number,
        pay: string
    }

    interface UserGetRecharge {
        status: string
    }

    interface UserBindUserPhone {

    }

    interface UserGetCaptcha {
        id: string,
        /** 头像图片 base64 */
        b64s: string
    }

    interface UserChangeNickname {
        
    }

    interface UserChangePassword {

    }

    type PostGetPost = Item.PostProps

    interface PostGetPosts {
        /** 帖子列表 */
        list: Item.PostProps[],
        /** 页码信息 */
        pager: Item.PagerProps
    }

    interface PostLockPost {
        /** 锁定状态：0为未锁定，1为锁定 */
        lock_status: 0 | 1
    }

    interface PostStickPost {
        /** 置顶状态：0为未置顶，1为置顶 */
        top_status: 0 | 1
    }

    interface PostGetPostStar {
        status: boolean
    }

    interface PostPostStar {
        status: boolean
    }

    interface PostGetPostCollection {
        status: boolean
    }

    interface PostPostCollection {
        status: boolean
    }

    type PostGetTags = Item.TagProps[]

    interface PostGetPostComments {
        /** 评论列表 */
        list: Item.CommentProps[],
        /** 页码信息 */
        pager: Item.PagerProps
    }

    type PostCreatePost = Item.PostProps

    interface PostDeletePost {
        
    }

    type PostCreateComment = Item.CommentProps

    interface PostDeleteComment {
        
    }

    type PostCreateCommentReply = Item.ReplyProps

    interface PostDeleteCommentReply{
        
    }

}
