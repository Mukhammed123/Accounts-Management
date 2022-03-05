<template>
    <div :class="`message-block bg-${type} p-2 text-white d-flex justify-content-between align-items-center`">
        <div class="text-container">
          <h4 class="m-0">
            <i class="bi bi-info-circle"></i>
            {{message}}
          </h4 >
        </div>
        <button type="button" class="btn text-white" @click="removeMessage">
          <i class="bi bi-x-lg"></i>
        </button >
    </div>
</template>

<script lang="ts">
export default {
  name: 'MessageBlock',
  props: {
    message: {type: String, default: ''},
    type: {type: String, default: 'danger'}
  },
  emits: ['remove-message'],
  setup(props, context) {
    const removeMessage = () => {
      context.emit('remove-message');
    };
    return {
      removeMessage
    };
  }
    
};
</script >
