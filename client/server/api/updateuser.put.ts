import { sendError, getSession, defineEventHandler, useSession } from 'h3'
import { fetcher } from '../utils/fetcher';
import { sessionUpdater } from '../utils/sessionHandler'
import type { ServerResponse, User } from '~/types';

export default defineEventHandler(async (event) => {
    const body = await readBody(event)
    const token = event.context.token

    const sessionServer = await getSession(event, {
        password: "5ec0312f-223f-4cc0-aa0f-303ff39fe1b2",
        name: "server-store",
    })
    
    if (sessionServer.data.sessionToken != token) {
        return sendError(event, createError({
            statusCode: 400,
            statusMessage: 'No user session available'
        }))
    } else {
        try {
            body["password"] = sessionServer.data.userInfos.password
            body["avatarImage"] = sessionServer.data.userInfos.avatarImage
            const result = await fetcher(`${process.env.BACKEND_URL}`+"/edituser", "PUT", JSON.stringify(body), token)
            await sessionUpdater(token, result.data, event)
            // console.log(result)
            const userUpdated = result.data
            const { password: _password, ...userWithoutPassword } = userUpdated;
            const cleanInfos = {
                message: result.message,
                user: userWithoutPassword,
            }
            return cleanInfos
        } catch (err) {
            console.log(err)
            return sendError(event, createError({
                statusCode: 500,
                statusMessage: 'Internal server error' + err
        }))
        }
    }
});


