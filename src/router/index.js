import Vue from 'vue'
import Router from 'vue-router'
import Login from '../pages/login'
import List from '../pages/list'

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
      name: 'list',
      component: List
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
      next("/")
    } else {
      next()
    }
  }
})


export default router