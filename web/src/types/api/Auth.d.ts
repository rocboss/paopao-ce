declare namespace Api {

    namespace Auth {

        interface Api {
            post: {
                /** 用户登录 */
                login: (params: NetParams.AuthUserLogin) => Promise<NetReq.AuthUserLogin>;
                /** 用户注册 */
                register: (params: NetParams.AuthUserRegister) => Promise<NetReq.AuthUserRegister>;
            };
        }
        
        namespace NetParams {
            interface AuthUserLogin {
                /** 用户名 */
                username: string;
                /** 密码 */
                password: string;
            }

            interface AuthUserRegister extends AuthUserLogin {

            }
        }

        namespace NetReq {
            interface AuthUserLogin {
                token: string;
            }

            interface AuthUserRegister {
                /** 用户UID */
                id: number;
                /** 用户名 */
                username: string;
            }
        }

    }

}