<template>
  <el-card>
    <!-- 搜索和添加按钮的布局 -->
    <div
      style="display: flex; justify-content: space-between; margin-bottom: 20px"
    >
      <!-- 搜索框在左上角 -->
      <el-input
        v-model="searchQuery"
        placeholder="搜索标题或者来源"
        style="width: 300px"
        @input="filterMovies"
      />
      <!-- 添加按钮在右上角 -->
      <el-button type="primary" @click="openAddMovieDialog">添加咨询</el-button>
    </div>

    <!-- 电影列表表格 -->
    <el-table :data="filteredMovies" style="width: 100%">
      <el-table-column prop="title" label="标题"></el-table-column>
      <el-table-column
        prop="source"
        label="来源"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column prop="releaseDate" label="发布时间"></el-table-column>
      <el-table-column label="操作" width="180">
        <template #default="scope">
          <el-button type="danger" size="small" @click="deleteMovie(scope.row)"
            >删除</el-button
          >
        </template>
      </el-table-column>
    </el-table>
    <EditConsultation
      ref="atDialog"
      @close="handleDialogClose"
    ></EditConsultation>
  </el-card>
</template>

<script lang="ts" setup>
import { ref, computed, onMounted } from "vue";
import {
  handleError,
  ApigetConsultation,
  ApideltConsultation,
} from "@/api/enter";
import { ElMessage, ElMessageBox } from "element-plus";
import EditConsultation from "@/components/EditConsultation.vue";
// 电影列表
const movies = ref([
  {
    image: "",
    title: "",
    description: "",
    source: "",
    releaseDate: "",
    url: "",
  },
]);

const atDialog = ref();
// 搜索相关
const searchQuery = ref("");
const filteredMovies = computed(() => {
  // 根据电影名称或导演进行过滤
  if (!searchQuery.value) return movies.value;
  return movies.value.filter(
    (movie) =>
      movie.title.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      movie.description.toLowerCase().includes(searchQuery.value.toLowerCase())
  );
});

// 打开添加电影的弹窗
const openAddMovieDialog = () => {
  atDialog.value.openAddMovieDialog();
};
// 关闭添加电影
const handleDialogClose = () => {
  init(); // 触发 init 函数
};
// 删除电影
const deleteMovie = (movie: any) => {
  ElMessageBox.confirm(`确定要删除电影咨询${movie.title} 吗？`, "提示", {
    type: "warning",
  })
    .then(() => {
      ApideltConsultation(movie.title)
        .then(() => {
            ElMessage.success("电影删除成功");
            init();
        })
        .catch((error) => {
          handleError(error);
        });
    })
    .catch(() => {
      ElMessage.info("取消删除");
    });
};

// 筛选电影
const filterMovies = () => {
  // searchQuery 的变化会自动触发 filteredMovies 的重新计算
};
function init() {
  ApigetConsultation()
    .then((response) => {
      movies.value = response.data.consultations;
    })
    .catch((error) => {
      handleError(error);
    });
}
onMounted(() => {
  init();
});
</script>
