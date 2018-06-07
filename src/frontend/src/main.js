// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import Vuetify from 'vuetify'
// [Begin] https://vuetifyjs.com/en/getting-started/quick-start
import 'vuetify/dist/vuetify.min.css'
import 'material-design-icons-iconfont/dist/material-design-icons.css'
// [End]
// [Begin] https://fontawesome.com/how-to-use/use-with-node-js
import fontawesome from '@fortawesome/fontawesome'
import faUser from '@fortawesome/fontawesome-free-solid/faUser'
import faCircle from '@fortawesome/fontawesome-free-regular/faCircle'
import faGithub from '@fortawesome/fontawesome-free-brands/faGithub'

fontawesome.library.add(faUser)
fontawesome.library.add(faCircle)
fontawesome.library.add(faGithub)

Vue.use(fontawesome)
// [End]
// https://vuetifyjs.com/en/getting-started/quick-start
Vue.use(Vuetify)

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
