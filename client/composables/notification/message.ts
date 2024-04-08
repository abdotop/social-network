import type { Notif } from "./notification";
import { useClearNotif } from "./notification"

export const messages = useState<Notif[]>(() => []);

export const addMessage = (notif: any) => {
    messages.value.push({
        id: messages.value.length + 1,
        notifId: notif.id,
        type: notif.type,
        message: notif.message,
        user: notif.user,
        created_at: new Date(notif.created_at),
    });
}


export const deleteMessages = (id: string) => {
    messages.value.forEach((n, index) => {
        if (n.notifId == id) {
            messages.value.splice(index, 1)
        }
    })
}

export const clearMessages = async (notif: Notif | undefined, action: string, message?: string) => {
    let type = 'clear'
    let notifId = notif?.notifId || ''
    if (action === 'all') {
        type = 'clear_all'
    }
    const res = await useClearNotif(notifId, type, "new_message");

    if (action === 'redirect' && res.message) {
        deleteMessages(notif!.notifId)
        navigateTo("/messages");
    } else if (action === 'all' && messages.value.length > 0 && res.message) {
        messages.value = []
        messages.value.push(
            {
                id: 1,
                notifId: '',
                type: "clear",
                message: "All messages have been cleared",
                user: 'accepted',
                created_at: new Date()
            }
        )
        setTimeout(() => {
            messages.value = []
        }, 5000);
    } else if (res.message) {
        messages.value[notif!.id - 1].message = message!
        messages.value[notif!.id - 1].type = "clear"
        setTimeout(() => {
            deleteMessages(notif!.notifId)
        }, 5000);
    }
}