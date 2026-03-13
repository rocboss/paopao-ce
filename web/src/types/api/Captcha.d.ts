declare namespace Api {

    namespace Captcha {

        interface Api {
            get: {
                /** 获取验证码 */
                _self: (params: NetParams.UserGetCaptcha) => Promise<NetReq.UserGetCaptcha>;
            },
            post: {
                /** 发送验证码 */
                _self: (data: NetParams.UserPostCaptcha) => Promise<void>;
            }
        }

        namespace NetParams {

            interface UserGetCaptcha {}

            interface UserPostCaptcha {
                phone: string;
                img_captcha: string;
                img_captcha_id: string;
            }

        }

        namespace NetReq {
              interface UserGetCaptcha {
                id: string;
                /** 头像图片 base64 */
                b64s: string;
            }
        }

    }

}