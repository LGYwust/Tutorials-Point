<template>
  <div class="carousel-container">
    <el-carousel indicator-position="outside" height="auto" autoplay>
      <el-carousel-item
        v-for="item in carouselItems"
        :key="item.id"
        style="height: 350px"
      >
        <img :src="item.url" style="width: 100%; height: 100%" />
      </el-carousel-item>
    </el-carousel>
  </div>
  <div class="movie">
    <el-card class="left-card">
      <template #header>
        <div class="panel-header">
          <h2 class="title">正在热播</h2>
          <a class="goto">
            <el-link type="danger" class="all" @click="handleClick">
              全部
              <el-icon>
                <ArrowRightBold />
              </el-icon>
            </el-link>
          </a>
        </div>
      </template>
      <div class="movie-grid">
        <div
          class="movie-item"
          v-for="movie in movies"
          :key="movie.name"
          @click="goto(movie)"
        >
          <img :src="movie.image" alt="Movie Poster" class="movie-poster" />
          <div class="movie-info">
            <div class="movie-title">{{ movie.name }}</div>
            <div class="movie-data" :title="movie.director">
              {{ movie.director }}
            </div>
            <a class="movie-data">{{ movie.category }}</a>
            <a class="movie-data">{{ movie.country }}</a>
            <div class="movie-rating">
              <div class="stars">
                <el-icon
                  v-for="n in 5"
                  :key="n"
                  :class="{ filled: n <= Math.round(Number(movie.rating) / 2) }"
                >
                  <StarFilled />
                </el-icon>
                <span class="rating-number">{{ movie.rating }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </el-card>
<!-- 右侧导航栏 -->
    <el-card class="right-card">
      <!-- 搜索电影 -->
      <el-autocomplete
        v-model="search"
        :fetch-suggestions="querySearchAsync"
        placeholder="搜索电影"
        @select="handleSelect"
        value-key="name"
      >
        <template #suffix>
          <el-icon class="el-input__icon">
            <Search />
          </el-icon>
        </template>
        <template #default="{ item }">
          <div class="movie-ritem" style="margin: 0px" @click="goto(item)">
            <img :src="item.image" alt="Movie Poster" class="movie-rposter" />
            <div class="movie-rinfo">
              <div class="movie-title">{{ item.name }}</div>
              <a class="movie-data">{{ item.category }}</a>
            </div>
          </div>
        </template>
      </el-autocomplete>

      <el-divider style="width: 100%" content-position="left">
        <a class="movie-recommend" @click="getrecommendmovies()">推荐电影</a>
      </el-divider>
      <div
        class="movie-ritem"
        v-for="movie in recommendmovies"
        :key="movie.name"
        @click="goto(movie)"
      >
        <img :src="movie.image" alt="Movie Poster" class="movie-rposter" />
        <div class="movie-rinfo">
          <div class="movie-title">{{ movie.name }}</div>
          <div class="movie-data" :title="movie.director">
            {{ movie.director }}
          </div>
          <a class="movie-data">{{ movie.category }}</a>
        </div>
      </div>
    </el-card>

  </div>
</template>

<script setup lang="ts">
import { inject, onMounted, ref } from "vue";
import {
  handleError,
  ApigetHotMovies,
  ApigetRecommendMovies,
  ApigetAllMovies,
} from "@/api/enter";
import { Search } from "@element-plus/icons-vue";
import { Movie } from "@/type/movietype";
import {goto} from "@/tool/other"
const setActiveButton = inject<(button: string) => void>("setActiveButton");
const handleClick = () => {
  if (setActiveButton) {
    // 如果注入成功，才调用该方法
    setActiveButton("allmovies");
  } else {
    console.error("未提供该方法");
  }
};
const movies = ref<Movie[]>();
const recommendmovies = ref<Movie[]>();
const search = ref("");
const carouselItems = ref([
  {
    id: 1,
    url: "https://static-cse.canva.cn/blob/234790/b33338ccb83c3ca921a5c9d34773213.jpg",
  },
  { id: 2, url: "https://photo.16pic.com/00/46/10/16pic_4610785_b.jpg" },
  {
    id: 3,
    url: "https://blog.hkmovie6.com/wp-content/uploads/2020/01/Star-Wars-The-Force-Awakens-Aladdin-Dark-Phoenix-movie-posters-are-very-similar.jpg",
  },
  {
    id: 4,
    url: "https://th.bing.com/th/id/R.4c107930a0f6431cf20d3326a0e40b88?rik=6phACYC5NHM9BA&riu=http%3a%2f%2feskipaper.com%2fimages%2fawesome-movie-wallpaper-1.jpg&ehk=n2dC51IYg%2bAAWb2B%2bMFkbluC3DWwH7mZiZWh0gD1QWs%3d&risl=&pid=ImgRaw&r=0",
  },
]);
//得到热映
function gethotmovies() {
  ApigetHotMovies()
    .then((response) => {
      movies.value = response.data.movies;
    })
    .catch((error) => {
      handleError(error);
    });
}

//得到推荐电影
function getrecommendmovies() {
  ApigetRecommendMovies()
    .then((response) => {
      recommendmovies.value = response.data.movies;
    })
    .catch((error) => {
      handleError(error);
    });
}
// 异步获取建议数据
const querySearchAsync = (queryString: string, callback: any) => {
  if (queryString) {
    ApigetAllMovies()
      .then((response) => {
        const filteredSuggestions = response.data.movies.filter(
          (item: { name: string }) => {
            return item.name.toLowerCase().includes(queryString.toLowerCase());
          }
        );
        callback(filteredSuggestions);
      })
      .catch((error) => {
        handleError(error);
      });
  } else {
    callback([]); // 如果输入为空，清空建议列表
  }
};

// 选择某个建议项后的处理方法
const handleSelect = (item: any) => {
  console.log("Selected item:", item);
  // 处理选中的建议项，比如更新输入框内容或执行其他逻辑
};

onMounted(() => {
  gethotmovies();
  getrecommendmovies();
});
</script>

<style src="../styles/front.css" scoped />
