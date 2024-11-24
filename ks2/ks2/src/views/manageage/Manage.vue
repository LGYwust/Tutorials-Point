<template>
    <div class="container">
    <el-tabs v-model="activeName" class="demo-tabs" @tab-click="handleClick">
      <el-tab-pane label="用户管理" name="first"><Usermanage></Usermanage></el-tab-pane>
      <el-tab-pane label="电影管理" name="second"><Moviemanage></Moviemanage></el-tab-pane>
      <el-tab-pane label="咨询管理" name="third"><Consultmanage></Consultmanage></el-tab-pane>
    </el-tabs>
    </div>
  </template>
  <script lang="ts" setup>
  import { ref } from 'vue'
  import type { TabsPaneContext } from 'element-plus'
  import Usermanage from './Usermanage.vue';
  import Moviemanage from './Moviemanage.vue';
  import Consultmanage from './Consultmanage.vue';
  const activeName = ref('first')
  
  const handleClick = (tab: TabsPaneContext, event: Event) => {
    console.log(tab, event)
  }
  </script>
  
  <style>
  .demo-tabs > .el-tabs__content {
    padding: 20px;
    color: #6b778c;
    font-size: 32px;
    font-weight: 600;
  }

  .container{
    margin:-5px 80px;
  }
  
  </style>
  