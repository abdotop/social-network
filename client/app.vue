<script setup lang="ts">
import usePostStore from "~/stores/usePostStore.js";
const postStore = usePostStore()
useHead({
  meta: [
    { charset: "utf-8" },
    { name: "viewport", content: "width=device-width, initial-scale=1" },
  ],
  link: [
    { rel: "icon", href: "/favicon.ico" },
  ],
  htmlAttrs: {
    lang: "en",
  },
});

useSeoMeta({
  titleTemplate: "%s - Social Network",
});

onMounted(() => {
  postStore.getUserFeed()
  if (
    localStorage.theme === "dark" ||
    (!("theme" in localStorage) &&
      window.matchMedia("(prefers-color-scheme: dark)").matches)
  ) {
    document.documentElement.classList.add("dark");
  } else {
    document.documentElement.classList.remove("dark");
  }
  localStorage.theme = "light";
  localStorage.theme = "dark";
  localStorage.removeItem("theme");
});
</script>

<template>
  <Html>
  <NuxtLoadingIndicator />
  <NuxtLayout>
    <NuxtPage />
  </NuxtLayout>
  <UNotifications />

  </Html>
</template>
