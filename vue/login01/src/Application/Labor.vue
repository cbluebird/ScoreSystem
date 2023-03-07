<template>
    <el-card class="labor-box-card">
      <div slot="header" class="clearfix">
        <span>劳育素质申报</span>
      </div>
      <div>
    <el-row :gutter="20">
     <div >
      <el-col :span="6">
  <el-input :disabled="true" placeholder="“文明寝室”创建、寝室风采展等活动加分"></el-input>
</el-col>
      <el-col :span="6"> <el-select v-model="applilist.admin"  placeholder="请选择申报老师" @visible-change='selectTeacher'>
      <el-option
      v-for="item in options"
      :key="item.TeacherId"
      :label="item.TeacherName"
      :value="item.TeacherId">
    </el-option>
  </el-select>
</el-col>
<el-col :span="6">
  <el-input placeholder="请输入分数" v-model.number="applilist.score" type="number"></el-input>
</el-col>
    </div>
  </el-row>
    <el-input
  type="textarea"
  :rows="10"
  placeholder="请输入内容"
  v-model="applilist.description"
  class="el-input">
</el-input>
<div>
  <el-button style="float: left; padding: 3px 0" type="text" @click="submitApplication()">点击提交</el-button></div>

  </div>
    </el-card>
    </template>
    
    <script>
  export default {
      data(){
        return{
        applilist:{
        userid:0,
        age: 0,  
        module:4,   
        description:'',
        admin:'',
        class:2,
        score:'',
        },

        options: [],
        TeacherId:0
        
      }
},
methods: {

          async submitApplication(){
    
            var date = new Date();
        this.applilist.age = date.getFullYear();    //  返回的是年份
        var month = date.getMonth() + 1;  //  返回的月份上个月的月份，记得+1才是当月
        if(month<=9) this.applilist.age=this.applilist.age-1
    
        this.applilist.userid=JSON.parse(window.sessionStorage.getItem('userid'))
           console.log(this.applilist)
           const{data : res}  = await  this.$axios.post('/api/student/add/application',this.applilist);
           console.log(res)
          },
          async selectTeacher(e){
          
          if(e){
            const {data:res}= await this.$axios.get(
            '/api/admin/inquireAdmin')
            
            this.options=res.data.teacherlist
            console.log(this.options)
            
          }
    
          }
        }
    }
    </script>
    
    <style scoped>
     .text {
        font-size: 14px;
      }
    
      .item {
        margin-bottom: 18px;
      }
    
      .clearfix:before,
      .clearfix:after {
        display: table;
        content: "";
      }
      .clearfix:after {
        clear: both
      }
    
      .labor-box-card {
        width: 1100px;
        margin-left:auto;
        margin-right:auto;
        margin-top:10px;
        height:550px;
      }
      .el-input{
    margin-top:10px;
  }
    </style>