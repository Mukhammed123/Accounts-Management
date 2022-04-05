<template>
  <div class="app-body" style="width: 100%; height: 100%">
    <header class="p-2" style="background-color: #056b80">
      <div class="wrapper">
        <nav class="d-flex justify-content-between align-items-center">
          <h2>
            <RouterLink to="/" class="text-white text-decoration-none"
              >Accounts Management</RouterLink
            >
          </h2>
          <div class="text-white">
            {{ store.username }}
            <button
              v-if="route.path !== '/login'"
              class="btn text-white"
              @click="logout"
            >
              Logout
            </button>
          </div>
        </nav>
      </div>
      <user-operation-toast
        :show-message="showMessage"
        :content="toastContent"
        :type="toastType"
        @close="closeToast"
      />
    </header>
    <RouterView
      v-if="store.isSignedIn"
      :key="route.path"
      @show-message="showMessageFunc"
    ></RouterView>
    <login-dialog @show-message="showMessageFunc" />
  </div>
</template>

<script lang="ts">
import { ref } from 'vue';
import { useStore } from '@/store/useStore';
import { useRoute } from 'vue-router';
import LoginDialog from '@/components/dialogs/LoginDialog.vue';
import UserOperationToast from '@/components/toasts/UserOperationToast.vue';

export default {
  name: 'App',
  components: { LoginDialog, UserOperationToast },
  setup() {
    const showMessage = ref(false);
    const toastType = ref(null);
    const toastContent = ref(null);
    const route = useRoute();
    const store = useStore();
    const accessToken = localStorage.getItem('accessToken') || '';

    const logout = () => {
      console.log('logout is called');
      localStorage.removeItem('accessToken');
      localStorage.removeItem('username');
      store.saveAccessToken('');
      store.saveUsername('');
      store.setIsSignedIn(false);
    };

    if (accessToken.length === 0) {
      logout();
    } else {
      const username = localStorage.getItem('username');
      store.saveUsername(username ?? '');
      store.saveAccessToken(accessToken);
      store.setIsSignedIn(true);
    }

    const showMessageFunc = (obj) => {
      showMessage.value = true;
      toastType.value = obj.type;
      toastContent.value = obj.content;
    };

    const closeToast = () => {
      showMessage.value = false;
      toastType.value = null;
      toastContent.value = null;
    };

    return {
      route,
      store,
      showMessage,
      toastType,
      toastContent,
      logout,
      closeToast,
      showMessageFunc,
    };
  },
};
</script>

<style>
@import '@/assets/base.css';
</style>
