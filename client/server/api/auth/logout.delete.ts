import type { ServerResponse } from "~/types";
import { sessionDeleter } from "~/server/utils/sessionHandler";

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig();

  if (!event.context.token) {
    return createError({
      statusCode: 401,
      message: "You don't have the rights to access this resource",
    });
  }
  
  const token = event.context.token;
  const response = await $fetch<ServerResponse<{}>>(`${process.env.BACKEND_URL}`+"/logout", {
    method: "DELETE",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
  });

  deleteCookie(event, config.cookieName, {
    httpOnly: true,
    path: "/",
    sameSite: "strict",
    secure: process.env.NODE_ENV === "production",
  });
  await sessionDeleter(event)

  return {
    user: {},
  };
});