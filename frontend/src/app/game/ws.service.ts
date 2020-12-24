import { Injectable } from '@angular/core';
import { WebSocketSubject, webSocket } from 'rxjs/webSocket';
import { environment } from 'src/environments/environment';
import { GameMessage, MessageTypes } from './wsModels/gameMessage';
import { ClientJoinBody } from './wsModels/bodies/join';

@Injectable({
  providedIn: 'root',
})
export class WsService {
  ws: WebSocketSubject<GameMessage>;
  lastMessageReceivedAt: number;
  sid: string;

  constructor() {}

  connect(
    sid: string,
    handler: (data: GameMessage, s: WebSocketSubject<GameMessage>) => void
  ) {
    this.ws = webSocket({
      url: environment.wsUrl + 'connect',
      deserializer: (e: MessageEvent) => JSON.parse(e.data),
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
}
