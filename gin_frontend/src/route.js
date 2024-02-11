// NOTE Using default import as following will cause export default error
// in the end of this file
// import  VueRouter from 'vue-router'
import { createRouter, createWebHashHistory } from 'vue-router'

import AppHome from "./components/AppHome.vue"
import HelloWorld from "./components/HelloWorld.vue"


const routes = [
  { path: '/', component: AppHome },
  { path: '/about', component: HelloWorld },
]


const router = createRouter({
  history: createWebHashHistory(),
  routes,
})


export default router