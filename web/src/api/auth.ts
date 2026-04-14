/**
 * 这两个接口暂时保留着，因为比较特殊，所以暂时不放进createApi里，后续如果有需要再考虑重构
 * 当然，也可以作为另一种API定义方式的示例，供后续参考
 */

import { request } from '@/utils/request';

/** 用户信息 */
export const userInfo = (
	token: NetParams.AuthUserInfo = '',
): Promise<NetReq.AuthUserInfo> => {
	return request({
		method: 'get',
		url: '/v1/user/info',
		headers: {
			Authorization: `Bearer ${token}`,
		},
	});
};

/** 修改用户密码，该接口暂时未使用 */
export const updateUserPassword = (
	data: NetParams.AuthUpdateUserPassword,
): Promise<NetReq.AuthUpdateUserPassword> => {
	return request({
		method: 'post',
		url: '/v1/api/user/password',
		data,
	});
};
