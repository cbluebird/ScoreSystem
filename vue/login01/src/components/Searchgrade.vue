<template>
  <div>
    <el-breadcrumb separator="/" >
  <el-breadcrumb-item :to="{ path: '/Main' }">首页</el-breadcrumb-item>
  <el-breadcrumb-item>查询往年综测分数</el-breadcrumb-item>

</el-breadcrumb>
    <el-card class="box-card"  >
  <div slot="header" class="clearfix">
    <span>查询往年综测分数</span>
    <el-button style="float: right; padding: 3px 0" type="text">操作按钮</el-button>
  </div>
  <el-select v-model="value" placeholder="请选择年份"  @change="selectChanged">
    <el-option
      v-for="item in options"
      :key="item.value"
      :label="item.label"
      :value="item.value"
      >
    </el-option>
  </el-select>
  <el-descriptions  direction="vertical" :column="4" border v-if="isShow">
    <el-descriptions-item label="年份" >{{ value }}</el-descriptions-item>
    <el-descriptions-item label="分数" >{{ score }}</el-descriptions-item>
    <el-descriptions-item label="详细分数">
      <el-button icon="el-icon-search" type="text" @click="searchDetail(value)">点击查看详细分数</el-button>
    </el-descriptions-item>

</el-descriptions>

        
</el-card>


  <!--底部区-->
  <span slot="footer" class="dialog-footer">
  </span>

  </div>
</template>

<script>
  export default {
    data() {
      return {
        isShow:false,
        options: [{
          value: '2019',
          label: '2019'
        }, {
          value: '2020',
          label: '2020'
        }, {
          value: '2021',
          label: '2021'
        }, {
          value: '2022',
          label: '2022'
        }],
        value:'',
        score:0,
        userid:0,
        detailList:[],
      
          }
      },

    methods: {
  async selectChanged(value) {
    this.userid=JSON.parse(window.sessionStorage.getItem('userid'))
    const {data:res}= await this.$axios.get(
        '/api/student/get/grade',{params:{year:value,userid:this.userid}}
        
       )
       this.isShow=true
       this.score=res.data.Score
         },
    searchDetail(){
        window.sessionStorage.setItem("year",this.value)
      this.$router.push({path:'/Detail'});

 }
}
  }
</script>

<style>
.el-descriptions{
  margin-top: 15px;
}
</style>