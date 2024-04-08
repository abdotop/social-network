import type { ServerResponse, Post } from "~/types";

export const useGroupPost = () => {
  async function getAllGroupPosts(gid: string | undefined): Promise<Post[]> {
    const response = await $fetch<ServerResponse<Post[]>>("/api/groups/posts", {
      method: "GET",
      headers: useRequestHeaders(["cookie"]) as HeadersInit,
      query: {
        gid
      }
    });
    return response.data as Post[];
  }

  return { getAllGroupPosts }
}
