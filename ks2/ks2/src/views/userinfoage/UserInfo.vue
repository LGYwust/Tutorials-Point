<template>
  <div class="form-container">
    <div
      class="avatar-wrapper"
      @mouseenter="hovering = true"
      @mouseleave="hovering = false"
      @click="goToProfile"
    >
      <el-avatar :src="avatarUrl" :size="180" class="avatar"></el-avatar>
      <div v-if="hovering" class="avatar-hover-overlay">
        <span class="avatar-hover-text">更换头像</span>
      </div>
    </div>
    <el-form
      :model="form"
      label-width="auto"
      style="min-width: 600px"
      size="large"
      label-position="left"
    >
      <el-form-item label="账号">
        <el-input v-model="account" disabled />
      </el-form-item>
      <el-form-item label="昵称">
        <el-input v-model="form.nickname" />
      </el-form-item>
      <el-form-item label="出生日期">
        <el-col :span="11">
          <el-date-picker
            v-model="form.time"
            type="date"
            placeholder="选择日期"
            style="width: 100%"
          />
        </el-col>
      </el-form-item>
      <el-form-item label="性别">
        <el-radio-group v-model="form.sex">
          <el-radio value="man">先生</el-radio>
          <el-radio value="woman">女士</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="个人简介">
        <el-input v-model="form.desc" type="textarea" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">保存</el-button>
        <el-button @click="onReset">重置</el-button>
      </el-form-item>
    </el-form>
  </div>
  <AvatarDialog ref="atDialog"  @close="handleDialogClose"></AvatarDialog>
</template>

<script lang="ts" setup>
import { reactive, ref, onMounted } from "vue";
import { handleError, ApigetAvatar, ApisaveUserinfo,ApiGetUserinfo } from "@/api/enter";
import AvatarDialog from "../../components/AvatarUpload.vue";
import { ElMessage, ElMessageBox } from "element-plus";
import {userInfoStore} from '@/stores/userinfo/userInfo';
const userStore = userInfoStore()
const account = ref(userStore.account);
const form = reactive({
  nickname: "",
  time: "",
  sex: "",
  desc: "",
});

const initialFormState = {
  nickname: "",
  time: "",
  sex: "",
  desc: "",
}; // 保存初始表单状态，账号不重置
const avatarUrl = ref("");
const hovering = ref(false);
const atDialog = ref();
function init() {
  ApigetAvatar()
    .then((response) => {
      console.log(response.data.path)
      avatarUrl.value = response.data.path;
    })
    .catch((error) => {
      handleError(error);
    });
  ApiGetUserinfo()
  .then((response) => {
    console.log(response)
    form.nickname = response.data.data.nickname;
    form.sex =  response.data.data.sex;
    form.time = response.data.data.time;
    form.desc = response.data.data.desc;
    })
    .catch((error) => {
      handleError(error);
    });
}

const handleDialogClose = () => {
  init(); // 触发 init 函数
  location.reload()
};
const goToProfile = () => {
  atDialog.value.open();
};

const onSubmit = () => {
  ElMessageBox.confirm("确认保存？", "提示", {
    confirmButtonText: "确认",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(() => {
      ApisaveUserinfo(form)
        .then(() => {
          ElMessage({
            type: "success",
            message: "保存成功",
          });
        })
        .catch((error) => {
          ElMessage({
            type: "error",
            message: "保存失败",
          });
        });
    })
    .catch(() => {
      ElMessage({
        type: "warning",
        message: "取消保存",
      });
    });
};

const onReset = () => {
  ElMessageBox.confirm("是否重置信息？", "提示", {
    confirmButtonText: "确认",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(() => {
      ElMessage({
        type: "success",
        message: "重置成功",
      });
      Object.assign(form, { ...form, ...initialFormState });
    })
    .catch(() => {
      ElMessage({
        type: "warning",
        message: "取消重置",
      });
    });
};

onMounted(() => {
  init();
});
</script>

<style scoped>
.form-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 20px;
}

.avatar-wrapper {
  position: relative;
  cursor: pointer;
  width: 180px;
  height: 180px;
  margin-bottom: 30px;
}

.avatar {
  transition: opacity 0.3s ease;
}

.avatar-wrapper:hover .avatar {
  opacity: 0.6;
}

.avatar-hover-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5); /* 背景加深 */
  display: flex;
  align-items: center;
  justify-content: center;
  color: white; /* 白色文字 */
  border-radius: 50%;
}

.avatar-hover-text {
  font-size: 16px;
  font-weight: bold;
}
</style>
