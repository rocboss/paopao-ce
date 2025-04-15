import { request } from '@/utils/request';

/** 获取站点概要信息 */
export const getSiteProfile = (): Promise<NetReq.SiteProfile> => {
  return request({
    method: 'get',
    url: '/v1/site/profile',
  });
};
