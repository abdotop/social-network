<template>
  <main id="site__main"
    class="2xl:ml-[--w-side] xl:ml-[--w-side-sm] p-2.5 h-[calc(100vh-var(--m-top))] mt-[--m-top] overflow-y-auto">
    <div class="max-w-[1065] mx-auto bg-slate-500">
      <div class="bg-white shadow lg:rounded-b-2xl lg:-mt-10 dark:bg-dark2">
        <div class="relative overflow-hidden w-full lg:h-72 h-48">
          <img src="assets/images/group-banner.jpg" alt="" class="h-full w-full object-cover inset-0" />
          <div class="w-full bottom-0 absolute left-0 bg-gradient-to-t from-black/60 pt-10 z-10" />
          <div class="absolute bottom-0 right-0 m-a z-20">
            <div class="flex items-center gap-3 px-2">

              <UButton v-if="isMember" :to="`/groups/${id}/chat`">
                Chat
              </UButton>
            </div>
          </div>
        </div>
        <div class="lg:px-10 md:p-5 p-3">
          <div class="flex flex-col justify-center">
            <div class="flex lg:items-center justify-between max-md:flex-col">
              <div class="flex-1">
                <h5 class="md:text-2xl font-bold">
                  {{ group?.Title }}
                </h5>
                <p class="font-normal text-gray-500 mt-2 flex gap-2 flex-wrap dark:text-white/80">
                  <span class="max-lg:hidden"> Private group </span>
                  <span class="max-lg:hidden"> • </span>
                  <span>
                    <b class="font-medium text-black dark:text-white">{{
                group?.GroupMembers.length
              }}</b>
                    member(s)
                  </span>
                </p>
              </div>
              <div>
                <div class="flex items-center gap-2 mt-1">
                  <div class="flex -space-x-4 mr-3">
                    <NuxtImg src="assets/images/avatars/avatar-2.jpg" alt=""
                      class="w-10 rounded-full border-4 border-white dark:border-slate-800" />
                    <NuxtImg src="assets/images/avatars/avatar-3.jpg" alt=""
                      class="w-10 rounded-full border-4 border-white dark:border-slate-800" />
                    <NuxtImg src="assets/images/avatars/avatar-7.jpg" alt=""
                      class="w-10 rounded-full border-4 border-white dark:border-slate-800" />
                    <NuxtImg src="assets/images/avatars/avatar-4.jpg" alt=""
                      class="w-10 rounded-full border-4 border-white dark:border-slate-800" />
                    <NuxtImg src="assets/images/avatars/avatar-5.jpg" alt=""
                      class="w-10 rounded-full border-4 border-white dark:border-slate-800" />
                  </div>
                  <div>
                    <button type="button" class="rounded-lg bg-slate-100 flex px-2.5 py-2 dark:bg-dark2">
                      <UIcon name="i-heroicons-ellipsis-horizontal" class="text-xl" />
                    </button>
                    <div class="w-[100px] shadow-lg"
                      uk-dropdown="pos: bottom-right; animation: uk-animation-scale-up uk-transform-origin-top-right; animate-out: true; mode: click;offset:10">
                      <nav class="bg-slate-200">
                        <a href="#">
                          <UIcon class="text-xl" name="i-heroicons-link" />
                          Copy link
                        </a>
                        <a href="#" class="text-red-400 hover:!bg-red-50 dark:hover:!bg-red-500/50">
                          <UIcon class="text-xl" name="i-heroicons-no-symbol" />
                          Block
                        </a>
                      </nav>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="pl-2 text-sm w-2/3 h-fit max-h-96 overflow-y-auto">
              {{ group?.Description }}
            </div>
          </div>
        </div>
        <div v-if="isMember"
          class="flex items-center justify-between border-t border-gray-100 p-2 dark:border-slate-700">
          <nav>
            <ul id="tabSwitcher"
              class="uk-subnav uk-subnav-pill flex gap-0.5  overflow-hidden -mb-px text-gray-500 font-medium text-sm overflow-x-auto dark:text-white"
              data-uk-switcher="connect: #group-menus ; animation: uk-animation-slide-right-medium, uk-animation-slide-left-medium">
              <li id="posts-tab">
                <a href="#" class="inline-block py-3 leading-8 rounded-md px-3.5">Posts</a>
              </li>
              <li id="events-tab" @click="refresh('')">
                <a href="#" class="inline-block py-3 leading-8 rounded-md px-3.5">Events</a>
              </li>
              <li id="members-tab" @click="refresh('')">
                <a href="#" class="inline-block py-3 leading-8 rounded-md px-3.5">Members</a>
              </li>
              <li id="requests-tab" @click="refresh('')" v-if="group?.CreatorID === user?.id">
                <a href="#" class="inline-block py-3 leading-8 rounded-md px-3.5">Requests</a>
              </li>
            </ul>
          </nav>
          <div class="flex items-center gap-1 text-sm p-3 bg-blue py-2 mr-2 rounded-xl max-md:hidden dark:bg-white/5">
            <UIcon name="i-heroicons-magnifying-glass" class="text-lg" />
            <input placeholder="Search .."
              class="!bg-transparent outline-none focus:outline-none focus:border-b-blue-500" />
          </div>
        </div>
        <div v-else class="text-center">
          <p class="text-xl">You are not part of this group</p>
          <UButton v-if="!isRequester" @click="handleJoin(group)" class="bg-blue-500 m-3">Request to Join
          </UButton>
          <p v-else class="text-blue-500 m-3">Request sent</p>
        </div>
      </div>
      <div id="group-menus" class="uk-switcher flex 2xl:gap-12 gap-10 mt-8 max-lg:flex-col">
        <!-- post tab-->
        <FeedStory style="width: 70vw;" />
        <!-- events tab-->
        <div class="w-full">
          <div class="px-5 flex flex-row justify-between">
            <h5 class="p-0 m-0">Events, {{ events?.length }}</h5>
            <div class="h-fit self-center">
              <UButton uk-toggle="target: #create-event-overlay" @click="createTabOpend = true" class="bg-blue-500">New
              </UButton>
            </div>
          </div>
          <!-- <hr class="border-slate-500 border-2 mb-3" /> -->
          <EventCard v-for="event in events" :event="event" />
        </div>
        <!-- members tab-->
        <div class="w-full">
          <div class=" flex flex-row justify-between px-5">
            <h5>Members: {{ group?.GroupMembers.length }}</h5>
            <div uk-toggle="target: #group-invite-overlay" class="h-fit self-center">
              <UButton class="bg-blue-500">invite</UButton>
            </div>
          </div>
          <hr class="border-slate-500 border-2 mb-3" />
          <div class="flex flex-col w-full gap-2">
            <GroupMemberListItem v-for="member in group?.GroupMembers" :isAdmin="group?.CreatorID === user?.id"
              :member="member" />
          </div>
          <div id="group-invite-overlay" class="hidden lg:p-20" ref="modal" uk-modal="">
            <div
              class=" tt relative overflow-hidden mx-auto bg-white shadow-xl rounded-lg md:w-[520px] w-full dark:bg-slate-800">
              <div class="text-center py-4 border-b mb-0 dark:border-slate-700">
                <h2 class="text-sm font-medium text-black">
                  Invite friends
                </h2>

                <!-- close button -->
                <button type="button" class="button-icon absolute top-0 right-0 m-2.5 uk-modal-close">
                  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                    stroke="currentColor" class="w-6 h-6">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
                  </svg>
                </button>
              </div>
              <div class="w-4/5 h-96 overflow-scroll ">
                <div v-if="!followers">No Friends</div>
                <GroupFollowerListItem v-else-if="group?.GroupMembers.some(member => member.User.id === user?.id)"
                  v-for="user in followers" :user :group-id="group.ID" />
              </div>
            </div>
          </div>
        </div>
        <!-- request tab-->
        <div class="w-full">
          <GroupRequestListItem v-for="member in joinRequests" :member="member" />
        </div>
      </div>
    </div>
    <EventCreateModal v-model="createTabOpend" :opened="createTabOpend" :groupId="group?.ID ?? ''" @event-created="refresh('events')" />
  </main>
