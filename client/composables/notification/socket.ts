import { addMessage } from "./message";
import { addNotif } from "./notification";

export let ws: WebSocket | undefined

export const connNotifSocket = async (id: string = "") => {
    console.log(location.host)
    const isSecure = location.protocol === "https:";
    const url = (isSecure ? "wss://" : "ws://") + location.host + `/api/socket?userId=${id}&channel=notif`;
    if (ws) {
        ws.close();
    }

    ws = new WebSocket(url);
    ws!.onopen = () => {
        // console.log("Connected to notif socket");
        ws!.addEventListener('message', (event) => {
            const notif = JSON.parse(event.data)
            if (notif.type === "new_message") {
                addMessage(notif)
            } else {
                addNotif(notif)
            }
        });
    };
}

// console.log('new_message');
