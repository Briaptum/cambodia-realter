<template>
  <MarketingLayout>
    <!-- Hero Section -->
    <section class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-16">
      <div class="text-center">
        <h1 class="text-4xl font-bold text-gray-900 sm:text-5xl mb-6">
          About Go-Vue Base
        </h1>
        <p class="text-xl text-gray-600 max-w-3xl mx-auto">
          A comprehensive full-stack template designed to accelerate your web development projects
          with modern technologies and best practices.
        </p>
      </div>
    </section>

    <!-- Content Section -->
    <section class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-16">
      <div class="prose prose-lg mx-auto">
        <h2>Why Go-Vue Base?</h2>
        <p>
          Building modern web applications requires setting up authentication, database connections,
          API endpoints, and frontend components. Go-Vue Base eliminates this repetitive setup by
          providing a production-ready template with all the essential features pre-configured.
        </p>

        <h2>Technology Stack</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-8 not-prose">
          <div>
            <h3 class="text-lg font-semibold text-gray-900 mb-4">Backend</h3>
            <ul class="space-y-2 text-gray-600">
              <li>• Go 1.22+ with Gin framework</li>
              <li>• GORM for database operations</li>
              <li>• PostgreSQL database</li>
              <li>• Auth0 integration</li>
              <li>• SendGrid email service</li>
              <li>• Docker containerization</li>
            </ul>
          </div>
          <div>
            <h3 class="text-lg font-semibold text-gray-900 mb-4">Frontend</h3>
            <ul class="space-y-2 text-gray-600">
              <li>• Vue.js 3 with Composition API</li>
              <li>• Vue Router for navigation</li>
              <li>• Tailwind CSS for styling</li>
              <li>• FontAwesome icons</li>
              <li>• Axios for API calls</li>
              <li>• Vite for fast development</li>
            </ul>
          </div>
        </div>

        <h2>Getting Started</h2>
        <p>
          Clone the repository, configure your environment variables, and you're ready to build.
          The template includes comprehensive documentation and example implementations to guide
          your development process.
        </p>
      </div>
    </section>
  </MarketingLayout>
</template>

<script>
import MarketingLayout from '../layouts/MarketingLayout.vue'

export default {
  name: 'AboutPage',
  components: {
    MarketingLayout
  }
}
</script>
