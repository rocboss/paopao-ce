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

    interface PostGetPosts {
        list: any[],
        pager: AnyObject
    }

}