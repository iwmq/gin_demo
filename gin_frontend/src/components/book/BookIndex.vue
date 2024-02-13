<template>
    <div>
        <table class="table w-full max-w-xs">
            <thead>
                <tr>
                    <th>Title</th>
                    <th>Author</th>
                    <th></th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="b in books">
                    <td>{{ b.title }}</td>
                    <td>{{ b.author }}</td>
                    <td></td>
                </tr>
            </tbody>
        </table>
    </div>
    <router-link to="/book_create" class="btn btn-link">
        New
    </router-link>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const BOOK_API = "/books/api"

const books = ref([])

const get_all = async () => {
    fetch(BOOK_API)
    .then(res => res.json())
    .then(data => {
        books.value = data.data
    }).catch(err => {
        console.log(err)
    })
}

onMounted(async () => {
 await get_all()
})
</script>