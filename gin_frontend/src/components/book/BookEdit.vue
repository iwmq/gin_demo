<template>
    <div class="flex flex-col gap-4">
        <label class="form-control w-full max-w-xs">
            <div class="label">
                <span class="label-text">Book Title?</span>
            </div>
            <input
                type="text"
                placeholder="Type here"
                class="input input-bordered input-info w-full max-w-xs"
                v-model="book_data.title"
            />
        </label>
        <label class="form-control w-full max-w-xs">
            <div class="label">
                <span class="label-text">Author Name?</span>
            </div>
            <input
                type="text"
                placeholder="Type here"
                class="input input-bordered input-info w-full max-w-xs"
                v-model="book_data.author"
            />
        </label>
        <button
            class="btn btn-warning w-full max-w-xs"
            :disabled="!data_valid"
            @click="update_book"
        >
            Update
        </button>
        <button
            class="btn btn-error w-full max-w-xs"
            @click="delete_book"
        >
            Delete
        </button>
        <router-link
            class="btn btn-neutral w-full max-w-xs"
            to="/books"
        >
            Return
        </router-link>
    </div>
</template>

<script setup>
import { computed, onMounted, reactive, ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'

import { BOOK_API } from '@/composables/constants'

const router = useRouter()
const route = useRoute()


const book_data = reactive({
    title: '',
    author: ''
})

const data_valid = computed(() => {
    return book_data.title.trim() !== '' && book_data.author.trim() !== '';
})

const url = ref("")
const set_url = () => {
    url.value = `${BOOK_API}/${route.params.id}`
}

const get_detail = async () => {
    fetch(url.value)
    .then(res => res.json())
    .then(data => {
        book_data.title = data.data.title
        book_data.author = data.data.author
    }).catch(err => {
        console.log(err)
    })
}

const update_book = async () => {
    fetch(url.value, {
        method: "PATCH",
        headers: { "content-type": "application/json"},
        body: JSON.stringify({
            title: book_data.title,
            author: book_data.author
        })
    }).then(function(res) {
        router.push({name: "book_home"})
    }).catch(err => {
        console.log(err)
    })
}

const delete_book = async () => {
    fetch(url.value, {
        method: "DELETE",
        headers: { "content-type": "application/json"}
    }).then(function(res) {
        router.push({name: "book_home"})
    }).catch(err => {
        console.log(err)
    })
}

onMounted(async () => {
    set_url()
    await get_detail()
})
</script>