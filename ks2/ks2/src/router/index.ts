import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path:'/',
      redirect:'/login'
    },
    {
      path:'/login',
      name:'Login',
      component:() =>import('../views/log&reg/Login.vue'),  
    },
    {
      path:'/register',
      name:'Register',
      component:()=>import('../views/log&reg/Register.vue')
    },
    {
      path:'/layout',
      name:'Layout',
      redirect:'/layout/front',
      component:()=>import('../views/layout/Layout.vue'),
      children: [
        {
          path: 'front',
          name: 'Front',
          component: () => import('../views/frontpage/Front.vue'),
        },
        {
          path: 'allmovies',
          name: 'AllMovies',
          component: () => import('../views/allmoviesage/Allmovies.vue'),
        },
        {
          path: 'consultation',
          name: 'Consultation',
          component: () => import('../views/consultationage/Consultation.vue'),
        },
        {
          path: 'userinfo',
          name: 'UserInfo',
          component: () => import('../views/userinfoage/UserInfo.vue'),
        },
        {
          path: 'manage',
          name: 'Manage',
          component: () => import('../views/manageage/Manage.vue'),
        },
      ],
    },
    {
      path: '/voideo',
      name: 'Voideo',
      component: () => import('../views/voideoage/Voideoage.vue')
    },
  ]
})
// 导航守卫
router.beforeEach((to, from, next) => {
  const isAuthenticated = !!localStorage.getItem('token'); // 检查是否已登录
  if (to.matched.some(record => record.meta.requiresAuth) && !isAuthenticated) {
    // 如果路由需要登录状态且用户未登录，重定向到登录页面
    next({ name: 'Login' });
  } else {
    // 否则，继续导航
    next();
  }
});

export default router
