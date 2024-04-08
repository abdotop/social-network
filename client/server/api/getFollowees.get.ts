export default defineEventHandler(async (event) => {
    const token = event.context.token;
    if (!token) {
        return {
            status: 401,
            body: 'Unauthorized',
        };
    }    
    
    const response = await fetch(`${process.env.BACKEND_URL}`+"/getAllFollowees", {
        method: 'GET',
        headers: {
            Authorization: `Bearer ${token}`,
        },
    }).then(async (res) => {
        return res.json()
    }).catch((err) => {
        console.log(err);
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