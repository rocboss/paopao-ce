declare namespace Api {

    namespace Suggest {

        interface Api {
            get: {
                /** 获取推荐用户 */
                users: (params: NetParams.GetSuggestUsers) => Promise<NetReq.UserGetSuggestUsers>;
                /** 获取推荐用户 */
                tags: (params: NetParams.GetSuggestTags) => Promise<NetReq.UserGetSuggestTags>;
            }
        }

        namespace NetParams {
            interface GetSuggestUsers {
                k: string
            }

            interface GetSuggestTags extends GetSuggestUsers {}
        }

        namespace NetReq {
            interface UserGetSuggestUsers {
                suggest: string[];
            }

            interface UserGetSuggestTags {
                suggest: string[];
            }
        }

    }

}