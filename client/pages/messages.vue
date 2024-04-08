<template>
  <main id="site__main" class="2xl:ml-[--w-side]  xl:ml-[--w-side-sm] p-2.5 h-[calc(100vh-var(--m-top))] mt-[--m-top]">
    <nuxt-child />
    <div class="relative overflow-hidden border -m-2.5 dark:border-slate-700">
      <div class="flex bg-white dark:bg-dark2">
        <!-- sidebar -->
        <div class="md:w-[360px] relative border-r dark:border-slate-700">
          <div id="side-chat"
            class="top-0 left-0 max-md:fixed max-md:w-5/6 max-md:h-screen bg-white z-50 max-md:shadow max-md:-translate-x-full dark:bg-dark2">
            <!-- heading title -->
            <div class="p-4 border-b dark:border-slate-700">
              <div class="flex mt-2 items-center justify-between">
                <h2 class="text-2xl font-bold text-black ml-1 dark:text-white">
                  Chats
                </h2>

                <!-- right action buttons -->
                <div class="flex items-center gap-2.5">
                  <button class="group">
                    <i class='bx bx-cog text-2xl flex group-aria-expanded:rotate-180'></i>
                  </button>
                  <div class="md:w-[270px] w-full"
                    uk-dropdown="pos: bottom-left; offset:10; animation: uk-animation-slide-bottom-small">
                    <nav>
                      <a href="#"> <i class='bx bx-check text-2xl shrink-0 -ml-1'></i>
                        Mark all as read </a>
                      <a href="#"> <i class='bx bx-bell text-2xl shrink-0 -ml-1'></i>
                        notifications setting </a>
                      <a href="#"> <i class='bx bx-volume-mute text-xl shrink-0 -ml-1'></i>
                        Mute notifications </a>
                    </nav>
                  </div>

                  <button class="">
                    <i class='bx bx-check-circle text-2xl flex'></i>
                  </button>

                  <!-- mobile toggle menu -->
                  <button type="button" class="md:hidden"
                    uk-toggle="target: #side-chat ; cls: max-md:-translate-x-full">
                    <i class='bx bx-chevron-down'></i>
                  </button>
                </div>
              </div>

              <!-- search -->
              <div class="relative mt-4">
                <div class="absolute left-3 bottom-1/2 translate-y-1/2 flex">
                  <i class='bx bx-search text-xl'></i>
                </div>
                <input type="text" placeholder="Search" class="w-full !pl-10 !py-2 !rounded-lg">
              </div>
            </div>


            <!-- users list -->
            <div v-for="user in users" :key="user.id" class="space-y-2 p-2 overflow-y-auto">

              <a @click="selectUser(user.id)" href="#"
                class="relative flex items-center gap-4 p-2 duration-200 rounded-xl hover:bg-secondery">
                <div class="relative w-14 h-14 shrink-0">
                  <img :src="'http://localhost:8081/' + user.avatar" alt=""
                    class="object-cover w-full h-full rounded-full">
                </div>
                <div class="flex-1 min-w-0">
                  <div class="flex items-center gap-2 mb-1.5">
                    <div class="mr-auto text-sm text-black dark:text-white font-medium">{{ user.firstName + " "
              + user.lastName }}</div>
                    <div v-if="user.lastMessageTime" class="text-xs font-light text-gray-500 dark:text-white/70">{{
              formatTimeAgo(user.lastMessageTime) }}</div>
                    <div :style="user.online === true ? 'background-color: green;' : ''"
                      class="w-2.5 h-2.5 bg-blue-600 rounded-full dark:bg-slate-700"></div>
                  </div>
                  <div style="padding-left: 0.5rem;"
                    class="font-medium overflow-hidden text-ellipsis text-gray-400 text-sm whitespace-nowrap">{{
              user.lastMessage }}</div>
                </div>
              </a>

            </div>

          </div>

          <!-- overly -->
          <div id="side-chat"
            class="bg-slate-100/40 backdrop-blur w-full h-full dark:bg-slate-800/40 z-40 fixed inset-0 max-md:-translate-x-full md:hidden"
            uk-toggle="target: #side-chat ; cls: max-md:-translate-x-full"></div>

        </div>

        <!-- message center -->
        <div class="flex-1">

          <!-- chat heading -->
          <div
            class="flex items-center justify-between gap-2 w- px-6 py-3.5 z-10 border-b dark:border-slate-700 uk-animation-slide-top-medium">

              <div class="flex items-center sm:gap-4 gap-2">
                <!-- toggle for mobile -->
                <button type="button" class="md:hidden" uk-toggle="target: #side-chat ; cls: max-md:-translate-x-full">
                  <i class='bx bx-chevron-left text-2xl -ml-4'></i>
                </button>

              <div class="relative cursor-pointer max-md:hidden" uk-toggle="target: .rightt ; cls: hidden">
                <img :src="'http://localhost:8081/' + selectedUser.avatar" alt="" class="w-8 h-8 rounded-full shadow">
                <div v-if="selectedUser.online" class="w-2 h-2 bg-teal-500 rounded-full absolute right-0 bottom-0 m-px">
                </div>
              </div>
              <div class="cursor-pointer" uk-toggle="target: .rightt ; cls: hidden">
                <div class="text-base font-bold"> {{ selectedUser.firstName + " " + selectedUser.lastName }} </div>
                <div v-if="selectedUser.online" class="text-xs text-green-500 font-semibold"> Online</div>
                <div v-if="!selectedUser.online" class="text-xs text-red-500 font-semibold"> Offline</div>
              </div>

            </div>

            <!-- <div class="flex items-center gap-2">
              <button type="button" class="button__ico">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" class="w-6 h-6">
                  <path fill-rule="evenodd"
                    d="M2 3.5A1.5 1.5 0 013.5 2h1.148a1.5 1.5 0 011.465 1.175l.716 3.223a1.5 1.5 0 01-1.052 1.767l-.933.267c-.41.117-.643.555-.48.95a11.542 11.542 0 006.254 6.254c.395.163.833-.07.95-.48l.267-.933a1.5 1.5 0 011.767-1.052l3.223.716A1.5 1.5 0 0118 15.352V16.5a1.5 1.5 0 01-1.5 1.5H15c-1.149 0-2.263-.15-3.326-.43A13.022 13.022 0 012.43 8.326 13.019 13.019 0 012 5V3.5z"
                    clip-rule="evenodd" />
                </svg>
              </button>
              <button type="button" class="hover:bg-slate-100 p-1.5 rounded-full">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                  stroke="currentColor" class="w-6 h-6">
                  <path stroke-linecap="round"
                    d="M15.75 10.5l4.72-4.72a.75.75 0 011.28.53v11.38a.75.75 0 01-1.28.53l-4.72-4.72M4.5 18.75h9a2.25 2.25 0 002.25-2.25v-9a2.25 2.25 0 00-2.25-2.25h-9A2.25 2.25 0 002.25 7.5v9a2.25 2.25 0 002.25 2.25z" />
                </svg>
              </button>
              <button type="button" class="hover:bg-slate-100 p-1.5 rounded-full"
                uk-toggle="target: .rightt ; cls: hidden">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                  stroke="currentColor" class="w-6 h-6">
                  <path stroke-linecap="round" stroke-linejoin="round"
                    d="M11.25 11.25l.041-.02a.75.75 0 011.063.852l-.708 2.836a.75.75 0 001.063.853l.041-.021M21 12a9 9 0 11-18 0 9 9 0 0118 0zm-9-3.75h.008v.008H12V8.25z" />
                </svg>
              </button>
            </div> -->

          </div>

          <!-- chats bubble -->
          <div id="chatsbubble" class="w-full p-5 py-10 overflow-y-auto md:h-[calc(100vh-204px)] h-[calc(100vh-195px)]">

            <div class="py-10 text-center text-sm lg:pt-8">
              <img :src="'http://localhost:8081/' + selectedUser.avatar" class="w-24 h-24 rounded-full mx-auto mb-3"
                alt="">
              <div class="mt-8">
                <div class="md:text-xl text-base font-medium text-black dark:text-white">
                  {{ selectedUser.firstName + " "
              + selectedUser.lastName }} </div>
                <div class="text-gray-500 text-sm   dark:text-white/80"> @{{ selectedUser.nickName }} </div>
              </div>
              <div class="mt-3.5">
                <a :href="'/profile/' + selectedUser.nickName"
                  class="inline-block rounded-lg px-4 py-1.5 text-sm font-semibold bg-secondery">View profile</a>
              </div>
            </div>

            <div v-for="message in messages" :key="message.id" class="text-sm font-medium space-y-6">
              <!-- time -->
              <!-- <div :key="message.id" v-if="message.created_at.getDate() > lastDate.getDate()" class="flex justify-center ">
                <div class="font-medium text-gray-500 text-sm dark:text-white/70">
                  {{changeDate(message!.created_at)}}
                </div>
              </div> -->
              <!-- received -->
              <div style="margin-top: 10px;" v-if="message.user.id === selectedUser.id" class="flex gap-3 items-end">
                <img :src="'http://localhost:8081/' + message.user.avatar" alt="" class="w-5 h-5 rounded-full shadow">
                <div class="px-4 py-2 rounded-[20px] max-w-sm bg-secondery"> {{ message.message }} </div>
              </div>
              <!-- sent -->
              <div style="margin-top: 10px;" v-if="message.user.id === currentUser!.id"
                class="flex  gap-2 flex-row-reverse items-end">
                <img :src="'http://localhost:8081/' + message.user.avatar" alt="" class="w-5 h-5 rounded-full shadow">
                <div
                  class="px-4 py-2 rounded-[20px] max-w-sm bg-gradient-to-tr from-sky-500 to-blue-500 text-white shadow">
                  {{ message.message }}</div>
              </div>



            </div>

          </div>

          <!-- sending message area -->
          <div class="flex items-center md:gap-4 gap-2 md:p-3 p-2 overflow-hidden">

            <div id="message__wrap" class="flex items-center gap-2 h-full dark:text-white -mt-1.5">
              <button type="button" class="shrink-0">
                <i style="font-size: 30px;" class='bx bx-happy-alt'></i>
              </button>
              <div class="dropbar p-2"
                uk-drop="stretch: x; target: #message__wrap ;animation: uk-animation-scale-up uk-transform-origin-bottom-left ;animate-out: true; pos: top-left ; offset:2; mode: click ; duration: 200 ">

                <div class="sm:w-60 bg-white shadow-lg border rounded-xl  pr-0 dark:border-slate-700 dark:bg-dark3">
                  <span class="text-sm font-semibold p-3 pb-0" style="display: flex;justify-content: space-evenly;">
                    <button @click="changeCategory(-1)"><i style="font-size: 20px;" class='bx bx-chevron-left'></i></button>
                    <h4 class="">{{ currentCategory }}</h4>
                    <button @click="changeCategory(1)"><i style="font-size: 20px;" class='bx bx-chevron-right'></i></button>
                  </span>
                  <div class="grid grid-cols-5 overflow-y-auto max-h-44 p-3 text-center text-xl">
                    <div v-for="(emoji, index) in emojis[currentCategory as keyof typeof emojis]" :key="index"
                      @click="addEmoji(emoji, index)"
                      class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"> {{ emoji
                      }}
                    </div>
                  </div>


                </div>

              </div>

            </div>

            <div class="relative flex-1">
              <textarea v-model="message" placeholder="Write your message" rows="1"
                class="w-full resize-none bg-secondery rounded-full px-4 p-2"></textarea>

              <button @click="send" type="button" class="text-white shrink-0 p-2 absolute right-0.5 top-0">
                <i style="font-size: 24px;" class='bx bx-send'></i>
              </button>

            </div>

          </div>

        </div>

        <!-- user profile right info -->
        <div class="rightt w-full h-full absolute top-0 right-0 z-10 hidden transition-transform">
          <div
            class="w-[360px] border-l shadow-lg h-screen bg-white absolute right-0 top-0 uk-animation-slide-right-medium delay-200 z-50 dark:bg-dark2 dark:border-slate-700">

            <div class="w-full h-1.5 bg-gradient-to-r to-purple-500 via-red-500 from-pink-500 -mt-px"></div>

            <div class="py-10 text-center text-sm pt-20">
              <img :src="'http://localhost:8081/' + selectedUser.avatar" class="w-24 h-24 rounded-full mx-auto mb-3"
                alt="">
              <div class="mt-8">
                <div class="md:text-xl text-base font-medium text-black dark:text-white"> {{ selectedUser.firstName +
              " "
              + selectedUser.lastName }} </div>
                <div class="text-gray-500 text-sm mt-1 dark:text-white/80">@{{ selectedUser.nickName }}</div>
              </div>
              <div class="mt-5">
                <a :href="'/profile/' + selectedUser.nickName"
                  class="inline-block rounded-full px-4 py-1.5 text-sm font-semibold bg-secondery">View profile</a>
              </div>
            </div>

            <hr class="opacity-80 dark:border-slate-700">

            <!-- <ul class="text-base font-medium p-3">
              <li>
                <div class="flex items-center gap-5 rounded-md p-3 w-full hover:bg-secondery">
                  <ion-icon name="notifications-off-outline" class="text-2xl"></ion-icon> Mute Notification
                  <label class="switch cursor-pointer ml-auto"> <input type="checkbox" checked><span
                      class="switch-button !relative"></span></label>
                </div>
              </li>
              <li> <button type="button" class="flex items-center gap-5 rounded-md p-3 w-full hover:bg-secondery">
                  <ion-icon name="flag-outline" class="text-2xl"></ion-icon> Report </button></li>
              <li> <button type="button" class="flex items-center gap-5 rounded-md p-3 w-full hover:bg-secondery">
                  <ion-icon name="settings-outline" class="text-2xl"></ion-icon> Ignore messages </button> </li>
              <li> <button type="button" class="flex items-center gap-5 rounded-md p-3 w-full hover:bg-secondery">
                  <ion-icon name="stop-circle-outline" class="text-2xl"></ion-icon> Block </button> </li>
              <li> <button type="button"
                  class="flex items-center gap-5 rounded-md p-3 w-full hover:bg-red-50 text-red-500"> <ion-icon
                    name="trash-outline" class="text-2xl"></ion-icon> Delete Chat </button> </li>
            </ul> -->

            <!-- close button -->
            <button type="button" class="absolute top-0 right-0 m-4 p-2 bg-secondery rounded-full"
              uk-toggle="target: .rightt ; cls: hidden">
              <i style="font-size: 24px;" class='bx bx-x'></i>
            </button>

          </div>

          <!-- overly -->
          <div class="bg-slate-100/40 backdrop-blur absolute w-full h-full dark:bg-slate-800/40"
            uk-toggle="target: .rightt ; cls: hidden"></div>

        </div>

      </div>

    </div>

  </main>
