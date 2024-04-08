<template>
  <main id="site__main" class="ml-48 2xl:ml-0 py-5 px-6 h-[calc(100vh-var(--m-top))] mt-[--m-top] overflow-y-auto">
    <div class="max-w-[680px] mx-auto">
      <div class="md:max-w-[580px] mx-auto flex-1 xl:space-y-6 space-y-3">
        <div class="bg-white rounded-xl shadow-sm md:p-4 p-2 space-y-4 text-sm font-medium border1 dark:bg-dark2">
          <div class="flex items-center md:gap-3 gap-1">

            <div class="flex-1 bg-slate-100 hover:bg-opacity-80 transition-all rounded-lg cursor-pointer dark:bg-dark3"
              uk-toggle="target: #create-status">
              <div class="py-2.5 text-center dark:text-white"> What do you have in mind? </div>
            </div>
            <div
              class="cursor-pointer hover:bg-opacity-80 p-1 px-1.5 rounded-xl transition-all bg-pink-100/60 hover:bg-pink-100 dark:bg-white/10 dark:hover:bg-white/20"
              uk-toggle="target: #create-status">
              <svg xmlns="http://www.w3.org/2000/svg" class="w-8 h-8 stroke-pink-600 fill-pink-200/70"
                viewBox="0 0 24 24" stroke-width="1.5" stroke="#2c3e50" fill="none" stroke-linecap="round"
                stroke-linejoin="round">
                <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                <path d="M15 8h.01" />
                <path d="M12 3c7.2 0 9 1.8 9 9s-1.8 9 -9 9s-9 -1.8 -9 -9s1.8 -9 9 -9z" />
                <path d="M3.5 15.5l4.5 -4.5c.928 -.893 2.072 -.893 3 0l5 5" />
                <path d="M14 14l1 -1c.928 -.893 2.072 -.893 3 0l2.5 2.5" />
              </svg>
            </div>
          </div>
        </div>
        <post-card-img v-for="post in postStore.posts" :post="post" />
      </div>
    </div>
  </main>
  <Post-text-preview />
</template>

<script setup lang="ts">

import usePostStore from "~/stores/usePostStore.js";
const postStore = usePostStore()
const currentUser = useAuthUser();


useHead({
  title: "Home",
})
onMounted(() => {
  postStore.getUserFeed()
})

definePageMeta({
  alias: ["/"],
  middleware: ["auth-only"],
});
</script>