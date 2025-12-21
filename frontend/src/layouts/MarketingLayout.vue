<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Header -->
    <header class="bg-white shadow-sm sticky top-0 z-50">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center py-6">
          <!-- ANHELM Logo -->
          <router-link to="/" class="flex items-center">
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
          </router-link>
          
          <!-- Navigation -->
          <nav class="hidden md:flex items-center space-x-8">
            <router-link 
              to="/" 
              class="text-sm text-gray-600 hover:text-gray-900 font-medium transition-colors"
              :class="{ 'text-cyan-600': $route.path === '/' }"
            >
              Home
            </router-link>
            <router-link 
              to="/about" 
              class="text-sm text-gray-600 hover:text-gray-900 font-medium transition-colors"
              :class="{ 'text-cyan-600': $route.path === '/about' }"
            >
              About
            </router-link>
            <router-link 
              to="/contact" 
              class="text-sm text-gray-600 hover:text-gray-900 font-medium transition-colors"
              :class="{ 'text-cyan-600': $route.path === '/contact' }"
            >
              Contact
            </router-link>
            <router-link 
              to="/admin" 
              class="text-sm text-cyan-600 hover:text-cyan-800 font-medium transition-colors"
            >
              Admin
            </router-link>
          </nav>

          <!-- Mobile menu button -->
          <button 
            @click="mobileMenuOpen = !mobileMenuOpen"
            class="md:hidden p-2 rounded-md text-gray-400 hover:text-gray-500 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-cyan-500"
          >
            <span class="sr-only">Open main menu</span>
            <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
            </svg>
          </button>
        </div>

        <!-- Mobile menu -->
        <div v-show="mobileMenuOpen" class="md:hidden pb-6">
          <div class="space-y-1">
            <router-link 
              to="/" 
              @click="mobileMenuOpen = false"
              class="block px-3 py-2 text-base font-medium text-gray-600 hover:text-gray-900 hover:bg-gray-50 rounded-md transition-colors"
              :class="{ 'text-cyan-600 bg-cyan-50': $route.path === '/' }"
            >
              Home
            </router-link>
            <router-link 
              to="/about" 
              @click="mobileMenuOpen = false"
              class="block px-3 py-2 text-base font-medium text-gray-600 hover:text-gray-900 hover:bg-gray-50 rounded-md transition-colors"
              :class="{ 'text-cyan-600 bg-cyan-50': $route.path === '/about' }"
            >
              About
            </router-link>
            <router-link 
              to="/contact" 
              @click="mobileMenuOpen = false"
              class="block px-3 py-2 text-base font-medium text-gray-600 hover:text-gray-900 hover:bg-gray-50 rounded-md transition-colors"
              :class="{ 'text-cyan-600 bg-cyan-50': $route.path === '/contact' }"
            >
              Contact
            </router-link>
            <router-link 
              to="/admin" 
              @click="mobileMenuOpen = false"
              class="block px-3 py-2 text-base font-medium text-cyan-600 hover:text-cyan-800 hover:bg-cyan-50 rounded-md transition-colors"
            >
              Admin
            </router-link>
          </div>
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <main class="flex-1">
      <slot />
    </main>

    <!-- Footer -->
    <footer class="bg-white border-t mt-16">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
        <div class="grid grid-cols-1 md:grid-cols-4 gap-8">
          <!-- Company Info -->
          <div class="col-span-1 md:col-span-2">
            <div class="flex items-center mb-4">
              <div class="w-6 flex flex-col items-center justify-center mr-2 mb-0.5">
                <!-- Top two dots -->
                <div class="flex gap-1 mb-0.5">
                  <div class="w-1.5 h-1.5 bg-cyan-400 rounded-full"></div>
                  <div class="w-1.5 h-1.5 bg-cyan-400 rounded-full"></div>
                </div>
                <!-- Bottom dot -->
                <div class="w-1.5 h-1.5 bg-cyan-400 rounded-full"></div>
              </div>
              <span class="text-xl font-light text-zinc-700 tracking-wider leading-tight">ANHELM</span>
            </div>
            <p class="text-gray-600 text-sm max-w-md">
              Modern full-stack template with Go + Gin + GORM backend and Vue.js frontend, 
              featuring Auth0 authentication and admin panel.
            </p>
          </div>

          <!-- Quick Links -->
          <div>
            <h3 class="text-sm font-semibold text-gray-900 tracking-wider uppercase mb-4">Quick Links</h3>
            <ul class="space-y-2">
              <li>
                <router-link to="/" class="text-sm text-gray-600 hover:text-gray-900 transition-colors">
                  Home
                </router-link>
              </li>
              <li>
                <router-link to="/about" class="text-sm text-gray-600 hover:text-gray-900 transition-colors">
                  About
                </router-link>
              </li>
              <li>
                <router-link to="/contact" class="text-sm text-gray-600 hover:text-gray-900 transition-colors">
                  Contact
                </router-link>
              </li>
            </ul>
          </div>

          <!-- Admin -->
          <div>
            <h3 class="text-sm font-semibold text-gray-900 tracking-wider uppercase mb-4">Admin</h3>
            <ul class="space-y-2">
              <li>
                <router-link to="/admin" class="text-sm text-gray-600 hover:text-gray-900 transition-colors">
                  Dashboard
                </router-link>
              </li>
              <li>
                <router-link to="/login" class="text-sm text-gray-600 hover:text-gray-900 transition-colors">
                  Login
                </router-link>
              </li>
            </ul>
          </div>
        </div>

        <!-- Bottom Footer -->
        <div class="mt-8 pt-8 border-t border-gray-200">
          <div class="flex flex-col md:flex-row justify-between items-center">
            <p class="text-sm text-gray-500">
              © {{ currentYear }} ANHELM. Built with Go-Vue Base Template.
            </p>
            <div class="mt-4 md:mt-0">
              <p class="text-xs text-gray-400">
                Powered by Go + Vue.js + Auth0
              </p>
            </div>
          </div>
        </div>
      </div>
    </footer>
  </div>
</template>

<script>
export default {
  name: 'MarketingLayout',
  data() {
    return {
      mobileMenuOpen: false
    }
  },
  computed: {
    currentYear() {
      return new Date().getFullYear()
    }
  },
  mounted() {
    // Close mobile menu when clicking outside
    document.addEventListener('click', (e) => {
      if (!this.$el.contains(e.target)) {
        this.mobileMenuOpen = false
      }
    })
  }
}
</script>
