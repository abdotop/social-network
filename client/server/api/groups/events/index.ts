import type { Event } from "~/types";

export default defineEventHandler(async (event) => {
    if (!event.context.token) {
        return createError({
            statusCode: 401,
            message: "You don't have the rights to access this resource",
        });
    }

    const token = event.context.token;
    const queryObject = getQuery(event)
    console.log(queryObject);
    
    const isParticipantNeeded = queryObject.participants === '1'
    const isUserNeeded = queryObject.user === '1'
    const group_id = queryObject.gid

    const response = await $fetch(`${process.env.BACKEND_URL}`+"/get-all-event-group", {
        method: "GET",
        headers: {
            Accept: "application/json",
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
        },
        query: {
            isParticipantNeeded,
            isUserNeeded,
            group_id
        }
    });

    return JSON.parse(response as string) as Event[];
});