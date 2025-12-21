<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8 relative">
    <!-- ANHELM Watermark Logo - Bottom Left -->
    <div class="absolute bottom-6 left-6 flex items-center">
      <div class="w-4 h-4 flex flex-col items-center justify-center mr-2">
        <!-- Top two dots -->
        <div class="flex gap-1 mb-0.5">
          <div class="w-1 h-1 bg-cyan-400 rounded-full"></div>
          <div class="w-1 h-1 bg-cyan-400 rounded-full"></div>
        </div>
        <!-- Bottom dot -->
        <div class="w-1 h-1 bg-cyan-400 rounded-full"></div>
      </div>
      <span class="text-xs font-light text-zinc-500 tracking-wider">ANHELM</span>
    </div>

    <div class="max-w-md w-full space-y-6">
      <!-- Customer Logo Section -->
      <div class="text-center">
        <div class="mx-auto flex items-center justify-center mb-8">
          <!-- ANHELM Logo (dots + text) -->
          <div class="flex items-center">
            <div class="w-8 h-8 flex flex-col items-center justify-center mr-3">
              <!-- Top two dots -->
              <div class="flex gap-1.5 mb-1">
                <div class="w-1.5 h-1.5 bg-cyan-400 rounded-full"></div>
                <div class="w-1.5 h-1.5 bg-cyan-400 rounded-full"></div>
              </div>
              <!-- Bottom dot -->
              <div class="w-1.5 h-1.5 bg-cyan-400 rounded-full"></div>
            </div>
            <h1 class="text-2xl font-light text-zinc-700 tracking-wider">ANHELM</h1>
          </div>
        </div>
        <div class="mb-2">
          <h2 class="text-lg font-medium text-zinc-600">Welcome back</h2>
          <p class="text-sm text-zinc-500 mt-1">Sign in to your account</p>
        </div>
      </div>
      <form class="mt-8 space-y-6" @submit.prevent="handleLogin">
        <div class="rounded-md shadow-sm -space-y-px">
          <div>
            <label for="email" class="sr-only">Email address</label>
            <input
              id="email"
              v-model="email"
              name="email"
              type="email"
              autocomplete="email"
              required
              class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-cyan-500 focus:border-cyan-500 focus:z-10 sm:text-sm"
              placeholder="Email address"
            />
          </div>
          <div>
            <label for="password" class="sr-only">Password</label>
            <input
              id="password"
              v-model="password"
              name="password"
              type="password"
              autocomplete="current-password"
              required
              class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-cyan-500 focus:border-cyan-500 focus:z-10 sm:text-sm"
              placeholder="Password"
            />
          </div>
        </div>

        <div v-if="error" class="bg-red-50 border border-red-200 rounded-md p-3">
          <p class="text-sm text-red-800">{{ error }}</p>
        </div>

        <div>
          <button
            type="submit"
            :disabled="loading"
            class="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-cyan-600 hover:bg-cyan-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-cyan-500 disabled:opacity-50 disabled:cursor-not-allowed"
          >
            <span v-if="loading">Signing in...</span>
            <span v-else>Sign in</span>
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'LoginPage',
  data() {
    return {
      email: '',
      password: '',
      loading: false,
      error: ''
    }
  },
  methods: {
    async handleLogin() {
      this.loading = true
      this.error = ''

      try {
        const apiUrl = import.meta.env.VITE_API_URL || 'http://localhost:8080'
        const response = await axios.post(`${apiUrl}/api/auth/login`, {
          email: this.email,
          password: this.password
        })

        if (response.data.token) {
          // Store token and user info
          localStorage.setItem('auth_token', response.data.token)
          localStorage.setItem('user', JSON.stringify(response.data.user))
          
          // Emit login success
          this.$emit('login-success', response.data)
          
          // Redirect to admin dashboard
          this.$router.push('/admin/dashboard')
        }
      } catch (error) {
        console.error('Login error:', error)
        this.error = error.response?.data?.error || 'Login failed. Please try again.'
      } finally {
        this.loading = false
      }
    }
  }
}
</script>
