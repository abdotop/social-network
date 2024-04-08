import { sendError } from "h3"
import type { Data } from "ws";
import { date, string } from "zod";

export let ws: WebSocket | undefined
type Message = { id: string, user: user, message: string, created_at: Date }
export const messages = useState<Message[]>(() => []);
export type user = { id: string; firstName: string; lastName: string; nickName: string; avatar: string; online: boolean; lastMessage: string; lastMessageTime?: Date; }
export const users = ref<user[]>([]);
export let selectedUser = {
    id: "",
    firstName: "",
    lastName: "",
    nickName: "",
    avatar: "",
    online: false,
    lastMessage: "",
} as user

// sort users last message time ,online status and alphabetically
const sort = (us: user[]) => {
    us.sort((a, b) => {
        if (a.online && !b.online) {
            return -1
        }
        if (!a.online && b.online) {
            return 1
        }
        if (a.lastMessageTime && b.lastMessageTime) {
            return b.lastMessageTime.getTime() - a.lastMessageTime.getTime()
        } else if (a.lastMessageTime && !b.lastMessageTime) {
            return -1
        } else if (!a.lastMessageTime && b.lastMessageTime) {
            return 1
        }
        return a.firstName.localeCompare(b.firstName)
    })
}
export const scroll = () => {
    if (messages.value.length) {
        nextTick(() => {
            const chatsbubble = document.getElementById("chatsbubble")
            chatsbubble!.scrollTo(0, chatsbubble!.scrollHeight)
        })
    }
}
export const getUsers = async (user: any) => {
    const response = await fetch('/api/user/usersByFollow', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
        }
    }).then(async (res) => await res.json()).catch((err) => {
        return {
            status: 500,
            body: 'Internal server error',
        };
    })
    if (response.error) {
        return new Error(response.error)
    }

    ws = await connMessageSocket(user.id)
    ws.onmessage = (event) => {
        const message = JSON.parse(event.data)
        if (message.type === 'online') {
            message.users.forEach((user: user) => {
                const u = users.value.find((u) => u.id === user.id)
                if (u) {
                    u.online = user.online
                }
            })
        } else if (message.type === 'new_message') {
            const data = message.data
            users.value.map((u) => {
                if (u.id === data.SenderID) {
                    u.lastMessage = data.Content
                    u.lastMessageTime = new Date(data.CreatedAt)
                }
            })

            if (selectedUser.id === data.SenderID) {
                messages.value.push({
                    id: data.ID,
                    user: users.value.find((u) => u.id === data.SenderID) as user,
                    message: data.Content,
                    created_at: new Date(data.CreatedAt),
                })
                scroll()
            } else if (user.id === data.SenderID) {
                users.value.map((u) => {
                    if (u.id === data.SenderID) {
                        u.lastMessage = data.Content
                        u.lastMessageTime = new Date(data.CreatedAt)
                    }
                })
                messages.value.push({
                    id: data.ID,
                    user: {
                        id: user.id,
                        firstName: user.firstName,
                        lastName: user.lastName,
                        nickName: user.nickname,
                        avatar: user.avatarImage,
                        online: true,
                        lastMessage: "string",
                    } as user,
                    message: data.Content,
                    created_at: new Date(data.CreatedAt),
                })
                scroll()
            }
            // console.log(message.data);
        }

    }
    users.value = []
    function isOlderThan50Years(date: Date): boolean {
        const fiftyYearsAgo = new Date();
        fiftyYearsAgo.setFullYear(fiftyYearsAgo.getFullYear() - 50);
        return date < fiftyYearsAgo;
    }
    response.data.forEach((user: any) => {
        let date
        if (!isOlderThan50Years(new Date(user.lastMessage.CreatedAt))) {
            date = new Date(user.lastMessage.CreatedAt)
        }
        users.value.push({
            id: user.user.id,
            firstName: user.user.firstName,
            lastName: user.user.lastName,
            nickName: user.user.nickname,
            avatar: user.user.avatarImage,
            online: false,
            lastMessage: user.lastMessage.Content,
            lastMessageTime: date,
        })
    })
    ws!.onopen = () => {
        ws!.send(JSON.stringify({ type: 'online', users: users.value }))
    }
    sort(users.value)
    const res = await getSelectUserMessage(user, users.value![0]!.id)
    // console.log(users,"response.data");
    return res
}


export const connMessageSocket = async (id: string = "") => {
    const isSecure = location.protocol === "https:";
    const url = (isSecure ? "wss://" : "ws://") + location.host + `/api/socket?userId=${id}&channel=message`;
    if (ws) {
        ws.close();
    }

    ws = new WebSocket(url);
    return ws
}

export const getSelectUserMessage = async (currentuser: any, receiverId: string) => {
    messages.value = []
    selectedUser = users.value.find((u) => u.id === receiverId) as user
    const response = await fetch('/api/getReceiver', {
        body: JSON.stringify({
            receiver_id: receiverId,
        }),
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        }
    }).then(async (res) => await res.json()).catch((err) => {
        return {
            status: 500,
            body: 'Internal server error',
        };
    })


    if (response.error) {
        return new Error(response.error)
    }

    if (response.body) {
        response.body.forEach((message: any) => {
            let u = users.value.find((u) => u.id === message.SenderID) as user
            if (!u) {
                u = {
                    id: currentuser.id,
                    firstName: currentuser.firstName,
                    lastName: currentuser.lastName,
                    nickName: currentuser.nickname,
                    avatar: currentuser.avatarImage,
                    online: true,
                    lastMessage: "string",
                }
            }
            messages.value.push({
                id: message.ID,
                user: u,
                message: message.Content,
                created_at: new Date(message.CreatedAt),
            })
        })
    }
    return true
}

export const sendMessage = async (userId: string, message: string) => {
    const newMessage = {
        "type": "private_message",
        "message": {
            "sender_id": userId,
            "receiver_id": selectedUser.id,
            "content": message
        }
    }

    ws!.send(JSON.stringify(newMessage))

    // if (!ws) {
    //     return sendError("Socket is not connected")
    // }
    // ws.send(JSON.stringify({ type: 'new_message', data: { Content: message, ReceiverID: selectedUser.id } }))
}
