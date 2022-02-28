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
            <button>Delete</button>
            <button @click="createUser">Create</button >
            <button>Update</button>
        </div>
    </div>
</template>

<script lang="ts">
import { ref } from 'vue';
import CollectionTitle from '@/components/CollectionTitle.vue';
import { createUsersAPI } from '@/services/api';

export default {
    name: "UserOperations",
    components: {CollectionTitle},
    setup() {
        const username = ref('');
        const password = ref('');
        const idNumber = ref('');
        const role = ref('');
        const email = ref('');
        const fullName = ref('');

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
        return {
            username,
            password,
            idNumber,
            role,
            email,
            fullName,
            createUser
        }
    }
}
</script>
