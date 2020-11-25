// Today (26 Nov 2020) vuejs 3 exists, but this project is done for version 2.
// This is because version 3 is not mature enough in terms of examples, configuration problems solving, etc
// I prefer to wait before using vuejs 3

import Vue from 'vue'
import App from './App.vue'

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')
