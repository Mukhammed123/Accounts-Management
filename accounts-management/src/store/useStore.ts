import { defineStore } from 'pinia';

export const useStore = defineStore('store', {
  state: () => ({
    accessToken: '',
    username: '',
    isSignedIn: false,
  }),
  actions: {
    reset() {},
    saveAccessToken(accToken: string) {
      this.accessToken = accToken;
    },
    saveUsername(name: string) {
      this.username = name;
    },
    setIsSignedIn(signedIn: boolean) {
      this.isSignedIn = signedIn;
    },
  }
});