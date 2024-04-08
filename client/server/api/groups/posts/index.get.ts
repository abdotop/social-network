import type { ServerResponse, Post } from "~/types";

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
    
    const group_id = queryObject.gid

    const response = await $fetch<ServerResponse<Post[]>>(`${process.env.BACKEND_URL}`+"/get-all-post-group", {
        method: "GET",
        headers: {
            Accept: "application/json",
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
        },
        query: {
            group_id
        }
    });

    console.log(response);
    

    return response;
});