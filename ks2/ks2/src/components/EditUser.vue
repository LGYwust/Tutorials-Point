<template>
        <!-- 添加/编辑用户对话框 -->
        <el-dialog title="编辑账号" v-model="dialogVisible">
      <el-form :model="userForm" ref="userFormRef" label-width="80px">
        <el-form-item label="账号" required>
          <el-input v-model="userForm.Account" disabled />
        </el-form-item>
        <el-form-item label="邮箱" required>
          <el-input v-model="userForm.Email" disabled />
        </el-form-item>
        <el-form-item label="手机号" required>
          <el-input v-model="userForm.Phone" disabled />
        </el-form-item>
        <el-form-item label="权限" required>
          <el-select v-model="userForm.RoleID" placeholder="请选择权限">
            <el-option label="管理员" value="1"></el-option>
            <el-option label="普通用户" value="0"></el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="saveUser">保存</el-button>
      </template>
    </el-dialog>
</template>
<script lang="ts" setup>
import { ApieditUser,handleError } from '@/api/enter';
import { ElMessage } from 'element-plus';
import { reactive, ref } from 'vue';
import { userInfoStore } from "@/stores/userinfo/userInfo";
const userStore = userInfoStore();
const dialogVisible = ref(false);
const userForm = reactive({
  Account: "",
  Email: "",
  Phone: "",
  RoleID: "",
});

// 保存用户
const saveUser = () => {
  if (userForm.Account == userStore.account) {
    ElMessage.error("不能修改自己的权限");
    return;
  }
    ApieditUser(userForm)
      .then(() => {
        ElMessage.success("用户权限修改成功");
      })
      .catch((error) => {
        handleError(error);
      });
  dialogVisible.value = false;
};

const open = ()=> {
    dialogVisible.value = true;
}

defineExpose({ open });
</script>