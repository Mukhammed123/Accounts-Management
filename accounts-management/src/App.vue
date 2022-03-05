<template>
  <header class="p-2" style="background-color: #056b80;">
    <div class="wrapper">
      <nav class="d-flex justify-content-between align-items-center">
        <h2><RouterLink to="/accounts-list" class="text-white text-decoration-none">Accounts Management</RouterLink></h2 >
        <div class="text-white">
          {{store.username}}
          <RouterLink v-if="route.path !== '/login'" to="/login" class="text-white" @click="logout">Logout</RouterLink >
        </div>
      </nav>
    </div>
  </header>
  <message-block v-if="messageTrigger" :message="message" :type="type" @remove-message="removeMessage"/>
  <RouterView :key="route.path" @show-message="showMessage"></RouterView >
</template>

<script lang="ts">
import { ref } from 'vue';
import {useStore} from '@/store/useStore';
import { useRoute, useRouter } from 'vue-router';
import MessageBlock from './components/MessageBlock.vue';

export default {
  name: 'App',
  components: { MessageBlock },
  setup() {
    const route = useRoute();
    const store = useStore();
    const router = useRouter();
    const accessToken = localStorage.getItem('accessToken');
    
    const messageTrigger = ref(false);
    const message = ref('');
    const type = ref('danger');

    const showMessage = (obj) => {
      message.value = obj.message;
      type.value = obj.type;
      messageTrigger.value = true;
    };

    const removeMessage = () => {
      messageTrigger.value = false;
    };

    const logout = () => {
      localStorage.removeItem('accessToken');
      store.saveAccessToken('');
      store.saveUsername('');
    };

    if((accessToken || '').length === 0) {
      logout();
      router.push({path: '/login'});
    } else {
      store.saveAccessToken(accessToken);
    }
    return {
      route,
      message,
      type,
      messageTrigger,
      store,
      showMessage,
      removeMessage,
      logout
    };
  },
};


</script>

<style>
@import '@/assets/base.css';
</style>

