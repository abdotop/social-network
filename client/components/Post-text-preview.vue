<template>
  <div class=" lg:p-20 max-lg:!items-start " id="text_preview_modal" uk-modal="">

    <div
      class="lg:w-[600px] w-full bg-white h-full relative  overflow-y-auto shadow-xl dark:bg-dark2 flex flex-col justify-between">
      <div class="p-5 pb-0">
        <!-- story heading -->
        <div class="flex gap-3 text-sm font-medium">
          <img :src="'http://localhost:8081/' + postPreviewContent.userAvatarImageUrl" alt=""
            class="w-9 h-9 rounded-full" />
          <div class="flex-1">
            <h4 class="text-black font-medium dark:text-white"> {{ postPreviewContent.userCompleteName }} </h4>
            <div class="text-gray-500 text-xs dark:text-white/80"> {{ postPreviewContent.createdAt }}</div>
          </div>

          <!-- dropdown -->
          <div class="-m-1">
            <button type="button"
              class="bg-white rounded-full p-2 absolute right-0 top-0 m-3 uk-animation-slide-right-medium z-10 dark:bg-slate-600 uk-modal-close">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                stroke="currentColor" class="w-6 h-6">
                <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
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
</template>

<script>
</script>

<style></style>