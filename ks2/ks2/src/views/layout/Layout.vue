<template>
  <div class="common-layout center-container">
    <el-container>
      <el-header class="header">
        电 影 院 线
        <div class="block">
          <el-dropdown placement="bottom">
            <el-avatar :size="46" :src="circleUrl" />
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item @click="openDrop">更换头像</el-dropdown-item>
                <el-dropdown-item  @click="setActiveButton('个人信息')">个人信息</el-dropdown-item>
                <el-dropdown-item @click="Logout">退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </el-header>
      <AvatarDialog ref="atDialog" @close="handleDialogClose"></AvatarDialog>
      <div class="button-box">
      <div class="button-container">
        <el-button :class="{ 'custom-button': true, 'active-button': activeButton === 'front' }"
          @click="setActiveButton('front')">首页</el-button>
        <el-button :class="{ 'custom-button': true, 'active-button': activeButton === 'allmovies' }"
          @click="setActiveButton('allmovies')">全部电影</el-button>
        <!-- <el-button :class="{ 'custom-button': true, 'active-button': activeButton === '推荐电影' }"
          @click="setActiveButton('推荐电影')">推荐电影</el-button> -->
        <el-button :class="{ 'custom-button': true, 'active-button': activeButton === 'consultation' }"
          @click="setActiveButton('consultation')">电影咨询</el-button>
        <el-button :class="{ 'custom-button': true, 'active-button': activeButton === 'userinfo' }"
          @click="setActiveButton('userinfo')">个人信息</el-button>
           <!-- 当 role == 1 时显示 系统管理 按钮 -->
        <el-button v-if="role === 1||role === 2" :class="{ 'custom-button': true, 'active-button': activeButton === 'manage' }"
      @click="setActiveButton('manage')">系统管理</el-button>
      </div>
      <el-divider />
    </div>
      <loading ref="load"></loading>
      <router-view></router-view>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import { onMounted, provide, ref } from 'vue'
import AvatarDialog from '../../components/AvatarUpload.vue';
import { useRouter } from 'vue-router';
import { handleError,ApigetAvatar } from '@/api/enter';
import {userInfoStore} from '@/stores/userinfo/userInfo'
import { ElMessage, ElMessageBox } from 'element-plus';
import loading from '../../components/loading.vue';
import srouter from '../../router';
const load = ref()
const router = useRouter()
const userStore = userInfoStore()
const circleUrl = ref<string>('')
const activeButton = ref<string>('front')
const atDialog = ref()
const role = ref(userStore.roleid);

const setActiveButton = (buttonName: string) => {
  router.push(`/layout/${buttonName}`);
  activeButton.value = buttonName
  localStorage.setItem('activeButton', buttonName);  // 保存到LocalStorage
}
provide('setActiveButton', setActiveButton);
//获取头像图片
function getavatar(){
  ApigetAvatar()
    .then((response) => {
      circleUrl.value = response.data.path
      userStore.setavatarpath(response.data.path)
    })
    .catch((error) => {
      handleError(error,router)
    });
}
const handleDialogClose = () => {
  getavatar()
};
//打开更换头像窗口
function openDrop() {
  atDialog.value.open()
}
//退出登录
function Logout(){
  ElMessageBox.confirm("是否退出登录？", "提示", {
    confirmButtonText: "确认",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(() => {
      localStorage.removeItem('token');
      ElMessage({
        type: "success",
        message: "退出成功",
      });
      location.reload()
    })
    .catch(() => {
      ElMessage({
        type: "warning",
        message: "取消退出",
      });
    });
}
onMounted(() => {
  load.value.out()
  // 注册全局路由守卫
  srouter.beforeEach((to, from, next) => {
      load.value.enter(next);
  });
  getavatar()
  const savedButton = localStorage.getItem('activeButton');
  if (savedButton) {
    activeButton.value = savedButton;
  }
})

</script>

<style scoped>
.center-container {
  display: flex;
  justify-content: center;
}

.header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 6vh;
  min-height: 50px;
  background: linear-gradient(to right, #001f3f, #0074D9);
  color: white;
  /* 设置字体颜色为白色 */
  font-size: 1.8em;
  /* 设置字体大小稍微放大，根据需要调整具体数值 */
  font-weight: bold;
  /* 设置字体为黑体加粗 */
  letter-spacing: 2px;
  /* 增加字体间的间距，根据需要调整具体数值 */
}

.block {
  margin-right: 30px;
}
.button-box{
  position: sticky;
  top:0px;
  z-index: 9999;
  background-color: rgb(255, 255, 255);
}

.button-container {
  display: flex;
  justify-content: center;
  /* 水平居中 */
  align-items: center;
  /* 垂直居中 */
  gap: 20px;
  /* 按钮之间的间距 */
  margin-top: 20px;
  /* 按钮容器与标题之间的间距 */
}

.custom-button {
  background-color: white !important;
  color: black !important;
  font-size: 18px;
  width: 120px;
  height: 45px;
  border: 1px solid rgb(255, 255, 255) !important;
}

.custom-button:hover {
  background-color: rgba(7, 100, 223, 0.1) !important;
}

.active-button {
  background-color: rgb(7, 100, 223) !important;
  color: white !important;
}

.active-button:hover {
  background-color: rgb(52, 149, 240) !important;
  color: black !important;
}
</style>
