<template>
  <MarketingLayout>
    <!-- 404 Content -->
    <section class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-16">
      <div class="text-center">
        <!-- Large 404 -->
        <div class="mb-8">
          <h1 class="text-9xl font-bold text-gray-200 select-none">404</h1>
        </div>
        
        <!-- Error Message -->
        <div class="mb-8">
          <h2 class="text-3xl font-bold text-gray-900 mb-4">Page Not Found</h2>
          <p class="text-xl text-gray-600 max-w-2xl mx-auto">
            Sorry, we couldn't find the page you're looking for. 
            The page might have been moved, deleted, or you entered the wrong URL.
          </p>
        </div>

        <!-- Action Buttons -->
        <div class="flex flex-col sm:flex-row gap-4 justify-center mb-12">
          <router-link 
            to="/" 
            class="inline-flex items-center px-6 py-3 border border-transparent text-base font-medium rounded-md text-white bg-cyan-600 hover:bg-cyan-700 transition-colors"
          >
            <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
            </svg>
            Go Home
          </router-link>
          
          <button 
            @click="goBack" 
            class="inline-flex items-center px-6 py-3 border-2 border-gray-300 text-base font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 transition-colors"
          >
            <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18" />
            </svg>
            Go Back
          </button>
        </div>

        <!-- Helpful Links -->
        <div class="bg-gray-50 rounded-lg p-8 max-w-2xl mx-auto">
          <h3 class="text-lg font-semibold text-gray-900 mb-4">Looking for something specific?</h3>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <router-link 
              to="/" 
              class="flex items-center p-3 text-gray-600 hover:text-cyan-600 hover:bg-white rounded-md transition-colors"
            >
              <svg class="w-5 h-5 mr-3 text-cyan-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
              </svg>
              <div>
                <div class="font-medium">Homepage</div>
                <div class="text-sm text-gray-500">Template features and overview</div>
              </div>
            </router-link>
            
            <router-link 
              to="/about" 
              class="flex items-center p-3 text-gray-600 hover:text-cyan-600 hover:bg-white rounded-md transition-colors"
            >
              <svg class="w-5 h-5 mr-3 text-cyan-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
              <div>
                <div class="font-medium">About</div>
                <div class="text-sm text-gray-500">Learn about the template</div>
              </div>
            </router-link>
            
            <router-link 
              to="/contact" 
              class="flex items-center p-3 text-gray-600 hover:text-cyan-600 hover:bg-white rounded-md transition-colors"
            >
              <svg class="w-5 h-5 mr-3 text-cyan-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 4.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
              </svg>
              <div>
                <div class="font-medium">Contact</div>
                <div class="text-sm text-gray-500">Get in touch with us</div>
              </div>
            </router-link>
            
            <router-link 
              to="/admin" 
              class="flex items-center p-3 text-gray-600 hover:text-cyan-600 hover:bg-white rounded-md transition-colors"
            >
              <svg class="w-5 h-5 mr-3 text-cyan-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
              </svg>
              <div>
                <div class="font-medium">Admin Panel</div>
                <div class="text-sm text-gray-500">Access admin dashboard</div>
              </div>
            </router-link>
          </div>
        </div>
      </div>
    </section>
  </MarketingLayout>
</template>

<script>
import MarketingLayout from '../layouts/MarketingLayout.vue'

export default {
  name: 'NotFoundPage',
  components: {
    MarketingLayout
  },
  methods: {
    goBack() {
      // Go back in history, or to home if no history
      if (window.history.length > 1) {
        this.$router.go(-1)
      } else {
        this.$router.push('/')
      }
    }
  }
}
</script>
