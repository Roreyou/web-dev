import Vue from 'vue'
import VueRouter from 'vue-router'
// 创建路由组件
import Create from '../views/Create.vue'
import Manage from '../views/Manage.vue'
import Record from '../views/Record.vue'
import Upload from '../views/Upload.vue'
import Change from '../views/Change.vue'
import Main from '../views/Main.vue'
import Login from '../views/Login.vue'
import Usermanage from '../views/Usermanage.vue'
import GPUmanage from '../views/GPUmanage.vue'
import Contmanage from '../views/Contmanage.vue'



// 将组件和路由进行映射
const routes = [
    //登录
    {
        name:'Login',
        component:Login,
        path:'/'
    },

    // 主路由
    {
        path: '/',
        component: Main,
        redirect: '/manage',
        children: [//配置的嵌套路由
            { path: 'create', component: Create },
            { path: 'manage', component: Manage },
            { path: 'record', component: Record },
            { path: 'upload', component: Upload },
            { path: 'change', component: Change },
        ]
    },

    {
        name:'GPUmanage',
        component: GPUmanage,
        path: '/gpumanage'
    },

    {
        name:'Usermanage',
        component: Usermanage,
        path: '/usermanage'
    },

    {
        name:'Contmanage',
        component: Contmanage,
        path: '/contmanage'
    },
    
]

// 创建router实例
const router = new VueRouter({
    routes // (缩写) 相当于 routes: routes
})

export default router

Vue.use(VueRouter)

