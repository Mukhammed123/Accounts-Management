
<template>
    <div class="collection-main-title py-3 pr-3 pl-0" style="font-size: 1.5rem;">
        Users
    </div>
    <div class="mb-3 p-2" style="background-color: #eaf2fa;">
        <input v-model="search" type="text" class="form-control">
    </div>
    <div class="delete-container">
        <span style="margin-right: 1em;">Action:</span >
        <span style="margin-right: 1em;">
            <select id="inputGroupSelect01" class="custom-select">
            <option selected>Choose...</option>
            <option value="1">Users</option>
        </select>
        </span >
        <span style="margin-right: 1em; padding: .1em">
            <button type="button" class="btn btn-secondary py-0 px-1">delete</button >
        </span>
        <span>Info</span>
    </div>
    <div class="table-wrapper-scroll-x my-custom-scrollbar" style="height: auto;">
        <table class="table table-striped">
            <thead>
                <tr>
                    <th scope="col">username</th>
                    <th scope="col">id number</th>
                    <th scope="col">full name</th>
                    <th scope="col">role</th>
                    <th scope="col">email</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(user, userId) in displayData" :key="user.idNumber">
                    <td><RouterLink :to="`/user-detail/${userId}`" :user="user">{{user.username}}</RouterLink ></td >
                    <td>{{user.idNumber}}</td>
                    <td>{{user.fullName}}</td>
                    <td>{{user.role}}</td>
                    <td>{{user.email}}</td>

                </tr >
            </tbody>
        </table>
    </div>
    <div>
        total users
        {{displayData ? Object.keys(displayData).length : 0}}
    </div>
</template>

<script lang="ts">
import { defineComponent, ref, toRef, computed } from 'vue';

export default defineComponent({
  name: 'CollectionMain',
  props: {
    users: { type: Object, default: null}
  },
  setup(props) {
    const search = ref('');
    const data = toRef(props, 'users');

    const displayData = computed(() => {
      return Object.keys(data.value).reduce((result, currentVal) => {
        if (Object.keys(data.value[currentVal]).some((userKey) => {
          return data.value[currentVal][userKey].includes(search.value);
        }))
          result[currentVal] = data.value[currentVal];
        return result;
      }, {});
    });

    return {
      search,
      displayData
    };
  }
});
</script>

<style>
.my-custom-scrollbar {
position: relative;
height: 200px;
overflow: auto;
}
.table-wrapper-scroll-y {
display: block;
}
</style>
 
