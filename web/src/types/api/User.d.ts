declare namespace Api {

    namespace User {

        interface Api {
            get: {
                /** 获取用户基础信息 */
                profile: (params: NetParams.UserGetUserProfile) => Promise<NetReq.UserGetUserProfile>;
                /** 获取关注列表 */
                follows: (params: NetParams.GetUserFollows) => Promise<NetReq.GetContacts>;
                /** 获取粉丝列表 */
                followings: (params: NetParams.GetUserFollowings) => Promise<NetReq.GetContacts>;
                msgcount: {
                    /** 获取未读消息数 */
                    unread: (params: NetParams.UserGetUnreadMsgCount) => Promise<NetReq.UserGetUnreadMsgCount>;
                };
                /** 获取消息列表 */
                messages: (params: NetParams.UserGetMessages) => Promise<NetReq.UserGetMessages>;
                /** 获取收藏列表 */
                collections: (params: NetParams.UserGetCollections) => Promise<NetReq.UserGetCollections>;
                /** 获取用户帖子列表 */
                posts: (params: NetParams.UserGetUserPosts) => Promise<NetReq.UserGetUserPosts>;
                wallet: {
                    /** 获取账单列表 */
                    bills: (params: NetParams.UserGetBills) => Promise<NetReq.UserGetBills>;
                }
                /** 获取充值状态 */
                recharge: (params: NetParams.UserGetRecharge) => Promise<NetReq.UserGetRecharge>;
            },
            post: {
                /** 发送用户私信 */
                whisper: (data: NetParams.UserWhisper) => Promise<NetParams.UserWhisper>;
                /** 关注用户 */
                follow: (data: NetParams.FollowUserReq) => Promise<NetReq.FollowUserResp>;
                /** 取消关注用户 */
                unfollow: (data: NetParams.UnfollowUserReq) => Promise<NetReq.UnfollowUserResp>;
                /** 获取联系人列表 */
                contacts: (params: NetParams.GetContacts) => Promise<NetReq.GetContacts>;
                /** 绑定用户手机 */
                phone: (data: NetParams.UserBindUserPhone) => Promise<NetParams.UserBindUserPhone>;
                /** 激活码激活 */
                activate: (data: NetParams.UserActivation) => Promise<NetParams.UserActivation>;
                /** 更改密码 */
                password: (data: NetParams.UserChangePassword) => Promise<NetReq.UserChangePassword>;
                /** 更改昵称 */
                nickname: (data: NetParams.UserChangeNickname) => Promise<NetReq.UserChangeNickname>;
                /** 更改头像 */
                avatar: (data: NetParams.UserChangeAvatar) => Promise<void>;
                /** 发起充值请求 */
                recharge: (data: NetParams.UserReqRecharge) => Promise<NetReq.UserReqRecharge>;
            },
            /** 信息相关 */
            message: {
                post: {
                    /** 阅读消息 */
                    read: (data: NetParams.ReadMessageReq) => Promise<NetReq.ReadMessageResp>;
                    /** 阅读所有消息 */
                    readall: () => Promise<NetReq.ReadAllMessageResp>;
                }
            }
        }

        namespace NetParams {

            interface UserGetUserProfile {
                username: string;
            }

            interface UserWhisper {
                user_id: number;
                content: string;
            }

            interface FollowUserReq {
                user_id: number;
            }

            interface UnfollowUserReq extends FollowUserReq {}

            interface GetUserFollows {
                username: string;
                page: number;
                page_size: number;
            }

            interface GetUserFollowings extends GetUserFollows {}

            interface GetContacts {
                page: number;
                page_size: number;
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

            interface UserChangeAvatar {
                /** 头像图片 */
                avatar: any;
            }

            interface UserGetUnreadMsgCount {}

            interface UserGetMessages {
                style: 'all' | 'system' | 'whisper' | 'requesting' | 'unread';
                page: number;
                page_size: number;
            }

            interface ReadMessageReq {
                id: number;
            }

            interface UserGetCollections {
                page: number;
                page_size: number;
            }

            interface UserGetUserPosts {
                username: string;
                style: string;
                page: number;
                page_size: number;
            }

            interface UserGetBills {
                page: number;
                page_size: number;
            }

            interface UserReqRecharge {
                amount: number;
            }

            interface UserGetRecharge {
                id: number;
            }

        }

        namespace NetReq {

            type UserGetUserProfile = Item.UserInfo;

            interface FollowUserResp {}

            interface UnfollowUserResp {}

            interface GetContacts {
                /** 评论列表 */
                list: Item.ContactItemProps[];
                /** 页码信息 */
                pager: Item.PagerProps;
            }

            interface UserChangePassword {}
            interface UserChangeNickname {}

            interface UserGetUnreadMsgCount {
                count: number;
            }

            interface UserGetMessages {
                /** 消息列表 */
                list: Item.MessageProps[];
                /** 页码信息 */
                pager: Item.PagerProps;
            }

            interface ReadMessageResp {}
            interface ReadAllMessageResp {}

            interface UserGetCollections {
                /** 帖子列表 */
                list: Item.PostProps[];
                /** 页码信息 */
                pager: Item.PagerProps;
            }

            interface UserGetUserPosts {
                /** 帖子列表 */
                list: Item.PostProps[];
                /** 页码信息 */
                pager: Item.PagerProps;
            }

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
        }

    }

}