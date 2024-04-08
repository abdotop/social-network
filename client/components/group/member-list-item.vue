<template>
  <UCard class="dark:bg-slate-800">
    <div class="px-10 h-12 w-full flex flex-row justify-between items-center">
      <div class="h-full text-center  bg-blue-700">
        <div class="flex flex-row items-center gap-3">
          <img src="http://localhost:8081/uploads/default-avatar.png" class="w-10 h-8" />
          <div>
            <div>{{ `${props.member.User?.firstName} ${props.member.User?.lastName}` }}</div>
          </div>
        </div>
      </div>
      <div class="h-full flex flex-col items-center  text-left">
        {{ props.member?.Role }}
      </div>
      <div class="h-full  w-max">
        <UButton v-if="props.member?.Role !== 'admin' && props.isAdmin"
          class="w-10 rounded-full text-red-500 border-4 border-white dark:border-slate-800">
          Delete
        </UButton>
      </div>
    </div>
  </UCard>
</template>

<script lang="ts" setup>
import type { PropType } from 'vue';
import type { GroupMember } from '~/types';

const props = defineProps({
  member: {
    type: Object as PropType<GroupMember>,
    required: true
  },
  isAdmin:{
    type: Boolean,
    required:false
  }
})

</script>

<style></style>