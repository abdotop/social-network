export default defineNuxtRouteMiddleware(async () => {
  const user = useAuthUser();

  if (!user.value?.isLoggedIn) return navigateTo("/login");
});