<template>
  <div class="about">
    <div class="loader">
      <svg class="circular" viewBox="25 25 50 50">
        <circle
          class="path"
          cx="50"
          cy="50"
          r="20"
          fill="none"
          stroke-width="4"
        />
      </svg>
    </div>
    <div class="text">LOADING</div>
  </div>
</template>
<script setup>
function out() {
  let con = document.querySelector(".about");
  con.classList.add("loading_out");
}
function enter(next) {
  let con = document.querySelector(".about");
  con.classList.remove("loading_out");
  setTimeout(()=>{
    next()
    out()
  },1000)
}
defineExpose({ out,enter });
</script>

<style scoped>
.about {
  position: fixed;
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #ffffff;
  transition:all 1s;
  z-index: 10000;
}
.text {
  position: absolute;
}
.loading_out{
    transform: translateY(100%);
}

.loader {
  border: 5px solid #f3f3f3; /* Light grey */
  border-top: 5px solid #3498db; /* Blue */
  border-radius: 50%;
  width: 300px;
  height: 300px;
  margin: 200px auto;
  animation: spin 2s linear infinite;
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}

.circular {
  animation: rotate 2s linear infinite;
  height: 100%;
  transform-origin: center center;
}

.path {
  stroke-dasharray: 1, 150;
  stroke-dashoffset: 0;
  animation: dash 1.5s ease-in-out infinite, color 6s ease-in-out infinite;
  stroke: #42db34;
  stroke-linecap: round;
}

@keyframes rotate {
  100% {
    transform: rotate(360deg);
  }
}

@keyframes dash {
  0% {
    stroke-dasharray: 1, 150;
    stroke-dashoffset: 0;
  }
  50% {
    stroke-dasharray: 89, 150;
    stroke-dashoffset: -35px;
  }
  100% {
    stroke-dasharray: 89, 150;
    stroke-dashoffset: -124px;
  }
}

@keyframes color {
  0%,
  100% {
    stroke: #3498db;
  }
  40% {
    stroke: #9b59b6;
  }
  66% {
    stroke: #e74c3c;
  }
  80%,
  90% {
    stroke: #3498db;
  }
}
</style>
