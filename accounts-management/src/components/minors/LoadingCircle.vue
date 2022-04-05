<template>
  <div class="overlay d-flex justify-content-center align-items-center">
    <div class="">
      <div
        class="spinner-border text-light"
        role="status"
        style="width: 3rem; height: 3rem; z-index: 20; opacity: 1"
      ></div>
    </div>
  </div>
</template>

<script>
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'LoadingCircle',
  setup() {},
});
</script>

<style scoped>
.overlay {
  position: fixed;
  width: 100%;
  height: 100%;
  z-index: 1056;
  left: 0;
  opacity: 0.6;
  background-color: rgb(41, 40, 40);
  filter: alpha(opacity=50);
}
</style>
