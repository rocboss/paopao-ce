import request from '@/utils/request';

/**
 * 获取动态列表
 * @param {Object} params
 * @returns Promise
 */
export const getPosts = params => {
    return request({
        method: 'get',
        url: '/posts',
        params
    });
};

/**
 * 获取标签列表
 * @param {Object} params
 * @returns Promise
 */
export const getTags = params => {
    return request({
        method: 'get',
        url: '/tags',
        params
    });
};

/**
 * 获取动态详情
 * @param {Object} params
 * @returns Promise
 */
export const getPost = params => {
    return request({
        method: 'get',
        url: '/post',
        params
    });
};

/**
 * 获取动态点赞状态
 * @param {Object} params
 * @returns Promise
 */
export const getPostStar = params => {
    return request({
        method: 'get',
        url: '/post/star',
        params
    });
};

/**
 * 动态点赞
 * @param {Object} data
 * @returns Promise
 */
export const postStar = data => {
    return request({
        method: 'post',
        url: '/post/star',
        data
    });
};

/**
 * 获取动态收藏状态
 * @param {Object} params
 * @returns Promise
 */
export const getPostCollection = params => {
    return request({
        method: 'get',
        url: '/post/collection',
        params
    });
};

/**
 * 动态收藏
 * @param {Object} data
 * @returns Promise
 */
export const postCollection = data => {
    return request({
        method: 'post',
        url: '/post/collection',
        data
    });
};

/**
 * 获取动态评论列表
 * @param {Object} params
 * @returns Promise
 */
export const getPostComments = params => {
    return request({
        method: 'get',
        url: '/post/comments',
        params
    });
};

/**
 * 发布动态
 * @param {Object} data
 *  - @param {array} contents 内容
 *  - @param {array} users at用户
 *  - @param {array} tags 话题
 * @returns Promise
 */
export const createPost = data => {
    return request({
        method: 'post',
        url: '/post',
        data
    });
};

/**
 * 删除动态
 * @param {Object} data
 *  - @param {number} id 
 * @returns Promise
 */
export const deletePost = data => {
    return request({
        method: 'delete',
        url: '/post',
        data
    });
};

/**
 * 锁定/解锁动态
 * @param {Object} data
 *  - @param {number} id 
 * @returns Promise
 */
export const lockPost = data => {
    return request({
        method: 'post',
        url: '/post/lock',
        data
    });
};

/**
 * 发布动态评论
 * @param {Object} data
 *  - @param {array} contents 内容
 *  - @param {array} users at用户
 * @returns Promise
 */
export const createComment = data => {
    return request({
        method: 'post',
        url: '/post/comment',
        data
    });
};

/**
 * 删除评论
 * @param {Object} data
 *  - @param {number} id 
 * @returns Promise
 */
export const deleteComment = data => {
    return request({
        method: 'delete',
        url: '/post/comment',
        data
    });
};

/**
 * 发布评论回复
 * @param {Object} data
 *  - @param {string} content 内容
 *  - @param {number} comment_id 评论ID
 *  - @param {number} at_user_id at用户ID
 * @returns Promise
 */
export const createCommentReply = data => {
    return request({
        method: 'post',
        url: '/post/comment/reply',
        data
    });
};

/**
 * 删除评论回复
 * @param {Object} data
 *  - @param {number} id 评论ID
 * @returns Promise
 */
export const deleteCommentReply = data => {
    return request({
        method: 'delete',
        url: '/post/comment/reply',
        data
    });
};