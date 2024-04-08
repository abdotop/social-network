import WebSocket from 'ws';
export class WebSocketClient {
    private socket: WebSocket | null = null;
    private readonly url: string;
    private readonly reconnectInterval: number;
    private seters: { [key: string]: (data: any) => void } = {}

    constructor(url: string, reconnectInterval = 5000) {
        this.url = url;
        this.reconnectInterval = reconnectInterval;
        this.connect();
    }

    private connect() {
        this.socket = new WebSocket(this.url);

        this.socket.onopen = (event: WebSocket.Event) => {
            console.log('WebSocket connection established');
            this.message();
        };

        this.socket.onerror = (error: WebSocket.Event) => {
            console.error('WebSocket error:', error);
        };

        this.socket.onclose = (event: WebSocket.CloseEvent) => {
            console.log('WebSocket connection closed. Reconnecting...');
            setTimeout(() => this.connect(), this.reconnectInterval);
        };
    }

    private message() {
        if (this.socket && this.socket.readyState === WebSocket.OPEN) {
            this.socket.onmessage = (event) => {

                try {
                    const data = JSON.parse(event.data.toString());
                                        
                    const type = data.type;

                    if (this.seters[type]) {
                        this.seters[type](data.data)
                    }
                } catch (error) {
                    console.error("Error parsing JSON:", error);
                    return;
                }
            };
        } else {
            console.error('Cannot set message handler, WebSocket is not open');
        }
    }
    public onmessage(type: string, seter: (data: any) => void) {
        this.seters[type] = seter
    }
    public send(data: any) {
        // console.log(data);
        
        if (this.socket && this.socket.readyState === WebSocket.OPEN) {
            this.socket.send(JSON.stringify(data));
        } else {
            console.error('Cannot send message, WebSocket is not open');
        }
    }
}