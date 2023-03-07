<template>
  <div>
    <el-breadcrumb separator="/" >
  <el-breadcrumb-item :to="{ path: '/Main' }">首页</el-breadcrumb-item>
  <el-breadcrumb-item :to="{ path: '/Searchgrade' }">查询往年综测分数</el-breadcrumb-item>
  <el-breadcrumb-item>查询细分</el-breadcrumb-item>

</el-breadcrumb>

 <el-card  class="box-card" >

  <el-table 
    style="width: 100%"
    :data="detailList">
    <el-table-column
      prop="Class"
      label="模块"
      width="220" 
      >
    </el-table-column>
    <el-table-column
      prop="Score"
      label="分数"
      width="180">                                        
    </el-table-column>
    <el-table-column
      prop="Description"
      label="描述"
      >
    </el-table-column>
    
  </el-table>

    
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
        detailList:[],
        value:0,
        userid:0
      }
    },
  
    created(){
     
     this.searchDetail(window.sessionStorage.getItem('year'))
    },
    methods: {

   async searchDetail(value) {
    this.userid=JSON.parse(window.sessionStorage.getItem('userid'))
    const {data:res}= await this.$axios.get(
        '/api/student/get/detailscore',{params:{year:value,userid:this.userid}})
        console.log(res)
        this.detailList=res.data.detaillist  
        console.log(this.detailList)
        let length = this.detailList.length
     for(let i=0;i<length;i++){
      switch(this.detailList[i].Module){
        case 1 : 
              switch(this.detailList[i].Class) {
                case 1 :
                this.detailList[i].Class = '美育素质--文化艺术实践成绩'
                break;
                case 2:
                this.detailList[i].Class = '美育素质--文化艺术竞赛获奖得分'
                break;
              }
        break;
        case 2 :
        this.detailList[i].Class = '智育素质--智育平均学分绩点'
        break;
        case 3 :
              switch(this.detailList[i].Class){
                case 1 : this.detailList[i].Class = '创新与实践素质--创新创业竞赛获奖得分'
                break;
                case 2 : this.detailList[i].Class = '创新与实践素质--水平等级考试'
                break;
                case 3 : this.detailList[i].Class = '创新与实践素质--社会实践活动'
                break;
                case 4 : this.detailList[i].Class = '创新与实践素质--社会工作'
                break;
          }
        break;
        case 4 : 
              switch(this.detailList[i].Class){
                case 1 : this.detailList[i].Class = '劳育素质--寝室日常考核基本分'
                break;
                case 2 : this.detailList[i].Class = '劳育素质--“文明寝室”创建、寝室风采展等活动加分'
                break;
                case 3 : this.detailList[i].Class = '劳育素质--寝室行为表现与卫生状况加减分'
                break;
                case 4 : this.detailList[i].Class = '劳育素质--志愿服务分'
                break;
                case 5 : this.detailList[i].Class = '劳育素质--实习实训'
                break;
              }
        break;
        case 5 :
              switch(this.detailList[i].Class){
                case 1 : this.detailList[i].Class = '德育素质--基本评定分'
                break;
                case 2 : this.detailList[i].Class = '德育素质--集体评定等级分'
                break;
                case 3 : this.detailList[i].Class = '德育素质--社会责任记实分'
                break;
                case 4 : this.detailList[i].Class = '德育素质--思政学习加减分'
                break;
                case 5 : this.detailList[i].Class = '德育素质--违纪违规扣分'
                break;
                case 6 : this.detailList[i].Class = '德育素质--学生荣誉称号加减分'
                break;
              }
        break;
        case 6 :
              switch(this.detailList[i].Class){
                case 1 : this.detailList[i].Class = '体育素质--体育课程成绩'
                break;
                case 2 : this.detailList[i].Class = '体育素质--课外体育活动成绩'
                break;
                case 3 : this.detailList[i].Class = '体育素质--体育竞赛获奖得分'
                break;
                case 4 : this.detailList[i].Class = '体育素质--早锻炼得分'
                break;
              }
        break;
      }
      }
     }
   }
  }
  
</script>

<style>

</style>