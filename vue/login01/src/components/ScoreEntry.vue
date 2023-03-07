<template>
    <div>
    <el-breadcrumb separator="/">
  <el-breadcrumb-item :to="{ path: '/TeacherHome' }">首页</el-breadcrumb-item>
  <el-breadcrumb-item>成绩录入与管理</el-breadcrumb-item>
</el-breadcrumb>
    <!--卡片视图-->
    <el-card>
        <!--搜索与添加区域-->
     <el-row :gutter="20">
      <el-col :span="8">
        <el-input placeholder="请输入内容" 
        v-model="queryInfo.query" clearable
        @clear="getStudentList">
         <el-button slot="append" icon="el-icon-search" @click="getStudentList"></el-button>
       </el-input>
      </el-col>
      <el-col :span="4">
        <el-button type="primary" @click="goAddscore">录入学生成绩</el-button>
      </el-col>
     </el-row >
     <!--学生列表区-->
     <el-table :data="studentsList" border stripe>
      <el-table-column label="#" type="index"></el-table-column>
      <el-table-column label="姓名" prop="name"></el-table-column>
      <el-table-column label="美育素质" prop="score1"></el-table-column>
      <el-table-column label="智育素质" prop="score2"></el-table-column>
      <el-table-column label="创新与实践素质" prop="score3"></el-table-column>
      <el-table-column label="劳育素质" prop="score4"></el-table-column>
      <el-table-column label="德育素质" prop="score5"></el-table-column>
      <el-table-column label="体育素质" prop="score6"></el-table-column>
      <el-table-column label="操作" width="180px">
        <template slot-scope="scope">
           <!-- 修改按钮 -->
           <el-button type="primary" icon="el-icon-edit" size="mini" @click="showdialog(scope.row.studentid)"></el-button>
            <!-- 删除按钮 -->
            <el-button type="danger" icon="el-icon-delete" size="mini" @click="removeStudentById(scope.row.studentid)"></el-button>
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
      <!--弹窗-->
      <el-dialog
         title="提示"
         :visible.sync="dialogVisible"
         width="50%"
         :before-close="handleClose">
         <!--内容-->
         <div class="block">
         <span class="demonstration">选择分项</span>
         <el-cascader
            v-model="value"
           :options="options"
           :props="{ expandTrigger: 'hover' }"
            @change="handleChange"></el-cascader>
          </div>
          <el-input v-model="score" ></el-input>
            <span slot="footer" class="dialog-footer">
          <el-button @click="dialogVisible = false">取 消</el-button>
            <el-button type="primary" @click="Submit(newid)">确 定</el-button>
          </span>
      </el-dialog>
    </el-card>


    </div>
  </template>
  
  <script>
  export default {
    data(){
      return {
        //获取列表参数对象
        queryInfo: {
          query: '',
          pagenum: 1,
          pagesize: 5
        },
        dialogVisible: false,
        studentsList: [],
        score: 0,
        total: 0,
        newid: 0,
        modifyList: [],
        module: 0,
        type: 0,
        value: [],
        options:[{
          value:'1',
          label:'美育素质',
          children:[{
            value:'1',
            label:'文化艺术实践成绩'},
            {
              value:'2',
              label:'文化艺术竞赛获奖得分'
            }]
        },
        {
          value:'2',
          label:'智育素质',
          children:[{
            value:'1',
            label:'智育平均学分绩点'
          }]
        },
        {
          value:'3',
          label:'创新与实践素质',
          children:[{
            value:'1',
            label:'创新创业竞赛获奖得分'
          },{
            value:'2',
            label:'水平等级考试'
          },{
            value:'3',
            label:'社会实践活动'
          },{
            value:'4',
            label:'社会工作'
          }]
        },
        {
          value:'4',
          label:'劳育素质',
          children:[{
            value:'1',
            label:'寝室日常考核基本分'
          },{
            value:'2',
            label:'“文明寝室”创建、寝室风采展等活动加分'
          },{
            value:'3',
            label:'寝室行为表现与卫生状况加减分'
          },{
            value:'4',
            label:'志愿服务分'
          },{
            value:'5',
            label:'实习实训'
          }]
        },{
          value:'5',
          label:'德育素质' ,
          children:[{
            value:'1',
            label:'基本评定分'
          },{
            value:'2',
            label:'集体评定等级分'
          },{
            value:'3',
            label:'社会责任记实分'
          },{
            value:'4',
            label:'思政学习加减分'
          },{
            value:'5',
            label:'违纪违规扣分'
          },{
            value:'6',
            label:'学生荣誉称号加减分'
          }]       
        },{
          value:'6',
          label:'体育素质',
          children:[{
            value:'1',
            label:'体育课程成绩'
          },{
            value:'2',
            label:'课外体育活动成绩'
          },{
            value:'3',
            label:'体育竞赛获奖得分'
          },{
            value:'4',
            label:'早锻炼得分'
          }
        ]
        }
      ]
      }
    },
    created(){
      this.getStudentList()
    },
    methods: {
      async getStudentList(){
      const { data : res } = await  this.$axios.get('/api/admin/instructor/scoresearch',{ 
        params : this.queryInfo 
      });
      const result  = await  this.$axios.get('/api/admin/instructor/scoresearch',{ 
        params : this.queryInfo 
      })
      ///instructor/scoresearch
      if (result.status !== 200) return this.$message.error('获取学生列表失败')
      this.studentsList = res.data.students
      this.total = res.data.total
      console.log(res)
      },
      // 监听 pagesize 改变事件 每页显示的个数
    handleSizeChange(newSize) {
      console.log(newSize)
      this.queryInfo.pagesize = newSize
      this.getStudentList()
    },
    // 监听 页码值 改变的事件 当前页面值
    handleCurrentChange(newPage) {
      console.log(newPage)
      this.queryInfo.pagenum = newPage
      this.getStudentList()
    },
    goAddscore(){
      this.$router.push('/TeacherAddscore')
    },
    handleClose(done) {
        this.$confirm('确认关闭？')
          .then(_ => {
            done();
          })
          .catch(_ => {});
      },
      handleChange(value) {
        console.log(value);
      },
      async Submit(id){
          this.modifyList.studentname=id;
          this.modifyList.module=this.value[0];
          this.modifyList.type = this.value[1]
          this.modifyList.score=this.score;
          console.log(this.modifyList);
          const result  = await  this.$axios.post('/api/admin/putscore',this.modifyList);
          console.log(result);
          this.dialogVisible = false;
      },
      showdialog(id){
        this.dialogVisible = true;
        this.newid=id;
        console.log(id)
      },
      async removeStudentById(id){
        //弹框询问是否删除
        const confirmResult = await this.$confirm('此操作将永久删除该学生成绩, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).catch(err => err)


        //确认删除，返回confirm字符串
        //取消删除，返回cancel
        //console.log(confirmResult)
        if(confirmResult !== 'confirm') {
          return this.$message.info('已取消删除')
        }

       const res = await this.$http.post('admin/deletestudent', id)
        if(res.status != 200) {
          return this.$message.error('删除失败')
        }
        this.$message.success('删除成功')
        this.getStudentList()
      }
  }
  };
  </script>
  
  <style>
  .el-input{
    margin:10px;
  }
  .el-breadcrumb {
  margin-bottom: 15px;
  font-size: 12px;
}
  </style>