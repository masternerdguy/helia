import { Injectable } from '@angular/core';
import { WebSocketSubject, webSocket } from 'rxjs/webSocket';
import { environment } from 'src/environments/environment';
import { GameMessage, MessageTypes } from './wsModels/gameMessage';
import { ClientJoinBody } from './wsModels/bodies/join';
import * as pako from 'pako';

@Injectable({
  providedIn: 'root',
})
export class WsService {
  ws: WebSocketSubject<GameMessage>;
  lastMessageReceivedAt: number;
  sid: string;
  ackToken: number;

  constructor() {}

  connect(
    sid: string,
    handler: (data: GameMessage, s: WebSocketSubject<GameMessage>) => void
  ) {
    this.ws = webSocket({
      url: environment.wsUrl + 'connect',
      deserializer: (e: MessageEvent) => {
        const decompressed = pako.inflate(this.base64ToArrayBuffer(e.data), {
          to: 'string',
        });
        return JSON.parse(decompressed);
      },
      serializer: (value: GameMessage) => JSON.stringify(value),
    });

    this.ws.asObservable().subscribe((data) => {
      this.lastMessageReceivedAt = Date.now();
      handler(data, this.ws);
    });

    // send initial join message
    const b = new ClientJoinBody();
    b.sid = sid;
    this.sid = sid;

    this.sendMessage(MessageTypes.Join, b);
  }

  sendMessage(type: number, body: any) {
    // package body as GameMessage
    const g = new GameMessage();
    g.body = JSON.stringify(body);
    g.type = type;

    // send message to server
    this.ws.next(g);
  }

  isStale(): boolean {
    return Date.now() - this.lastMessageReceivedAt > 5000;
  }

  base64ToArrayBuffer(base64: string) {
    var bs = window.atob(base64);
    var len = bs.length;
    var bytes = new Uint8Array(len);

    for (var i = 0; i < len; i++) {
      bytes[i] = bs.charCodeAt(i);
    }

    return bytes.buffer;
  }
}
