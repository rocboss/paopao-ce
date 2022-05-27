import request from '@/utils/request';

/**
 * 获取动态列表
 * @param {Object} params
 * @returns Promise
 */
export const getPosts = (params: any) => {
    return request({
        method: 'get',
        url: '/posts',
        params
    }) as unknown as Promise<NetReq.PostGetPosts>;
};

/**
 * 获取标签列表
 * @param {Object} params
 * @returns Promise
 */
export const getTags = (params: any) => {
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
export const getPost = (params: any) => {
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
export const getPostStar = (params: any) => {
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
export const postStar = (data: any) => {
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
export const getPostCollection = (params: any) => {
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
export const postCollection = (data: any) => {
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
export const getPostComments = (params: any) => {
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
export const createPost = (data: any) => {
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
export const deletePost = (data: any) => {
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
export const lockPost = (data: any) => {
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
export const createComment = (data: any) => {
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
export const deleteComment = (data: any) => {
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
export const createCommentReply = (data: any) => {
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
export const deleteCommentReply = (data: any) => {
    return request({
        method: 'delete',
        url: '/post/comment/reply',
        data
    });
};