<template>

<el-card class="moral-box-card">
  <div slot="header" class="clearfix">
    <span>德育素质申报</span>
    <span class="text item">
    <el-dropdown @command="handleCommand">
  <span class="el-dropdown-link">
    选择项目<i class="el-icon-arrow-down el-icon--right"></i>
  </span>
  <el-dropdown-menu slot="dropdown">
    <el-dropdown-item command="a">集体评分等级分</el-dropdown-item>
    <el-dropdown-item command="b">社会责任纪实分</el-dropdown-item>
    <el-dropdown-item command="c">学业荣誉加减分</el-dropdown-item>
  </el-dropdown-menu>
</el-dropdown>

  </span>

  </div>
  <div v-if="isShow1">
    <el-row :gutter="20">
     <div >
      <el-col :span="6">
  <el-input :disabled="true" placeholder="集体评分等级分"></el-input>
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
  <div v-if="isShow2">
    <el-row :gutter="20">
     <div >
      <el-col :span="6">
  <el-input :disabled="true" placeholder="社会责任纪实分"></el-input>
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
  <div v-if="isShow3">
    <el-row :gutter="20">
     <div >
      <el-col :span="6">
  <el-input :disabled="true" placeholder="学业荣誉加减分"></el-input>
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
    module:5,   
    description:'',
    admin:'',
    class:1,
    score:'',
    },

    isShow1:false,
    isShow2:false,
    isShow3:false,
    options: [],
    TeacherId:0
    
  }
},

  methods: {
      handleCommand(command) {
        if(command=='a'){
        this.isShow1=true;this.isShow2=false;this.isShow3=false;this.applilist.class=2;
        }
        if(command=='b'){
        this.isShow1=false;this.isShow2=true;this.isShow3=false;this.applilist.class=3;
        }
        if(command=='c'){
        this.isShow1=false;this.isShow2=false;this.isShow3=true;this.applilist.class=6;

        }
      },
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

  .moral-box-card {
    width: 1100px;
    margin-left:auto;
    margin-right:auto;
    margin-top:10px;
    height:550px;
  }
  .el-dropdown-link {
    cursor: pointer;
    color: #409EFF;
  }
  .el-icon-arrow-down {
    font-size: 12px;
  }
  .el-dropdown-link{
    margin-left:10px;
  }
  .el-input{
    margin-top:10px;

  }


</style>