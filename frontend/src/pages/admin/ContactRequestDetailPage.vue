<template>
  <div>
    <!-- Back Button -->
    <div class="mb-6">
      <button
        @click="$router.push('/admin/contact-requests')"
        class="inline-flex items-center text-sm text-cyan-600 hover:text-cyan-800 font-medium"
      >
        <font-awesome-icon icon="arrow-left" class="mr-2 h-4 w-4" />
        Back to Contact Requests
      </button>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="flex items-center justify-center py-12">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-cyan-600"></div>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="bg-red-50 border border-red-200 rounded-lg p-4">
      <p class="text-sm text-red-800">{{ error }}</p>
    </div>

    <!-- Contact Request Detail -->
    <div v-else-if="request" class="bg-white shadow rounded-lg">
      <div class="px-4 py-5 sm:p-6">
        <!-- Header -->
        <div class="flex items-center justify-between mb-6">
          <div>
            <h1 class="text-2xl font-semibold text-zinc-900">Contact Request #{{ request.id }}</h1>
            <p class="mt-1 text-sm text-zinc-600">Submitted {{ formatDate(request.created_at) }}</p>
          </div>
          <button
            @click="handleDelete"
            class="inline-flex items-center px-3 py-2 border border-red-300 shadow-sm text-sm leading-4 font-medium rounded-md text-red-700 bg-white hover:bg-red-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
          >
            <font-awesome-icon icon="trash" class="mr-2 h-4 w-4" />
            Delete
          </button>
        </div>

        <!-- Contact Information -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
          <div>
            <h3 class="text-lg font-medium text-gray-900 mb-4">Contact Information</h3>
            <dl class="space-y-3">
              <div>
                <dt class="text-sm font-medium text-gray-500">Name</dt>
                <dd class="mt-1 text-sm text-gray-900">{{ request.name }}</dd>
              </div>
              <div>
                <dt class="text-sm font-medium text-gray-500">Email</dt>
                <dd class="mt-1 text-sm text-gray-900">
                  <a :href="`mailto:${request.email}`" class="text-cyan-600 hover:text-cyan-800">
                    {{ request.email }}
                  </a>
                </dd>
              </div>
              <div v-if="request.phone">
                <dt class="text-sm font-medium text-gray-500">Phone</dt>
                <dd class="mt-1 text-sm text-gray-900">
                  <a :href="`tel:${request.phone}`" class="text-cyan-600 hover:text-cyan-800">
                    {{ request.phone }}
                  </a>
                </dd>
              </div>
            </dl>
          </div>

          <div>
            <h3 class="text-lg font-medium text-gray-900 mb-4">Technical Details</h3>
            <dl class="space-y-3">
              <div>
                <dt class="text-sm font-medium text-gray-500">Submitted</dt>
                <dd class="mt-1 text-sm text-gray-900">{{ formatFullDate(request.created_at) }}</dd>
              </div>
              <div v-if="request.ip_address">
                <dt class="text-sm font-medium text-gray-500">IP Address</dt>
                <dd class="mt-1 text-sm text-gray-900">{{ request.ip_address }}</dd>
              </div>
              <div v-if="request.user_agent">
                <dt class="text-sm font-medium text-gray-500">User Agent</dt>
                <dd class="mt-1 text-sm text-gray-900 break-all">{{ request.user_agent }}</dd>
              </div>
            </dl>
          </div>
        </div>

        <!-- Message -->
        <div>
          <h3 class="text-lg font-medium text-gray-900 mb-4">Message</h3>
          <div class="bg-gray-50 rounded-lg p-4 border-l-4 border-cyan-500">
            <p class="text-sm text-gray-900 whitespace-pre-wrap">{{ request.message }}</p>
          </div>
        </div>

        <!-- Actions -->
        <div class="mt-8 flex space-x-4">
          <a
            :href="`mailto:${request.email}?subject=Re: Your Contact Request`"
            class="inline-flex items-center px-4 py-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-cyan-600 hover:bg-cyan-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-cyan-500"
          >
            <font-awesome-icon icon="reply" class="mr-2 h-4 w-4" />
            Reply via Email
          </a>
          <a
            v-if="request.phone"
            :href="`tel:${request.phone}`"
            class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-cyan-500"
          >
            <font-awesome-icon icon="phone" class="mr-2 h-4 w-4" />
            Call {{ request.phone }}
          </a>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'ContactRequestDetailPage',
  data() {
    return {
      request: null,
      loading: true,
      error: null
    }
  },
  async mounted() {
    await this.fetchRequest()
  },
  methods: {
    getApiUrl() {
      return import.meta.env.VITE_API_URL || 'http://localhost:8080'
    },
    getToken() {
      return localStorage.getItem('auth_token')
    },
    async fetchRequest() {
      this.loading = true
      this.error = null
      
      try {
        const id = this.$route.params.id
        const response = await axios.get(`${this.getApiUrl()}/api/contact-requests/${id}`, {
          headers: {
            'Authorization': `Bearer ${this.getToken()}`
          }
        })
        this.request = response.data.request
      } catch (error) {
        console.error('Error fetching contact request:', error)
        this.error = error.response?.data?.error || 'Failed to load contact request'
      } finally {
        this.loading = false
      }
    },
    formatDate(dateString) {
      if (!dateString) return 'N/A'
      const date = new Date(dateString)
      const now = new Date()
      const diffTime = Math.abs(now - date)
      const diffDays = Math.floor(diffTime / (1000 * 60 * 60 * 24))
      
      if (diffDays === 0) {
        return 'today'
      } else if (diffDays === 1) {
        return 'yesterday'
      } else if (diffDays < 7) {
        return `${diffDays} days ago`
      } else {
        return date.toLocaleDateString('en-US', {
          month: 'long',
          day: 'numeric',
          year: 'numeric'
        })
      }
    },
    formatFullDate(dateString) {
      if (!dateString) return 'N/A'
      const date = new Date(dateString)
      return date.toLocaleDateString('en-US', {
        weekday: 'long',
        year: 'numeric',
        month: 'long',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit'
      })
    },
    async handleDelete() {
      if (!confirm(`Are you sure you want to delete this contact request from ${this.request.name}? This action cannot be undone.`)) {
        return
      }

      try {
        const response = await axios.delete(`${this.getApiUrl()}/api/contact-requests/${this.request.id}`, {
          headers: {
            'Authorization': `Bearer ${this.getToken()}`
          }
        })

        if (response.status === 200) {
          alert('Contact request deleted successfully')
          this.$router.push('/admin/contact-requests')
        }
      } catch (error) {
        console.error('Error deleting contact request:', error)
        alert(error.response?.data?.error || 'Failed to delete contact request')
      }
    }
  }
}
</script>
