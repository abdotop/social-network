import { sendError, getSession, useSession } from 'h3'
import { sessionUpdater } from '../../utils/sessionHandler'
import { fetcher } from '../../utils/fetcher'

export default defineEventHandler(async (event) => {
    const reader = await readMultipartFormData(event);
    if (!reader) return { status: 400, body: 'Bad request', ok: false }

    const body = new FormData();
    for await (const part of reader) {
        if (part.name === "file") {
            body.append('file', new Blob([part.data]), part.filename);
        }
    }

    const token = event.context.token
    const session = await getSession(event, {
        password: "5ec0312f-223f-4cc0-aa0f-303ff39fe1b2",
        name: "server-store",
        // generateId: () => { return '' }
    })

    if (session.data.sessionToken != token) {
        return sendError(event, createError({
            statusCode: 400,
            statusMessage: 'No user session available'
        }))
    } else {
        try {
            const _response = await fetch(`${process.env.BACKEND_URL}` + "/upload", {
                method: "POST",
                headers: {
                    Authorization: `Bearer ${token}`,
                },
                body: body
            });

            const response = JSON.parse(await _response.text()) as { imageurl: string };
            
            if (!response.imageurl) {
                return sendError(event, createError({
                    statusCode: 400,
                    statusMessage: "Error while updloading avatar image"
                }))
            }
            
            const updateValue = {
                email: session.data.userInfos.email,
                avatarImage: response.imageurl
            }
            console.log(updateValue)
            const _response2 = await fetch(`${process.env.BACKEND_URL}` + "/updateavatar", {
                method: "PUT",
                headers: {
                    Accept: "application/json",
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${token}`,
                },
                body: JSON.stringify(updateValue)
            });

            const response3 = JSON.parse(await _response2.text()) as { status: number, message: string };
            if (response3.status != 200) {
                return sendError(event, createError({
                    statusCode: response3.status,
                    statusMessage: response3.message
                }))
            }
            const newUserAvatar = session.data.userInfos
            newUserAvatar.avatarImage = response.imageurl
            await sessionUpdater(token, newUserAvatar, event)
            return {imageurl: response.imageurl, message: "Avatar update successfully"}
        } catch (err) {
            console.log(err)
            return sendError(event, createError({
                statusCode: 500,
                statusMessage: 'Error: Internal server error'
            }))
        }
    }
});


