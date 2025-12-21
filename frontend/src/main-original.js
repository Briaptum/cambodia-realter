import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router'

// FontAwesome
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { 
  faBars, 
  faTachometerAlt, 
  faEnvelope, 
  faArrowLeft, 
  faTrash, 
  faReply, 
  faPhone,
  faCalendarDay,
  faClock,
  faUserCheck
} from '@fortawesome/free-solid-svg-icons'

// Add icons to library
library.add(
  faBars, 
  faTachometerAlt, 
  faEnvelope, 
  faArrowLeft, 
  faTrash, 
  faReply, 
  faPhone,
  faCalendarDay,
  faClock,
  faUserCheck
)

const app = createApp(App)

app.use(router)
app.component('font-awesome-icon', FontAwesomeIcon)

app.mount('#app')
