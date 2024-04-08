

export default defineEventHandler(async (event) => {

    const token = event.context.token;
    if (!token) {
        return {
            status: 401,
            body: 'Unauthorized',
        };
    }
    const response = await fetch(`${process.env.BACKEND_URL}`+'/usersByFollow', {
        headers: {
            Authorization: `Bearer ${token}`,
        }
    });
    // console.log(response);
    return response
})
