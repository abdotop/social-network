import type { ServerResponse, User } from "~/types";
import { serialize, sign } from "~/server/utils/cookie";
import { sessionCreator } from "~/server/utils/sessionHandler";

export default defineEventHandler(async (event) => {
  const body = await readBody<{ username: string; password: string; rememberMe: boolean }>(event);

  const { username, password, rememberMe } = body;

  const response = await $fetch<ServerResponse<User>>(`${process.env.BACKEND_URL}`+"/login", {
    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: {
      email: username,
      password,
    },
  });

  // console.log(">>>>>>>>>>>>>>>>>>>>>>>", response);

  if (response.status !== "200" || response.data === undefined) {
    return sendError(event, createError({
      statusCode: 400 as number,
      statusMessage: response.message
    }))
  }

  // response.data.avatarImage = process.env.BACKEND_URL + response.data.avatarImage

  const config = useRuntimeConfig();
  const userWithPassword = response.data;
  const session = serialize({ session: response.session });
  const signedSession = sign(session, config.cookieSecret);

  setCookie(event, config.cookieName, signedSession, {
    httpOnly: true,
    path: "/",
    sameSite: "strict",
    secure: false,
    expires: rememberMe
      ? new Date(Date.now() + config.cookieRememberMeExpires)
      : new Date(Date.now() + config.cookieExpires),
  });
  
  const { password: _password, ...userWithoutPassword } = userWithPassword;
  const sessionServer = await sessionCreator(response.session, userWithPassword, event)
  return {
    user: userWithoutPassword,
  };
});
