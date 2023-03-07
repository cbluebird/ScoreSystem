<template>
    <el-container class="home-container">
    <!-- 头部区域 -->
      <el-header >
        <div>
        <span>学生综合测试评价平台(辅导员端)</span>
        </div>
        <el-button type="info" @click="logout">退出</el-button>
      </el-header>
    <!-- 页面主体 -->
      <el-container>
    <!-- 侧边栏 -->
        <el-aside width="200px" >
        <!-- <div class="toggle-button" @click="toggleCollpase">|||</div> -->
          <el-menu
             background-color="#545c64"
             text-color="#fff"
             active-text-color="#409EFF"
             router
             unique-opened :collapse="false"
             >
             <!-- 一级菜单 -->
      <el-submenu index="score">
        <template slot="title">
          <i class="el-icon-location"></i>
          <span>成绩录入与管理</span>
        </template>
        <!-- 二级菜单 -->
          <el-menu-item index="scoreentry">成绩录入</el-menu-item>
      </el-submenu>
      <el-menu-item index="teachermodify">
        <i class="el-icon-setting"></i>
        <span slot="title">修改密码</span>
      </el-menu-item>
      <el-menu-item index="talkteacher">
        <i class="el-icon-s-comment"></i>
        <span slot="title">话题广场</span>
      </el-menu-item>
      <el-menu-item index="teacherapply">
        <i class="el-icon-s-claim"></i>
        <span slot="title">审批申请</span>
      </el-menu-item>
      <el-menu-item index="teacherreapprovedapplication">
        <i class="el-icon-refresh"></i>
        <span slot="title">重新审批</span>
      </el-menu-item>
    </el-menu>
        </el-aside>
        
    <!-- 右侧内容主体 -->
        <el-main>
          <router-view>
          </router-view>
        </el-main>
      </el-container>
    </el-container>
</template>

<script>

export  default {
    methods:{
      logout() {
      window.sessionStorage.clear()
      this.$router.push('/login')
    },
    toggleCollpase(){
    
    },
  }
}
</script>

<style lang="less" scoped>
.home-container {
  height: 100vh;
}
.el-header {
  display: flex;
  justify-content: space-between;
  padding-left: 0;
  background-color: #373d3f;
  align-items: center;
  color: #fff;
  font-size: 20px;
  > div {
    display: flex;
    align-items: center;
    span {
      margin-left: 15px;
    }
  }
}

.el-aside {
  background-color: #333744;
  .el-menu {
    border-right: none;
  }
}

.el-main {
  background-color: #eaedf1;
}
.iconfont {
  padding-right: 10px;
}
.toggle-button {
  background: #4a5064;
  font-size: 10px;
  text-align: center;
  line-height: 24px;
  color: #fff;
  letter-spacing: 0.2em;
  cursor: pointer;
}
</style>