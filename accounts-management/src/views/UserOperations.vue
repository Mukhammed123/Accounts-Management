<template>
    <div class="crud-operation-container" style="height: 100%;">
        <div class="row px-3 d-flex justify-content-between">
            <div class="col-3">
                <collection-title />
            </div>
            <div class="col-9">
                <div class="title">User</div>
                <div class="row">
                    <label for="username">Username: </label>
                    <input  v-model="username" :readonly="route.path !== '/add-user'" type="text">
                </div>
                <div class="row">
                    <label for="Id Number:">Id Number: </label>
                    <input v-model="idNumber" :readonly="route.path !== '/add-user'" type="text">
                </div>
                <div v-if="route.path === '/add-user'" class="row">
                    <label for="Password:">Password: </label>
                    <input v-model="password" type="password">
                </div>
                <div class="row">
                    <label for="Full Name:">Full Name: </label>
                    <input v-model="fullName" type="text">
                </div>
                <div class="row">
                    <label for="Role:">Role: </label>
                    <input v-model="role" type="text">
                </div>
                <div class="row">
                    <label for="Email:">Email: </label>
                    <input v-model="email" type="text">
                </div>
            </div>

            <!-- Modal -->
            <div id="exampleModal" class="modal fade" tabindex="-1" role="dialog" aria-labelledby="exampleModalLabel" aria-hidden="true">
                <div class="modal-dialog" role="document">
                    <div class="modal-content">
                    <div class="modal-header">
                        <h5 id="exampleModalLabel" class="modal-title">Modal title</h5>
                        <button type="button" class="close" data-bs-dismiss="modal" aria-label="Close">
                        <span aria-hidden="true">&times;</span>
                        </button>
                    </div>
                    <div class="modal-body">
                        <p>Are you sure you want to delete this user?</p>
                    </div>
                    <div class="modal-footer">
                        <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
                        <button type="button" class="btn btn-primary" data-bs-dismiss="modal" @click="deleteUser">Save changes</button>
                    </div>
                    </div>
                </div>
            </div>
        </div>
        <div  v-if="currentPath !== '/add-user'" class="d-flex justify-content-between mt-3 px-3">
            <button type="button" class="btn btn-danger" data-bs-toggle="modal" data-bs-target="#exampleModal">Delete</button >
            <button type="button" class="btn btn-warning" @click="updateUser">Update</button >
        </div>
        <div v-else class="d-flex justify-content-end mt-3 px-3">
            <button type="button" class="btn text-white" style="background-color: #056b80;" @click="createUser">Create</button  >
        </div>
    </div>
</template>

<script lang="ts">
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import CollectionTitle from '@/components/CollectionTitle.vue';
import { createUsersAPI, getUsersAPI, updateUserAPI, deleteUserAPI } from '@/services/api';
import router from '@/router';

export default {
  name: 'UserOperations',
  components: {CollectionTitle},
  emits: ['show-message'],
  setup(props, context) {
    const route = useRoute();
    const currentPath = route.path;
    const username = ref('');
    const password = ref('');
    const idNumber = ref('');
    const role = ref('');
    const email = ref('');
    const fullName = ref('');

    let getResponse;
    onMounted( () => {
      getUser();
    });

    const getUser = async () => {
      if((route.params.id || '').length > 0) {
        getResponse = await getUsersAPI(route.params.id);
        if(getResponse.status === 200) {
          username.value = getResponse.data.username;
          idNumber.value = getResponse.data.idNumber;
          role.value = getResponse.data.role;
          email.value = getResponse.data.email;
          fullName.value = getResponse.data.fullName;
        }
      }
    };

    const createUser = async () => {
      const response = await createUsersAPI(
        {
          username: username.value,
          password: password.value,
          fullName: fullName.value,
          idNumber: idNumber.value,
          role: role.value
        }
      );
      console.log(response);

      if(response.status < 300 && response.status >= 200) {
        context.emit('show-message', {
          message: `Successfully created user ${username.value}!`,
          type: 'success'
        });
        router.push({path: '/accounts-list'});
      } else {
        context.emit('show-message', {
          message: `Failed to create user ${username.value}! Status code is ${response.status}`,
          type: 'danger'
        });
      }
    };

    const updateUser = async () => {
      const data = {
        username: username.value,
        fullName: fullName.value,
        idNumber: idNumber.value,
        role: role.value
      };
      console.log(data);
      const response = await updateUserAPI(
        data,
        route.params.id
      );
      console.log(response);
      if(response.status < 300 && response.status >= 200) {
        context.emit('show-message', {
          message: `Successfully updated user ${username.value}!`,
          type: 'success'
        });
        getUser();
      } else {
        context.emit('show-message', {
          message: `Failed to update user ${username.value}! Status code is ${response.status}`,
          type: 'danger'
        });
      }
    };

    const deleteUser = async () => {
      const response = await deleteUserAPI(
        route.params.id
      );
      if(response.status < 300 && response.status >= 200) {
        context.emit('show-message', {
          message: `Successfully deleted user ${username.value}!`,
          type: 'success'
        });
        router.push({path: '/accounts-list'});
      } else {
        context.emit('show-message', {
          message: `Failed to delete user ${username.value}! Status code is ${response.status}`,
          type: 'danger'
        });
      }
    };

    return {
      username,
      password,
      idNumber,
      role,
      email,
      fullName,
      currentPath,
      route,
      createUser,
      updateUser,
      deleteUser,
    };
  }
};
</script>
