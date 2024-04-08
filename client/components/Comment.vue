<template>
  <div class="flex items-start gap-3 relative">
    <nuxt-link
      :to="useAuthUser().value.nickname == comment.userOwnerNickname ? `/profile` : `/profile/${comment.userOwnerNickname}`">
      <img :src="'http://localhost:8081/' + comment.userAvatarImageUrl" alt="" class="w-6 h-6 mt-1 rounded-full" />
    </nuxt-link>
    <div class="flex-1">
      <nuxt-link
        :to="useAuthUser().value.nickname == comment.userOwnerNickname ? `/profile` : `/profile/${comment.userOwnerNickname}`"
        class="text-black font-medium inline-block dark:text-white">
        {{ comment.userCompleteName }} </nuxt-link>
      <p class="mt-0.5"> {{ comment.content }} </p>
      <img v-if="comment.imageUrl" class="w-16 h-16 mt-1 " :src="'http://localhost:8081/' + comment.imageUrl" ></img>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    comment: {
      type: Object,
      required: true
    }
  }
}
</script>

<style></style>
