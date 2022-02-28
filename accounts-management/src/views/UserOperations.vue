<template>
    <div class="crud-operation-container">
        <div class="row px-3" style="height: 100%;">
            <div class="col">
                <CollectionTitle />
            </div>
            <div class="col-7">
                <div class="title">User</div>
                <div class="row">
                    <label for="username">Username: </label>
                    <input  v-model="username" type="text">
                </div>
                <div class="row">
                    <label for="Id Number:">Id Number: </label>
                    <input v-model="idNumber" type="text">
                </div>
                <div class="row">
                    <label for="Password:">Password: </label>
                    <input v-model="password" type="text">
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
            <button v-if="currentPath !== '/add-user'" @click="deleteUser">Delete</button >
            <button v-else @click="createUser">Create</button  >
            <button v-if="currentPath !== '/add-user'" @click="updateUser">Update</button >
        </div>
    </div>
</template>

<script lang="ts">
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import CollectionTitle from '@/components/CollectionTitle.vue';
import { createUsersAPI, getUsersAPI, updateUserAPI, deleteUserAPI } from '@/services/api';

export default {
    name: "UserOperations",
    components: {CollectionTitle},
    setup() {
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
            getUsers();
        });

        const getUsers = async () => {
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
        }

        console.log(getResponse);

        const createUser = async () => {
            const response = await createUsersAPI(
                {
                    username: username.value,
                    password: password.value,
                    fullName: fullName.value,
                    idNumber: idNumber.value,
                    role: role.value
                }
            )
            console.log(response)
        }

        const updateUser = async () => {
            const data = {
                    username: username.value,
                    fullName: fullName.value,
                    idNumber: idNumber.value,
                    role: role.value
                }
                console.log(data);
            const response = await updateUserAPI(
                data,
                route.params.id
            )
            console.log(response)
            if(response.status < 300 && response.status >= 200) {
                getUsers();
            }
        }

        const deleteUser = async () => {
            const response = await deleteUserAPI(
                route.params.id
            )
            console.log(response)
            if(response.status < 300 && response.status >= 200) {
                getUsers();
            }
        }

        return {
            username,
            password,
            idNumber,
            role,
            email,
            fullName,
            currentPath,
            createUser,
            updateUser,
            deleteUser
        }
    }
}
</script>
