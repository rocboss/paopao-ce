declare namespace Api {

    namespace Admin {

        interface Api {
            post: {
                user: {
                    /** 管理·用户禁言/解禁 */
                    status: (params: NetParams.UserStatusReq) => Promise<NetReq.UserChangeStatus>;
                }
            },
            get: {
                site: {
                    /** 获取站点状态信息 */
                    status: () => Promise<NetReq.SiteInfoResp>;
                }
            }
        }

        namespace NetParams {

            interface UserStatusReq {
                id: number;
                status: number;
            }

        }

        namespace NetReq {
            interface UserChangeStatus {}

            interface SiteInfoResp {
                register_user_count: number;
                online_user_count: number;
                history_max_online: number;
                server_up_time: number;
            }
        }

    }

}