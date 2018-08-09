import Vue from 'vue'
import Router from 'vue-router'
import Main from '@/components/Main'
import Shop from '@/components/Shop'
import Discount from '@/components/Discount'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Main',
      component: Main
    },
    {
      path: '/shop',
      name: 'Shop',
      component: Shop
    },
    {
      path: '/discount',
      name: 'Discount',
      component: Discount
    }
  ]
})
