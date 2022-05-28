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
        user: UserInfo,
        at_user_id: number,
        at_user: UserInfo,
        content: string,
        ip_loc: string,
        created_on: number
    }

    interface PostProps {
        id: number,
        type: number,
        user: UserInfo,
        attachment_price: number,
        ip_loc: string,
        created_on: number,
        latest_replied_on: number,
        upvote_count: number,
        comment_count: number,
        collection_count: number,
        content: string,
        contents: PostProps[],
        is_lock: number,
        is_top: boolean,
        texts?: PostProps[],
        imgs?: PostProps[],
        videos?: PostProps[],
        links?: PostProps[],
        attachments?: PostProps[],
        charge_attachments?: PostProps[]
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

    interface AttachmentProps {
        id: number,
        type: number
    }

    interface TagProps {
        id: number,
        quote_num: number,
        tag: string,
        user: UserInfo
    }

    interface PagerProps {
        total_rows: number
    }

    interface BillProps {
        id: number,
        reason: string,
        change_amount: number,
        created_on: number
    }

}