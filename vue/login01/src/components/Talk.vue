<template>
  <div>
 <div>
    <el-breadcrumb separator="/">
  <el-breadcrumb-item :to="{ path: '/TeacherMain' }">首页</el-breadcrumb-item>
  <el-breadcrumb-item>话题广场</el-breadcrumb-item>
</el-breadcrumb>

<el-card class="box-card"> 

<el-button type="primary" @click="addDialogVisible=true" class="send-post" >发起话题</el-button>
<el-table
      :data="wordlist">
      <el-table-column
      width="80">
        <div>
      <el-avatar icon="el-icon-user-solid"></el-avatar>
    </div>
      </el-table-column>
      <el-table-column
        label="言论"
        prop="Word">
      </el-table-column>
    </el-table>
    <el-pagination
        @current-change="handleCurrentChange"
        :current-page="queryInfo.page_num"
        @size-change="handleSizeChange"
        :page-size="queryInfo.page_size"
        :page-sizes="[1,2,5,10]"
        layout="total, sizes, prev, pager, next, jumper"
        :total="total"
      >
      </el-pagination>
</el-card>
<!--发帖时的对话框-->
<el-dialog
  title="发起话题"
  :visible.sync="addDialogVisible"
  width="50%">
  <!--内容主题区-->
  <el-input
  type="textarea"
  :rows="6"
  placeholder="请输入内容"
  v-model="word"
  maxlength="200"
  show-word-limit>

</el-input> 

  <!--底部区-->
  <span slot="footer" class="dialog-footer">
    <el-button @click="addDialogVisible = false">取 消</el-button>
    <el-button type="primary" @click="publishWord">确 定</el-button>
  </span>
</el-dialog>



</div>




</div>
</template>

<script>
export default {
data(){
  return{
    //控制添加话题
    addDialogVisible:false,
    wordlist:[],
    word:'',
    userid:0,
    queryInfo:{
      page_num:1,
      page_size:2
    },
    total:0
  }
},
created(){
this.getWord()
},
methods:{
  async getWord() {
    const {data:res}= await this.$axios.get(
        '/api/square/iquire',{
          params:this.queryInfo
        })
      console.log(res)
      this.wordlist=res.data.square
      this.total=res.data.total


    },
  async publishWord(){
    this.userid=JSON.parse(window.sessionStorage.getItem('userid'))
      const {data:res}=await this.$axios.post('/api/square/pulish',{word:this.word,userid:this.userid})
      console.log(res)
      this.addDialogVisible = false
      this.getWord()
    },
    handleSizeChange(newSize){
    this.queryInfo.page_size=newSize
    this.getWord()
    },
    handleCurrentChange(newPage){
      this.queryInfo.page_num=newPage
      this.getWord()
    }
  }
}
</script>

<style scoped>
  .send-post{
width:120px;
margin:auto;
margin-top:10px;
} 
 .post-page{
   text-align:right ;
   margin-bottom: 36px;
}

  .box-card {
    width: 100%;
    height:700px;
    box-shadow: 0 2px 2px rgba(0,0,0,0.15) !important;
  
  }
  .el-pagination{
   margin-bottom:10px;
  }
</style>