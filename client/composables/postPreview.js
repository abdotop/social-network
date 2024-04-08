import usePostStore from "~/stores/usePostStore.js";
export const postPreviewContent = ref({
    imageUrl: "",
    comments: []
})
export const passDataOnPostPreviewContent = (postID, group_id) => {
    if (group_id != "00000000-0000-0000-0000-000000000000") {
        for (let post of usePostStore().groupPost[group_id]) {
            if (post.id === postID) {
                postPreviewContent.value = post
                break
            }
        }
        return
    }
    for (let post of usePostStore().posts) {
        if (post.id === postID) {
            postPreviewContent.value = post
            break
        }
    }
}