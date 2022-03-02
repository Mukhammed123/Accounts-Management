<template>
  <div class="modal d-flex align-items-center" style="display: block;" tabindex="-1" aria-hidden="false" aria-labelledby="modal-title">
    <div class="modal-dialog" role="document">
      <div class="modal-content" style="width: 400px;">
        <div class="modal-header d-flex justify-content-center" style="background-color: #056b80;">
          <h5 id="modal-title" class="modal-title text-white">Accounts Management</h5 >
        </div>
        <div class="modal-body">
          <form>
            <div class="">
              <label for="Login">Login: </label>
            </div>
            <div>
              <input v-model="login" type="text" class="form-control">
            </div>
            <label for="Password">Password: </label>
            <br>
            <input v-model="password" type="password" class="form-control">
          </form>
        </div>
        <div class="modal-footer d-flex justify-content-center">
          <button type="button" class="btn text-white" style="background-color: #056b80;" @click="myEvent">Save changes</button>
        </div>
      </div>
    </div>
  </div >
  <button @click="myEvent">Hey</button>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { signInAPI } from '@/services/api';
import {useStore} from '@/store/useStore';

const login = ref(null);
const password = ref(null);
const router = useRouter();
const myEvent = async () => {
  const store = useStore();
  const signInRes = await signInAPI({username: login.value, password: password.value});
  if(signInRes.status < 300 && signInRes.status >= 200) {
    localStorage.setItem('accessToken', signInRes.data.accessToken);
    store.saveAccessToken(signInRes.data.accessToken);
    router.push({path: '/accounts-list'});
  }
}; 

</script >