<template>
  <div class="selectheader">
    <div class="con-row">
      <span class="name">分类：</span>
      <div class="val">
        <a
          v-for="category in categories"
          :key="category"
          :class="['a', { active: active[0] === category }]"
          @click="setActiveCategory(category)"
          :title="category"
        >
          {{ category }}
        </a>
      </div>
    </div>
    <div class="con-row">
      <span class="name">地区：</span>
      <div class="val">
        <a
          v-for="area in areas"
          :key="area"
          :class="['a', { active: active[1] === area }]"
          @click="setArea(area)"
          :title="area"
        >
          {{ area }}
        </a>
      </div>
    </div>
    <div class="con-row" style="align-self: flex-end; margin-right: 50px">
      <el-radio-group v-model="hotornew" size="large">
        <el-radio-button label="最热" value="Hot" />
        <el-radio-button label="最新" value="New" />
      </el-radio-group>
    </div>
  </div>
  <div class="movie-grid">
    <el-empty
      v-if="!movies || movies.length == 0"
      :image-size="300"
      class="empty-container"
    />
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
</template>

<script setup lang="ts">
import { handleError, Apiselectmovie } from "@/api/enter";
import { Movie } from "@/type/movietype";
import {goto} from "@/tool/other"
import { onMounted, ref, watch } from "vue";
const active = ref(["全部", "全部", "Hot"]); // 默认为全部
const categories = [
  "全部",
  "动作",
  "悬疑",
  "科幻",
  "剧情",
  "儿童",
  "惊悚",
  "爱情",
  "喜剧",
  "其他",
];
const areas = [
  "全部",
  "英国",
  "美国",
  "中国",
  "日本",
  "韩国",
  "德国",
  "瑞士",
  "印度",
  "其他",
];
const hotornew = ref("Hot");
const movies = ref<Movie[]>();
const setActiveCategory = (category: string) => {
  active.value[0] = category;
  getMovies(active.value);
};
const setArea = (area: string) => {
  active.value[1] = area;
  getMovies(active.value);
};
const getMovies = (active: string[]) => {
  Apiselectmovie(active)
    .then((response) => {
      movies.value = response.data.movies;
    })
    .catch((error) => {
      handleError(error);
    });
};
// 使用 watch 监听 hotornew 的变化
watch(hotornew, () => {
  // 当 hotornew 变化时，重新调用 getMovies 函数
  active.value[2] = hotornew.value;
  getMovies(active.value);
});
onMounted(() => {
  getMovies(active.value);
});
</script>

<style src="../styles/allmovies.css" scoped />
