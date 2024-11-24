<template>
  <!-- 添加电影的弹窗 -->
  <el-dialog title="添加电影" v-model="addMovieDialogVisible">
    <el-form :model="newMovie" label-width="100px" :rules="rules">
      <el-form-item label="电影名称" prop="name" required>
        <el-input v-model="newMovie.name"></el-input>
      </el-form-item>
      <el-form-item label="电影类型" required>
        <el-select
          v-model="newMovie.category"
          multiple
          placeholder="选择电影类型"
        >
          <el-option
            v-for="genre in genres"
            :key="genre"
            :label="genre"
            :value="genre"
          ></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="地区" required>
        <el-select v-model="newMovie.country" multiple placeholder="选择地区">
          <el-option
            v-for="region in regions"
            :key="region"
            :label="region"
            :value="region"
            default-first-option
          ></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="上映时间" required>
        <el-date-picker
          v-model="newMovie.time"
          type="date"
          placeholder="选择日期"
        ></el-date-picker>
      </el-form-item>
      <el-form-item label="导演" prop="director" required>
        <el-input v-model="newMovie.director"></el-input>
      </el-form-item>
      <el-form-item label="简述" required>
        <el-input v-model="newMovie.quote" type="textarea"></el-input>
      </el-form-item>
      <el-form-item label="资源链接" prop="url" required>
        <el-input v-model="newMovie.url"></el-input>
      </el-form-item>
      <el-form-item label="电影封面">
        <div class="avatar-uploader" @click="handleClick">
          <input
            type="file"
            ref="fileInput"
            @change="handleFileChange"
            style="display: none"
          />
          <img v-if="imageUrl" :src="imageUrl" class="avatar" />
          <i v-else class="el-icon-plus avatar-uploader-icon"></i>
        </div>
      </el-form-item>
    </el-form>

    <div slot="footer" class="dialog-footer">
      <el-button @click="addMovieDialogVisible = false">取消</el-button>
      <el-button type="primary" @click="addMovie">确定</el-button>
    </div>
  </el-dialog>
</template>
<script lang="ts" setup>
import { ref } from "vue";
import {
  handleError,
  ApiaddMovie,
  ApichangeMovie,
} from "@/api/enter";
import { ElMessage } from "element-plus";
// 新电影表单
const newMovie = ref({
  name: "",
  image: "",
  director: "",
  time: "",
  country:[],
  category: [],
  quote: "",
  url:""
});
const putMovie = ref({
  name: "",
  image: "",
  director: "",
  time: "",
  country:"",
  category: "",
  quote: "",
  url:""
});

// 弹窗控制
const addMovieDialogVisible = ref(false);

// 电影类型和地区选项
const genres = ['动作', '悬疑', '科幻', '剧情', '儿童', '惊悚','爱情','喜剧','其他'];
const regions = [ '英国', '美国', '中国', '日本', '韩国', '德国', '瑞士', '印度', '其他'];

// 打开添加电影的弹窗
const openAddMovieDialog = () => {
  addMovieDialogVisible.value = true;
};
// 自定义验证函数，确保输入的是合法的URL格式
const validateUrl = (rule: any, value: string, callback: (arg0: Error | undefined) => void) => {
  const urlRegex =  /^www\.[^ "]+$/; // 正则表达式检查URL
  if (value && !urlRegex.test(value)) {
    callback(new Error('请输入有效的网页格式'));
  } else {
    callback(undefined);
  }
};

// 自定义验证函数：检查字段是否为空或仅包含空格
const validateNotEmpty = (rule: any, value: string, callback: (arg0: Error | undefined) => void) => {
  if (!value || value.trim() === '') {
    callback(new Error('此字段不能为空或全为空格'));
  } else {
    callback(undefined);
  }
};
const rules = ref({
  name: [
    { required: true, message: '请输入电影名', trigger: 'blur' },
    { validator: validateNotEmpty, trigger: 'blur' }
  ],
  director: [
    { required: true, message: '请输入导演', trigger: 'blur' },
    { validator: validateNotEmpty, trigger: 'blur' }
  ],
  url: [
    { required: true, message: '请输入资源链接', trigger: 'blur' },
    { validator: validateUrl, trigger: 'blur' }
  ]
});
// 添加电影
const addMovie = async () => {
  const trimmedName = newMovie.value.name
  const trimmedDirector = newMovie.value.director
  if (trimmedName && trimmedDirector && selectedFile.value && newMovie.value.time) {
    // 提取年份
    putMovie.value.time = new Date(newMovie.value.time)
      .getFullYear()
      .toString();
    putMovie.value.name = newMovie.value.name
    putMovie.value.director = newMovie.value.director
    putMovie.value.quote = newMovie.value.quote
    putMovie.value.url = newMovie.value.url
      putMovie.value.category = Array.isArray(newMovie.value.category)
      ? newMovie.value.category.join(" ")
      : newMovie.value.category;
      putMovie.value.country = Array.isArray(newMovie.value.country)
      ? newMovie.value.country.join(" ")
      : newMovie.value.country;
    try {
      // 首先添加电影信息
      await ApiaddMovie(putMovie.value);
      if (selectedFile.value) {
        const formData = new FormData();
        formData.append('name', putMovie.value.name);  // 电影名称
        formData.append("image", selectedFile.value);
        // 然后更改电影图片
        await ApichangeMovie(formData);
        ElMessage({
          message: "电影添加成功",
          type: "success",
        });
        selectedFile.value = null;
        imageUrl.value = null;
      } else {
        ElMessage.error("请先选择图片");
      }
      addMovieDialogVisible.value = false;
      // 清空表单
      newMovie.value = {
        name: "",
        image: "",
        director: "",
        time: "",
        country: [],
        category: [],
        quote: "",
        url: ""
      };
    } catch (error) {
      handleError(error);
    }
  } else {
    ElMessage.error("请填写完整的电影信息");
  }
};
const imageUrl = ref<string | null>("");
const fileInput = ref<HTMLInputElement | null>(null);
const selectedFile = ref<File | null>(null);
const handleAvatarSuccess = (file: File) => {
  imageUrl.value = URL.createObjectURL(file);
  selectedFile.value = file;
};

const beforeAvatarUpload = (file: File) => {
  const isJPG = file.type === "image/jpeg" || file.type === "image/png";
  const isLt2M = file.size / 1024 / 1024 < 2;
  if (!isJPG) {
    ElMessage.error("上传头像图片只能是 JPG 或 PNG 格式!");
  }
  if (!isLt2M) {
    ElMessage.error("上传头像图片大小不能超过 2MB!");
  }
  return isJPG && isLt2M;
};

const handleFileChange = (event: Event) => {
  const input = event.target as HTMLInputElement;
  if (input.files && input.files.length > 0) {
    const file = input.files[0];
    if (beforeAvatarUpload(file)) {
      handleAvatarSuccess(file);
    }
  }
};

const handleClick = () => {
  fileInput.value?.click(); // 打开文件选择对话框
};

defineExpose({ openAddMovieDialog });
</script>

<style scoped>
.dialog-footer {
  text-align: right;
}
.avatar-uploader {
  width: 150px;
  height: 200px;
  border: 1px dashed #d9d9d9;
  overflow: hidden;
  background-color: #e8e7e7;
  cursor: pointer;
  object-fit: cover;
}
.avatar-uploader img {
  width: 150px;
  height: 200px;
  object-fit: cover; /* 裁剪图片适应容器 */
}
</style>
