<template>
  <!--  post image-->
  <div class="bg-white rounded-xl shadow-sm text-sm font-medium border1 dark:bg-dark2 ">

    <!-- post heading -->
    <div class="flex gap-3 sm:p-4 p-2.5 text-sm font-medium">
      <nuxt-link
        :to="useAuthUser().value.nickname == post.userOwnerNickname ? `/profile` : `/profile/${post.userOwnerNickname}`">
        <img :src="'http://localhost:8081/' + post.userAvatarImageUrl" alt="" class="w-9 h-9 rounded-full" />
      </nuxt-link>
      <div class="flex-1">
        <nuxt-link
          :to="useAuthUser().value.nickname == post.userOwnerNickname ? `/profile` : `/profile/${post.userOwnerNickname}`">
          <h4 class="text-black dark:text-white"> {{ post.userCompleteName }}</h4>
        </nuxt-link>
        <div class="text-xs text-gray-500 dark:text-white/80"> {{ post.createdAt }}</div>
      </div>

      <!-- <div class="-mr-1">
        <button type="button" class="button-icon w-8 h-8"> <ion-icon class="text-xl"
            name="ellipsis-horizontal"></ion-icon> </button>
        <div class="w-[245px]"
          uk-dropdown="pos: bottom-right; animation: uk-animation-scale-up uk-transform-origin-top-right; animate-out: true; mode: click">
          <nav>
            <a href="#"> <ion-icon class="text-xl shrink-0" name="bookmark-outline"></ion-icon> Add to favorites
            </a>
            <a href="#"> <ion-icon class="text-xl shrink-0" name="notifications-off-outline"></ion-icon> Mute
              Notification </a>
            <a href="#"> <ion-icon class="text-xl shrink-0" name="flag-outline"></ion-icon> Report this post </a>
            <a href="#"> <ion-icon class="text-xl shrink-0" name="share-outline"></ion-icon> Share your profile </a>
            <hr>
            <a href="#" class="text-red-400 hover:!bg-red-50 dark:hover:!bg-red-500/50">
              <ion-icon class="text-xl shrink-0" name="stop-circle-outline"></ion-icon> Unfollow </a>
          </nav>
        </div>
      </div> -->
    </div>
    <div class="sm:px-4 p-2.5 pt-0">
      <p class="font-normal"> {{ post.content }} </p>
    </div>
    <!-- post image -->

    <a @click="showPostPreview" v-if="post.imageUrl">
      <div class="relative w-full lg:h-96 h-full sm:px-4">
        <img :src="'http://localhost:8081/' + post.imageUrl" class="sm:rounded-lg w-full h-full object-cover" />
      </div>
    </a>

    <!-- post icons -->
    <div class="sm:p-4 p-2.5 flex items-center gap-4 text-xs font-semibold">

      <div class="flex items-center gap-3">
        <button type="button" class="button-icon bg-slate-200/70 dark:bg-slate-700">
          <i class="text-lg bx bxs-message-dots"></i> </button>
        <span>{{ post.comments ? post.comments.length : null }}</span>
      </div>

    </div>

    <!-- comments -->
    <Comments-snippets :post="post" />
    <!-- add comment -->
    <Comment-input :postId="post.id" />

  </div>
</template>
<script>
useAuthUser();
import { passDataOnPostPreviewContent } from '~/composables/postPreview';
export default {
  mounted() {
    console.log(useAuthUser().value.nickname)
    // console.log('posssssssst',this.props.post);
  },
  props: {
    post: {
      type: Object,
      required: true
    }
  },
  methods: {
    showPostPreview() {
      passDataOnPostPreviewContent(this.post.id)
      UIkit.modal("#preview_modal").show()
    }
  }
}
</script>