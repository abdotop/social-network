
export default defineEventHandler(async (event) => {
  
    const token = getHeader(event, 'Authorization')
    const groups = await $fetch(`${process.env.BACKEND_URL}`+'/get-groups', {
      headers: {
        Authorization: `${token}`
      }
  
    })
    return { groups }
  })
  