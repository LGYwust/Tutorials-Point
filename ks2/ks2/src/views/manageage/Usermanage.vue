<template>
  <el-card class="user-management">
    <div class="header">
      <!-- 搜索框 -->
      <el-input
        v-model="searchQuery"
        placeholder="搜索账号"
        clearable
        class="search-input"
        @input="handleSearch"
      />
    </div>

    <!-- 用户列表 -->
    <el-table :data="filteredUsers" style="width: 100%">
      <el-table-column prop="account" label="账号" />
      <!-- 动态显示权限 -->
      <el-table-column label="权限">
        <template #default="scope">
          <span>{{
            scope.row.roleid == "0"
              ? "普通用户"
              : scope.row.roleid == "1"
              ? "管理员"
              : "超级管理员"
          }}</span>
        </template>
      </el-table-column>

      <el-table-column label="操作" width="180">
        <template #default="scope">
          <el-button @click="editUser(scope.row)" type="primary" size="small"
            >编辑</el-button
          >
          <el-button @click="deleteUser(scope.row)" type="danger" size="small"
            >删除</el-button
          >
        </template>
      </el-table-column>
    </el-table>

    <!-- 添加/编辑用户对话框 -->
    <el-dialog :title="dialogTitle" v-model="dialogVisible">
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
  </el-card>
</template>

<script lang="ts" setup>
import { reactive, ref, computed, onMounted } from "vue";
import { ElMessageBox, ElMessage } from "element-plus";
import {
  handleError,
  ApigetAllUser,
  ApieditUser,
  ApideltUser,
} from "@/api/enter";
import { userInfoStore } from "@/stores/userinfo/userInfo";
const userStore = userInfoStore();
// 初始化用户数据
const users = ref([
  {
    Account: "",
    Email: "",
    Phone: "",
    RoleID: "",
  },
]);

const searchQuery = ref("");
const dialogVisible = ref(false);
const dialogTitle = ref("添加用户");
const isEdit = ref(false);
const userForm = reactive({
  Account: "",
  Email: "",
  Phone: "",
  RoleID: "",
});

// 计算过滤后的用户列表
const filteredUsers = computed(() => {
  if (!searchQuery.value) {
    return users.value;
  }
  return users.value.filter((user) => user.Account.includes(searchQuery.value));
});

// 打开编辑用户对话框
const editUser = (user: any) => {
  dialogTitle.value = "编辑用户";
  userForm.Account = user.account;
  userForm.Email = user.email;
  userForm.Phone = user.phone;
  userForm.RoleID =
    user.roleid == 0 ? "普通用户" : user.roleid == 1 ? "管理员" : "超级管理员";
  isEdit.value = true;
  dialogVisible.value = true;
};

// 保存用户
const saveUser = () => {
  if (userForm.Account == userStore.account) {
    ElMessage.error("不能修改自己的权限");
    return;
  }

  if (isEdit.value) {
    ApieditUser(userForm)
      .then(() => {
        init();
        ElMessage.success("用户权限修改成功");
      })
      .catch((error) => {
        handleError(error);
      });
  }
  dialogVisible.value = false;
};

// 删除用户
const deleteUser = (user: any) => {
  if (user.account != userStore.account) {
    ElMessageBox.confirm(`确定要删除账号${user.account} 吗？`, "提示", {
      type: "warning",
    })
      .then(() => {
        ApideltUser(user.account)
          .then(() => {
            ElMessage.success("用户删除成功");
            init();
          })
          .catch((error) => {
            handleError(error);
          });
      })
      .catch(() => {
        ElMessage.info("取消删除");
      });
  } else {
    ElMessage.error("不能删除自己");
  }
};
// 搜索用户
const handleSearch = () => {
  // 搜索逻辑通过 computed 实现
};

function init() {
  ApigetAllUser()
    .then((response) => {
      users.value = response.data.data;
    })
    .catch((error) => {
      handleError(error);
    });
}
onMounted(() => {
  init();
});
</script>

<style scoped>
.header {
  margin-bottom: 10px;
  display: flex;
  justify-content: space-between;
}

.search-input {
  width: 250px;
}
</style>
