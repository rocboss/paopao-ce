declare module NetParams {

    interface AuthUserLogin {
        username?: string,
        password?: string
    }

    interface UserGetCollections {
        list: any[],
        pager: AnyObject
    }

    interface PostGetPosts {
        list: any[],
        pager: AnyObject
    }

}