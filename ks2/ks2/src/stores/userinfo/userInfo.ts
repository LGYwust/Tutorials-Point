import { defineStore } from "pinia";
import { reactive } from "vue";

export const userInfoStore = defineStore("userInfo", {
  state: () => {
    return  reactive({
      avatarpath: "",
      nickname: "",
      email: "",
      phone: "",
      old: "",
      account:"",
      roleid:0
    });
  },
  actions: {
    setnickname(value: string) {
      this.nickname = value;
    },
    setemail(value: string) {
      this.email = value;
    },
    setphone(value: string) {
      this.phone = value;
    },
    setold(value: string) {
      this.old = value;
    },
    setavatarpath(value: string) {
      this.avatarpath = value;
    },
    setaccount(value: string){
      this.account = value
    },
    setroleid(value:number){
      this.roleid = value
    }
  },
  //开启数据持久化缓存 默认`sessionStorage` 储存
  persist: {
    enabled: true,
    strategies: [
      {
        paths: ["avatarpath", "nickname","email","phone","old","account","roleid"], //存储属性
      },
    ],
  },
});
