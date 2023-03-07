<template>
  <el-container class="main-container">
    <!--头部区域-->
  <el-header>
    <div>
      <span>学生综测成绩管理系统</span>
    </div>
  <el-button type="info" @click="logout">退出</el-button>
</el-header>
  <el-container>
    <!--左侧侧边栏-->
    <el-aside width="200px">
      <el-menu background-color="#545c64" text-color="#fff" active-text-color="#409EFF" router>

      <el-menu-item index="application">
        <i class="el-icon-menu"></i>
        <span slot="title">综测申报</span>
      </el-menu-item>
      <el-menu-item index="modify">
        <i class="el-icon-setting"></i>
        <span slot="title">修改密码</span>
      </el-menu-item>
      <el-menu-item index="talk">
        <i class="el-icon-s-comment"></i>
        <span slot="title">话题广场</span>
      </el-menu-item>
      <el-menu-item index="research">
        <i class="el-icon-zoom-in"></i>
        <span slot="title">查询申报结果</span>
      </el-menu-item>
      <el-menu-item index="suggest">
        <i class="el-icon-edit"></i>
        <span slot="title">建议箱</span>
      </el-menu-item>
      <el-menu-item index="searchgrade">
        <i class="el-icon-search"></i>
        <span slot="title">查看往年所有分数</span>
      </el-menu-item>
    </el-menu>
    </el-aside>
    <!--右侧主题区域-->
    <el-main>
      <router-view>

      </router-view>
    </el-main>
  </el-container>
</el-container>

</template>

<script>
export default {
//created(){
  //this.getMenuList()
//},
methods:{
    logout(){
      this.$router.push('/login')
      //清除session
      window.sessionStorage.clear()
    },

  }
}
</script>

<style scoped>
.main-container{
  height:100vh;
  /*vh为相对与视窗的高度，视窗被分为100单位的vh，同理还有相对于视窗的宽度vw，height：100%当元素没有内容的时候不会被撑开*/
}

.el-header{
background-color: #373d41;
display:flex;/*flex弹性布局，一般用于垂直居中布局 */
justify-content:space-between;/*justify-content是水平布局，详见css3中justify-content属性的教程 */
align-items:center;/*与justify-content相对应，为垂直布局 */
color:#fff;
font-size:30px;
}
.el-aside{
  background-color:#333744;
}
.el-main{
  background-color:#eaedf1;
}


</style>