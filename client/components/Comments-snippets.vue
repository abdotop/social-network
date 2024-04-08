<template>
  <div v-if="post.comments?.length"
    class="sm:p-4 p-2.5 border-t border-gray-100 font-normal space-y-3 relative dark:border-slate-700/40">
    <Comment v-for="comment in post.comments.slice(-2) " :comment="comment" />
    <button v-if="post.comments.length > 2" type="button"
      class="flex items-center gap-1.5 text-gray-500 hover:text-blue-500 mt-2" @click="showMoreHandler"> <i
        class='bx bx-chevron-down ml-auto duration-200 group-aria-expanded:rotate-180'></i>
      More Comment
    </button>
  </div>
</template>

<script>
export default {
  props: {
    post: {
      type: Object,
      required: false,
    }
  },
  methods: {
    showMoreHandler() {
      passDataOnPostPreviewContent(this.post.id, this.post.group_id)
      if (this.post.imageUrl) {
        UIkit.modal("#preview_modal").show()
        return
      }
      UIkit.modal("#text_preview_modal").show(  )
    }
  }
}
</script>

<style></style>