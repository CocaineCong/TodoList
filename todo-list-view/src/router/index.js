//配置路由相关的信息
import VueRouter from 'vue-router'  //导入vue插件  导入之前安装
import Vue from 'vue'
//导入vue  vue来安装vue-router
//1.vue的use安装这个插件
Vue.use(VueRouter)

const login = () => import("../components/Login")
const register = () => import("../components/Register")
const Home = ()=>import("../components/Home")
//2.创建routes对象参数

const routes = [
    {
        path: '/',
        redirect: '/login'
    },
    {
        path: '/login',
        component: login,
        meta: {
            title: '登陆'
        }
    },{
        path: '/register',
        component: register,
        meta: {
            title: '注册'
        }
    },{
        path: '/home',
        component: Home,
        meta:{
            title: '主页'
        }
    }
    ]



//导出对象并设置为HTML5的history模式
const router =  new VueRouter({
    routes,
    mode: 'history'
})
// router.beforeEach((to, from, next) => {
//     console.log(to + from +next)
//     next()
//     /* 必须调用 `next` */
// })


export default router;