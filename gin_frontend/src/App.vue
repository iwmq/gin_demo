<script setup>
import HelloWorld from './components/HelloWorld.vue'
</script>

<template>
  <div class="w-[80vw] mx-auto my-8">
    <h1 class="text-5xl text-bold">
      Gin Demo
    </h1>
    <div>
      <router-link to="/" class="btn btn-link">Go to Home</router-link>
      <router-link to="/about" class="btn btn-link">About</router-link>
    </div>
    <div>
      <router-view></router-view>
    </div>
  </div>
</template>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
