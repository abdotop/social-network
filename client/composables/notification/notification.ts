import { getNotifications } from "./getNotifications";
import { addMessage, messages } from "./message";

export const useClearNotif = async (id: string, type: string = "clear", action: string) => {
    const data = { type: type, id: id, action: action };
    const response = await fetch("/api/notification/clearnotification", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    }).then(async (res) => await res.json()).catch((err) => {
        return {
            status: 500,
            body: "Internal server error",
        };
    });
    return response;
}

export const notifications = useState<Notif[]>(() => []);

export type Notif = {
    id: number;
    notifId: string;
    type: string;
    message: string;
    user: any;
    created_at: Date;
    group_id?: string;
    member_id?: string;
    is_invite?: boolean;
};

export const addNotif = (notif: any) => {
    console.log(notif);
    
    removeLasNotifType(notif)
    notifications.value.push({
        id: notifications.value.length + 1,
        notifId: notif.id,
        type: notif.type,
        message: notif.message,
        user: notif.user,
        created_at: new Date(notif.created_at),
        group_id: notif.group_id,
        member_id: notif.member_id,
        is_invite: notif.is_invite
    });
}

const removeLasNotifType = (notif: any) => {
    notifications.value.forEach((n, index) => {
        if (((notif.type === "follow_request"
            || notif.type === "follow_accepted"
            || notif.type === "follow_declined"
            || notif.type === "unfollow")
            && (
                n.type === "follow_request"
                || n.type === "follow_accepted"
                || n.type === "follow_declined"
                || n.type === "unfollow"))
            && notif.user.id === n.user.id) {
            notifications.value.splice(index, 1)
        }
    })
}

export const deleteNotif = (id: string) => {
    notifications.value.forEach((n, index) => {
        if (n.notifId == id) {
            notifications.value.splice(index, 1)
        }
    })
}

export const clearNotif = async (notif: Notif | undefined, action: string, message?: string) => {
    let type = 'clear'
    let notifId = notif?.notifId || ''
    if (action === 'all') {
        type = 'clear_all'
    }
    const res = await useClearNotif(notifId, type, "");

    if (action === 'redirect') {
        deleteNotif(notif!.notifId)
        navigateTo("/profile/" + notif!.user.nickname);
    } else if (action === 'delete') {
        deleteNotif(notif!.notifId)
    }else if (action === 'all' && notifications.value.length > 0) {
        notifications.value = []
        notifications.value.push(
            {
                id: 1,
                notifId: '',
                type: "clear",
                message: "All notifications have been cleared",
                user: 'accepted',
                created_at: new Date()
            }
        )
        setTimeout(() => {
            notifications.value = []
        }, 5000);
    } else if (res.message) {
        notifications.value[notif!.id - 1].message = message!
        notifications.value[notif!.id - 1].type = "clear"
        setTimeout(() => {
            deleteNotif(notif!.notifId)
        }, 5000);
    }
}

async function fetchNotifications() {
    const notifs = await getNotifications();
    return notifs
}

if (!notifications.value.length) {
    fetchNotifications().then((notifs) => {
        if (!notifs.error && Array.isArray(notifs)) {
            Array.from(notifs).forEach((notif: any) => {
                if (notif.type === "new_message") {
                    addMessage(notif)
                } else {
                    addNotif(notif)
                }
            })
            messages.value.sort((a, b) => b.created_at.getTime() - a.created_at.getTime());
            notifications.value.sort((a, b) => b.created_at.getTime() - a.created_at.getTime());
        }
    });
}
