<template>
  <MarketingLayout>
    <!-- Hero Section -->
    <section class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-16">
      <div class="text-center">
        <h1 class="text-4xl font-bold text-gray-900 sm:text-5xl mb-6">
          Contact Us
        </h1>
        <p class="text-xl text-gray-600 max-w-3xl mx-auto">
          Have questions about the Go-Vue Base template or need help getting started? 
          We'd love to hear from you.
        </p>
      </div>
    </section>

    <!-- Contact Form Section -->
    <section class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 pb-16">
      <div class="max-w-2xl mx-auto">
        <div class="bg-white rounded-lg shadow-sm border p-8">
          <ContactForm />
        </div>
        
        <!-- Additional Info -->
        <div class="mt-12 text-center">
          <h2 class="text-2xl font-bold text-gray-900 mb-6">Other Ways to Reach Us</h2>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
            <div class="bg-gray-50 rounded-lg p-6">
              <h3 class="text-lg font-semibold text-gray-900 mb-2">Documentation</h3>
              <p class="text-gray-600">
                Check out our comprehensive documentation and setup guides for detailed instructions.
              </p>
            </div>
            <div class="bg-gray-50 rounded-lg p-6">
              <h3 class="text-lg font-semibold text-gray-900 mb-2">GitHub Issues</h3>
              <p class="text-gray-600">
                Found a bug or have a feature request? Open an issue on our GitHub repository.
              </p>
            </div>
          </div>
        </div>
      </div>
    </section>
  </MarketingLayout>
</template>

<script>
import MarketingLayout from '../layouts/MarketingLayout.vue'
import ContactForm from '../components/ContactForm.vue'

export default {
  name: 'ContactPage',
  components: {
    MarketingLayout,
    ContactForm
  }
}
</script>
