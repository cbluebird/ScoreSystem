<template>
    <div>
    <el-breadcrumb separator="/">
    <el-breadcrumb-item :to="{ path: '/TeacherHome' }">首页</el-breadcrumb-item>
    <el-breadcrumb-item :to="{ path: '/scoreentry' }">成绩录入与管理</el-breadcrumb-item>
    <el-breadcrumb-item>成绩录入</el-breadcrumb-item>
    </el-breadcrumb>
    <el-card>
        
        <div class="block">
         <span class="demonstration">选择分项</span>
         <el-cascader
            v-model="value"
           :options="options"
           :props="{ expandTrigger: 'hover' }"
            @change="handleChange"></el-cascader>
          </div>
          <span>
            <div>请输入学生姓名</div>
            <el-input v-model="name" ></el-input>
            <div>请输入分数</div>
          <el-input v-model.number="score" type="number"></el-input>
            <el-button type="primary" @click="Submit()">确 定</el-button>
          </span>
    </el-card>
</div>
    </template>
    
    <script>
    
    export  default {
        data(){
            return{
                value: [],
                score: '',
                studentList: {},
                name: '',
        options:[{
          value:1,
          label:'美育素质',
          children:[{
            value:1,
            label:'文化艺术实践成绩'},
            {
              value:2,
              label:'文化艺术竞赛获奖得分'
            }]
        },
        {
          value:2,
          label:'智育素质',
          children:[{
            value:1,
            label:'智育平均学分绩点'
          }]
        },
        {
          value:3,
          label:'创新与实践素质',
          children:[{
            value:1,
            label:'创新创业竞赛获奖得分'
          },{
            value:2,
            label:'水平等级考试'
          },{
            value:3,
            label:'社会实践活动'
          },{
            value:4,
            label:'社会工作'
          }]
        },
        {
          value:4,
          label:'劳育素质',
          children:[{
            value:1,
            label:'寝室日常考核基本分'
          },{
            value:2,
            label:'“文明寝室”创建、寝室风采展等活动加分'
          },{
            value:3,
            label:'寝室行为表现与卫生状况加减分'
          },{
            value:4,
            label:'志愿服务分'
          },{
            value:5,
            label:'实习实训'
          }]
        },{
          value:5,
          label:'德育素质' ,
          children:[{
            value:1,
            label:'基本评定分'
          },{
            value:2,
            label:'集体评定等级分'
          },{
            value:3,
            label:'社会责任记实分'
          },{
            value:4,
            label:'思政学习加减分'
          },{
            value:5,
            label:'违纪违规扣分'
          },{
            value:6,
            label:'学生荣誉称号加减分'
          }]       
        },{
          value:6,
          label:'体育素质',
          children:[{
            value:1,
            label:'体育课程成绩'
          },{
            value:2,
            label:'课外体育活动成绩'
          },{
            value:3,
            label:'体育竞赛获奖得分'
          },{
            value:4,
            label:'早锻炼得分'
          }
        ]
        }
      ]
            }
        },
        methods: {
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
      go(){
        this.$router.push('/scoreentry')
      },
      async Submit(){
        this.userid=JSON.parse(window.sessionStorage.getItem('userid'))
        var date = new Date();
    this.studentList.age = date.getFullYear();    //  返回的是年份
    var month = date.getMonth() + 1;  //  返回的月份上个月的月份，记得+1才是当月
    if(month<=9) this.studentList.age=this.studentList.age-1

          this.studentList.studentname=this.name;
          this.studentList.module=this.value[0];
          this.studentList.class = this.value[1];
          this.studentList.score=this.score;
          this.studentList.userid=this.userid;
        this.studentList.type=1;
          console.log(this.studentList);
          const result  = await this.$axios.post('/api/admin/putscore',this.studentList);
          console.log(result);


          this.$message.success('成功');
        this.name=' '
        this.score=''
      },
        }
    }
    </script>
    
    <style lang="less" scoped>
      .el-breadcrumb {
  margin-bottom: 15px;

}
    </style>