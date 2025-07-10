<template>
    <div>
        <table class="table w-full max-w-xl">
            <thead>
                <tr>
                    <th>Title</th>
                    <th>Author</th>
                    <th></th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="b in books" :key="b.id">
                    <td>{{ b.title }}</td>
                    <td>{{ b.author }}</td>
                    <td>
                        <router-link
                            class="btn btn-sm btn-warning"
                            :to="{name: 'book_edit', params: {id: b.id}}"
                        >
                            Edit
                        </router-link>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>

    <div class="mt-8">
        <router-link
            to="/book_create"
            class="btn btn-primary"
        >
            New
        </router-link>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

import { BOOK_API } from '@/composables/constants'


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