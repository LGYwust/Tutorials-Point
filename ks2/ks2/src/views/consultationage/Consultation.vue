<template>
  <div class="movie-container">
    <el-empty
      v-if="currentMovies.length === 0"
      :image-size="300"
      description="暂无咨询信息"
    ></el-empty>
    <!-- 电影列表 -->
    <div v-else>
      <div
        v-for="movie in currentMovies"
        :key="movie.id"
        class="movie-item"
        @click="goto(movie)"
      >
        <img :src="movie.image" alt="Movie Cover" class="movie-cover" />
        <div class="movie-info">
          <h3 class="movie-title">{{ movie.title }}</h3>
          <p class="movie-desc">{{ movie.description }}</p>
          <div class="movie-meta">
            <span>出自：{{ movie.source }}</span> |
            <span>发布时间：{{ movie.releaseDate }}</span>
          </div>
        </div>
      </div>
    </div>
    <!-- 分页组件容器，居中分页栏 -->
    <div class="pagination-container">
      <el-pagination
        v-model:current-page="currentPage"
        :page-size="10"
        :total="movies.length"
        layout="prev, pager, next"
        @current-change="handlePageChange"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { handleError, ApigetConsultation } from "@/api/enter";
import { ElMessageBox } from "element-plus";
// 模拟电影数据
interface Movie {
  id: number;
  image: string;
  title: string;
  description: string;
  source: string; // 新增：来源
  releaseDate: string; // 新增：时间
}

const movies = ref<Movie[]>([]);

// 当前页码
const currentPage = ref(1);
const pageSize = 20;

// 计算当前页的电影数据
const currentMovies = computed(() => {
  const start = (currentPage.value - 1) * pageSize;
  return movies.value.slice(start, start + pageSize);
});

// 处理页码变更
const handlePageChange = (page: number) => {
  currentPage.value = page;
};

const getConsultations = () => {
  ApigetConsultation()
    .then((response) => {
      movies.value = response.data.consultations;
    })
    .catch((error) => {
      handleError(error);
    });
};

const goto = (movie: any) => {
  if (movie.url == "") {
    ElMessageBox.confirm("暂无资源请联系管理员", "错误", {
      confirmButtonText: "OK",
      cancelButtonText: "Cancel",
      type: "error",
    });
  } else {
    window.open("http://" + movie.url, "_blank");
  }
};

onMounted(() => {
  getConsultations();
});
</script>

<style scoped>
.movie-container {
  width: 100%;
  margin: 20px auto;
}

.movie-item {
  display: flex;
  align-items: flex-start; /* 修改为顶部对齐 */
  border-bottom: 1px solid #ddd;
  padding-bottom: 15px;
  margin: 15px 80px;
  cursor: pointer;
}

.movie-cover {
  width: 300px;
  height: 180px;
  margin-right: 20px;
}

.movie-info {
  flex: 1;
}

.movie-title {
  margin: 0;
  font-size: 1.4em; /* 调整字体大小 */
  font-weight: bold; /* 加粗标题 */
}

.movie-desc {
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 4; /* 限制显示4行 */
  overflow: hidden;
  text-overflow: ellipsis;
  height: 100px;
  max-width: 1100px;
  min-width: 500px;
  font-size: 14px;
  color: #7a7a7a;
  line-height: 1.5em; /* 行高 */
  max-height: calc(1.5em * 4); /* 根据行高计算最大高度 */
}
.movie-info p {
  margin: 5px 0;
  color: #666;
}

.movie-meta {
  bottom: 5px;
  color: #999;
  font-size: 0.9em;
}

.pagination-container {
  display: flex;
  justify-content: center;
  margin-top: 20px;
}
</style>
