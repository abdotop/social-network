import type { Group, GroupMessage, ServerResponse } from "~/types";

export const useGroups = () => {
  const groups = ref<Group[]>([]);

  const createGroup = async (group: { title: string, description: string }) => {
    const response = await $fetch<ServerResponse<Group>>("/api/groups", {
      method: "POST",
      headers: useRequestHeaders(["cookie"]) as HeadersInit,
      body: {
        title: group.title,
        description: group.description
      },
    });

    if (response.data) {
      groups.value = [...groups.value, response.data];
    }
  };

  const getAllGroups = async () => {
    const response = await $fetch<ServerResponse<Group[]>>("/api/groups", {
      headers: useRequestHeaders(["cookie"]) as HeadersInit,
    });
    if (response.data) {
      groups.value = response.data;
    }
  };

  const getGroupByID = async (id: string) => {
    const response = await $fetch<ServerResponse<Group>>("/api/groups/" + id, {
      headers: useRequestHeaders(["cookie"]) as HeadersInit,
    });

    if (response.data) {
      return response.data;
    }

    return null;
  };

  const getAllMessagesByGroup = async (id: string) => {
    const response = await $fetch<ServerResponse<GroupMessage[]>>("/api/groups/messages/" + id, {
      headers: useRequestHeaders(["cookie"]) as HeadersInit,
    });

    if (response.data) {
      return response.data;
    }

    return null;
  };

  return {
    groups,
    createGroup,
    getAllGroups,
    getGroupByID,
    getAllMessagesByGroup
  };
}
