<script lang="ts" setup>
const colorMode = useColorMode()

function toggleMode() {
  colorMode.preference = colorMode.value === 'light' ? 'dark' : 'light'
}
</script>

<template>
  <UButton :icon="colorMode.value === 'dark'? 'i-ph-moon-stars-duotone' : 'i-ph-sun-dim-duotone' " @click="toggleMode" variant="ghost">
  </UButton>
</template>