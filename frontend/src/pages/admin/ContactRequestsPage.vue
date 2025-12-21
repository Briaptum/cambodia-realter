<template>
  <div>
    <!-- Page Header -->
    <div class="mb-8">
      <h1 class="text-2xl font-semibold text-zinc-900">Contact Requests</h1>
      <p class="mt-1 text-sm text-zinc-600">View and manage contact form submissions</p>
    </div>
      
      <!-- Loading State -->
      <div v-if="loading" class="flex items-center justify-center py-12">
        <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-cyan-600"></div>
      </div>

      <!-- Error State -->
      <div v-else-if="error" class="bg-red-50 border border-red-200 rounded-lg p-4 mb-6">
        <p class="text-sm text-red-800">{{ error }}</p>
      </div>

      <!-- Contact Requests List -->
      <div v-else-if="requests.length > 0" class="bg-white border border-gray-200 rounded-lg overflow-hidden">
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Name</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Email</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Phone</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Message</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">IP Address</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Submitted</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Actions</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr 
                v-for="request in requests" 
                :key="request.id" 
                @click="$router.push(`/admin/contact-requests/${request.id}`)"
                class="hover:bg-gray-50 cursor-pointer transition-colors"
              >
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                  {{ request.name }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-600">
                  <a 
                    :href="`mailto:${request.email}`" 
                    @click.stop
                    class="text-cyan-600 hover:text-cyan-800"
                  >
                    {{ request.email }}
                  </a>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-600">
                  <a 
                    v-if="request.phone"
                    :href="`tel:${request.phone}`" 
                    @click.stop
                    class="text-cyan-600 hover:text-cyan-800"
                  >
                    {{ request.phone }}
                  </a>
                  <span v-else class="text-gray-400">N/A</span>
                </td>
                <td class="px-6 py-4 text-sm text-gray-600 max-w-md">
                  <p>{{ truncateMessage(request.message, 100) }}</p>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {{ request.ip_address || 'N/A' }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {{ formatDate(request.created_at) }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  <button
                    @click.stop="handleDelete(request.id, request.name)"
                    class="text-red-600 hover:text-red-800 font-medium transition-colors"
                    title="Delete contact request"
                  >
                    Delete
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Empty State -->
      <div v-else class="bg-white p-6 border border-gray-200 rounded-lg">
        <div class="text-center py-12">
          <font-awesome-icon icon="envelope" class="mx-auto h-16 w-16 text-zinc-300 mb-4" />
          <h3 class="text-lg font-medium text-zinc-900 mb-2">No Contact Requests</h3>
          <p class="text-sm text-zinc-600">Contact form submissions will appear here</p>
        </div>
      </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'ContactRequestsPage',
  data() {
    return {
      requests: [],
      loading: true,
      error: null,
      deleting: false
    }
  },
  async mounted() {
    await this.fetchRequests()
  },
  methods: {
    getApiUrl() {
      return import.meta.env.VITE_API_URL || 'http://localhost:8080'
    },
    async fetchRequests() {
      this.loading = true
      this.error = null
      
      try {
        const response = await axios.get(`${this.getApiUrl()}/api/contact-requests`, {
          headers: {
            'Authorization': `Bearer ${this.getToken()}`
          }
        })
        this.requests = response.data.requests || []
      } catch (error) {
        console.error('Error fetching contact requests:', error)
        this.error = error.response?.data?.error || 'Failed to load contact requests'
      } finally {
        this.loading = false
      }
    },
    formatDate(dateString) {
      if (!dateString) return 'N/A'
      const date = new Date(dateString)
      return date.toLocaleDateString('en-US', {
        year: 'numeric',
        month: 'short',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit'
      })
    },
    getToken() {
      return localStorage.getItem('auth_token')
    },
    truncateMessage(message, maxLength) {
      if (!message) return ''
      if (message.length <= maxLength) return message
      return message.substring(0, maxLength) + '...'
    },
    async handleDelete(id, name) {
      if (!confirm(`Are you sure you want to delete the contact request from ${name}? This action cannot be undone.`)) {
        return
      }

      this.deleting = true
      try {
        const response = await axios.delete(`${this.getApiUrl()}/api/contact-requests/${id}`, {
          headers: {
            'Authorization': `Bearer ${this.getToken()}`
          }
        })

        if (response.status === 200) {
          // Remove the deleted request from the list
          this.requests = this.requests.filter(req => req.id !== id)
          // Show success message (you could add a toast notification here)
          alert('Contact request deleted successfully')
        }
      } catch (error) {
        console.error('Error deleting contact request:', error)
        alert(error.response?.data?.error || 'Failed to delete contact request')
      } finally {
        this.deleting = false
      }
    }
  }
}
</script>
