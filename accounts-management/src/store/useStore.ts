import { defineStore } from 'pinia';

export const useStore = defineStore('store', {
  state: () => ({
    accessToken: '',
  }),
  getters: {},
  actions: {
    reset() {},
    saveAccessToken(accToken :string) {
      this.accessToken = accToken;
    }
  }
});