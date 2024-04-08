import type { Group, ServerResponse } from "~/types";

export default defineEventHandler(async (event) => {
  if (!event.context.token) {
    return createError({
      statusCode: 401,
      message: "You don't have the rights to access this resource",
    });
  }

  const token = event.context.token;
  const body = await readBody<{ title: string, description: string }>(event);
  const response = await $fetch<ServerResponse<Group>>(`${process.env.BACKEND_URL}`+"/create-group", {
    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
    body: {
      title: body.title,
      description: body.description
    },
  });

  // console.log(response);

  return response;
});