import type { ServerResponse, GroupMessage } from "~/types";

export default defineEventHandler(async (event) => {
  if (!event.context.token) {
    return createError({
      statusCode: 401,
      message: "You don't have the rights to access this resource",
    });
  }

  const token = event.context.token;
  const groupId = getRouterParam(event, "id");

  const response = await $fetch<ServerResponse<GroupMessage>>(`${process.env.BACKEND_URL}/group/messages/new?group_id=${groupId}`, {
    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
  });

  return response;
});