import { WebSocketClient } from "~/server/utils/server_socket";
import { setServerSocket, server_socket } from "~/stores/socketCon";

export default defineTask({
    async run() {
        if (!server_socket) {
            let url = process.env.BACKEND_URL?.split('http')[1]

            let ws = new WebSocketClient(`ws${url}/socket?key=socket`);
            setServerSocket(ws)
        }
        return { result: 'ok' }
    }
})
