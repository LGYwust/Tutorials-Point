<template>
  <el-card>
    <!-- 搜索和添加按钮的布局 -->
    <div
      style="display: flex; justify-content: space-between; margin-bottom: 20px"
    >
      <!-- 搜索框在左上角 -->
      <el-input
        v-model="searchQuery"
        placeholder="搜索电影名称或导演"
        style="width: 300px"
        @input="filterMovies"
      />
      <!-- 添加按钮在右上角 -->
      <el-button type="primary" @click="openAddMovieDialog">添加电影</el-button>
    </div>

    <!-- 电影列表表格 -->
    <el-table :data="filteredMovies" style="width: 100%">
      <el-table-column
        prop="name"
        label="电影名称"
        width="300"
      ></el-table-column>
      <el-table-column
        prop="director"
        label="导演"
        width="300"
        show-overflow-tooltip
      ></el-table-column>
      <el-table-column prop="time" label="上映时间"></el-table-column>
      <el-table-column prop="category" label="电影类型"></el-table-column>
      <el-table-column prop="country" label="地区"></el-table-column>
      <el-table-column label="操作" width="180">
        <template #default="scope">
          <el-button
            type="danger"
            size="small"
            @click="deleteMovie(scope.row)"
            >删除</el-button
          >
        </template>
      </el-table-column>
    </el-table>
    <EditMovie ref="atDialog" @close="handleDialogClose"></EditMovie>
  </el-card>
</template>

<script lang="ts" setup>
import { ref, computed, onMounted } from "vue";
import { handleError, ApideltMovie, ApigetAllMovies } from "@/api/enter";
import { ElMessage, ElMessageBox } from "element-plus";
import EditMovie from "@/components/EditMovie.vue";
// 电影列表
const movies = ref([
  {
    name: "",
    image: "",
    director: "",
    time: "",
    country: "",
    category: "",
    quote: "",
    url:""
  },
]);

const atDialog = ref()
// 搜索相关
const searchQuery = ref("");
const filteredMovies = computed(() => {
  // 根据电影名称或导演进行过滤
  if (!searchQuery.value) return movies.value;
  return movies.value.filter(
    (movie) =>
      movie.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      movie.director.toLowerCase().includes(searchQuery.value.toLowerCase())
  );
});

// 打开添加电影的弹窗
const openAddMovieDialog = () => {
    atDialog.value.openAddMovieDialog()
};
// 关闭添加电影
const handleDialogClose = () => {
  init(); // 触发 init 函数
};
// 删除电影
const deleteMovie = (movie: any) => {
    ElMessageBox.confirm(`确定要删除电影${movie.name} 吗？`, "提示", {
      type: "warning",
    })
      .then(() => {
        ApideltMovie(movie.name)
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
}


// 筛选电影
const filterMovies = () => {
  // searchQuery 的变化会自动触发 filteredMovies 的重新计算
};
function init() {
  ApigetAllMovies()
    .then((response) => {
      movies.value = response.data.movies;
    })
    .catch((error) => {
      handleError(error);
    });
}
onMounted(() => {
  init();
});
</script>

