<template>
    <div class="collection-main-title py-3 pr-3 pl-0" style="font-size: 1.5rem;">
        Users
    </div>
    <div class="input-group input-group-sm mb-3 p-2" style="background-color: #eaf2fa;">
        <input type="text" class="form-control" placeholder="" aria-label="" aria-describedby="basic-addon1">
        <div class="input-group-prepend">
            <button class="btn btn-outline-secondary" type="button">Search</button>
        </div>
    </div>
    <div class="delete-container">
        <span style="margin-right: 1em;">Action:</span >
        <span style="margin-right: 1em;">
            <select class="custom-select" id="inputGroupSelect01">
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
                <tr v-for="(user) in users" :key="user.idNumber">
                    <td>{{user.username}}</td>
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
        {{users ? Object.keys(users).length : 0}}
    </div>
</template>

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

<script lang="ts">
import { defineComponent, toRefs, toRef } from 'vue';

export default defineComponent({
    name: 'CollectionMain',
    props: ['users'],
    setup(props) {
        // const displayData = toRefs(props)
        // const dsiplayData = toRef(props, 'usersData');
        // console.log(dsiplayData);
        // return {
        //     dsiplayData
        // }
    }
})
</script >
