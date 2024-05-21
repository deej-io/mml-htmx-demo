import { EditableNetworkedDOM } from "networked-dom-server";
import { ObservableDOM, JSDOMRunner } from "@mml-io/observable-dom";
import { WebSocketServer } from "ws";

const API_URL = process.env.API_URL ?? "http://localhost:8080/mml"

const document = new EditableNetworkedDOM(
  API_URL,
  (observableDOMParameters, callback) => {
    return new ObservableDOM(
      observableDOMParameters,
      callback,
      (htmlPath, htmlContents, params, callback) => {
        return new JSDOMRunner(htmlPath, htmlContents, params, callback, {
          // Configure the JSDOMRunner using this optional config
          allowResourceLoading: ["https://unpkg.com/htmx.org@1.9.12"],
        });
      },
    );
  },
);

// Could the library should probably load the body itself?
fetch(API_URL).then(res => res.text()).then(body => document.load(body));

const server = new WebSocketServer({
  port: 8081
});

server.on('connection', socket => {
  document.addWebSocket(socket) // FIXME: create a proper websocket abstraction
  socket.on('close', () => document.removeWebSocket(socket))
})

console.log(`mml-server: listening on port :${server.options.port}`)
