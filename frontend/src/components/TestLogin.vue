<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50">
    <div class="max-w-md w-full space-y-8">
      <div class="text-center">
        <h2 class="text-3xl font-bold text-gray-900">Test Login Page</h2>
        <p class="mt-2 text-sm text-gray-600">This is a simple test login page</p>
        <button 
          @click="testLogin" 
          class="mt-4 bg-blue-500 text-white px-4 py-2 rounded"
        >
          Test Login
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'TestLogin',
  methods: {
    testLogin() {
      // Store fake token for testing
      localStorage.setItem('auth_token', 'test-token')
      localStorage.setItem('user', JSON.stringify({ email: 'test@example.com' }))
      this.$router.push('/admin/dashboard')
    }
  }
}
</script>
