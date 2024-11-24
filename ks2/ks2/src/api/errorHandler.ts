import { StatusCodeOperations } from '@/tool/StatusCode';
import { ElMessage } from 'element-plus';
import { Router } from 'vue-router';

export const handleError = (error:any,router?:Router) => {
  if (error.response) {
    const statusCode = error.response.status; // 获取错误状态码
    StatusCodeOperations(statusCode,router)
    ElMessage({
      message: error.response.data.error || 'An error occurred',
      type: 'error',
    });
  } else {
    console.error('Error:', error.message);
    alert(`Error: ${error.message}`);
  }
};