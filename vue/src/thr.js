import Vue from 'vue'
import thr from './thr.vue'

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#thr',
  render: h => h(thr)
})