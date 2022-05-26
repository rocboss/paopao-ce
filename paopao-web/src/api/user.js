import request from '@/utils/request';

/**
 * 获取验证码
 * @param {Object} params
 * @returns Promise
 */
export const getCaptcha = params => {
    return request({
        method: 'get',
        url: '/captcha',
        params
    });
};

/**
 * 发送短信验证码
 * @param {Object} data
 * @returns Promise
 */
export const sendCaptcha = data => {
    return request({
        method: 'post',
        url: '/captcha',
        data
    });
};

/**
 * 绑定用户手机
 * @param {Object} data
 * @returns Promise
 */
export const bindUserPhone = data => {
    return request({
        method: 'post',
        url: '/user/phone',
        data
    });
};

/**
 * 更改密码
 * @param {Object} data
 * @returns Promise
 */
export const changePassword = data => {
    return request({
        method: 'post',
        url: '/user/password',
        data
    });
};

/**
 * 更改昵称
 * @param {Object} data
 * @returns Promise
 */
export const changeNickname = data => {
    return request({
        method: 'post',
        url: '/user/nickname',
        data
    });
};

/**
 * 更改头像
 * @param {Object} data
 * @returns Promise
 */
export const changeAvatar = data => {
    return request({
        method: 'post',
        url: '/user/avatar',
        data
    });
};

/**
 * 获取未读消息数
 * @param {Object} params
 * @returns Promise
 */
export const getUnreadMsgCount = params => {
    return request({
        method: 'get',
        url: '/user/msgcount/unread',
        params
    });
};

/**
 * 获取消息列表
 * @param {Object} params
 * @returns Promise
 */
export const getMessages = params => {
    return request({
        method: 'get',
        url: '/user/messages',
        params
    });
};

/**
 * 阅读消息
 * @param {Object} data
 * @returns Promise
 */
export const readMessage = data => {
    return request({
        method: 'post',
        url: '/user/message/read',
        data
    });
};

/**
 * 获取收藏列表
 * @param {Object} params
 * @returns Promise
 */
export const getCollections = params => {
    return request({
        method: 'get',
        url: '/user/collections',
        params
    });
};

/**
 * 获取点赞列表
 * @param {Object} params
 * @returns Promise
 */
export const getStars = params => {
    return request({
        method: 'get',
        url: '/user/stars',
        params
    });
};

/**
 * 获取用户基础信息
 * @param {Object} params
 * @returns Promise
 */
export const getUserProfile = params => {
    return request({
        method: 'get',
        url: '/user/profile',
        params
    });
};

/**
 * 获取点赞列表
 * @param {Object} params
 * @returns Promise
 */
export const getUserPosts = params => {
    return request({
        method: 'get',
        url: '/user/posts',
        params
    });
};

/**
 * 获取账单列表
 * @param {Object} params
 * @returns Promise
 */
export const getBills = params => {
    return request({
        method: 'get',
        url: '/user/wallet/bills',
        params
    });
};

/**
 * 发起充值请求
 * @param {Object} data
 * @returns Promise
 */
export const reqRecharge = data => {
    return request({
        method: 'post',
        url: '/user/recharge',
        data
    });
};

/**
 * 获取充值状态
 * @param {Object} params
 * @returns Promise
 */
export const getRecharge = params => {
    return request({
        method: 'get',
        url: '/user/recharge',
        params
    });
};

/**
 * 获取推荐用户
 * @param {Object} params
 * @returns Promise
 */
export const getSuggestUsers = params => {
    return request({
        method: 'get',
        url: '/suggest/users',
        params
    });
};

/**
 * 获取推荐用户
 * @param {Object} params
 * @returns Promise
 */
export const getSuggestTags = params => {
    return request({
        method: 'get',
        url: '/suggest/tags',
        params
    });
};

/**
 * 获取附件
 * @param {Object} params
 * @returns Promise
 */
export const precheckAttachment = params => {
    return request({
        method: 'get',
        url: '/attachment/precheck',
        params
    });
};

/**
 * 获取附件
 * @param {Object} params
 * @returns Promise
 */
export const getAttachment = params => {
    return request({
        method: 'get',
        url: '/attachment',
        params
    });
};
