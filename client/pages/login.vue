<template>
  <div class="flex w-screen h-screen overflow-hidden">
    <div class="flex items-center justify-center px-6 mx-auto w-full lg:w-1/2">
      <div class="flex flex-col lg:w-4/6 w-full min-w-[250px]">
        <div class="flex w-full items-start flex-col">
          <div class="flex justify-center mx-auto">
            <NuxtLink to="/">
              <img class="w-auto h-10 mb-2" src="assets/images/logo-icon.png" alt="">
            </NuxtLink>
          </div>

          <p class="mt-3 font-bold text-4xl mb-4">Sign in</p>
          <p>Welcome back, please enter your credentials!</p>
        </div>

        <div class="mt-8 w-full flex items-center justify-center flex-col">
          <UForm ref="formEl" :schema="schema" :state="state" class="space-y-5 w-full" @submit="onSubmit">
            <UFormGroup label="Email" name="email">
              <UInput size="lg" v-model="state.email" placeholder="Enter your email" />
            </UFormGroup>

            <UFormGroup label="Password" name="password">
              <UInput size="lg" v-model="state.password" type="password" placeholder="•••••••••" autocomplete="on" />
            </UFormGroup>

            <div class="flex w-full justify-between">
              <UCheckbox v-model="state.remember" label="Remember me" name="remember" />
              <ULink class="font-bold text-sm">Forgot your password?</ULink>
            </div>

            <UButton type="submit" size="lg"  class="button bg-primary text-white w-full cursor-pointer" block>
              Sign in
            </UButton>
          </UForm>
          <p class="mt-8 text-md text-center text-secondary">Don't have an account yet? <ULink to="/register"
              active-class="text-primary">Register</ULink>.</p>
        </div>
      </div>
    </div>
    <div class="bg-cover hidden lg:block lg:w-1/2 bg-blue-600">
      <div class="flex items-center h-full px-20 bg-gray-900 bg-opacity-40">
        <div class="flex w-full flex-col items-center justify-center">
          <div class="h-70 w-[100%] max-w-[500px] bg-white/40 mb-20 rounded-xl"></div>

          <h2 class="text-2xl font-bold text-white sm:text-2xl">Have access to your dashboard</h2>

          <p class="max-w-xl mt-3 text-gray-300">
            Lorem ipsum dolor sit, amet consectetur adipisicing elit. In
            autem ipsa, nulla laboriosam dolores, repellendus perferendis libero.
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { z } from 'zod'
import type { FormSubmitEvent } from '#ui/types'

const { login } = useAuth();
const formEl = ref();

useHead({
  title: "Sign in",
})

definePageMeta({
  alias: ["/login"],
  layout: "auth",
  middleware: ["guest-only"],
});

const schema = z.object({
  email: z.string().email('Invalid email'),
  remember: z.boolean().optional(),
  password: z.string().min(8, 'Must be at least 8 characters')
})

type Schema = z.output<typeof schema>

const state = reactive({
  email: undefined,
  remember: undefined,
  password: undefined
})

async function onSubmit(event: FormSubmitEvent<Schema>) {
  try {
    await login(event.data.email, event.data.password, event.data.remember ? true : false)
    const redirect = '/'
    await navigateTo(redirect);
  } catch (error) {
    console.log(error);

    formEl.value.setErrors([
      {
        message: "Invalid email or password",
        path: "email",
      }, {
        message: "Invalid email or password",
        path: "",
      }
    ])
  }
}
</script>