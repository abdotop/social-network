<template>
  <div class="flex w-screen h-screen">
    <div class="flex items-center justify-center p-8 mx-auto overflow-y-auto w-full lg:w-1/2">
      <div class="flex flex-col lg:w-4/6 w-full h-full min-w-[250px]">
        <div class="flex w-full items-start flex-col">
          <div class="flex justify-center mx-auto">
            <NuxtLink to="/">
              <img class="w-auto h-10 mb-2" src="assets/images/logo-icon.png" alt="">
            </NuxtLink>
          </div>

          <p class="mt-3 font-bold text-4xl mb-4">Register</p>
          <p>Sign up to get started</p>
        </div>

        <!-- <div class="mt-8 w-full flex items-center justify-center flex-col"> -->

        <div class="space-y-7 text-sm text-black font-medium dark:text-white"
          uk-scrollspy="target: > *; cls: uk-animation-scale-up; delay: 100 ;repeat: true">
          <!-- <h2 class="space-y-7 text-xl text-red-500 font-bold dark:text-red-500">{{ data.registerError }}</h2> -->
          <!-- <div class="grid grid-cols-2 gap-4 gap-y-7"> -->
          <UForm ref="formEl" :schema="schema" :state="state" class="space-y-5 w-full" @submit="handleRegister">
            <div class="flex w-full space-x-2">
              <UFormGroup label="Firstname" name="firstname" class="mt-2.5">
                <UInput size="md" v-model="state.firstName" placeholder="Enter your firstname"
                  class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5" />
              </UFormGroup>

              <UFormGroup label="Lastname" name="lastname" class="mt-2.5">
                <UInput size="md" v-model="state.lastName" placeholder="Enter your lastname"
                  class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5" />
              </UFormGroup>
            </div>

            <div class="flex w-full space-x-2">
              <UFormGroup label="Nickname" name="nickname" class="mt-2.5">
                <UInput size="md" v-model="state.nickname" placeholder="Enter your nickname"
                  class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5" />
              </UFormGroup>
              <UFormGroup label="Email" name="email" class="mt-2.5">
                <UInput size="md" v-model="state.email" placeholder="Enter your email"
                  class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5" />
              </UFormGroup>
            </div>

            <div class="flex w-full space-x-2">
              <UFormGroup label="Password" name="password" class="mt-2.5">
                <UInput size="md" v-model="state.password" type="password" placeholder="•••••••••"
                  class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5" />
              </UFormGroup>

              <UFormGroup label="Confirm Password" name="confirmPassword" class="mt-2.5">
                <UInput size="md" v-model="state.confirmPassword" type="password" placeholder="•••••••••"
                  class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5" />
              </UFormGroup>
            </div>

            <div class="flex w-full space-x-2">
              <UFormGroup label="Date Of Birth" name="dateOfBirth" class="mt-2.5">
                <UInput size="md" v-model="state.dateOfBirth" type="date" placeholder="01/01/2000"
                  class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5" />
              </UFormGroup>

              <UFormGroup label="Avatar Image" class="mt-2.5">
                <UInput type="file" size="md" @change="handleImageSelected" />
                <div v-if="imageUrl" class="col-span-1">
                  <img :src="imageUrl" alt="Uploaded Image">
                </div>
              </UFormGroup>
            </div>

            <UFormGroup label="About Me" name="aboutMe" class="col-span-2">
              <UTextarea size="lg" v-model="state.aboutMe" placeholder="Short introduction about yourself ..." />
            </UFormGroup>

            {{ state.registerError }}
            <UButton type="submit" size="lg" class="button bg-primary text-white w-full cursor-pointer" block>
              Register
            </UButton>
          </UForm>
          <!-- </div> -->
          <p class="mt-8 text-md text-center text-secondary">Already with an account? <ULink to="/login"
              active-class="text-primary">Sign in</ULink>.</p>
        </div>
      </div>
    </div>
    <div class="bg-cover hidden lg:block lg:w-1/2 bg-blue-600">
      <div class="flex items-center h-full px-20 bg-gray-900 bg-opacity-40">
        <div class="flex w-full flex-col items-center justify-center">
          <div class="h-70 w-[100%] max-w-[500px] bg-white mb-20 rounded-xl"></div>

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

const { register } = useAuth();
const formEl = ref();

useHead({
  title: "Register",
})

definePageMeta({
  alias: ["/register"],
  layout: "auth",
  middleware: ["guest-only"],
});

const schema = z.object({
  firstName: z.string().min(1, 'First name is required'),
  lastName: z.string().min(1, 'Last name is required'),
  email: z.string().email('Invalid email address'),
  nickname: z.string().optional(), //min(1, 'Nickname is required'),
  password: z.string().min(8, 'Password must be at least 8 characters'),
  confirmPassword: z.string(),
  dateOfBirth: z.string().min(1, 'Date of birth is required'),
  aboutMe: z.string().optional(), //.min(1, 'About me is required'),
  avatarImg: z.any().optional(),
}).refine((data) => data.confirmPassword === data.password, {
  message: 'Passwords do not match',
})

type Schema = z.output<typeof schema>

const state = reactive({
  firstName: '',
  lastName: '',
  email: '',
  nickname: '',
  password: '',
  confirmPassword: '',
  dateOfBirth: '',
  aboutMe: '',
  avatarImg: File,
  avatarLocalUrl: '',
  registerError: '',
  loading: false
})
const imageUrl = ref("")

const handleImageSelected = (event: Event) => {
  const input = event.target as HTMLInputElement;
  if (!input.files?.length) return;

  const imageFile = input.files[0];
  imageUrl.value = URL.createObjectURL(imageFile)
  state.avatarLocalUrl = imageUrl.value
  // @ts-ignore
  state.avatarImg = imageFile;
}

async function handleRegister(event: FormSubmitEvent<Schema>) {
  state.loading = true
  state.registerError = ""
  try {
    const fomData = JSON.stringify({
      firstName: state.firstName.trim(),
      lastName: state.lastName.trim(),
      email: state.email.trim(),
      nickname: state.nickname.trim(),
      password: state.password.trim(),
      repeatPassword: state.confirmPassword.trim(),
      dateOfBirth: state.dateOfBirth,
      aboutMe: state.aboutMe.trim(),
    })
    await register(state.avatarImg, fomData)
    const redirect = '/'
    await navigateTo(redirect);
  } catch (error) {
    // @ts-ignore
    state.registerError = error.statusMessage
    state.loading = false
  } finally {
    state.loading = false
  }
}
</script>