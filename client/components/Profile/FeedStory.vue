<template>
    <div class="flex 2xl:gap-12 gap-10 mt-8 max-lg:flex-col" id="js-oversized">

        <div class="flex-1 xl:space-y-6 space-y-3">
            <div class="md:max-w-[580px] mx-auto flex-1 xl:space-y-6 space-y-3">

                <div v-if="!potentialUser || root.includes('groups')"
                    class="bg-white rounded-xl shadow-sm md:p-4 p-2 space-y-4 text-sm font-medium border1 dark:bg-dark2">

                    <div class="flex items-center md:gap-3 gap-1">

                        <div class="flex-1 bg-slate-100 hover:bg-opacity-80 transition-all rounded-lg cursor-pointer dark:bg-dark3"
                            uk-toggle="target: #create-status">
                            <div class="py-2.5 text-center dark:text-white"> What do you have in mind? </div>
                        </div>
                        <div class="cursor-pointer hover:bg-opacity-80 p-1 px-1.5 rounded-xl transition-all bg-pink-100/60 hover:bg-pink-100 dark:bg-white/10 dark:hover:bg-white/20"
                            uk-toggle="target: #create-status">
                            <svg xmlns="http://www.w3.org/2000/svg" class="w-8 h-8 stroke-pink-600 fill-pink-200/70"
                                viewBox="0 0 24 24" stroke-width="1.5" stroke="#2c3e50" fill="none"
                                stroke-linecap="round" stroke-linejoin="round">
                                <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                                <path d="M15 8h.01" />
                                <path d="M12 3c7.2 0 9 1.8 9 9s-1.8 9 -9 9s-9 -1.8 -9 -9s1.8 -9 9 -9z" />
                                <path d="M3.5 15.5l4.5 -4.5c.928 -.893 2.072 -.893 3 0l5 5" />
                                <path d="M14 14l1 -1c.928 -.893 2.072 -.893 3 0l2.5 2.5" />
                            </svg>
                        </div>

                    </div>

                </div>

                <post-card-img
                    v-for="post in 
                    potentialUser ? postStore.getUserPosts(potentialUser) :
                    !potentielGroupId? postStore.userPosts: postStore.groupPost[potentielGroupId]"
                    :post="post" />



            </div>
        </div>
        <!-- posts component  -->

        <Detail v-if="status == 'about'" :data="data" />

    </div>
</template>

<script setup>

import usePostStore from "~/stores/usePostStore.js";
let root = location.pathname
let potentialUser = !root.includes('profile')? null: root.split('/')[2]
let potentielGroupId  =  !root.includes('groups')? null :root.split('/')[2]
console.log(potentielGroupId, potentialUser);


const postStore = usePostStore()
// console.log(Nickname)
// onMounted(async () => {
//     await postStore.getUserFeed()
// });

const props = defineProps({
    status: {
        type: String,
        required: true
    },
    data: {
        type: Object,
        required: true
    }
})
</script>