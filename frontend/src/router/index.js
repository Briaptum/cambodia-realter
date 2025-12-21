import { createRouter, createWebHistory } from 'vue-router'

// Import components
import HomePage from '../pages/HomePage.vue'
import LoginPage from '../components/LoginPage.vue'
import AdminLayout from '../layouts/AdminLayout.vue'
import DashboardPage from '../pages/admin/DashboardPage.vue'
import ContactRequestsPage from '../pages/admin/ContactRequestsPage.vue'
import ContactRequestDetailPage from '../pages/admin/ContactRequestDetailPage.vue'

// Auth guard with backend validation
const requireAuth = async (to, from, next) => {
  console.log('Auth guard triggered for:', to.path)
  const token = localStorage.getItem('auth_token')
  
  if (!token) {
    console.log('No token, redirecting to login')
    next('/login')
    return
  }
  
  // Validate token with backend
  try {
    console.log('Validating token with backend...')
    const apiUrl = import.meta.env.VITE_API_URL || 'http://localhost:8080'
    const response = await fetch(`${apiUrl}/api/profile`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json'
      }
    })
    
    if (response.ok) {
      const userData = await response.json()
      console.log('Token valid, user:', userData.email)
      // Update user data in localStorage
      localStorage.setItem('user', JSON.stringify(userData))
      next()
    } else {
      console.log('Token invalid, clearing auth data and redirecting to login')
      // Clear fake/invalid auth data
      localStorage.removeItem('auth_token')
      localStorage.removeItem('user')
      next('/login')
    }
  } catch (error) {
    console.log('Auth validation failed:', error.message)
    // Clear auth data on network/validation error
    localStorage.removeItem('auth_token')
    localStorage.removeItem('user')
    next('/login')
  }
}

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomePage
    },
    {
      path: '/login',
      name: 'login',
      component: LoginPage
    },
    {
      path: '/admin',
      component: AdminLayout,
      beforeEnter: requireAuth,
      props: () => {
        const userData = localStorage.getItem('user')
        if (userData) {
          try {
            return { user: JSON.parse(userData) }
          } catch (error) {
            return { user: { email: 'Unknown User' } }
          }
        }
        return { user: { email: 'Unknown User' } }
      },
      children: [
        {
          path: '',
          redirect: 'dashboard'
        },
        {
          path: 'dashboard',
          name: 'dashboard',
          component: DashboardPage,
          beforeEnter: requireAuth
        },
        {
          path: 'contact-requests',
          name: 'contact-requests',
          component: ContactRequestsPage,
          beforeEnter: requireAuth
        },
        {
          path: 'contact-requests/:id',
          name: 'contact-request-detail',
          component: ContactRequestDetailPage,
          beforeEnter: requireAuth
        }
      ]
    }
  ],
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    }
    return { top: 0, behavior: 'smooth' }
  }
})

// Global navigation guard with backend validation
router.beforeEach(async (to, from, next) => {
  console.log('Global guard - navigating to:', to.path)
  const token = localStorage.getItem('auth_token')
  
  // If going to login and have token, validate it first
  if (to.path === '/login' && token) {
    try {
      console.log('Validating existing token before redirecting...')
      const apiUrl = import.meta.env.VITE_API_URL || 'http://localhost:8080'
      const response = await fetch(`${apiUrl}/api/profile`, {
        method: 'GET',
        headers: {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json'
        }
      })
      
      if (response.ok) {
        console.log('Token valid, redirecting to dashboard')
        next('/admin/dashboard')
        return
      } else {
        console.log('Token invalid, clearing and staying on login')
        localStorage.removeItem('auth_token')
        localStorage.removeItem('user')
        next()
        return
      }
    } catch (error) {
      console.log('Token validation failed, clearing and staying on login')
      localStorage.removeItem('auth_token')
      localStorage.removeItem('user')
      next()
      return
    }
  }
  
  // If going to any admin route without token, redirect to login
  if (to.path.startsWith('/admin') && !token) {
    console.log('No token for admin route, redirecting to login')
    next('/login')
    return
  }
  
  console.log('Global guard - allowing navigation')
  next()
})

export default router
