<template>
  <div>
    <!-- Page Header -->
    <div class="mb-8">
      <h1 class="text-2xl font-semibold text-zinc-900">Dashboard</h1>
      <p class="mt-1 text-sm text-zinc-600">Welcome to the admin dashboard</p>
    </div>

    <!-- Stats Cards -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
      <div class="bg-white overflow-hidden shadow rounded-lg">
        <div class="p-5">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <font-awesome-icon icon="envelope" class="h-6 w-6 text-cyan-400" />
            </div>
            <div class="ml-5 w-0 flex-1">
              <dl>
                <dt class="text-sm font-medium text-gray-500 truncate">Total Contact Requests</dt>
                <dd class="text-lg font-medium text-gray-900">{{ stats.totalRequests || 0 }}</dd>
              </dl>
            </div>
          </div>
        </div>
      </div>

      <div class="bg-white overflow-hidden shadow rounded-lg">
        <div class="p-5">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <font-awesome-icon icon="calendar-day" class="h-6 w-6 text-green-400" />
            </div>
            <div class="ml-5 w-0 flex-1">
              <dl>
                <dt class="text-sm font-medium text-gray-500 truncate">This Week</dt>
                <dd class="text-lg font-medium text-gray-900">{{ stats.thisWeek || 0 }}</dd>
              </dl>
            </div>
          </div>
        </div>
      </div>

      <div class="bg-white overflow-hidden shadow rounded-lg">
        <div class="p-5">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <font-awesome-icon icon="clock" class="h-6 w-6 text-yellow-400" />
            </div>
            <div class="ml-5 w-0 flex-1">
              <dl>
                <dt class="text-sm font-medium text-gray-500 truncate">Today</dt>
                <dd class="text-lg font-medium text-gray-900">{{ stats.today || 0 }}</dd>
              </dl>
            </div>
          </div>
        </div>
      </div>

      <div class="bg-white overflow-hidden shadow rounded-lg">
        <div class="p-5">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <font-awesome-icon icon="user-check" class="h-6 w-6 text-blue-400" />
            </div>
            <div class="ml-5 w-0 flex-1">
              <dl>
                <dt class="text-sm font-medium text-gray-500 truncate">Unique Contacts</dt>
                <dd class="text-lg font-medium text-gray-900">{{ stats.uniqueEmails || 0 }}</dd>
              </dl>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Recent Contact Requests -->
    <div class="bg-white shadow rounded-lg">
      <div class="px-4 py-5 sm:p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg leading-6 font-medium text-gray-900">Recent Contact Requests</h3>
          <router-link
            to="/admin/contact-requests"
            class="text-sm text-cyan-600 hover:text-cyan-800 font-medium"
          >
            View all
          </router-link>
        </div>

        <!-- Loading State -->
        <div v-if="loading" class="flex items-center justify-center py-8">
          <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-cyan-600"></div>
        </div>

        <!-- Error State -->
        <div v-else-if="error" class="bg-red-50 border border-red-200 rounded-lg p-4">
          <p class="text-sm text-red-800">{{ error }}</p>
        </div>

        <!-- Recent Requests List -->
        <div v-else-if="recentRequests.length > 0" class="flow-root">
          <ul class="-my-5 divide-y divide-gray-200">
            <li v-for="request in recentRequests" :key="request.id" class="py-4">
              <div class="flex items-center space-x-4">
                <div class="flex-shrink-0">
                  <div class="h-8 w-8 bg-cyan-100 rounded-full flex items-center justify-center">
                    <span class="text-sm font-medium text-cyan-600">{{ request.name.charAt(0).toUpperCase() }}</span>
                  </div>
                </div>
                <div class="flex-1 min-w-0">
                  <p class="text-sm font-medium text-gray-900 truncate">{{ request.name }}</p>
                  <p class="text-sm text-gray-500 truncate">{{ request.email }}</p>
                </div>
                <div class="flex-shrink-0 text-sm text-gray-500">
                  {{ formatDate(request.created_at) }}
                </div>
              </div>
            </li>
          </ul>
        </div>

        <!-- Empty State -->
        <div v-else class="text-center py-8">
          <font-awesome-icon icon="envelope" class="mx-auto h-12 w-12 text-gray-300 mb-4" />
          <h3 class="text-sm font-medium text-gray-900 mb-2">No contact requests yet</h3>
          <p class="text-sm text-gray-500">Contact form submissions will appear here</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'DashboardPage',
  data() {
    return {
      stats: {
        totalRequests: 0,
        thisWeek: 0,
        today: 0,
        uniqueEmails: 0
      },
      recentRequests: [],
      loading: true,
      error: null
    }
  },
  async mounted() {
    await this.fetchDashboardData()
  },
  methods: {
    getApiUrl() {
      return import.meta.env.VITE_API_URL || 'http://localhost:8080'
    },
    getToken() {
      return localStorage.getItem('auth_token')
    },
    async fetchDashboardData() {
      this.loading = true
      this.error = null
      
      try {
        const response = await axios.get(`${this.getApiUrl()}/api/contact-requests`, {
          headers: {
            'Authorization': `Bearer ${this.getToken()}`
          }
        })
        
        const requests = response.data.requests || []
        this.recentRequests = requests.slice(0, 5) // Show only 5 most recent
        
        // Calculate stats
        this.calculateStats(requests)
      } catch (error) {
        console.error('Error fetching dashboard data:', error)
        this.error = error.response?.data?.error || 'Failed to load dashboard data'
      } finally {
        this.loading = false
      }
    },
    calculateStats(requests) {
      const now = new Date()
      const today = new Date(now.getFullYear(), now.getMonth(), now.getDate())
      const weekAgo = new Date(today.getTime() - 7 * 24 * 60 * 60 * 1000)
      
      this.stats.totalRequests = requests.length
      
      this.stats.today = requests.filter(req => {
        const reqDate = new Date(req.created_at)
        return reqDate >= today
      }).length
      
      this.stats.thisWeek = requests.filter(req => {
        const reqDate = new Date(req.created_at)
        return reqDate >= weekAgo
      }).length
      
      // Count unique email addresses
      const uniqueEmails = new Set(requests.map(req => req.email))
      this.stats.uniqueEmails = uniqueEmails.size
    },
    formatDate(dateString) {
      if (!dateString) return 'N/A'
      const date = new Date(dateString)
      const now = new Date()
      const diffTime = Math.abs(now - date)
      const diffDays = Math.floor(diffTime / (1000 * 60 * 60 * 24))
      
      if (diffDays === 0) {
        return 'Today'
      } else if (diffDays === 1) {
        return 'Yesterday'
      } else if (diffDays < 7) {
        return `${diffDays} days ago`
      } else {
        return date.toLocaleDateString('en-US', {
          month: 'short',
          day: 'numeric',
          year: date.getFullYear() !== now.getFullYear() ? 'numeric' : undefined
        })
      }
    }
  }
}
</script>
