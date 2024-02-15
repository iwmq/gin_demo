<template>
  <div class="w-[80vw] mx-auto my-8">
    <h1 class="text-5xl text-bold">
      Gin Demo
    </h1>
    <div>
      <router-link to="/" class="btn btn-link">Go to Home</router-link>
      <router-link to="/books" class="btn btn-link">Books</router-link>
      <router-link to="/about" class="btn btn-link">About</router-link>
    </div>
    <div>
      <router-view></router-view>
    </div>
  </div>
</template>
