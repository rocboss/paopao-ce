import request from '@/utils/request';

/**
 * 获取动态列表
 * @param {Object} params
 * @returns Promise
 */
export const getPosts = (params: NetParams.PostGetPosts) => {
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
export const getTags = (params: NetParams.PostGetTags) => {
    return request({
        method: 'get',
        url: '/tags',
        params
    }) as unknown as Promise<NetReq.PostGetTags>;
};

/**
 * 获取动态详情
 * @param {Object} params
 * @returns Promise
 */
export const getPost = (params: NetParams.PostGetPost) => {
    return request({
        method: 'get',
        url: '/post',
        params
    }) as unknown as Promise<NetReq.PostGetPost>;
};

/**
 * 获取动态点赞状态
 * @param {Object} params
 * @returns Promise
 */
export const getPostStar = (params: NetParams.PostPostStar) => {
    return request({
        method: 'get',
        url: '/post/star',
        params
    }) as unknown as Promise<NetReq.PostGetPostStar>;
};

/**
 * 动态点赞
 * @param {Object} data
 * @returns Promise
 */
export const postStar = (data: NetParams.PostPostStar) => {
    return request({
        method: 'post',
        url: '/post/star',
        data
    }) as unknown as Promise<NetReq.PostPostStar>;
};

/**
 * 获取动态收藏状态
 * @param {Object} params
 * @returns Promise
 */
export const getPostCollection = (params: NetParams.PostGetPostCollection) => {
    return request({
        method: 'get',
        url: '/post/collection',
        params
    }) as unknown as Promise<NetReq.PostGetPostCollection>;
};

/**
 * 动态收藏
 * @param {Object} data
 * @returns Promise
 */
export const postCollection = (data: NetParams.PostPostCollection) => {
    return request({
        method: 'post',
        url: '/post/collection',
        data
    }) as unknown as Promise<NetReq.PostPostCollection>;
};

/**
 * 获取动态评论列表
 * @param {Object} params
 * @returns Promise
 */
export const getPostComments = (params: NetParams.PostGetPostComments) => {
    return request({
        method: 'get',
        url: '/post/comments',
        params
    }) as unknown as Promise<NetReq.PostGetPostComments>;
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
export const lockPost = (data: NetParams.PostLockPost) => {
    return request({
        method: 'post',
        url: '/post/lock',
        data
    }) as unknown as Promise<NetReq.PostLockPost>;
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