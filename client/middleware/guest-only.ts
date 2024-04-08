export default defineNuxtRouteMiddleware(async () => {
  const user = useAuthUser();

  if (user.value?.isLoggedIn) {
    if (process.server) return navigateTo("/");

    return abortNavigation();
  }
});