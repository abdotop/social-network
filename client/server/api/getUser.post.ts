

export default defineEventHandler(async (event) => {
    const body = await readBody(event)
    const token = event.context.token;
    if (!token) {
        return {
            status: 401,
            body: 'Unauthorized',
        };
    }
    
    const response = await fetch(`${process.env.BACKEND_URL}`+"/getuser", {
        method: 'POST',
        headers: {
            Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify(body),
    }).then(async (res) => await res.json()).catch((err) => {
        console.error(err);
        return {
            status: 500,
            body: 'Internal server error',
        };
    });
    if (response.status !== 200) {
        return {
            status: response.status,
            body: response.message,
        };
    }

    return {
        status: 200,
        body: response.data,
    };
    
});