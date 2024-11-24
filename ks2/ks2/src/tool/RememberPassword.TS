// 设置存储用户名和密码的函数
export function rememberPassword(username: string, password: string) {
    localStorage.setItem('username', username);
    localStorage.setItem('password', password);
    localStorage.setItem('remember', 'true');
  }
  
  // 删除存储的用户名和密码
export function forgetPassword() {
    localStorage.removeItem('username');
    localStorage.removeItem('password');
    localStorage.removeItem('remember');
  }