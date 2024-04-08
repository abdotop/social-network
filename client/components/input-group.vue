<template>
  <UForm ref="formEl" :schema="schema" :state="state" class="space-y-5 w-full" @submit="handleCreateGroup">
    <UFormGroup label="Title" name="title">
      <UInput size="md" v-model="state.title" placeholder="Enter group title" />
    </UFormGroup>

    <UFormGroup label="Description" name="description">
      <UTextarea size="md" v-model="state.description" placeholder="Enter group description" />
    </UFormGroup>

    <UButton type="submit" size="lg" block>
      Create Group
    </UButton>
  </UForm>
</template>

<script setup lang="ts">
import { z } from 'zod'
import type { FormSubmitEvent } from '#ui/types'

const { createGroup } = useGroups();

const schema = z.object({
  title: z.string().min(8, 'Must be at least 8 characters'),
  description: z.string().min(8, 'Must be at least 8 characters'),
})

type Schema = z.output<typeof schema>

const state = reactive({
  title: "",
  description: ""
})
const formEl = ref();

async function handleCreateGroup(event: FormSubmitEvent<Schema>) {
  await createGroup(event.data)
}

</script>