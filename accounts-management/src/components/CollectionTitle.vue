<template>
  <div class="collection-title-container  ml-3">
    <div class="title-container mt-1 text-warning" style="background-color: #056b80;">
        <h6 class="px-2 m-0">Accounts</h6 >
    </div>
    <div class="list-container">
        <ul class="list-group">
        <li class="list-group-item d-flex justify-content-between px-1">
            <div>Users</div >
            <div><button>Add</button ></div >
        </li >
        </ul>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
export default defineComponent({
    name: 'CollectionTitle',
    setup() {
        console.log("Hello from Collection Title")
    }
}); 
</script >