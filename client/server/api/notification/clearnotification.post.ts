export default defineEventHandler(async (event) => {
    const data = await readBody(event)
    // console.log(data);
    const token = event.context.token;
    if (!token) {
        return {
            status: 401,
            body: 'Unauthorized',
        };
    }

    const response = await fetch(`${process.env.BACKEND_URL}`+'/clearnotifications', {
        method: 'POST',
        headers: {
            Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify(data),
    }).then(async (res) => await res.json()).catch((err) => {
        return {
            status: 500,
            body: 'Internal server error',
        };
    });
    return response
})