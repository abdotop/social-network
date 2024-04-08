export default defineEventHandler(async (event) => {
    const reader = await readMultipartFormData(event);
    const token = event.context.token;
    
    if (!reader) return { status: 400, body: 'Bad request' };
    let file;
    for (let part of reader) {
        if (part.filename) { // assuming that the file part has a filename property
            file = part;
            break;
        }
    }

    if (!file) {
        return {
            status: 400,
            body: 'No file uploaded',
        };
    }

    // Create a new FormData instance
    const body = new FormData();
    // Append the file to the FormData instance
    body.append('file', new Blob([file.data]), file.filename);

    if (!token) {
        return {
            status: 401,
            body: 'Unauthorized',
        };
    }

    const response = await fetch(`${process.env.BACKEND_URL}`+"/upload", {
        method: 'POST',
        headers: {
            Authorization: `Bearer ${token}`,
        },
        body,
    }).then(async (res) => await res.json()).catch((err) => {
        console.log(err);
        return {
            status: 500,
            body: 'Internal server error',
        };
    });    

    return {
        status: 200,
        data: response.imageurl,
    };
});