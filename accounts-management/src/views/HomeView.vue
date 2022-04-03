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
        <div v-else class="col-6">Loading...</div>
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
import CollectionTitle from '@/components/CollectionTitle.vue';
import CollectionMain from '@/components/CollectionMain.vue';

export default {
  name: 'HomeView',
  components: { CollectionTitle, CollectionMain },
  setup() {
    let usersData = ref([]);

    onMounted(async () => {
      const res = await getUsersAPI();
      usersData.value = res.data;
    });
    return { usersData };
  },
};
</script>
