<template lang="">
  <div id="create-event-overlay" class="hidden lg:p-20" uk-modal="" ref="modal">
      <UCard :ui="{
          base: 'flex flex-col',
          rounded: '',
          divide: 'divide-y divide-gray-100 dark:divide-gray-800',
          body: {
            base: 'grow'
          }
        }"
        class=" relative overflow-hidden mx-auto bg-white shadow-xl rounded-lg md:w-[520px] w-full dark:bg-dark2"
      >
        <div class="text-center py-4 border-b mb-0 dark:border-slate-700">
          <h2 class="text-sm font-medium text-black">
            Create Event
          </h2>
  
          <!-- close button -->
          <button type="button" class="button-icon absolute top-0 right-0 m-2.5 uk-modal-close">
            <svg
              xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
              stroke="currentColor" class="w-6 h-6"
            >
              <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
  
        <form ref="create_event_form" action="">
          <div class="space-y-5 mt-3 p-2">
            <input
              id="" label="Title"
              class="w-full !text-black placeholder:!text-black !bg-white !border-transparent focus:!border-transparent focus:!ring-transparent !font-normal !text-xl   dark:!text-white-100 dark:placeholder:!text-red-500 dark:!bg-slate-800" name="title" type="text" placeholder="title"
            />
            <textarea id="" class="resize-none w-full" cols="30" rows="5" name="description" placeholder="Description" />
            <input name="date_time" placeholder="01/01/2000" label="Date" type="datetime-local" required />
          </div>
  
  
          <div class="p-5 flex justify-between items-center">
            <div class="flex items-center gap-2">
              <button type="submit" class="button bg-blue-500 text-white py-2 px-12 text-[14px]">
                Create
              </button>
            </div>
          </div>
        </form>
      </UCard>
    </div>
    
  </template>

<script setup lang="ts">

const emit = defineEmits(['event-created'])

const uikit = useUIkit3()

// const { createEvent } = useEvents()

const create_event_form = ref(null)
const props = defineProps({
  groupId: {
    type: String,
    required: true
  },
  opened: {
    type: Boolean
  }
});
const isOpen = ref(false)

onMounted(() => {
  create_event_form.value?.addEventListener("submit", submitData);
});


async function submitData(e) {
  e.preventDefault()


  const formData = new FormData(e.target)
  const formObj = Object.fromEntries(formData.entries())
  formObj.date_time = new Date(formObj.date_time)

  const { data, error } = await useFetch("/api/groups/events/create", {
    method: "post",
    headers: useRequestHeaders(["cookie"]) as HeadersInit,
    body: JSON.stringify(formObj),
    query: {
      gid: props.groupId,
    },
  })

  if (!error.value) {
    emit("event-created")
    console.log('emitted');
    uikit?.modal("#create-event-overlay").hide()
    
  }

}

function close() {
  isOpen.value = false
}

</script>

<style></style>