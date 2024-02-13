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
            class="btn btn-primary w-full max-w-xs"
            :disabled="!data_valid"
            @click="create_book"
        >
            Create
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
import { computed, reactive } from 'vue'
import { useRouter } from 'vue-router'

import { BOOK_API } from '@/composables/constants'

const router = useRouter()


const book_data = reactive({
    title: '',
    author: ''
})

const data_valid = computed(() => {
    return book_data.title.trim() !== '' && book_data.author.trim() !== '';
})

const create_book = () => {
    fetch(BOOK_API, {
        method: "POST",
        headers: { "content-type": "application/json"},
        body: JSON.stringify({
            title: book_data.title,
            author: book_data.author
        })
    }).then(res => {
        router.push({name: "book_home"})
    }).catch(err => {
        console.log(err)
    })
}
</script>