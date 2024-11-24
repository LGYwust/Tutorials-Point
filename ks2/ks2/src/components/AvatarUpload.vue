<template>
    <el-dialog title="更换头像"  v-model="dialogVisible" width="500" center>
        <div class="avatar-uploader" @click="handleClick">
            <input type="file" ref="fileInput" @change="handleFileChange" style="display: none;" />
            <img v-if="imageUrl" :src="imageUrl" class="avatar">
            <i v-else class="el-icon-plus avatar-uploader-icon"></i>
        </div>
        <template #footer>
            <el-button @click="handleClose">取消</el-button>
            <el-button type="primary" @click="handleConfirm">确定</el-button>
        </template>
    </el-dialog>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { ElMessage } from 'element-plus';
import { handleError,ApichangeAvatar } from '@/api/enter';
import {userInfoStore} from '@/stores/userinfo/userInfo';
const userStore = userInfoStore()
const dialogVisible = ref(false);
const imageUrl = ref<string | null>('');
const fileInput = ref<HTMLInputElement | null>(null);
const selectedFile = ref<File | null>(null);
const handleAvatarSuccess = ( file: File) => {
    imageUrl.value = URL.createObjectURL(file);
    selectedFile.value = file;
};

const beforeAvatarUpload = (file: File) => {
    const isJPG = file.type === 'image/jpeg' || file.type === 'image/png';
    const isLt2M = file.size / 1024 / 1024 < 2;
    if (!isJPG) {
        ElMessage.error('上传头像图片只能是 JPG 或 PNG 格式!');
    }
    if (!isLt2M) {
        ElMessage.error('上传头像图片大小不能超过 2MB!');
    }
    return isJPG && isLt2M;
};

const handleFileChange = (event: Event) => {
    const input  = event.target as HTMLInputElement;
    if (input.files && input.files.length > 0) {
        const file = input.files[0];
        if (beforeAvatarUpload(file)) {
            handleAvatarSuccess(file);
        }
    }
};

const handleClick = () => {
    fileInput.value?.click();// 打开文件选择对话框
};

const handleConfirm = async () => {
    if (selectedFile.value) {
        const formData = new FormData();
        formData.append('image', selectedFile.value);
        ApichangeAvatar(formData)
            .then((response:any) => {
                userStore.setavatarpath(response.data.image_path)
                ElMessage({
                message: "头像更换成功",
                type: "success",
              });
              dialogVisible.value = false
            })
            .catch((error:any) => {
                handleError(error)
            });
    } else {
        ElMessage.error('请先选择图片');
    }
};
const handleClose = () =>{
    dialogVisible.value = false;
}

const open = ()=> {
    imageUrl.value = userStore.avatarpath
    dialogVisible.value = true;
}

defineExpose({ open });
</script>

<style scoped>
.avatar-uploader {
    width: 300px;
    height: 300px;
    border: 1px dashed #d9d9d9;
    border-radius: 50%;
    cursor: pointer;
    overflow: hidden;
    background-color: #f9f9f9;
    display: flex;
    justify-content: center;
    align-items: center;
    text-align: center;
    margin: 0 auto;
}

.avatar-uploader-icon {
    font-size: 3vw;
    color: #3c3f44;
}

.avatar {
    width: 100%;
    height: 100%;
    border-radius: 50%;
}
</style>