export default defineEventHandler(async (event) => {
    const token = event.context.token
  
    try {
      const payload = await readBody(event);
      const group_id = getQuery(event).gid
      const response = await $fetch(`${process.env.BACKEND_URL}`+"/create-event", {
        method: "POST",
        headers: {
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify(payload),
        query:{
          group_id
        }
      });
      return response;
    } catch (e: any) {
      console.log(e);
      send(event, { code: e.code, msg: e.message });
    }
  });
  