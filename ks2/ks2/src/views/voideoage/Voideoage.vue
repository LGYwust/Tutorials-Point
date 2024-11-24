<template>
  <el-header class="header"> 电 影 院 线 </el-header>
  <div class="video-player">
    <div class="video-info">
      <p class="title">标题</p>
      <p class="director">导演</p>
      <p class="summary">简介</p>
    </div>
    <div class="genres">
      <span v-for="genre in genres" :key="genre" class="genre">
        {{ genre }}
      </span>
    </div>
    <video
      ref="videoElement"
      :src="videoSrc"
      controls
      width="100%"
      @play="handlePlay"
      @pause="handlePause"
      @ended="handleEnd"
    ></video>
  </div>
</template>

<script setup>
import { ref } from "vue";

const videoSrc = "path_to_your_video.mp4"; // 视频文件路径
const videoElement = ref(null);
const isPlaying = ref(false);
const genres = ref(["剧情", "喜剧", "冒险"]); // 假设这是视频的类型

const handlePlay = () => {
  isPlaying.value = true;
};

const handlePause = () => {
  isPlaying.value = false;
};

const handleEnd = () => {
  isPlaying.value = false;
};
</script>

<style scoped>
.header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 6vh;
  min-height: 50px;
  background: linear-gradient(to right, #001f3f, #0074D9);
  color: white;
  /* 设置字体颜色为白色 */
  font-size: 1.8em;
  /* 设置字体大小稍微放大，根据需要调整具体数值 */
  font-weight: bold;
  /* 设置字体为黑体加粗 */
  letter-spacing: 2px;
  /* 增加字体间的间距，根据需要调整具体数值 */
}

.video-player {
  position: relative;
  width: 60vw;
  min-width: 600px;
  margin: 20px auto;
}

.video-info {
  margin: 10px;
}

.title {
  font-size: 24px;
  font-weight: bold;
  color: #333;
  margin-bottom: 10px;
}

.director {
  font-size: 16px;
  color: #666;
  margin-bottom: 5px;
}

.summary {
  font-size: 14px;
  color: #666;
}

.genres {
  position: absolute;
  top: 20px;
  right: 20px;
  display: flex;
  flex-wrap: wrap;
  gap: 10px; /* 间距 */
}

.genre {
  background-color: #fff;
  border: 1px solid #ddd;
  border-radius: 20px;
  padding: 5px 15px;
  font-size: 14px;
  color: #333;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}
</style>
