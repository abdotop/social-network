<template>
  <UCard class="dark:bg-slate-800">
    <div class="px-10 h-12 w-full flex flex-row justify-between items-center">
      <div class="h-full text-center  bg-blue-700">
        <div class="flex flex-row items-center gap-3">
          <img src="http://localhost:8081/uploads/default-avatar.png" class="w-10 h-8" />
          <div>
            <div>{{ `${props.user?.firstname} ${props.user?.lastname}` }}</div>
          </div>
        </div>
      </div>
      <div class="h-full  w-max">
        <UButton @click="handleInvitation" v-if="!invited" class="rounded bg-blue-500 border-4 border-white dark:border-slate-800">
          invite
        </UButton>
        <div v-else class="text-blue-500"> invited</div>
      </div>
    </div>
  </UCard>
</template>

<script lang="ts" setup>

const { sendInvitation } = useGroupRequest()

const props = defineProps({
  user: {
    type: Object,
    required: true
  },
  groupId: {
    type: String,
    required: true
  },
  isInvited: {
    type: Boolean,
    required: false
  }
})

const invited = ref(props.isInvited)

async function handleInvitation() {
  const { data, error } = await sendInvitation(props.groupId,props.user.id)
  console.log(data);
  if (data.value) {
    invited.value = true
  }
  
}

</script>