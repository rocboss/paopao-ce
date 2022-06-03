import axios, { AxiosRequestConfig, AxiosRequestHeaders, Method } from 'axios';

const service = axios.create({
    baseURL: import.meta.env.VITE_HOST,
    timeout: 30000,
});

service.interceptors.request.use(
    config => {
        // 鉴权Header
        if (localStorage.getItem('PAOPAO_TOKEN')) {
            (config.headers as any)['Authorization'] = 'Bearer ' + localStorage.getItem('PAOPAO_TOKEN');
        }

        return config;
    },
    error => {
        return Promise.reject(error);
    }
);

service.interceptors.response.use(
    response => {
        const { data = {}, code = 0 } = response?.data || {};
        if (+code === 0) {
            return data || {};
        } else {
            Promise.reject(response?.data || {});
        }
    },
    (error = {}) => {
        const { response = {} } = error || {};
        // 重定向
        if (+response?.status === 401) {
            localStorage.removeItem('PAOPAO_TOKEN');

            if (response?.data.code !== 10005) {
                window.$message.warning(response?.data.msg || '鉴权失败');
            } else {
                // 打开登录弹窗
                window.$store.commit('triggerAuth', true);
            }
        } else {
            window.$message.error(response?.data?.msg || '请求失败');
        }
        return Promise.reject(response?.data || {});
    }
);

export default service;

export function request<T, R>(config: AxiosRequestConfig<T>): Promise<R> {
    return service(config) as unknown as Promise<R>;
}
