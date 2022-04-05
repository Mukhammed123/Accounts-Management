<template>
  <loading-circle v-if="showLoading" />
  <div
    id="myModal"
    ref="loginRef"
    class="modal fade"
    tabindex="-1"
    aria-hidden="true"
    role="dialog"
    data-bs-backdrop="static"
    data-bs-keyboard="false"
  >
    <div class="modal-dialog modal-dialog-centered">
      <div class="modal-content" style="width: 400px">
        <div
          class="modal-header d-flex justify-content-center"
          style="background-color: #056b80"
        >
          <h5 id="modal-title" class="modal-title text-white">
            Accounts Management
          </h5>
        </div>
        <div class="modal-body">
          <form>
            <div class="">
              <label for="Login">Login: </label>
            </div>
            <div>
              <input v-model="login" type="text" class="form-control" />
            </div>
            <label for="Password">Password: </label>
            <br />
            <input v-model="password" type="password" class="form-control" />
          </form>
        </div>
        <div class="modal-footer d-flex justify-content-center">
          <button
            type="button"
            class="btn text-white"
            style="background-color: #056b80"
            @click="submit"
          >
            Submit
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { ref, onMounted, watchEffect } from 'vue';
import { Modal } from 'bootstrap';
import { signInAPI } from '@/services/api';
import { useStore } from '@/store/useStore';
import LoadingCircle from '../minors/LoadingCircle.vue';

export default {
  name: 'LoginDialog',
  components: { LoadingCircle },
  emits: ['show-message'],
  setup(props, context) {
    const store = useStore();
    const loginRef = ref(null);
    const login = ref(null);
    const password = ref(null);
    const showLoading = ref(false);
    let loginModal;

    onMounted(() => {
      loginModal = new Modal(loginRef.value);
      watchEffect(() => {
        if (!store.isSignedIn) {
          login.value = null;
          password.value = null;
          loginModal.show();
        } else loginModal.hide();
      });
    });

    const submit = async () => {
      showLoading.value = true;
      const store = useStore();
      const signInRes = await signInAPI({
        username: login.value ?? '',
        password: password.value ?? '',
      });
      if (signInRes.status < 300 && signInRes.status >= 200) {
        const username = login.value ?? '';
        localStorage.setItem('accessToken', signInRes.data.accessToken);
        localStorage.setItem('username', username);
        store.saveAccessToken(signInRes.data.accessToken);
        store.saveUsername(username);
        store.setIsSignedIn(true);
        if (loginRef.value) {
          loginModal.hide();
        }
      } else {
        context.emit('show-message', {
          content: `Failed to sign in! Status code is ${signInRes.status}`,
          type: 'danger',
        });
      }
      showLoading.value = false;
    };

    return {
      login,
      password,
      loginRef,
      submit,
      store,
      showLoading,
    };
  },
};
</script>
