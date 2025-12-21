<template>
  <div class="flex flex-col h-screen bg-gray-50">
    <!-- Top Bar -->
    <nav class="bg-gray-50 border-b border-gray-200 shadow-sm">
      <div class="px-4 sm:px-6 flex items-center justify-between h-12">
        <!-- Left: Hamburger + Logo -->
        <div class="flex items-center">
          <!-- Mobile Hamburger Menu -->
          <button
            @click="toggleSidebar"
            class="lg:hidden mr-3 p-2 text-zinc-600 hover:text-zinc-900 hover:bg-gray-100 rounded-md transition-colors"
            aria-label="Toggle menu"
          >
            <font-awesome-icon icon="bars" class="h-5 w-5" />
          </button>
          
          <!-- ANHELM Logo -->
          <div class="flex items-center">
            <div class="w-8 sm:w-10 flex flex-col items-center justify-center mr-1 mb-0.5">
              <!-- Top two dots -->
              <div class="flex gap-1 sm:gap-1.5 mb-0.5">
                <div class="w-1 sm:w-1.5 h-1 sm:h-1.5 bg-cyan-400 rounded-full"></div>
                <div class="w-1 sm:w-1.5 h-1 sm:h-1.5 bg-cyan-400 rounded-full"></div>
              </div>
              <!-- Bottom dot -->
              <div class="w-1 sm:w-1.5 h-1 sm:h-1.5 bg-cyan-400 rounded-full"></div>
            </div>
            <h1 class="text-xl sm:text-2xl font-light text-zinc-700 tracking-wider leading-tight">ANHELM</h1>
          </div>
        </div>
      </div>
    </nav>

    <div class="flex flex-1 overflow-hidden relative">
      <!-- Mobile Overlay -->
      <div
        v-if="sidebarOpen"
        @click="closeSidebar"
        class="lg:hidden fixed inset-0 bg-black bg-opacity-50 z-40"
      ></div>

      <!-- Left Sidebar -->
      <nav 
        :class="[
          'fixed lg:static inset-y-0 left-0 z-50 w-64 bg-gray-50 shadow-sm border-r border-gray-200 flex flex-col transform transition-transform duration-300 ease-in-out',
          sidebarOpen ? 'translate-x-0' : '-translate-x-full lg:translate-x-0'
        ]"
      >
        <!-- Navigation Menu -->
        <div class="flex-1 py-4 overflow-y-auto">
          <nav class="space-y-px">
            <router-link 
              to="/admin/dashboard"
              @click="closeSidebarOnMobile"
              :class="$route.name === 'dashboard' ? 'bg-white text-zinc-700 group flex items-center px-4 sm:px-6 py-3 text-sm font-medium border-l-4 border-cyan-500 transition-all duration-200 ease-in-out' : 'text-zinc-600 hover:bg-white hover:text-zinc-700 hover:border-l-4 hover:border-cyan-500 group flex items-center px-4 sm:px-6 py-3 text-sm font-medium border-l-4 border-transparent transition-all duration-200 ease-in-out'"
            >
              <font-awesome-icon icon="tachometer-alt" class="mr-3 h-5 w-5 text-cyan-500 transition-colors duration-200" />
              Dashboard
            </router-link>
            <router-link 
              to="/admin/contact-requests"
              @click="closeSidebarOnMobile"
              :class="['contact-requests', 'contact-request-detail'].includes($route.name) ? 'bg-white text-zinc-700 group flex items-center px-4 sm:px-6 py-3 text-sm font-medium border-l-4 border-cyan-500 transition-all duration-200 ease-in-out' : 'text-zinc-600 hover:bg-white hover:text-zinc-700 hover:border-l-4 hover:border-cyan-500 group flex items-center px-4 sm:px-6 py-3 text-sm font-medium border-l-4 border-transparent transition-all duration-200 ease-in-out'"
            >
              <font-awesome-icon icon="envelope" class="mr-3 h-5 w-5 text-cyan-500 transition-colors duration-200" />
              Contact Requests
            </router-link>
          </nav>
        </div>
        
        <!-- User Section -->
        <div class="border-t border-gray-200">
          <div class="p-4">
            <div class="flex items-center">
              <div class="flex-shrink-0">
                <div class="w-8 h-8 bg-cyan-100 rounded-full flex items-center justify-center">
                  <span class="text-sm font-medium text-cyan-600">{{ (user.email || 'U').charAt(0).toUpperCase() }}</span>
                </div>
              </div>
              <div class="ml-3 flex-1 min-w-0">
                <p class="text-sm font-medium text-zinc-700 truncate">{{ user.email || 'User' }}</p>
                <button
                  @click="$emit('logout')"
                  class="text-xs text-zinc-500 hover:text-zinc-700"
                >
                  Sign out
                </button>
              </div>
            </div>
          </div>
          
          <!-- ANHELM Watermark -->
          <div class="flex items-center justify-center pt-3 pb-4 border-t border-gray-200 px-4">
            <div class="w-5 flex flex-col items-center justify-center mr-1 mb-0.5">
              <!-- Top two dots -->
              <div class="flex gap-1 mb-0.5">
                <div class="w-1 h-1 bg-cyan-400 rounded-full"></div>
                <div class="w-1 h-1 bg-cyan-400 rounded-full"></div>
              </div>
              <!-- Bottom dot -->
              <div class="w-1 h-1 bg-cyan-400 rounded-full"></div>
            </div>
            <span class="text-xs font-light text-zinc-500 tracking-wider leading-tight">ANHELM</span>
          </div>
        </div>
      </nav>

      <!-- Main Content Area -->
      <div class="flex-1 flex flex-col overflow-hidden">
        <main class="flex-1 bg-white overflow-y-auto">
          <div class="py-4 sm:py-6">
            <div class="px-4 sm:px-6">
              <slot />
            </div>
          </div>
        </main>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'AdminLayout',
  props: {
    user: {
      type: Object,
      required: true
    }
  },
  data() {
    return {
      sidebarOpen: false
    }
  },
  methods: {
    toggleSidebar() {
      this.sidebarOpen = !this.sidebarOpen
    },
    closeSidebar() {
      this.sidebarOpen = false
    },
    closeSidebarOnMobile() {
      // Close sidebar on mobile after navigation
      if (window.innerWidth < 1024) {
        this.closeSidebar()
      }
    }
  },
  mounted() {
    // Close sidebar when clicking outside on mobile
    const handleResize = () => {
      if (window.innerWidth >= 1024) {
        this.sidebarOpen = false
      }
    }
    window.addEventListener('resize', handleResize)
    this.$options._resizeHandler = handleResize
  },
  beforeUnmount() {
    if (this.$options._resizeHandler) {
      window.removeEventListener('resize', this.$options._resizeHandler)
    }
  }
}
</script>
