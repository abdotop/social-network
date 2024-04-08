import { processParts } from "~/server/utils/processParts";
import { Register } from "~/server/models/register";
import type { ServerResponse, User } from "~/types";
import { sessionCreator } from "~/server/utils/sessionHandler";

export default defineEventHandler(async (event) => {
  const reader = await readMultipartFormData(event);
  if (!reader) return { status: 400, body: 'Bad request', ok: false }

  const { file, jsonData } = await processParts(reader);

  const register = new Register(jsonData);
  const [isValid, message] = register.validate();
  if (!isValid) {
    return sendError(event, createError({
      statusCode: 400,
      statusMessage: message
    }))
  }

  const _response = await fetch(`${process.env.BACKEND_URL}` + "/registration", {
    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify(register),
  });
  // @ts-ignore
  const response = JSON.parse(await _response.text()) as ServerResponse<User>;

  if (response.status !== "200") {
    return sendError(event, createError({
      statusCode: 400,
      statusMessage: response.message
    }))
  }

  const config = useRuntimeConfig();
  const session = serialize({ session: response.session });
  const signedSession = sign(session, config.cookieSecret);

  setCookie(event, config.cookieName, signedSession, {
    httpOnly: true,
    path: "/",
    sameSite: "strict",
    secure: process.env.NODE_ENV === "production",
    expires: true
      ? new Date(Date.now() + config.cookieRememberMeExpires)
      : new Date(Date.now() + config.cookieExpires),
  });
  const user = response.data;
  const { password: _password, ...userWithoutPassword } = user;

  if (!file) {
    // create a server Session
    await sessionCreator(response.session, response.data, event)
    return {
      status: 200,
      body: 'No file uploaded',
      session: response.session,
      user: userWithoutPassword,
      ok: true
    };
  }

  const body = new FormData();
  body.append('file', new Blob([file.data]), file.filename);

  const _response2 = await fetch(`${process.env.BACKEND_URL}` + "/upload", {
    method: "POST",
    headers: {
      Authorization: `Bearer ${response.session}`,
    },
    body
  });

  const response2 = JSON.parse(await _response2.text()) as { imageurl: string };

  if (!response2.imageurl) {
    return { status: 200, body: response.message, session: response.session, ok: false };
  }

  userWithoutPassword.avatarImage = response2.imageurl;
  register.avatarImage = response2.imageurl;

  const _response3 = await fetch(`${process.env.BACKEND_URL}` + "/updateuser", {
    method: "PUT",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
      Authorization: `Bearer ${response.session}`,
    },
    body: JSON.stringify(register)
  });

  const response3 = JSON.parse(await _response3.text()) as { status: number, message: string };

  if (response3.status !== 200) {
    return { status: 200, body: response3.message, session: response.session, ok: false }
  }

  // create a server Session
  let userWithAvatar =  {
    id: userWithoutPassword.id,
    email: userWithoutPassword.email,
    firstName: userWithoutPassword.firstName,
    lastName: userWithoutPassword.lastName,
    dateOfBirth: userWithoutPassword.dateOfBirth,
    avatarImage: userWithoutPassword.avatarImage,
    nickname: userWithoutPassword.nickname,
    aboutMe: userWithoutPassword.aboutMe,
    isPublic: userWithoutPassword.isPublic,
    password: response.data?.password
  }
  await sessionCreator(response.session, userWithAvatar, event)
  return {
    status: 200,
    body: "User registered successfully",
    session: response.session,
    user: userWithoutPassword,
    ok: true
  };
});