</template>

<script setup lang="ts">
import { formatDate, formatTimeAgo } from '@vueuse/core';
import { getSelectUserMessage, getUsers, users, messages, scroll } from '~/composables/getMessage';
import { emojis } from "~/composables/emoji/emoji"
const currentUser = useAuthUser();
let lastDate = new Date(0);
const message = ref<string>("");
const categorysEmoji = Object.keys(emojis)
const currentCategory = ref<string>("")
currentCategory.value = categorysEmoji[0]

// circle direction change infinte loop
const changeCategory = (direction: number) => {
  let index = categorysEmoji.indexOf(currentCategory.value)
  if (direction === 1) {
    if (index === categorysEmoji.length - 1) {
      index = 0
    } else {
      index++
    }
  } else {
    if (index === 0) {
      index = categorysEmoji.length - 1
    } else {
      index--
    }
  }
  currentCategory.value = categorysEmoji[index]
}

const changeDate = (date: Date) => {
  lastDate = date
  const d = formatDate(date, 'MMMM d,yyyy,h:mm a')
  return d
}

const addEmoji = (emoji: string, index: number) => {
  message.value += emoji
}

const send = () => {
  if (message.value.trim() !== "") {
    sendMessage(currentUser!.value!.id, message.value)
  }
  message.value = "";
};

useHead({
  title: "Message",
})

definePageMeta({
  alias: ["/messages"],
  middleware: ["auth-only"],
});

const selectUser = async (receiverId: string) => {
  const res = await getSelectUserMessage(currentUser!.value!, receiverId)
  if (res) {
    scroll()
  }
}

onMounted(async () => {
  const res = await getUsers(currentUser!.value!)
  if (res) {
    scroll()
  }
});
</script>