import type { Group, ServerResponse } from "~/types";

export default defineEventHandler(async (event) => {
  if (!event.context.token) {
    return createError({
      statusCode: 401,
      message: "You don't have the rights to access this resource",
    });
  }

  const token = event.context.token;
  const groupId = getRouterParam(event, "id");

  const response = await $fetch<ServerResponse<Group>>(`${process.env.BACKEND_URL}/get-group?group_id=${groupId}&isMemberNeeded=true&isUserNeeded=true`, {
    method: "GET",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
  });

  return response;
});