<template>
  <h1>Homess</h1>
  <p>
    {{ currentUser }}
  </p>
  <nuxt-link to="/">Back to home</nuxt-link>
</template>

<script setup lang="ts">
const currentUser = useAuthUser();

definePageMeta({
  alias: ["/home"],
  middleware: ["auth-only"],
});
</script>