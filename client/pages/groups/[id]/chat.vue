<template>
  <main id="site__main"
    class="2xl:ml-[--w-side] xl:ml-[--w-side-sm] p-2.5 h-[calc(100vh-var(--m-top))] mt-[--m-top] overflow-y-auto">
    <div class="h-screen flex flex-col justify-between">
      <div class="page-heading">
        <h1>🗨️Group Chat: </h1>
        <div class="flex flex-col" v-if="group">
          {{ group.Title }} : {{ group.Description }}
          <div v-if="messagesList">
            <div v-for="message in messagesList" :key="message.ID" class="text-sm font-medium m-3 space-y-6">
              <div>
                <div v-if="message.SenderID === currentUser?.id" class="flex gap-2 flex-row-reverse items-end">
                  <img v-if="currentUser && currentUser.avatarImage"
                    :src="'http://localhost:8081/' + currentUser.avatarImage" class="w-9 h-9 rounded-full shadow" />
                  <div
                    class="px-4 py-2 rounded-[20px] max-w-sm bg-gradient-to-tr from-sky-500 to-blue-500 text-white shadow">
                    {{ message.Content }}
                  </div>
                </div>
                <div v-else class="flex gap-3">
                  <img :src="'http://localhost:8081/' + message.Sender.avatarImage" class="w-9 h-9 rounded-full shadow" />
                  <div class="px-4 py-2 rounded-[20px] max-w-sm bg-secondery">
                    {{ message.Content }}
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="flex w-full items-center justify-center">
          <div id="message__wrap" class="flex items-center gap-2 dark:text-white -mt-1.5">
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
          <u-input class="w-full" v-model="message" placeholder="Type your message..." @keydown.enter="send" />
          <u-button @click="send" class="bg-blue-500">Send</u-button>
        </div>
        <nuxt-link :to="`/groups/${groupId}`" class="bg-red-500">Back Group</nuxt-link>
      </div>
    </div>
  </main>
</template>

<script lang="ts" setup>
import { emojis } from "~/composables/emoji/emoji";
import type { Group, GroupMessage } from '~/types';

const currentUser = useAuthUser();
const userId = currentUser.value?.id

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


const addEmoji = (emoji: string, index: number) => {
  message.value += emoji
}

const { getGroupByID, getAllMessagesByGroup } = useGroups();
const route = useRoute();
const groupId = route.params.id as string;
const group = ref<Group | null>(null);
const message = ref<string>("");
const messagesList = ref<GroupMessage[]>([]);
let ws: WebSocket | undefined;

useHead({
  title: "Chat " + group.value ? group.value?.Title : '',
})

definePageMeta({
  alias: ["/groups/:id/chat"],
  middleware: ["auth-only"],
});

const scroll = () => {
  nextTick(() => {
    console.log('scrooling')
    window.scrollTo(0, document.body.scrollHeight + 100);
  })
}

const log = (user: string, message: GroupMessage | string) => {
  if (typeof message === "string") {
    console.log("[ws]", user, message);
  } else {
    console.log("[ws]", user, message.Content);
    if (!message.CreatedAt) {
      message.CreatedAt = new Date().toISOString();
    }
    messagesList.value = [...messagesList.value, message];
    scroll();
  }
};

const send = () => {
  console.log("sending message...");
  if (message.value) {
    ws!.send(message.value);
  }
  message.value = "";
};

const connect = async () => {
  const url = "ws://" + location.host + "/api/groups/chat-ws?group_id=" + groupId;
  if (ws) {
    log("ws", "Closing previous connection before reconnecting...");
    ws.close();
  }

  log("ws", "Connecting to" + url + "...");
  ws = new WebSocket(url);

  ws.addEventListener("message", (event) => {
    const message = (event.data.startsWith("{")
      ? JSON.parse(event.data)
      : event.data) as GroupMessage;
    log(
      message.Sender.nickname,
      message,
    );
  });

  await new Promise((resolve) => ws!.addEventListener("open", resolve));
  log("ws", "Connected!");
};


onMounted(async () => {
  group.value = await getGroupByID(groupId)


  if (group.value) {
    const isMember = group.value?.GroupMembers.some(
      (member) => member.User.id === userId
    );
    if (!isMember) {
      navigateTo('/groups/' + groupId)
      return
    }
  }
  await connect();
  const _messages = await getAllMessagesByGroup(groupId)
  if (_messages) {
    messagesList.value = [...messagesList.value, ..._messages]
  }
  scroll();
});
</script>