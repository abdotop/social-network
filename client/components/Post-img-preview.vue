<template>
  <div class=" lg:p-20 max-lg:!items-start  " id="preview_modal" uk-modal="">

    <div
      class="tt relative mx-auto overflow-hidden shadow-xl rounded-lg lg:flex items-center ax-w-[86rem] w-full lg:h-[80vh]">

      <!-- image previewer -->
      <div class="lg:h-full lg:w-[calc(100vw-400px)] w-full h-96 flex justify-center items-center relative">

        <div class="relative z-10 w-full h-full">
          <img :src="'http://localhost:8081/' + postPreviewContent.imageUrl" alt=""
            class="w-full h-full object-cover absolute" />
        </div>

        <!-- close button -->
        <button type="button" ref= 'closeModal' 
          class="bg-white rounded-full p-2 absolute right-0 top-0 m-3 uk-animation-slide-right-medium z-10 dark:bg-slate-600 uk-modal-close">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
            stroke="currentColor" class="w-6 h-6">
            <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>

      </div>

      <!-- right sidebar -->
      <div
        class="lg:w-[400px] w-full bg-white h-full relative  overflow-y-auto shadow-xl dark:bg-dark2 flex flex-col justify-between">

        <div class="p-5 pb-0">

          <!-- story heading -->
          <div class="flex gap-3 text-sm font-medium">
            <nuxt-link @click= "this.$refs.closeModal.click()"  :to="useAuthUser().value.nickname == postPreviewContent.userOwnerNickname ? `/profile` : `/profile/${postPreviewContent.userOwnerNickname}`">
              <img  :src="'http://localhost:8081/' + postPreviewContent.userAvatarImageUrl" alt="" class="w-9 h-9 rounded-full" />
            </nuxt-link>
            <div class="flex-1">
              <nuxt-link @click= "this.$refs.closeModal.click()"
                :to="useAuthUser().value.nickname == postPreviewContent.userOwnerNickname ? `/profile` : `/profile/${postPreviewContent.userOwnerNickname}`">
                <h4 class="text-black dark:text-white"> {{ postPreviewContent.userCompleteName }}</h4>
              </nuxt-link>
              <div class="text-xs text-gray-500 dark:text-white/80"> {{ postPreviewContent.createdAt }}</div>
            </div>

            <!-- dropdown -->
            <!-- <div class="-m-1">
              <button type="button" class="button__ico w-8 h-8"> <ion-icon class="text-xl"
                  name="ellipsis-horizontal"></ion-icon> </button>
              <div class="w-[253px]"
                uk-dropdown="pos: bottom-right; animation: uk-animation-scale-up uk-transform-origin-top-right; animate-out: true">
                <nav>
                  <a href="#"> <ion-icon class="text-xl shrink-0" name="bookmark-outline"></ion-icon>
                    Add to favorites </a>
                  <a href="#"> <ion-icon class="text-xl shrink-0" name="notifications-off-outline"></ion-icon> Mute
                    Notification </a>
                  <a href="#"> <ion-icon class="text-xl shrink-0" name="flag-outline"></ion-icon>
                    Report this post </a>
                  <a href="#"> <ion-icon class="text-xl shrink-0" name="share-outline"></ion-icon>
                    Share your profile </a>
                  <hr>
                  <a href="#" class="text-red-400 hover:!bg-red-50 dark:hover:!bg-red-500/50">
                    <ion-icon class="text-xl shrink-0" name="stop-circle-outline"></ion-icon>
                    Unfollow </a>
                </nav>
              </div>
            </div> -->
          </div>
          <p class="font-normal text-sm leading-6 mt-4">{{ postPreviewContent.content }}</p>

          <div class="shadow relative -mx-5 px-5 py-3 mt-3">
            <div class="flex items-center gap-4 text-xs font-semibold">

                <div class="flex items-center gap-3">
                <button type="button" class="button-icon bg-slate-200/70 dark:bg-slate-700">
                  <i class="text-lg bx bxs-message-dots"></i> </button>
                <span>{{ postPreviewContent.comments.length }}</span>
              </div>
              <!-- <button type="button" class="button__ico ml-auto"> <ion-icon class="text-xl"
                  name="share-outline"></ion-icon> </button>
              <button type="button" class="button__ico"> <ion-icon class="text-xl" name="bookmark-outline"></ion-icon>
              </button> -->
            </div>
          </div>
        </div>
        <div class="p-5 h-full overflow-y-auto flex-1">

          <!-- comment list -->
          <div class="relative text-sm font-medium space-y-5">
            <Comment v-for="comment in postPreviewContent.comments " :comment="comment" />
          </div>
        </div>
        <Comment-input :postId="postPreviewContent.id" />
      </div>

    </div>

  </div>
</template>

<script setup>

</script>

<style></style>