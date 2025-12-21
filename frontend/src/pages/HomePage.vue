<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Header -->
    <header class="bg-white shadow-sm">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center py-6">
          <!-- ANHELM Logo -->
          <div class="flex items-center">
            <div class="w-8 flex flex-col items-center justify-center mr-2 mb-0.5">
              <!-- Top two dots -->
              <div class="flex gap-1.5 mb-0.5">
                <div class="w-1.5 h-1.5 bg-cyan-400 rounded-full"></div>
                <div class="w-1.5 h-1.5 bg-cyan-400 rounded-full"></div>
              </div>
              <!-- Bottom dot -->
              <div class="w-1.5 h-1.5 bg-cyan-400 rounded-full"></div>
            </div>
            <h1 class="text-2xl font-light text-zinc-700 tracking-wider leading-tight">ANHELM</h1>
          </div>
          
          <!-- Admin Link -->
          <router-link 
            to="/admin" 
            class="text-sm text-cyan-600 hover:text-cyan-800 font-medium"
          >
            Admin Dashboard
          </router-link>
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <main class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
      <div class="text-center">
        <h1 class="text-4xl font-bold text-gray-900 sm:text-5xl">
          Go-Vue Base Template
        </h1>
        <p class="mt-6 text-xl text-gray-600 max-w-3xl mx-auto">
          A modern, reusable full-stack template with Go + Gin + GORM backend and Vue.js frontend, 
          featuring Auth0 authentication, contact form system, and admin panel.
        </p>
        
        <!-- Features Grid -->
        <div class="mt-16 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
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

        <!-- Contact Form Section -->
        <div class="mt-16 bg-white rounded-lg shadow-sm border p-8">
          <h2 class="text-2xl font-bold text-gray-900 mb-6">Test Contact Form</h2>
          <div class="max-w-md mx-auto">
            <ContactForm />
          </div>
        </div>
      </div>
    </main>

    <!-- Footer -->
    <footer class="bg-white border-t">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div class="flex items-center justify-center">
          <div class="w-5 flex flex-col items-center justify-center mr-1 mb-0.5">
            <!-- Top two dots -->
            <div class="flex gap-1 mb-0.5">
              <div class="w-1 h-1 bg-cyan-400 rounded-full"></div>
              <div class="w-1 h-1 bg-cyan-400 rounded-full"></div>
            </div>
            <!-- Bottom dot -->
            <div class="w-1 h-1 bg-cyan-400 rounded-full"></div>
          </div>
          <span class="text-sm font-light text-zinc-500 tracking-wider leading-tight">ANHELM</span>
        </div>
      </div>
    </footer>
  </div>
</template>

<script>
import ContactForm from '../components/ContactForm.vue'

export default {
  name: 'HomePage',
  components: {
    ContactForm
  }
}
</script>
