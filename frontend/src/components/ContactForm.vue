<template>
  <form @submit.prevent="handleSubmit" class="space-y-4">
    <div>
      <label for="name" class="block text-sm font-medium text-gray-700 mb-1">Name</label>
      <input
        id="name"
        v-model="formData.name"
        type="text"
        required
        class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-cyan-500 focus:border-cyan-500"
      />
    </div>
    
    <div>
      <label for="email" class="block text-sm font-medium text-gray-700 mb-1">Email</label>
      <input
        id="email"
        v-model="formData.email"
        type="email"
        required
        class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-cyan-500 focus:border-cyan-500"
      />
    </div>
    
    <div>
      <label for="phone" class="block text-sm font-medium text-gray-700 mb-1">Phone (optional)</label>
      <input
        id="phone"
        v-model="formData.phone"
        type="tel"
        class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-cyan-500 focus:border-cyan-500"
      />
    </div>
    
    <div>
      <label for="message" class="block text-sm font-medium text-gray-700 mb-1">Message</label>
      <textarea
        id="message"
        v-model="formData.message"
        rows="4"
        required
        class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-cyan-500 focus:border-cyan-500"
        placeholder="Tell us about your project or question..."
      ></textarea>
    </div>

    <div v-if="submitError" class="bg-red-50 border border-red-200 rounded-md p-3">
      <p class="text-sm text-red-800">{{ submitError }}</p>
    </div>

    <div v-if="submitSuccess" class="bg-green-50 border border-green-200 rounded-md p-3">
      <p class="text-sm text-green-800">Thank you! Your message has been sent successfully.</p>
    </div>

    <button
      type="submit"
      :disabled="submitting"
      class="w-full bg-cyan-600 text-white py-2 px-4 rounded-md hover:bg-cyan-700 focus:outline-none focus:ring-2 focus:ring-cyan-500 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
    >
      <span v-if="submitting">Sending...</span>
      <span v-else>Send Message</span>
    </button>
  </form>
</template>

<script>
import axios from 'axios'

export default {
  name: 'ContactForm',
  data() {
    return {
      formData: {
        name: '',
        email: '',
        phone: '',
        message: ''
      },
      submitting: false,
      submitError: '',
      submitSuccess: false
    }
  },
  methods: {
    async handleSubmit() {
      this.submitting = true
      this.submitError = ''
      this.submitSuccess = false

      try {
        const apiUrl = import.meta.env.VITE_API_URL || 'http://localhost:8080'
        const response = await axios.post(`${apiUrl}/api/contact-requests`, {
          name: this.formData.name,
          email: this.formData.email,
          phone: this.formData.phone || undefined,
          message: this.formData.message
        })

        if (response.status === 201) {
          this.submitSuccess = true
          // Reset form
          this.formData = {
            name: '',
            email: '',
            phone: '',
            message: ''
          }
          // Clear success message after 5 seconds
          setTimeout(() => {
            this.submitSuccess = false
          }, 5000)
        }
      } catch (error) {
        console.error('Error submitting contact request:', error)
        this.submitError = error.response?.data?.error || 'Failed to send message. Please try again.'
      } finally {
        this.submitting = false
      }
    }
  }
}
</script>
