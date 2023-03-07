<template>
    <div>
    <el-breadcrumb separator="/">
    <el-breadcrumb-item :to="{ path: '/TeacherHome' }">首页</el-breadcrumb-item>
    <el-breadcrumb-item>重新审批</el-breadcrumb-item>
    </el-breadcrumb>
    <el-card>
      <el-table :data="applicationList" border stripe>
        <el-table-column label="#" type="index"></el-table-column>
      <el-table-column label="申请项" prop="Class"></el-table-column>
      <el-table-column label="申请理由" prop="Description"></el-table-column>
      <el-table-column label="申请分数" prop="Score"></el-table-column>
      <el-table-column label="操作" width="180px">
          <template slot-scope="scope">
             <!-- 通过按钮 -->
             <el-button type="primary" icon="el-icon-check" size="mini" class="onebutton" @click="returnResult(scope.row.ID)">通过</el-button>
              <!-- 不通过按钮 -->
              <el-button type="danger" icon="el-icon-close" size="mini" class="onebutton" @click="giveId(scope.row.ID)">不通过</el-button>
          </template>
        </el-table-column>
      </el-table>
       <!--分页-->
       <el-pagination
          @current-change="handleCurrentChange"
          :current-page="queryInfo.pagenum"
          @size-change="handleSizeChange"
          :page-size="queryInfo.pagesize"
          :page-sizes="[1,2,5,10]"
          layout="total, sizes, prev, pager, next, jumper"
          :total="total"
        >
        </el-pagination>
        <!--不通过理由弹窗-->
        <el-dialog
           title="提示"
           :visible.sync="dialogVisible"
           width="30%"
           :before-close="handleClose">
           <!--内容-->
           <el-input
           type="textarea"
           :rows="6"
           placeholder="请输入理由"
           v-model="result.reason"
           maxlength="200"
           show-word-limit></el-input>
              <span slot="footer" class="dialog-footer">
            <el-button @click="dialogVisible = false">取 消</el-button>
              <el-button type="primary" @click="submitResult()">确 定</el-button>
            </span>
        </el-dialog>
  
    </el-card>
    </div>
  </template>
    
    <script>
    export default {
      data(){
        return {
           queryInfo: {
            page: 1,
            page_size: 2
          },
          applicationList: [],
          total: 0,
          dialogVisible: false,
          result: { 
          reason: '',
          status: 0,
          userid: 0,
          id: 0
        },
          success: 1,
          false: 2,
        }
      },
      created(){
        this.getApplication()
      },
      methods: {
        async getApplication(){
          this.queryInfo.userid=JSON.parse(window.sessionStorage.getItem('userid'))
          const { data : res } = await  this.$axios.get('/api/admin/getreapprovedapplication',{ 
          params : this.queryInfo 
        });

        this.applicationList = res.data.application
          this.total=res.data.total
      var length = this.applicationList.length
      for(var i=0;i<length;i++){
      switch(this.applicationList[i].Module){
        case 1 : 
              switch(this.applicationList[i].Class) {
                case 1 :
                this.applicationList[i].Class = '美育素质--文化艺术实践成绩'
                break;
                case 2:
                this.applicationList[i].Class = '美育素质--文化艺术竞赛获奖得分'
              }
        break;
        case 2 :
        this.applicationList[i].Class = '智育素质--智育平均学分绩点'
        break;
        case 3 :
              switch(this.applicationList[i].Class){
                case 1 : this.applicationList[i].Class = '创新与实践素质--创新创业竞赛获奖得分'
                break;
                case 2 : this.applicationList[i].Class = '创新与实践素质--水平等级考试'
                break;
                case 3 : this.applicationList[i].Class = '创新与实践素质--社会实践活动'
                break;
                case 4 : this.applicationList[i].Class = '创新与实践素质--社会工作'
                break;
          }
        break;
        case 4 : 
              switch(this.applicationList[i].Class){
                case 1 : this.applicationList[i].Class = '劳育素质--寝室日常考核基本分'
                break;
                case 2 : this.applicationList[i].Class = '劳育素质--“文明寝室”创建、寝室风采展等活动加分'
                break;
                case 3 : this.applicationList[i].Class = '劳育素质--寝室行为表现与卫生状况加减分'
                break;
                case 4 : this.applicationList[i].Class = '劳育素质--志愿服务分'
                break;
                case 5 : this.applicationList[i].Class = '劳育素质--实习实训'
                break;
              }
        break;
        case 5 :
              switch(this.applicationList[i].Class){
                case 1 : this.applicationList[i].Class = '德育素质--基本评定分'
                break;
                case 2 : this.applicationList[i].Class = '德育素质--集体评定等级分'
                break;
                case 3 : this.applicationList[i].Class = '德育素质--社会责任记实分'
                break;
                case 4 : this.applicationList[i].Class = '德育素质--思政学习加减分'
                break;
                case 5 : this.applicationList[i].Class = '德育素质--违纪违规扣分'
                break;
                case 6 : this.applicationList[i].Class = '德育素质--学生荣誉称号加减分'
                break;
              }
        break;
        case 6 :
              switch(this.applicationList[i].Class){
                case 1 : this.applicationList[i].Class = '体育素质--体育课程成绩'
                break;
                case 2 : this.applicationList[i].Class = '体育素质--课外体育活动成绩'
                break;
                case 3 : this.applicationList[i].Class = '体育素质--体育竞赛获奖得分'
                break;
                case 4 : this.applicationList[i].Class = '体育素质--早锻炼得分'
                break;
              }
        break;
      }
    }
       console.log(this.applicationList)
      console.log(res)
        },
          // 监听 pagesize 改变事件 每页显示的个数
      handleSizeChange(newSize) {
        console.log(newSize)
        this.queryInfo.pagesize = newSize
        this.getApplication()
      },
      // 监听 页码值 改变的事件 当前页面值
      handleCurrentChange(newPage) {
        console.log(newPage)
        this.queryInfo.pagenum = newPage
        this.getApplication()
      },
      handleClose(done) {
          this.$confirm('确认关闭？')
            .then(_ => {
              done();
            })
            .catch(_ => {});
        },
        async giveId(ID){
          this.result.id = ID
          this.dialogVisible = true
        },
        async returnResult(ID) {
        this.result.userid=JSON.parse(window.sessionStorage.getItem('userid'))
        this.result.id = ID
        this.result.status = this.success
          this.result.reason="ok"
        const res = await this.$axios.post('/api/admin/judge',this.result)
        console.log('提交成功')

        this.getApplication()
      },
      async submitResult() {
        this.result.userid=JSON.parse(window.sessionStorage.getItem('userid'))

        this.result.status = this.false
        const res = await this.$axios.post('/api/admin/judge',this.result)
        console.log(this.result)
        this.dialogVisible = false
        this.getApplication()

      }
      
      }
    };
    </script>
    
    <style lang="less" scoped>
  .onebutton{
     margin: 10px;
  
  }
  .el-breadcrumb {
  margin-bottom: 15px;

}
    </style>