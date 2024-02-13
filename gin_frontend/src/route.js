// NOTE Using default import as following will cause export default error
// in the end of this file
// import  VueRouter from 'vue-router'
import { createRouter, createWebHashHistory } from 'vue-router'

import AppHome from "./components/AppHome.vue"
import HelloWorld from "./components/HelloWorld.vue"
import BookIndex from "./components/book/BookIndex.vue"
import BookCreate from "./components/book/BookCreate.vue"

const routes = [
  { path: '/', component: AppHome },
  { path: '/about', component: HelloWorld },
  { path: '/books', component: BookIndex },
  { path: '/book_create', component: BookCreate },
]


const router = createRouter({
  history: createWebHashHistory(),
  routes,
})


export default router