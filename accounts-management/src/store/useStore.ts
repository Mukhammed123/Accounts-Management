import { defineStore } from 'pinia';

export const useStore = defineStore('store', {
  state: () => ({
    accessToken: '',
    username: '',
    isSignedIn: false,
    showToast: true,
    toastContent: '',
    toastType: '',
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
    setShowToast(showToast: boolean) {
      this.showToast = showToast;
    },
    setToastContent(toastContent: string) {
      this.toastContent = toastContent;
    },
    setToastType(toastType: string) {
      this.toastType = toastType;
    },
  }
});