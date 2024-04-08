import { sendError, defineEventHandler } from 'h3'
import { fetcher } from '../utils/fetcher';
import type { ServerResponse, User } from '~/types';

export default defineEventHandler(async (event) => {
    const body = await readBody(event)
    const token = event.context.token

    if (!event.context.token) {
        return sendError(event, createError({
            statusCode: 400,
            statusMessage: 'No user session available'
        }))
    }

    try {
        body["password"] = event.context.user.password
        body["avatarImage"] = event.context.user.avatarImage
        const result = await fetcher(`${process.env.BACKEND_URL}` + "/edituser", "PUT", JSON.stringify(body), token)
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
});


