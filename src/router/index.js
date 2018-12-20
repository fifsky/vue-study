import Vue from 'vue'
import Router from 'vue-router'
import Login from '../pages/login'
import List from '../pages/list'
import Layout from '../pages/layout'
import User from '../pages/user'

import {getAccessToken} from "../utils";

Vue.use(Router)

let router = new Router({
  mode: 'history',
  routes: [
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/',
      component: Layout,
      children:[
        {
          path: '',
          name: 'list',
          component: List
        },
        {
          path: 'user',
          name: 'user',
          component: User
        }
      ]
      //   // route level code-splitting
      // // this generates a separate chunk (about.[hash].js) for this route
      // // which is lazy-loaded when the route is visited.
      // component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
    }
  ]
})

router.beforeEach(function (to, from, next) {
  if (!getAccessToken()) {
    if (to.path !== '/login') {
      next("/login")
    } else {
      next()
    }
  } else {
    if (to.path === "/login") {
      next("/list")
    } else {
      next()
    }
  }
})


export default router