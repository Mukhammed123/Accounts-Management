<template>
  <div>
    <h1>hello world</h1>
    <div
      class="toast"
      role="alert"
      aria-live="assertive"
      aria-atomic="true"
      autohide="false"
    >
      <div class="toast-header">
        <img src="..." class="rounded mr-2" alt="..." />
        <strong class="mr-auto">Bootstrap</strong>
        <small>11 mins ago</small>
        <button
          type="button"
          class="ml-2 mb-1 close"
          data-dismiss="toast"
          aria-label="Close"
        >
          <span aria-hidden="true">&times;</span>
        </button>
      </div>
      <div class="toast-body">Hello, world! This is a toast message.</div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'UserOperationToast',
  setup() {
    console.log('Toast is created');
  },
});
</script>
