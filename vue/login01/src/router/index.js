import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../components/Login.vue'
import Main from '../components/Main.vue'
import Welcome from '../components/Welcome.vue'
import Modify from '../components/Modify.vue'
import Application from '../components/Application.vue'
import Rule from '../components/Rule.vue'
import Talk from '../components/Talk.vue'
import Research from '../components/Research.vue'
import Searchgrade from '../components/Searchgrade.vue'
import Suggest from '../components/Suggest.vue'
import Detail from '../components/Detail.vue'
import TeacherMain from '../components/TeacherMain.vue'
import TalkTeacher from '../components/TalkTeacher.vue'
import TeacherModify from '../components/TeacherModify.vue'
import ScoreEntry from '../components/ScoreEntry.vue'
import TeacherApply from '../components/TeacherApply.vue'
import TeacherAddscore from '../components/TeacherAddscore.vue'
import TeacherReapprovedApplication from '../components/TeacherReapprovedApplication.vue'
import Moral from '../Application/Moral.vue'
import Sport from '../Application/Sport.vue'
import Art from '../Application/Art.vue'
import Labor from '../Application/Labor.vue'
import Creation from '../Application/Creation.vue'
Vue.use(VueRouter)
const originalPush = VueRouter.prototype.push
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}


const router = new VueRouter({
  routes:[
    {
      path: '/',
      redirect: '/login'
    },
    {
      path: '/login',//当用户访问path为/Login时
      name: 'login',
      component: Login//展示上方的Login组件
    },
    {
      path: '/TeacherMain',
      component: TeacherMain,
      children:[{
        path:'/TalkTeacher',
        component:TalkTeacher},
        {path:'/TeacherModify',
        component:TeacherModify},
        {path:'/ScoreEntry',
        component:ScoreEntry},
        {path:'/TeacherApply',
        component:TeacherApply},
        {path:'/TeacherReapprovedApplication',
        component:TeacherReapprovedApplication},
        {path:'/TeacherAddscore',
        component:TeacherAddscore}
      ]
    },
    {
      path: '/Main',
      component: Main,
      redirect:'/welcome',
      children:[{
        path:'/welcome',
        component:Welcome},
        {path:'/Modify',
        component:Modify},
        {path:'/Talk',
        component:Talk},
        {path:'/Research',
        component:Research},
        {path:'/Suggest',
        component:Suggest},
        {path:'/Searchgrade',
        component:Searchgrade},
        {path:'/Detail',
        component:Detail},
        {path:'/Application',
      component:Application,
      children:[{path:'/Moral',component:Moral},{path:'/Sport',component:Sport},{path:'/Art',component:Art},{path:'/Labor',component:Labor},{path:'/Creation',component:Creation}]
    },
    {path:'/Rule',
    component:Rule},
  

      ]
    }
  ]
  
})


router.beforeEach((to,from,next)=> {
      if (to.path === '/login') return next()
     const userid=window.sessionStorage.getItem('userid')
  if(!userid)return next('/login')
  next()
    }
)
export default router
