<template>
  <div class="p-8">
    <h1 class="text-2xl font-bold text-gray-900 mb-4">Test Dashboard</h1>
    <p class="text-gray-600">This is a simple test dashboard without FontAwesome icons.</p>
    <div class="mt-4">
      <button 
        @click="logout" 
        class="bg-red-500 text-white px-4 py-2 rounded"
      >
        Logout
      </button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'TestDashboard',
  methods: {
    logout() {
      localStorage.removeItem('auth_token')
      localStorage.removeItem('user')
      this.$router.push('/login')
    }
  }
}
</script>
