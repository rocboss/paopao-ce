declare module NetParams {

    interface AuthUserLogin {
        username?: string,
        password?: string
    }

    interface UserGetCollections {
        page: number,
        page_size: number
    }

    interface UserPrecheckAttachment {
        id: number
    }

    interface UserGetAttachment {
        id: number
    }

    interface UserGetUnreadMsgCount {

    }

    interface UserGetMessages {
        page: number,
        page_size: number
    }

    interface UserGetUserPosts {
        username: string,
        page: number,
        page_size: number
    }

    interface UserGetStars {
        page: number,
        page_size: number
    }

    interface UserGetUserProfile {
        username: string
    }

    interface UserGetBills {
        page: number,
        page_size: number
    }

    interface UserReqRecharge {
        amount: number
    }

    interface UserGetRecharge {
        id: number
    }

    interface UserBindUserPhone {
        phone: string,
        captcha: string
    }

    interface UserGetCaptcha {
        
    }

    interface PostGetPost {
        id: number
    }

    interface PostGetPosts {
        query: string | null,
        type: number,
        page: number,
        page_size: number
    }

    interface PostLockPost {
        id: number
    }

    interface PostGetPostStar {
        id: number
    }

    interface PostGetPostCollection {
        id: number
    }

    interface PostPostStar {
        id: number
    }

    interface PostPostCollection {
        id: number
    }

    interface PostGetTags {
        type: "hot" | string,
        num: number 
    }

    interface PostGetPostComments {
        id: number
    }

}