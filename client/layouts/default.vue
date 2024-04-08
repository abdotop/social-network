<template>

  <Body class="bg-15-23-42 overflow-hidden">
    <Header />
    <Sidebar />
    <slot/>
    
  </Body>
  <Post-img-preview />
  <Post-text-preview />
  <!-- create post  -->
  <PostInput />
</template>

<script setup>
</script>