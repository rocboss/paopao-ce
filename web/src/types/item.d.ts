declare module Item {

    interface UserInfo {
        /** 用户UID */
        id: number,
        /** 用户名 */
        username: string,
        /** 用户昵称 */
        nickname: string,
        /** 用户头像 */
        avatar: string,
        /** 用户手机号 */
        phone?: string,
        /** 是否为管理员 */
        is_admin: boolean,
        /** 用户余额（分） */
        balance?: number,
        /** 用户状态 */
        status?: 0 | 1
    }

    /** 评论内容 */
    interface CommentItemProps {
        /** 内容ID */
        id: number,
        /** 评论ID */
        comment_id: number,
        /** 评论者UID */
        user_id: number,
        /** 类别：1为标题，2为文字段落，3为图片地址，4为视频地址，5为语音地址，6为链接地址 */
        type: number,
        /** 内容 */
        content: string,
        /** 排序，越小越靠前 */
        sort: number,
        /** 创建时间 */
        created_on: number,
        /** 修改时间 */
        modified_on?: number,
        /** 删除时间 */
        deleted_on?: number,
        /** 是否删除，0为未删除，1为已删除 */
        is_del?: 0 | 1
    }

    /** 评论数据 */
    interface CommentProps {
        id: number,
        post_id: number,
        /** 评论者UID */
        user_id: number,
        /** 评论者用户信息 */
        user: UserInfo,
        /** 评论内容 */
        contents: CommentItemProps[],
        /** 回复列表 */
        replies: ReplyProps[],
        /** 评论者IP地址 */
        ip?: string,
        /** 评论者城市地址 */
        ip_loc: string,
        /** 创建时间 */
        created_on: number,
        /** 修改时间 */
        modified_on?: number,
        /** 删除时间 */
        deleted_on?: number,
        /** 是否删除，0为未删除，1为已删除 */
        is_del?: 0 | 1
    }

    interface CommentComponentProps extends CommentProps {
        /** 文字评论列表 */
        texts: CommentItemProps[],
        /** 图片评论列表 */
        imgs: CommentItemProps[]
    }

    /** 回复内容 */
    interface ReplyProps {
        /** 内容ID */
        id: number,
        /** 评论ID */
        comment_id: number,
        /** 回复人ID */
        user_id: number,
        /** 回复人用户数据 */
        user: UserInfo,
        /** 艾特人ID */
        at_user_id: number,
        /** 艾特人用户数据 */
        at_user: UserInfo,
        /** 内容 */
        content: string,
        /** 回复人IP地址 */
        ip?: string,
        /** 回复人城市地址 */
        ip_loc: string,
        /** 创建时间 */
        created_on: number,
        /** 修改时间 */
        modified_on?: number,
        /** 删除时间 */
        deleted_on?: number,
        /** 是否删除，0为未删除，1为已删除 */
        is_del?: 0 | 1
    }

    /** 帖子内容 */
    interface PostItemProps {
        /** 内容ID */
        id: number,
        /** 类型：1为标题，2为文字段落，3为图片地址，4为视频地址，5为语音地址，6为链接地址，7为附件资源，8为收费资源 */
        type: number,
        /** POST ID */
        post_id: number,
        /** 内容 */
        content: string,
        /** 排序，越小越靠前 */
        sort: number,
        /** 用户UID */
        user_id?: number,
        /** 创建时间 */
        created_on: number,
        /** 修改时间 */
        modified_on?: number,
        /** 删除时间 */
        deleted_on?: number,
        /** 是否删除，0为未删除，1为已删除 */
        is_del?: 0 | 1,
    }

    /** 帖子 */
    interface PostProps {
        id: number,
        /** 发帖人UID */
        user_id: number,
        /** 发帖人用户数据 */
        user: UserInfo,
        /** 附件价格（分） */
        attachment_price: number,
        /** 发帖时IP地址 */
        ip?: string,
        /** 发帖时城市地址 */
        ip_loc: string,
        /** 最新回复时间 */
        latest_replied_on: number,
        /** 创建时间 */
        created_on: number,
        /** 修改时间 */
        modified_on?: number,
        /** 删除时间 */
        deleted_on?: number,
        /** 点赞数 */
        upvote_count: number,
        /** 评论数 */
        comment_count: number,
        /** 收藏数 */
        collection_count: number,
        /** 内容列表 */
        contents: PostItemProps[],
        /** 标签列表 */
        tags: {[key: string]: number} | string,
        /** 是否锁定 */
        is_lock: number,
        /** 是否置顶 */
        is_top: number,
        /** 是否精华 */
        is_essence: number,
        /** 是否删除：0为未删除，1为已删除 */
        is_del?: 0 | 1
    }

    /** 组件用帖子 */
    interface PostComponentProps extends PostProps {
        /** 文字段落列表 */
        texts: PostItemProps[],
        /** 图片列表 */
        imgs: PostItemProps[],
        /** 视频列表 */
        videos: PostItemProps[],
        /** 链接列表 */
        links: PostItemProps[],
        /** 附件列表 */
        attachments: PostItemProps[],
        /** 收费附件列表 */
        charge_attachments: PostItemProps[]
    }

    interface MessageProps {
        id: number,
        /** 类型：1为动态，2为评论，3为回复，4为私信，99为系统通知 */
        type: 1 | 2 | 3 | 4 | 99,
        /** 摘要说明 */
        brief: string,
        /** 详细内容 */
        content: string,
        /** 是否已读：0为未读，1为已读 */
        is_read: 0 | 1,
        /** 发送人UID */
        sender_user_id: number,
        /** 发送人用户数据 */
        sender_user: UserInfo,
        /** 接收方UID */
        receiver_user_id: number,
        /** 帖子ID */
        post_id: number,
        /** 帖子内容 */
        post: PostProps,
        /** 评论ID */
        comment_id: number,
        /** 评论内容 */
        comment: CommentProps,
        /** 回复ID */
        reply_id: number,
        /** 回复内容 */
        replay: ReplyProps,
        /** 创建时间 */
        created_on: number,
        /** 修改时间 */
        modified_on?: number,
        /** 删除时间 */
        deleted_on?: number,
        /** 是否删除：0为未删除，1为已删除 */
        is_del?: 0 | 1
    }

    interface AttachmentProps {
        id: number,
        /** 类别：1为图片，2为视频，3为其他附件 */
        type: 1 | 2 | 3,
        /** 发布者用户UID */
        user_id: number,
        /** 发布者用户数据 */
        user: UserInfo,
        /** 文件大小 */
        file_size: number,
        /** 图片宽度 */
        img_width?: number,
        /** 图片高度 */
        img_height?: number,
        /** 内容 */
        content: string,
        /** 创建时间 */
        created_on: number,
        /** 修改时间 */
        modified_on?: number,
        /** 删除时间 */
        deleted_on?: number,
        /** 是否删除：0为未删除，1为已删除 */
        is_del?: 0 | 1
    }

    interface TagProps {
        id: number,
        /** 创建者UID */
        user_id: number,
        /** 创建者用户数据 */
        user: UserInfo,
        /** 标签名 */
        tag: string,
        /** 引用数 */
        quote_num: number,
        /** 创建时间 */
        created_on: number,
        /** 修改时间 */
        modified_on?: number,
        /** 删除时间 */
        deleted_on?: number,
        /** 是否删除：0为未删除，1为已删除 */
        is_del?: 0 | 1
    }

    interface PagerProps {
        /** 当前页码 */
        page: number,
        /** 每页条数 */
        page_size: number,
        /** 总条数 */
        total_rows: number
    }

    interface BillProps {
        id: number,
        reason: string,
        change_amount: number,
        created_on: number
    }

}