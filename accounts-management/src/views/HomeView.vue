<template>
  <main style="height: 100%">
    <div class="row px-3" style="height: 100%">
      <div class="col-3">
        <CollectionTitle />
      </div>
      <div class="col">
        <CollectionMain
          v-if="Object.keys(usersData).length > 0"
          :users="usersData"
        />
        <loading-circle />
      </div>
      <!-- <div class="col d-flex justify-content-end">
        <span><RouterLink to="/add-user">+ Add User</RouterLink></span>
      </div> -->
    </div>
  </main>
</template>

<script>
import { ref, onMounted } from 'vue';
import { getUsersAPI } from '@/services/api.ts';
import { useStore } from '@/store/useStore';
import CollectionTitle from '@/components/CollectionTitle.vue';
import CollectionMain from '@/components/CollectionMain.vue';
import LoadingCircle from '@/components/minors/LoadingCircle.vue';

export default {
  name: 'HomeView',
  components: { CollectionTitle, CollectionMain, LoadingCircle },
  setup() {
    let usersData = ref([]);
    const store = useStore();

    onMounted(async () => {
      const res = await getUsersAPI();
      usersData.value = res.data;
    });
    return { usersData, store };
  },
};
</script>
