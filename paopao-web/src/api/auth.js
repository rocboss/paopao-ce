import request from '@/utils/request';


/**
 * 用户登录
 * @param {Object} params
 *  - @param {string} username
 *  - @param {string} password
 * @returns Promise
 */
export const userLogin = (params = {}) => {
    return request({
        method: 'post',
        url: '/auth/login',
        data: params,
    });
};

/**
 * 注册用户
 * @param {Object} params
 *  - @param {string} username
 *  - @param {string} password
 * @returns Promise
 */
export const userRegister = (params = {}) => {
    return request({
        method: 'post',
        url: '/auth/register',
        data: params,
    });
};

/**
 * 用户信息
 * @param {Object} params
 * @returns Promise
 */
export const userInfo = (token = '') => {
    return request({
        method: 'get',
        url: '/user/info',
        headers: {
            Authorization: `Bearer ${token}`,
        },
    });
};

/**
 * 修改用户密码
 * @param {Object} params
 *  - @param {string} password 新密码
 *  - @param {string} old_password 旧密码
 * @returns Promise
 */
export const updateUserPassword = data => {
    return request({
        method: 'post',
        url: '/api/user/password',
        data,
    });
};
