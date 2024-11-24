import { ElMessageBox } from "element-plus";

export function goto (movie: any) {
    if (movie.url == "") {
      ElMessageBox.confirm("暂无资源请联系管理员", "错误", {
        confirmButtonText: "OK",
        cancelButtonText: "Cancel",
        type: "error",
      });
    } else {
      window.open("http://"+movie.url, '_blank');
    }
  };