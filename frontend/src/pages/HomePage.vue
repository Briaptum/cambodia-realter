<template>
  <MarketingLayout>
    <!-- Hero Section -->
    <section class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-16">
      <div class="text-center">
        <h1 class="text-4xl font-bold text-gray-900 sm:text-5xl">
          Go-Vue Base Template
        </h1>
        <p class="mt-6 text-xl text-gray-600 max-w-3xl mx-auto">
          A modern, reusable full-stack template with Go + Gin + GORM backend and Vue.js frontend, 
          featuring Auth0 authentication, contact form system, and admin panel.
        </p>
      </div>
    </section>

    <!-- Features Section -->
    <section id="features" class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-16">
      <div class="text-center mb-12">
        <h2 class="text-3xl font-bold text-gray-900 mb-4">Template Features</h2>
        <p class="text-xl text-gray-600 max-w-2xl mx-auto">
          Everything you need to build modern web applications with Go and Vue.js
        </p>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
          <div class="bg-white p-6 rounded-lg shadow-sm border">
            <div class="w-12 h-12 bg-cyan-100 rounded-lg flex items-center justify-center mx-auto mb-4">
              <span class="text-cyan-600 text-xl">🔐</span>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mb-2">Auth0 Integration</h3>
            <p class="text-gray-600">Modern authentication with role-based access control for admins and customers.</p>
          </div>
          
          <div class="bg-white p-6 rounded-lg shadow-sm border">
            <div class="w-12 h-12 bg-cyan-100 rounded-lg flex items-center justify-center mx-auto mb-4">
              <span class="text-cyan-600 text-xl">📧</span>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mb-2">Contact Form System</h3>
            <p class="text-gray-600">SendGrid email notifications with database storage and admin management.</p>
          </div>
          
          <div class="bg-white p-6 rounded-lg shadow-sm border">
            <div class="w-12 h-12 bg-cyan-100 rounded-lg flex items-center justify-center mx-auto mb-4">
              <span class="text-cyan-600 text-xl">👨‍💼</span>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mb-2">Admin Panel</h3>
            <p class="text-gray-600">Clean, responsive interface with ANHELM branding and contact management.</p>
          </div>
          
          <div class="bg-white p-6 rounded-lg shadow-sm border">
            <div class="w-12 h-12 bg-cyan-100 rounded-lg flex items-center justify-center mx-auto mb-4">
              <span class="text-cyan-600 text-xl">🏢</span>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mb-2">Multi-tenant Support</h3>
            <p class="text-gray-600">Site-specific access control via environment variables for multiple projects.</p>
          </div>
          
          <div class="bg-white p-6 rounded-lg shadow-sm border">
            <div class="w-12 h-12 bg-cyan-100 rounded-lg flex items-center justify-center mx-auto mb-4">
              <span class="text-cyan-600 text-xl">🐳</span>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mb-2">Docker Ready</h3>
            <p class="text-gray-600">Complete containerized development environment with hot reload.</p>
          </div>
          
          <div class="bg-white p-6 rounded-lg shadow-sm border">
            <div class="w-12 h-12 bg-cyan-100 rounded-lg flex items-center justify-center mx-auto mb-4">
              <span class="text-cyan-600 text-xl">📱</span>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mb-2">Responsive Design</h3>
            <p class="text-gray-600">Mobile-first approach for all interfaces with modern UI components.</p>
          </div>
      </div>
    </section>

    <!-- Call to Action Section -->
    <section class="bg-cyan-600 py-16">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <h2 class="text-3xl font-bold text-white mb-4">Ready to Get Started?</h2>
        <p class="text-xl text-cyan-100 max-w-2xl mx-auto mb-8">
          Clone the repository and start building your next project with our modern full-stack template.
        </p>
        <div class="flex flex-col sm:flex-row gap-4 justify-center">
          <router-link 
            to="/contact" 
            class="inline-flex items-center px-6 py-3 border border-transparent text-base font-medium rounded-md text-cyan-600 bg-white hover:bg-gray-50 transition-colors"
          >
            Get in Touch
          </router-link>
          <router-link 
            to="/about" 
            class="inline-flex items-center px-6 py-3 border-2 border-white text-base font-medium rounded-md text-white hover:bg-white hover:text-cyan-600 transition-colors"
          >
            Learn More
          </router-link>
        </div>
      </div>
    </section>
  </MarketingLayout>
</template>

<script>
import MarketingLayout from '../layouts/MarketingLayout.vue'

export default {
  name: 'HomePage',
  components: {
    MarketingLayout
  }
}
</script>
