import { server_socket } from "~/stores/socketCon";
import { notifUser } from "~/server/api/socket"

export default defineTask({
    async run() {
        server_socket?.onmessage("notification", notifUser);
        return { result: 'ok' }
    }
})