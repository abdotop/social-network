import { sendError } from 'h3'
import { fetcher } from '../utils/fetcher'

export default defineEventHandler(async (event) => {
    const body = await readBody(event)
    const token = event.context.token
    
    if (!token) {
        return sendError(event, createError({
            statusCode: 400,
            statusMessage: 'No user session available'
        }))
    } else {
        try {
            body['email'] = event.context.user.email
            const result = await fetcher(`${process.env.BACKEND_URL}`+"/updatepassword", "PUT", JSON.stringify(body), token)
            
            const { password: _password, ...userWithoutPassword } = result.data;
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


