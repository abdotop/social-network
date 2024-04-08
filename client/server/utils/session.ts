import type { H3Event } from "h3";
import type { User } from "~/types";

export async function getUserFromToken(event: H3Event | string) {
  const config = useRuntimeConfig();
  let cookie: string;
  if (typeof event !== "string") {
    const _cookie = getCookie(event, config.cookieName);
    if (!_cookie) return { user: null, token: null };
    cookie = _cookie;
  } else {
    cookie = event;
  }

  const unsignedToken = unsign(cookie, config.cookieSecret);
  if (!unsignedToken) return { user: null, token: null };

  const token = deserialize(unsignedToken);

  try {
    const response = await $fetch(`${process.env.BACKEND_URL}`+"/me", {
      method: "GET",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
        Authorization: `Bearer ${token.session}`,
      },
    });
    return { user: JSON.parse(response as string) as User, token: token.session };
  } catch (error) {
    if (typeof event !== "string") {
      deleteCookie(event, config.cookieName, {
        httpOnly: true,
        path: "/",
        sameSite: "strict",
        secure: process.env.NODE_ENV === "production",
      });
    }
    return { user: null, token: null };
  }
}