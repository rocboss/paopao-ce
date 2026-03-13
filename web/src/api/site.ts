/**
 * 这个接口暂时保留着，暂时不放进createApi里，后续如果有需要再考虑重构
 * 当然，也可以作为另一种API定义方式的示例，供后续参考
 */

import { request } from '@/utils/request';

/** 获取站点概要信息 */
export const getSiteProfile = (): Promise<NetReq.SiteProfile> => {
	return request({
		method: 'get',
		url: '/v1/site/profile',
	});
};
