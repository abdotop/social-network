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

  console.log("tooooooooken\n", cookie);

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
    console.log("okkk kk tooooooooken\n",token,"//", cookie);
    return { user: JSON.parse(response as string) as User, token: token.session };
  } catch (error) {
    console.log(error);
    
    if (typeof event !== "string") {
      deleteCookie(event, config.cookieName, {
        httpOnly: true,
        path: "/",
        sameSite: "strict",
        secure: false,
      });
    }
    return { user: null, token: null };
  }
}