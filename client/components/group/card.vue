<script setup lang="ts">
const { joinRequest } = useGroupRequest();

const props = defineProps({
  group: {
    type: Object,
    required: true
  },
  joined:{
    type: Boolean,
    default: false
  }
});

const joined = ref(props.joined)

const handleJoin = async ()=>{
    const {data} = await joinRequest(props.group.ID)

    if (data) {
      joined.value = true
    }
}
</script>

<template>
  <UCard class="card w-60 bg-gray-100 dark:bg-gray-950 p-0 dark:bg-dark2 dark:border-slate-800"
    :ui="{ header: { padding: 'p-0 rounded-t-lg overflow-hidden' } }">
    <NuxtLink :href="'/groups/' + props.group.ID">
      <div class="card-media h-24 rounded-t-lg">
        <NuxtImg :src="'http://localhost:8081/'+props.group.BannerURL" class="" alt="" />
      </div>
    </NuxtLink>
    <div class="card-body z-10 relative w-full">
      <NuxtLink :href="'/groups/' + props.group.ID">
        <h5 class="card-title text-xl h-min text-black font-bold text-ellipsis w-full">
          {{ props.group.Title }}
        </h5>
      </NuxtLink>
      <div class="flex mb-4 text-sm mt-2">
        <div class="md:block hidden" />
        <div>{{ props.group.GroupMembers.length }} member(s)</div>
      </div>
      <UButton v-if="!joined" @click="handleJoin"
        class="items-center bg-blue-500 justify-center flex-1 w-full">
        Join
      </UButton>
      <NuxtLink v-else :href="'/groups/' + props.group.ID">
        <UButton class="items-center bg-blue-500 justify-center w-full">
          View
        </UButton>
      </NuxtLink>
    </div>
  </UCard>
</template>