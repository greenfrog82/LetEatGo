import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import Shop from '@/components/Shop'
import Discount from '@/components/Discount'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'HelloWorld',
      component: HelloWorld
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
    },
  ]
})
