<template>
  <div id="app">
    <router-view @logout="handleLogout" @login-success="handleLoginSuccess" />
  </div>
</template>

<script>
export default {
  name: 'App',
  methods: {
    handleLogout() {
      // Clear authentication data
      localStorage.removeItem('auth_token')
      localStorage.removeItem('user')
      
      // Redirect to login
      this.$router.push('/login')
    },
    handleLoginSuccess(data) {
      // Store authentication data
      localStorage.setItem('auth_token', data.token)
      localStorage.setItem('user', JSON.stringify(data.user))
    }
  }
}
</script>
