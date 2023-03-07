<template>
  <div>
    <el-breadcrumb separator="/">
  <el-breadcrumb-item :to="{ path: '/Main' }">首页</el-breadcrumb-item>
  <el-breadcrumb-item>综测申报</el-breadcrumb-item>
</el-breadcrumb>
<div class="option">
<el-button type="primary" style="width:200px;height:60px;margin:auto" @click="moral">德育素质</el-button>
  <el-button type="success" style="width:200px;height:60px;margin:auto" @click="sport">体育素质</el-button>
  <el-button type="info" style="width:200px;height:60px;margin:auto"  @click="art">美育素质</el-button>
  <el-button type="warning" style="width:200px;height:60px;margin:auto" @click="labor">劳育素质</el-button>
  <el-button type="danger" style="width:200px;height:60px;margin:auto" @click="creation">创新与实践素质</el-button>
</div>
    <el-card class="box-card">
        <span>注意：正式申报前需明确学校各项规定。</span>
  
<span>
  <el-link type="primary" href="/#/Rule"><i class="el-icon-view el-icon--right"></i>点击此处查看各项规定 </el-link>
</span>
        <br/>
        <span> <el-checkbox v-model="checked">我已知晓</el-checkbox></span>
  <div class="text item">

  </div>
</el-card>
<div>
    <router-view>

    </router-view>
 </div>
</div>

</template>

<script>
export default {
    data(){
return{
  checked:false
}
    },
    methods:{
        moral(){
          if(this.checked==true)
            this.$router.push({path:'/Moral'});
            
        },
        sport(){
          if(this.checked==true)
            this.$router.push({path:'/Sport'});
        },
        art(){
          if(this.checked==true)
            this.$router.push({path:'/Art'});
        },
        labor(){
          if(this.checked==true)
            this.$router.push({path:'/Labor'});
        },
        creation(){
          if(this.checked==true)
            this.$router.push({path:'/Creation'});
        }
    }
}
</script>

<style scoped>
.option{
    margin-top: 20px;
    display: flex;
    width:1130px;
    margin-left:auto;
    margin-right:auto;
}
.text {
    font-size: 14px;
  }

  .item {
    padding: 0px 0;
  }

  .box-card {
    width: 1100px;
    margin-top:30px;
    margin-left:auto;
    margin-right:auto;
    height:80px;
  }

</style>