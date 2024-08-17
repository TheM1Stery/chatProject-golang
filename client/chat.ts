import { BACKEND_URL } from "./shared";

function createWebSocketConnection(url: string) {
    let ws = new WebSocket(url);

    let input = document.getElementById("input") as HTMLInputElement;
    let send = document.getElementById("send") as HTMLButtonElement;

    send.onclick = () => {
        ws.send(input.value);
        input.value = "";
    }


    ws.onmessage = (event) => {
        console.log(event.data);
    }
    ws.onclose = () => {
        console.log("Connection closed");
    }
}

createWebSocketConnection(`ws://${BACKEND_URL}/chat/connection`);