</template>
<style scoped>
li.uk-active {
  --tw-text-opacity: 1;
  color: rgb(37 99 235 / var(--tw-text-opacity));
}
</style>
<script lang="ts" setup>
import FeedStory from '~/components/Profile/FeedStory.vue';
import { useTitle } from '@vueuse/core';
import type { Group, GroupMember, Event, Post } from '~/types';
console.log
const uikit = useUIkit3()

definePageMeta({
  alias: ["/groups/[id]"],
  middleware: ["auth-only"],
});

const handleCreation = () => {
  console.log('created');
  
}
const createTabOpend = ref(false)

const { getGroupByID } = useGroups();
const { getJoinRequests, joinRequest } = useGroupRequest()
const { getAllEvents } = useEvents()
const postStore = usePostStore()
const user = useAuthUser()
const route = useRoute();

const group = ref<Group | null>(null);
const joinRequests = ref<GroupMember[] | null>(null)
const isMember = ref(false);
const isRequester = ref(false);
const events = ref<Event[]>([])
const posts = ref<Post[]>()
const title = computed(() => group.value?.Title)

const followers = ref()

useTitle(title, { titleTemplate: "%s | Social Network" })

const id = route.params.id as string;

postStore.getGroupPosts(id)
console.log(postStore.groupPost)


async function handleJoin(group: Group | null) {
  const { data } = await joinRequest(group?.ID)
  if (data) {
    isRequester.value = true;
  }
}

onMounted(async () => {

  uikit?.switcher("#tabSwitcher")?.show(2)

  group.value = await getGroupByID(id);
  if (group.value) {
    isMember.value = group.value?.GroupMembers.some(
      (member) => member.User.id === user.value?.id
    );
  }

  // await postStore.getGroupFeeds(group.value?.ID)
  // posts.value = postStore.groupPosts as Post[]

  joinRequests.value = await getJoinRequests(id) || []
  console.log("posts /////////// \n", posts.value);

  const { data, pending, error } = await getAllEvents(id)

  events.value = data.value.data

  console.log("events /////////// \n", events.value);

  isRequester.value =
    joinRequests.value?.some((member) => member.User.id === user.value?.id) || false;

  followers.value = await getFollowers()


})

async function refresh(tab: String) {
  if (tab === 'events') {
    const { data } = await getAllEvents(id)
    console.log(data);

    events.value = data.value.data
  } else {
    group.value = await getGroupByID(id)
  }
}




</script>
