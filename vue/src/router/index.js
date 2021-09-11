import Vue from 'vue'
import Router from 'vue-router'
import one from './../one.vue'
import two from './../two.vue'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/one',
      component: one
    },
    {
      path: '/two',
      component: two
    }
  ]
})

