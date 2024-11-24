<template>
  <div class="login-container">
    <el-card class="login-card">
      <h2 class="login-title">登录</h2>
      <el-form
        :model="loginForm"
        :rules="rules"
        label-width="80px"
        class="login-form"
        hide-required-asterisk
      >
        <el-form-item label="身份" prop="RoleID" class="item">
          <el-select
            v-model="loginForm.RoleID"
            placeholder="Select"
            size="large"
          >
            <el-option
              v-for="item in options"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="账号" prop="Account" class="item">
          <el-input
            v-model="loginForm.Account"
            placeholder="请输入账号"
            class="input-height"
          >
            <template #prefix>
              <el-icon>
                <User />
              </el-icon>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item label="密码" prop="Password" class="item">
          <el-input
            v-model="loginForm.Password"
            type="password"
            placeholder="请输入密码"
            class="input-height"
          >
            <template #prefix>
              <el-icon>
                <Lock />
              </el-icon>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item label="验证码" prop="captcha" class="item">
          <el-row>
            <el-col :span="14">
              <el-input
                v-model="loginForm.captcha"
                placeholder="请输入验证码"
                class="input-height"
              >
                <template #prefix>
                  <el-icon>
                    <CircleCheck />
                  </el-icon>
                </template>
              </el-input>
            </el-col>
            <el-col :span="10">
              <canvas
                ref="captchaCanvas"
                width="100"
                height="40"
                @click="getcanvas"
                class="captcha-canvas"
              ></canvas>
            </el-col>
          </el-row>
        </el-form-item>
        <div class="checkbox-link-container">
          <el-checkbox v-model="remember">记住密码</el-checkbox>
          <el-link class="forgot-password" @click="onForgotPassword"
            >忘记密码?</el-link
          >
        </div>
        <div>
          <el-button type="primary" @click="onSubmit" class="button"
            >登录</el-button
          >
        </div>
        <div>
          <el-button
            type="danger"
            @click="onRegister"
            class="button register-button"
            >注册</el-button
          >
        </div>
      </el-form>
    </el-card>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted, reactive, type Ref } from "vue";
import { User, Lock, CircleCheck } from "@element-plus/icons-vue";
import { useRouter } from "vue-router";
import { generateCaptcha } from "../../tool/Captcha";
import { rememberPassword, forgetPassword } from "../../tool/RememberPassword";
import { handleError, Apilogin } from "@/api/enter";
import { userInfoStore } from "@/stores/userinfo/userInfo";
const userStore = userInfoStore();
const router = useRouter();
const loginForm = reactive({
  RoleID: 0,
  Account: "",
  Password: "",
  captcha: "",
});
const remember = ref(false);
const rules = {
  Account: [{ required: true, message: "请输入账号", trigger: "blur" }],
  Password: [{ required: true, message: "请输入密码", trigger: "blur" }],
  RoleID: [{ required: true, message: "请选择身份", trigger: "change" }],
  captcha: [{ required: true, message: "请输入验证码", trigger: "blur" }],
};
const options = [
  {
    value: 0,
    label: "用户",
  },
  {
    value: 1,
    label: "管理员",
  },
];

const captchaText = ref("");
captchaText.value = generateCaptcha();
const getcanvas = () => {
  captchaText.value = generateCaptcha();
};
const onSubmit = () => {
  if (loginForm.captcha.toLowerCase() === captchaText.value) {
    Apilogin(loginForm)
      .then((response) => {
        const token = response.data.token; // 从响应头中获取 Token
        localStorage.setItem("token", token); // 将 Token 保存到 localStorage
        userStore.setaccount(loginForm.Account);
        userStore.setroleid(loginForm.RoleID);
        router.push({ name: "Layout" });
      })
      .catch((error) => {
        handleError(error);
      });
  } else {
    alert("验证码错误!");
    captchaText.value = generateCaptcha();
  }
};

const onRegister = () => {
  router.push({ name: "Register" });
};

const onForgotPassword = () => {
  alert("联系管理员");
};
onMounted(() => {
  captchaText.value = generateCaptcha(); //加载验证码
});
</script>

<style>
body{
  background-color: #f5f5f5;
}
.login-container {
  position: relative;
  margin: 5vh auto;
  display: flex;
  justify-content: center;
  align-items: center;
}

.login-card {
  width: 400px;
  padding: 40px;
}

.login-title {
  text-align: center;
  margin-bottom: 30px;
  font-size: 24px;
}

.forgot-password {
  margin-left: auto;
}

.input-height .el-input__inner {
  height: 40px;
}

.captcha-canvas {
  cursor: pointer;
  margin-left: 5px;
  border: 1px solid #ccc;
  height: 40px;
  /* Ensure the canvas height matches the input height */
}

.checkbox-link-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 30px;
}

.button {
  width: 100%;
  /* margin-right: -20px; */
  margin-bottom: 10px;
}

.register-button {
  background-color: red;
  color: white;
}
.item {
  margin-left: -30px;
  width: 360px;
}
</style>
