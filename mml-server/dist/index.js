var _a;
import { EditableNetworkedDOM, LocalObservableDOMFactory } from "networked-dom-server";
import * as Ws from 'ws';
var API_URL = (_a = process.env.API_URL) !== null && _a !== void 0 ? _a : "https://localhost:8080/";
var document = new EditableNetworkedDOM(API_URL, LocalObservableDOMFactory);
var server = new Ws.Server({
    port: 8081
});
server.on('connection', function (socket) {
    document.addWebSocket(socket); // FIXME: create a proper websocket abstraction
    socket.on('close', function () { return document.removeWebSocket(socket); });
});
