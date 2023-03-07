<template>
  <div>
  <el-breadcrumb separator="/">
<el-breadcrumb-item :to="{ path: '/Main' }">首页</el-breadcrumb-item>
<el-breadcrumb-item>修改密码</el-breadcrumb-item>

</el-breadcrumb>
  <el-card class="box-card">
<div slot="header" class="clearfix">
  <span>老师修改密码</span>
  <el-button style="float: right; padding: 3px 0" type="text">操作按钮</el-button>
</div>
<div>
  <el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">

  <el-form-item label="新密码" prop="pass">
  <el-input type="password" v-model="ruleForm.pass" autocomplete="off"></el-input>
</el-form-item>
<el-form-item label="确认密码" prop="checkPass">
  <el-input type="password" v-model="ruleForm.password" autocomplete="off"></el-input>
</el-form-item>
<el-form-item > 
  <el-button type="primary" @click="submitForm('ruleForm')" >提交</el-button>
  <el-button @click="resetForm('ruleForm')">重置</el-button>
</el-form-item>
</el-form>
</div>

</el-card>
</div>


</template>

<script>
export default {
  data() {
    var validatePass = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入密码'));
      } else {
        if (this.ruleForm.checkPass !== '') {
          this.$refs.ruleForm.validateField('checkPass');
        }
        callback();
      }
    };
    var validatePass2 = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请再次输入密码'));
      } else if (value !== this.ruleForm.pass) {
        callback(new Error('两次输入密码不一致!'));
      } else {
        callback();
      }
    };
    return {
      ruleForm: {
        pass: '',
        password: ''
      },
      userid:0,
      rules: {

        pass: [
          { validator: validatePass, trigger: 'blur' }
        ],
        password: [
          { validator: validatePass2, trigger: 'blur' }
        ]
      }
    };
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((async valid => {
        if (valid) {
          alert('submit!');
          this.userid=JSON.parse(window.sessionStorage.getItem('userid'))
          const result =await this.$axios.post("/api/user/bind/repassword",{password:this.ruleForm.password,userid:this.userid});
          console.log(result);
          //console.log(this.ruleForm.checkPass);
        } else {
          console.log('error submit!!');
          return false;
        }
      }));
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    }
  }
}
</script>

<style scoped>
.box-card{
  margin-top:20px;
  align-items:center;
  width:50%;
  margin-left:auto;
  margin-right:auto;
}





</style>