import {server_socket} from "~/stores/socketCon"

import { messageUser } from "../../api/socket";

export default defineTask({
    async run() {
        server_socket?.onmessage("private_message", messageUser);
        return { result: 'ok' }
    }
})