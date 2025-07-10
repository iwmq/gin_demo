<template>
    <div class="w-full md:w-1/2 flex flex-col gap-4">
        <div class="flex flex-col gap-2">
            <label class="input w-full">
                <span class="label">Book Title?</span>
                <input
                    type="text"
                    placeholder="Book Title?"
                    v-model="book_data.title"
                />
            </label>
            <label class="input w-full">
                <span class="label">Author Name?</span>
                <input
                    type="text"
                    placeholder="Author Name?"
                    v-model="book_data.author"
                />
            </label>
        </div>

        <div class="flex flex-row gap-4">
            <button
                class="btn btn-primary"
                :disabled="!data_valid"
                @click="create_book"
            >
                Create
            </button>
            <router-link
                class="btn btn-neutral"
                to="/books"
            >
                Return
            </router-link>
        </div>
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