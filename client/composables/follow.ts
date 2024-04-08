import { ServerResponse, type UserWithoutPassword } from "~/types";

export const useFollow = async (nickname: string, action: string) => {

    const response = await fetch(`/api/follow`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ action: action, nickname: nickname })
    }).then((res) => res.json()).catch((e) => {
        console.error(e)
    });
    if (response.status !== 200) {
        return {
            ok: false,
        }
    }
    return {
        ok: true,
        action: response.message
    }
}

export const getFollowers = async () => {
    const response = await $fetch("/api/getFollowers") as {status:String,body:any}
    return response.body
}
