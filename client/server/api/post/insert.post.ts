export default defineEventHandler(async (event) => {
    const body = await readBody(event)

    const token = event.context.token;
    if (!token) {
        return {
            status: 401,
            body: 'Unauthorized',
        };
    }
    let jsonBody = JSON.stringify(body);
    // console.log("jsonBody", jsonBody)

    const response = await fetch(`${process.env.BACKEND_URL}` + '/post/insert', {
        method: 'POST',
        headers: {
            Authorization: `Bearer ${token}`,
        },
        body: jsonBody,
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