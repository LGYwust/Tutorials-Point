<template>
  <div class="register-container">
    <el-card class="register-card">
      <el-button class="back-button" @click="goBack"><el-icon size="20">
          <CaretLeft />
        </el-icon>&nbsp;&nbsp;&nbsp;</el-button>
      <h2 class="register-title">注册</h2>
      <el-form :model="registerForm" :rules="rules" label-width="80px" class="register-form" hide-required-asterisk
        ref="registerFormRef">
        <el-form-item label="账号" prop="Account" class="item">
          <el-input v-model="registerForm.Account" placeholder="请输入账号" class="input-height">
            <template #prefix>
              <el-icon>
                <User />
              </el-icon>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item label="邮箱" prop="Email" class="item">
          <el-input v-model="registerForm.Email" placeholder="请输入邮箱" class="input-height">
            <template #prefix>
              <el-icon>
                <Message />
              </el-icon>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item label="手机号" prop="Phone" class="item">
          <el-input v-model="registerForm.Phone" placeholder="请输入手机号" class="input-height">
            <template #prefix>
              <el-icon>
                <Phone />
              </el-icon>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item label="密码" prop="Password" class="item">
          <el-input v-model="registerForm.Password" type="password" placeholder="请输入密码" class="input-height">
            <template #prefix>
              <el-icon>
                <Lock />
              </el-icon>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item label="确认密码" prop="Confirmpassword" class="item">
          <el-input v-model="registerForm.Confirmpassword" type="password" placeholder="请确认密码" class="input-height">
            <template #prefix>
              <el-icon>
                <Lock />
              </el-icon>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item label="验证码" prop="Captcha" class="item">
          <el-row>
            <el-col :span="14">
              <el-input v-model="registerForm.Captcha" placeholder="请输入验证码" class="input-height">
                <template #prefix>
                  <el-icon>
                    <CircleCheck />
                  </el-icon>
                </template>
              </el-input>
            </el-col>
            <el-col :span="10">
              <canvas ref="captchaCanvas" width="100" height="40" @click="getcanvas"
                class="captcha-canvas"></canvas>
            </el-col>
          </el-row>
        </el-form-item>

        <el-button type="primary" @click="onRegister" class="button register-button">注册</el-button>
      </el-form>
    </el-card>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted, reactive, watch } from "vue";
import { useRouter } from "vue-router";
import { FormInstance } from "element-plus";
import { ElMessage } from "element-plus";
import { generateCaptcha } from "../../tool/Captcha";
import { handleError, Apiregister } from "@/api/enter";

const router = useRouter();
const registerForm = reactive({
  Account: "",
  Email: "",
  Phone: "",
  Password: "",
  Confirmpassword: "",
  Captcha: "",
});

const validateConfirmpassword = (rule: any, value: string, callback: (arg0: Error | undefined) => void) => {
  if (value !== registerForm.Password) {
    callback(new Error("确认密码与密码不一致"));
  } else {
    callback(undefined);
  }
};

const validatePasswordLength = (rule: any, value: string, callback: (arg0: Error | undefined) => void) => {
  if (value.length < 8) {
    callback(new Error("密码长度至少8位"));
  }  else if(/\s/.test(value)){
    callback(new Error("密码不能存在空格"));
  }else{
    callback(undefined);
  }
};

const validateAccountFormat = (rule: any, value: string, callback: (arg0: Error | undefined) => void) => {
  if (value.length<6) {
    callback(new Error("账号长度至少大于6"));
  } else if(/\s/.test(value)){
    callback(new Error("账号不能存在空格"));
  }else{
    callback(undefined);
  }
};


const  validateEmailFormat= (rule: any, value: string, callback: (arg0: Error | undefined) => void) => {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  if (!emailRegex.test(value)) {
    callback(new Error("请输入有效的邮箱地址"));
  } else {
    callback(undefined);
  }
};
const getcanvas=()=>{
  captchaText.value = generateCaptcha()
}
const rules = {
  Account: [
    { required: true, message: "请输入账号", trigger: "blur" },
    {  validator: validateAccountFormat, trigger: "blur"}
  ],
  Email: [
    { required: true, message: "请输入邮箱", trigger: "blur" },
    { validator: validateEmailFormat, trigger: "blur" },
  ],
  Phone: [{ required: true, message: "请输入手机号", trigger: "blur" },  { pattern: /^1[0-9]{10}$/, message: '手机号格式不正确', trigger: 'blur' }],
  Password: [
    { required: true, message: "请输入密码", trigger: "blur" },
    { validator: validatePasswordLength, trigger: "blur" },
  ],
  Confirmpassword: [
    { required: true, message: "请确认密码", trigger: "blur" },
    { validator: validateConfirmpassword, trigger: "blur" },
  ],
  Captcha: [{ required: true, message: "请输入验证码", trigger: "blur" }],
};

const registerFormRef = ref<FormInstance | null>(null);
const captchaText = ref(""); // 验证码

const validateOnInput = () => {
  if (registerFormRef.value) {
    registerFormRef.value.validateField("Password");
    registerFormRef.value.validateField("Confirmpassword");
  }
};

watch(() => registerForm.Password, validateOnInput);
watch(() => registerForm.Confirmpassword, validateOnInput);

captchaText.value = generateCaptcha();

const onRegister = () => {
  if (registerFormRef.value) {
    registerFormRef.value.validate(async (valid) => {
      if (valid) {
        if (registerForm.Captcha.toLowerCase() === captchaText.value) {
          Apiregister(registerForm)
            .then((response) => {
              ElMessage({
                message: "注册成功",
                type: "success",
              });
              router.push("login");
            })
            .catch((error) => {
              handleError(error);
            });
        } else {
          alert("验证码错误!");
          captchaText.value = generateCaptcha();
        }
      } else {
        console.log("error submit!");
        return false;
      }
    });
  }
};

const goBack = () => {
  router.push({ name: "Login" });
};

onMounted(() => {
  captchaText.value = generateCaptcha();
});
</script>

<style >
.register-container {
  position: relative;
  margin: 100px auto;
  display: flex;
  justify-content: center;
  align-items: center;
}

.register-card {
  width: 400px;
  padding: 40px;
  position: relative;
  /* 使返回按钮定位相对于卡片 */
}

.register-title {
  text-align: center;
  margin-bottom: 30px;
  font-size: 24px;
}

.input-height .el-input__inner {
  height: 40px;
}

.captcha-canvas {
  cursor: pointer;
  margin-left: 5px;
  border: 1px solid #ccc;
  height: 40px;
  /* 确保画布高度与输入高度匹配 */
}

.button {
  width: 100%;
  margin-bottom: 10px;
}

.register-button {
  background-color: red;
  color: white;
}

/* 新增样式 */
.el-form-item__label {
  display: flex;
  justify-content: flex-end;
  width: 100px;
  /* 设定一个固定宽度 */
}

.back-button {
  position: absolute;
  top: 20px;
  left: 25px;
}

.item {
  margin-left: -20px;
  width: 360px;
}
</style>
