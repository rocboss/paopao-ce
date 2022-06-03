import { request } from '@/utils/request';

/** 获取动态列表 */
export const getPosts = (params: NetParams.PostGetPosts): Promise<NetReq.PostGetPosts> => {
    return request({
        method: 'get',
        url: '/v1/posts',
        params
    });
};

/** 获取标签列表 */
export const getTags = (params: NetParams.PostGetTags): Promise<NetReq.PostGetTags> => {
    return request({
        method: 'get',
        url: '/v1/tags',
        params
    });
};

/** 获取动态详情 */
export const getPost = (params: NetParams.PostGetPost): Promise<NetReq.PostGetPost> => {
    return request({
        method: 'get',
        url: '/v1/post',
        params
    });
};

/** 获取动态点赞状态 */
export const getPostStar = (params: NetParams.PostPostStar): Promise<NetReq.PostGetPostStar> => {
    return request({
        method: 'get',
        url: '/v1/post/star',
        params
    });
};

/** 动态点赞 */
export const postStar = (data: NetParams.PostPostStar): Promise<NetReq.PostPostStar> => {
    return request({
        method: 'post',
        url: '/v1/post/star',
        data
    });
};

/** 获取动态收藏状态 */
export const getPostCollection = (params: NetParams.PostGetPostCollection): Promise<NetReq.PostGetPostCollection> => {
    return request({
        method: 'get',
        url: '/v1/post/collection',
        params
    });
};

/** 动态收藏 */
export const postCollection = (data: NetParams.PostPostCollection): Promise<NetReq.PostPostCollection> => {
    return request({
        method: 'post',
        url: '/v1/post/collection',
        data
    });
};

/** 获取动态评论列表 */
export const getPostComments = (params: NetParams.PostGetPostComments): Promise<NetReq.PostGetPostComments> => {
    return request({
        method: 'get',
        url: '/v1/post/comments',
        params
    });
};

/** 发布动态 */
export const createPost = (data: NetParams.PostCreatePost): Promise<NetReq.PostCreatePost> => {
    return request({
        method: 'post',
        url: '/v1/post',
        data
    });
};

/** 删除动态 */
export const deletePost = (data: NetParams.PostDeletePost): Promise<NetReq.PostDeletePost> => {
    return request({
        method: 'delete',
        url: '/v1/post',
        data
    });
};

/** 锁定/解锁动态 */
export const lockPost = (data: NetParams.PostLockPost): Promise<NetReq.PostLockPost> => {
    return request({
        method: 'post',
        url: '/v1/post/lock',
        data
    });
};

/** 置顶/取消置顶动态 */
export const stickPost = (data: NetParams.PostStickPost): Promise<NetReq.PostStickPost> => {
    return request({
        method: 'post',
        url: '/v1/post/stick',
        data
    });
};

/** 发布动态评论 */
export const createComment = (data: NetParams.PostCreateComment): Promise<NetReq.PostCreateComment> => {
    return request({
        method: 'post',
        url: '/v1/post/comment',
        data
    });
};

/** 删除评论 */
export const deleteComment = (data: NetParams.PostDeleteComment): Promise<NetReq.PostDeleteComment> => {
    return request({
        method: 'delete',
        url: '/v1/post/comment',
        data
    });
};

/** 发布评论回复 */
export const createCommentReply = (data: NetParams.PostCreateCommentReply): Promise<NetReq.PostCreateCommentReply> => {
    return request({
        method: 'post',
        url: '/v1/post/comment/reply',
        data
    });
};

/** 删除评论回复 */
export const deleteCommentReply = (data: NetParams.PostDeleteCommentReply): Promise<NetReq.PostDeleteCommentReply> => {
    return request({
        method: 'delete',
        url: '/v1/post/comment/reply',
        data
    });
};
