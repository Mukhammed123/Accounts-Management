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
    <div class="table-wrapper-scroll-x my-custom-scrollbar">
        <table class="table table-striped">
            <thead>
                <tr>
                    <th scope="col">#</th>
                    <th scope="col">username</th>
                    <th scope="col">password</th>
                    <th scope="col">full name</th>
                    <th scope="col">id number</th>
                    <th scope="col">role</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <th>1</th>
                    <td>head_nurse</td>
                    <td>password</td>
                    <td>Head Nurse</td>
                    <td>A123456789</td>
                    <td>head nurse</td>
                </tr>
            </tbody>
        </table>
    </div>
    <div>
        total users
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
import { defineComponent } from 'vue';

export default defineComponent({
    name: 'CollectionMain',
    setup() {}
})
</script >
