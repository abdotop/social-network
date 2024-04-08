import type { User, ServerResponse } from '~/types'
const GetAllUsers = async () => {
    try {
        let response = await $fetch<ServerResponse<User[]>>("/api/user/chatlist", {
            method: "get"
        })
        console.log(response)
        return response
    } catch (err) {
        console.log(err)
    }
}

export { GetAllUsers }