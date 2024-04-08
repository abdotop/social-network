<script lang="ts" setup>
import { useDateFormat } from "@vueuse/core";
import type { PropType } from "vue";
import type { Event, EventParticipant } from "~/types";

const props = defineProps({
  event: {
    type: Object as PropType<Event>,
    required: true,
  },
});

const status = ref<string>()
const user = useAuthUser()

status.value = props.event.Participants.find((p)=>p.User.id === user.value?.id)?.response
const date = useDateFormat(new Date(props.event.date_time), 'DD,MMMM HH:MM')

const { respondEvent } = useEvents()

async function handleResponseClick(response: string) {
  const { data, error } = await respondEvent(props.event, response)

  if (!error.value) {
    console.log(data);

    const participant: EventParticipant = data.data

    console.log('participant', participant);

    status.value = participant.response
  }
}


</script>

<template>
  <UCard class="w-full flex flex-row justify-between dark:bg-slate-800">
    <div>
      <div class="text-lg text-blue-700">
        {{ props.event.title }}
      </div>
      <div>
        {{ props.event.description }}
      </div>
    </div>
    <div>
      <div>{{ date }}</div>
    </div>
    <div class="flex gap-1 flex-row">
      <UButton v-if="!(status === 'going')" class="bg-blue-500 " @click="handleResponseClick('going')">going</UButton>
      <div v-else class="text-blue-500 ">going</div>
      <UButton v-if="!(status === 'not_going')" class="bg-blue-500 " @click="handleResponseClick('not_going')">not going
      </UButton>
      <div v-else class="text-blue-500 ">not going</div>
    </div>


  </UCard>
</template>

<style></style>