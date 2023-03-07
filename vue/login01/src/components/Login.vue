<template>
    <body id="poster">
            <el-form class="login-container" label-position="left" ref="form" :model="loginForm" :rules="loginFormRules" label-width="0px">
              <!--model为数据据绑定，所有数据都会同步到model上--->
              <!--ref可以定义表单的实例对象，获取dom元素，可以使用resetFields--->
                <h3>
                    系统登陆
                </h3>

            <el-form-item label="" prop="account">
            <el-input type="text" v-model="loginForm.account" placeholder="账号"></el-input>
            </el-form-item>
            <el-form-item label="" prop="password">
            <el-input type="password" v-model="loginForm.password" placeholder="密码"></el-input>
            </el-form-item>
            <el-form-item style="width:100%">
            <el-button type="primary" style="background-color: #505458;border:none;width:100%;" @click="Login();getData()">立即登录</el-button>
            </el-form-item>
            </el-form>
    </body>
  </template>
  
  <script>

  
  
  export default {
    name: 'Login',
    data() {
      return {
        loginForm: {
          account: '',
          password: ''
        },
      
        loginFormRules:{
          account: [
            { required: true, message: '请输入账号', trigger: 'blur' },
          ],
          password:[
           { required: true, message: '请输入密码', trigger: 'blur' },
          ]
        }
      }
      
    },
    methods: {
      Login() {
        console.log('submit!');

        
      },
      getData(){

        this.$refs.form.validate(async valid=>{
          if(!valid)return;
      const  {data:res} =await this.$axios.post("/api/user/login",this.loginForm);
      console.log(res);
     
      if(res.code!=200) return this.$message({
         showClose: true,
          message: '登录失败',
          type: 'error'
        });
        console.log(res.data.admin)
      if(res.data.admin==0)this.$router.push({path:'/Main'});
      if(res.data.admin==1)this.$router.push({path:'/TeacherMain'});
      this.$message.success('登录成功');
      window.sessionStorage.setItem("userid",res.data.userid)
          window.sessionStorage.setItem("name",res.data.name)
          window.sessionStorage.setItem("profession",res.data.profession)
      
      });
      }
    }
  } ;

  
  </script>
  

<style scoped>
  #poster{
    background-position:center ;
    height: 100%;
    width:100%;
    background-size:cover;
    position:fixed;
    background-image:url(../assets/background01.jpg);
    background-repeat:no-repeat;
}
body{
    margin:0px;
    padding:0px ;
}
.login-container{
    border-radius:15px;
    background-clip:padding-box ;
    margin: 280px auto;
    width:350px;
    background: rgba(0,0,0,0.1);
    padding:35px 35px 15px 35px;
   border:1px solid #333;
    box-shadow:0 0 5px #333 ;

}
</style>