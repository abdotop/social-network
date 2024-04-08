<template>
  <form class="sm:px-4 sm:py-3 p-2.5 border-t border-gray-100 flex items-center gap-1 dark:border-slate-700/40"
    @submit="handleCommentSubmission">

    <img :src="'http://localhost:8081/' + useAuthUser().value?.avatarImage" alt="" class="w-6 h-6 rounded-full" />
    <div class="flex-1 relative overflow-hidden h-10">
      <textarea placeholder="Add Comment...." rows="1"
        class="w-full resize-none !bg-transparent px-4 py-2 focus:!border-transparent focus:!ring-transparent"
        name="content" ref="commentInput" @keydown.enter="handleEnterPress"></textarea>
      <input type="file" id="photo-input" name="photo" accept="image/*" style="display: none" ref="fileInput">
    </div>
    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-6 h-6 fill-sky-600"
      @click="() => $refs.fileInput.click()">
      <path fill-rule="evenodd"
        d="M1.5 6a2.25 2.25 0 012.25-2.25h16.5A2.25 2.25 0 0122.5 6v12a2.25 2.25 0 01-2.25 2.25H3.75A2.25 2.25 0 011.5 18V6zM3 16.06V18c0 .414.336.75.75.75h16.5A.75.75 0 0021 18v-1.94l-2.69-2.689a1.5 1.5 0 00-2.12 0l-.88.879.97.97a.75.75 0 11-1.06 1.06l-5.16-5.159a1.5 1.5 0 00-2.12 0L3 16.061zm10.125-7.81a1.125 1.125 0 112.25 0 1.125 1.125 0 01-2.25 0z"
        clip-rule="evenodd" />
    </svg>

    <button type="submit" class="text-sm rounded-full py-1.5 px-3.5 bg-secondery" ref="commentSubmit">
      Reply </button>
  </form>
</template>

<script setup>
import usePostStore from "~/stores/usePostStore.js";
const postStore = usePostStore()

const commentSubmit = ref(null)
const commentInput = ref(null)
const fileInput = ref(null)

const props = defineProps({
  postId: {
    type: String,
    required: true
  }
});

const handleCommentSubmission = async (e) => {
  e.preventDefault()
  let formData = new FormData(e.target)
  let res = await FormImgUploader(formData)
  let jsonFormObject = {
    post_id: props.postId,
    content: formData.get("content"),
    image_url: res.data ? res.data : ""
  }
  res = await fetch("/api/post/insertComment", { method: "POST", body: JSON.stringify(jsonFormObject) })
  if (res.status != 200) {
    return
  }
  let commentContent = await res.json().catch(err => ({ error: err }))
  if (commentContent.error) {
    console.log(commentContent.error)
    return
  }
  postStore.addComment(commentContent.body.data)
  commentInput.value.value = ""
  document.getElementById('photo-input').value = null
}

const handleEnterPress = (event) => {
  if (event.keyCode === 13 && !event.shiftKey) {
    event.preventDefault(); // empêche le saut de ligne
    commentSubmit.value.click() // soumet le commentaire
  }
}
</script>

<style></style>