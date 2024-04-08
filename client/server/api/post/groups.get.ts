export default defineEventHandler(async (event) => {
    const groupId  = event.headers.get('groupId')
    const token = event.context.token;
    console.log(groupId, token, "");
    

    if (!token) {
        return {
            status: 401,
            body: 'Unauthorized',
        };
    }
    const response = await fetch(`${process.env.BACKEND_URL}`+`/post/groups?id=${groupId}`, {
        method: 'GET',
        headers: {
            Authorization: `Bearer ${token}`,
        },
    }).then(async (res) => await res.json()).catch((err) => {
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
        body: response
    };

});