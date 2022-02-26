<template>
  <main>
    <div class="row px-3">
      <div class="col">
        <CollectionTitle />
      </div>
      <div class="col-6">
        <CollectionMain />
      </div>
      <div class="col">
        <button>+ Add User</button>
      </div>
    </div>
  </main >
</template>


<script>
import axios from 'axios';
import { signInAPI } from '@/services/api';
import CollectionTitle from '@/components/CollectionTitle.vue';
import CollectionMain from '@/components/CollectionMain.vue';


export default {
  name: 'HomeView',
  components: {CollectionTitle, CollectionMain},
}
</script>

