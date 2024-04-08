import { getUserFromToken } from "~/server/utils/session";

export default defineEventHandler(async (event) => {
  const { user, token } = await getUserFromToken(event);
  if (user) event.context.user = user;
  if (token) event.context.token = token;
});