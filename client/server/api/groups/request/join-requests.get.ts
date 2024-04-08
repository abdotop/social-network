export default defineEventHandler(async event => {
    try {
        if (!event.context.token) {
            return createError({
                statusCode: 401,
                message: "You don't have the rights to access this resource",
            });
        }

        const token = event.context.token;
        const queryObj = getQuery(event)
        const groupId = queryObj.gid


        const response = await $fetch(`${process.env.BACKEND_URL}`+"/get-all-access-demand", {
            method: 'GET',
            headers: {
                Authorization: `Bearer ${token}`,
            },
            query: {
                group_id: groupId,
            }
        })

        console.log(response);

        return response

    } catch (e: any) {
        console.log(e)
        setResponseStatus(event, 500, e.message)
    }

})