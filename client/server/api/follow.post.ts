

export default defineEventHandler(async (event) => {
    const body = await readBody(event);
    const token = event.context.token;
    if (!token) {
        return {
            status: 401,
            body: 'Unauthorized',
        };
    }

    const response = await fetcher(`${process.env.BACKEND_URL}`+"/follower", 'POST', JSON.stringify(body), token)
    return response
   
});