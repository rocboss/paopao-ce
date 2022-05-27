declare module Item {

    interface UserInfo {
        id: number,
        username: string,
        nickname: string,
        avatar: string
    }

    interface CommentItemProps {
        id: number | string,
        content: string
    }

    interface CommentProps extends CommentItemProps {
        type: number,
        user: UserInfo,
        contents: CommentProps[],
        replies: ReplyProps[],
        ip_loc: string,
        created_on: number
    }

    interface ReplyProps {
        id: number,
        user_id: number,
        user: UserInfo
    }

    interface PostProps {
        id: number,
    }

    interface MessageProps {
        id: number,
        type: number,
        is_read: number,
        sender_user: UserInfo
        post: PostProps,
        post_id: number,
        created_on: number,
        breif: string
    }

}