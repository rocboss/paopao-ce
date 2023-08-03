import { request } from "@/utils/request";

//获取用户的sharekey
// export const getShareKeys = (
//     params: NetParams.UserGetShareKeys
// ): Promise<NetReq.UserGetShareKeys> => {
//     return request({
//         method: "get",
//         url: "/v1/user/keys",
//         params,
//     });
// }
export const getShareKeys = (token: NetParams.AuthUserInfo = ""): Promise<NetReq.UserGetShareKeys> => {
    return request({
        method: "get",
        url: "/v1/user/keys",
        headers: {
            Authorization: `Bearer ${token}`,
        },
    });
}

export const deleteThisShareKey = (params: NetParams.UserDeleteShareKey,token: NetParams.AuthUserInfo = ""): Promise<NetReq.UserDeleteShareKey> => {
    return request({
        method: "post",
        url: "/v1/user/keys",
        headers: {
            Authorization: `Bearer ${token}`,
        },
        params,
    });
}