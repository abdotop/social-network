export default defineNuxtRouteMiddleware(async () => {
  const { setUser } = useAuth()
  const user = useAuthUser();

  fetch("/api/auth/me")
    .then(
      async (r) => {

        let response = await r.json()
        if (!response.user) {
          setUser(null)
          return navigateTo("/login", { redirectCode: 301 })
        }
      }
    )


  if (!user.value?.isLoggedIn) return navigateTo("/login", { redirectCode: 301 });
});