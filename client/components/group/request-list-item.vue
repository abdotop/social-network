<template>
  <div class="px-10  h-12 w-full flex flex-row  items-center justify-between">
    <div class="h-full flex-3 bg-blue-700">
      <div class="flex flex-row gap-3 items-center">
        <NuxtImg src="http://localhost:8081/uploads/default-avatar.png" class="w-10 h-8"/>
        <div>
          <div>{{ `${props.member.User?.firstName} ${props.member.User?.lastName}` }}</div>
        </div>
      </div>  
    </div>
    <div class="h-full  text-left">
      {{ props.member?.Role }}
    </div>
    <div class="h-full">
      <div v-if="status === 'requesting'">
        <UButton class="text-green-500" @click="handleAccept">
          Accept
        </UButton>
        <UButton class="text-red-400" @click="handleDecline">
          Decline
        </UButton>
      </div>
      <div v-else-if="status === 'accepted'">
        <div class="text-blue-500">
          Accepted
        </div>
      </div>
      <div v-else>
        <div class="text-red-400">
          Declined
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
const  { acceptJoinRequest, declineJoinRequest } = useGroupRequest();


const props = defineProps(['member'])
const status = ref<string>(props.member.Status)


async function handleAccept(){
   const response =  await acceptJoinRequest(props.member.GroupID,props.member.ID)
   if (response) {
    const data = JSON.parse(response.data)
    console.log(data);
    
    status.value = data.data.Status
   }
}

async function handleDecline(){
    const response =  await declineJoinRequest(props.member.GroupID,props.member.ID)
   if (response) {
    const data = JSON.parse(response.data)
    console.log(data);

    status.value = data.data.Status
   }
}

</script>

<style></style>