declare module NetReq {

    interface AuthUserLogin {
        token?: string
    }

    interface UserGetCollections {
        list: any[],
        pager: AnyObject
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
        list: Item.MessageProps[],
        pager: Item.PagerProps
    }

    interface UserGetUserPosts {
        list: Item.PostProps[],
        pager: Item.PagerProps
    }

    interface UserGetStars {
        list: Item.PostProps[],
        pager: Item.PagerProps
    }

    type UserGetUserProfile = Item.UserInfo

    interface UserGetBills {
        list: Item.BillProps[],
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
        b64s: string
    }

    type PostGetPost = Item.PostProps

    interface PostGetPosts {
        list: Item.PostProps[],
        pager: Item.PagerProps
    }

    interface PostLockPost {
        lock_status: number
    }

    interface PostStickPost {
        top_status: number
    }

    interface PostGetPostStar {
        status: boolean
    }

    interface PostGetPostCollection {
        status: boolean
    }

    interface PostPostStar {
        status: boolean
    }

    interface PostPostCollection {
        status: boolean
    }

    type PostGetTags = Item.TagProps[]

    interface PostGetPostComments {
        list: Item.PostProps[]
    }

    type PostCreatePost = Item.PostProps

}