import { Router } from "vue-router";

export function StatusCodeOperations(statusCode: number, router?: Router) {
  if (!!router) {
    switch (statusCode) {
      //用户未登录
      case 401:
        router.push({ name: "Login" });
        break;
      //服务器异常
      case 402:
        router.push({ name: "Login" });
        break;
    }
  }
}
