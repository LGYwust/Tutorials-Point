import axios from 'axios';

// 创建 Axios 实例
const axiosClient = axios.create({
  baseURL: 'http://127.0.0.1:8080', // 设置基础URL
  headers: {
    'Content-Type': 'application/json', // 默认内容类型
    Accept: 'application/json',
  },
  timeout: 10000, // 设置请求超时时间
});

// 添加请求拦截器
axiosClient.interceptors.request.use(
  (config) => {
    // 从 localStorage 或其他存储中获取 Token
    const token = localStorage.getItem('token');
    // 如果存在 Token，将其添加到请求头中
    if (token) {
      config.headers['Token'] = token;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// 添加响应拦截器
axiosClient.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    return Promise.reject(error);
  }
);

export default axiosClient;