<template>
  <header class="p-2" style="background-color: #056b80;">
    <div class="wrapper">
      <nav>
        <RouterLink to="/">Home</RouterLink>
      </nav>
    </div>
  </header>
  <RouterView></RouterView>
</template>

<style>
@import '@/assets/base.css';
</style>

<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router';

const isLoggedIn : boolean = (() => {
  try {
    return JSON.parse(localStorage.getItem('isLoggedIn') ?? 'false') === true;
  } catch {
    return false;
  }
})();

let accessToken : string;

</script >
