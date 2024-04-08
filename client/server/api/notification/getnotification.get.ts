export default defineEventHandler(async (event) => {
    const token = event.context.token;
    
    if (!token) {
        return {
            status: 401,
            body: 'Unauthorized',
        };
    }
    
    const response = await fetch(`${process.env.BACKEND_URL}`+'/getnotifications', {
        method: 'GET',
        headers: {
            Authorization: `Bearer ${token}`,
        },
    }).then(async (res) => await res.json()).catch((err) => {
        return {
            status: 500,
            body: 'Internal server error',
        };
    });    
    return response
})