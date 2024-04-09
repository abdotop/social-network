import { getUserFromToken } from "~/server/utils/session";
import { logout } from "../api/auth/logout.delete";

export default defineEventHandler(async (event) => {
  const { user, token } = await getUserFromToken(event);
const config  = useRuntimeConfig()

  if (!user) {
    await logout(token,event,config)
    return
  }
  if (user) event.context.user = user;
  if (token) event.context.token = token;
});