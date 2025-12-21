import { createRouter, createWebHistory } from 'vue-router'

// Import test components
import TestLogin from '../components/TestLogin.vue'
import TestDashboard from '../pages/admin/TestDashboard.vue'

// Auth guard
const requireAuth = (to, from, next) => {
  const token = localStorage.getItem('auth_token')
  if (!token) {
    next('/login')
  } else {
    next()
  }
}

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      redirect: '/admin/dashboard'
    },
    {
      path: '/login',
      name: 'login',
      component: TestLogin
    },
    {
      path: '/admin/dashboard',
      name: 'dashboard',
      component: TestDashboard,
      beforeEnter: requireAuth
    }
  ]
})

// Global navigation guard to check auth
router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('auth_token')
  
  // If going to login and already authenticated, redirect to dashboard
  if (to.path === '/login' && token) {
    next('/admin/dashboard')
    return
  }
  
  // If going to protected route without token, redirect to login
  if (to.path.startsWith('/admin') && !token) {
    next('/login')
    return
  }
  
  next()
})

export default router
