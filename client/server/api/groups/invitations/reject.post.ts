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
        const iId = queryObj.iId


        const response = await $fetch(`${process.env.BACKEND_URL}`+"/decline-invitation", {
            method: 'POST',
            headers: {
                Authorization: `Bearer ${token}`,
            },

            query: {
                invitation_id: iId
            }
        })
        return response

    } catch (e: any) {
        console.log(e)
        setResponseStatus(event, 500, e.message)
    }

})