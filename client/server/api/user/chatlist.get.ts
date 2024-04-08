
export default defineEventHandler(async (event) => {
    const token = event.context.token
    const users = await $fetch(`${process.env.BACKEND_URL}`+'/chatlist', {
      headers: {
        Authorization: `Bearer ${token}`
      }
  
    })
    console.log(users)
    return users
  })
  