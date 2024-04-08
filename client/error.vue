<script setup lang="ts">
defineProps({
  error: Object,
})

const handleError = () => clearError({ redirect: '/' })
</script>

<template>
  <NuxtLoadingIndicator />
  <NuxtLayout>
    <div>
      <p>
        {{ error?.statusCode }}
      </p>
      <p mt-7 text-6xl font-bold text-gray-800 z-10>
        Oops! Something went wrong.
      </p>
      <p mt-7 text-2xl z-10>
        {{ error?.message }}
      </p>
      <UButton variant="ghost" text-2l size="xl" mt-8 @click="handleError" z-10>
        Go to home
      </UButton>
    </div>
  </NuxtLayout>
</template>

<style lang="css" scoped>
.special:first-letter {
  @apply text-primary/20;
}
</style>