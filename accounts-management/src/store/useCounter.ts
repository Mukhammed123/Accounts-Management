import { defineStore } from 'pinia';

export const useCounter = defineStore('counter', {
    state: () => ({
        name: 'Mukhammed',
        counter: 0,
    }),
    getters: {
        returnDouble(state) {
            return state.counter * 2;
        }
    },
    actions: {
        reset() {
            this.counter = 0;
        },
        addOne() {
            this.counter += 1;
        }
    }
});