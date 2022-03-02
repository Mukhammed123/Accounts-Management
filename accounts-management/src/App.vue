<template>
  <header class="p-2" style="background-color: #056b80;">
    <div class="wrapper">
      <nav>
        <RouterLink to="/accounts-list">Home</RouterLink>
      </nav>
    </div>
  </header>
  <RouterView></RouterView>
</template>

<style>
@import '@/assets/base.css';
</style>

<script lang="ts">
import { RouterLink, RouterView, useRouter } from 'vue-router';

export default {
  name: 'App',
  setup() {
    const router = useRouter();
    console.log(localStorage.getItem('accessToken'))
    const isLoggedIn : boolean = (() => {
    try {
      return (localStorage.getItem('accessToken') || '').length > 0;
    } catch {
      return false;
    }
  })();

  if(!isLoggedIn) {
    router.push({path: '/login'});
  } 
  },
}


</script >
