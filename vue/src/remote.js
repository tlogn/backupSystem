import Vue from 'vue'
import remote from './remote.vue'

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#remote',
  render: h => h(remote)
})