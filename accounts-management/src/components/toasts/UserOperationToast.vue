<template>
  <!-- Then put toasts within -->
  <div class="d-flex justify-content-center">
    <div class="position-absolute">
      <div
        ref="toastRef"
        class="toast"
        role="alert"
        aria-live="assertive"
        aria-atomic="true"
      >
        <div :class="`toast-header bg-${type} text-white`">
          <strong class="me-auto">{{
            type === 'success' ? 'Success' : 'Error'
          }}</strong>
          <!-- <small>11 mins ago</small> -->
          <button
            type="button"
            class="btn-close"
            data-bs-dismiss="toast"
            aria-label="Close"
          ></button>
        </div>
        <div class="toast-body bg-white">
          {{ content }}
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref, toRef, watchEffect } from 'vue';
import { Toast } from 'bootstrap';

export default defineComponent({
  name: 'UserOperationToast',
  props: {
    showMessage: { type: Boolean, default: false },
    content: { type: String, default: '' },
    type: { type: String, default: '' },
  },
  setup(props, context) {
    console.log('Toast is created');
    const toastRef = ref(null);
    const showMessage = toRef(props, 'showMessage');
    onMounted(() => {
      if (toastRef.value) {
        const toast = new Toast(toastRef.value);

        watchEffect(() => {
          if (showMessage.value) toast.show();
          else toast.hide();

          toastRef.value.addEventListener('hidden.bs.toast', () => {
            context.emit('close');
          });
        });
      }
    });

    return {
      toastRef,
    };
  },
});
</script>

<style scoped>
.toast {
  z-index: 1057;
}
</style>
