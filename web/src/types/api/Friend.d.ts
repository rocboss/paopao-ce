declare namespace Api {

    namespace Friend {

        interface Api {
            post: {
                /** 添加朋友申请 */
                requesting: (data: NetParams.RequestingFriend) => Promise<NetReq.RequestingFriend>;
                /** 添加好友 */
                add: (data: NetParams.AddFriend) => Promise<NetReq.AddFriend>;
                /** 拒绝朋友申请 */
                reject: (data: NetParams.RejectFriend) => Promise<NetReq.RejectFriend>;
                /** 删除好友 */
                delete: (data: NetParams.DeleteFriend) => Promise<NetReq.DeleteFriend>;
            }
        }

        namespace NetParams {

            interface RequestingFriend {
                user_id: number;
                greetings: string;
            }

            interface AddFriend {
                user_id: number;
            }

            interface RejectFriend extends AddFriend {}
            interface DeleteFriend extends AddFriend {}
        }

        namespace NetReq {

            interface RequestingFriend {}

            interface AddFriend {}
            interface RejectFriend {}
            interface DeleteFriend {}
        }

    }

}