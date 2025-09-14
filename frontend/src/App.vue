<template>
  <div class="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center">
    <div class="bg-white rounded-lg shadow-xl p-8 max-w-md w-full mx-4">
      <div class="text-center">
        <h1 class="text-3xl font-bold text-gray-800 mb-4">
          Welcome {{ name || 'Loading...' }}!
        </h1>
        <p class="text-gray-600 mb-6">
          This is your management dashboard
        </p>
        <div class="flex justify-center">
          <div class="animate-pulse bg-blue-500 h-2 w-16 rounded"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'App',
  data() {
    return {
      name: ''
    }
  },
  async mounted() {
    try {
      const response = await axios.get('http://localhost:8080/api/name')
      this.name = response.data.name
    } catch (error) {
      console.error('Error fetching name:', error)
      this.name = 'Guest'
    }
  }
}
</script>
