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
  skew: number = 0;
  metrics = {};

  constructor() {}

  connect(
    sid: string,
    handler: (data: GameMessage, s: WebSocketSubject<GameMessage>) => void
  ) {
    this.ws = webSocket({
      url: environment.wsUrl + 'connect',
      deserializer: (e: MessageEvent) => {
        return this.parse(e);
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

  private parse(e: MessageEvent<any>): any {
    // deobfuscate
    let z = this.obfuscate(e.data, this.getKey(this.skew));

    if (!z.startsWith('H4sIAAAAAAAA')) {
      let skew = -5;

      while (!z.startsWith('H4sIAAAAAAAA')) {
        skew++;
        z = this.obfuscate(e.data, this.getKey(skew));

        if (skew > 5) {
          // not parsable
          const stub = new GameMessage();
          stub.body = '';
          stub.type = -999;

          return stub;
        }
      }

      this.skew = skew;
    }

    // decompress
    const decompressed = pako.inflate(this.base64ToArrayBuffer(z), {
      to: 'string',
    });

    // parse
    const json = JSON.parse(decompressed);

    // record metrics
    var gm = json as GameMessage;

    if (!this.metrics[gm.type]) {
      this.metrics[gm.type] = 0;
    }

    this.metrics[gm.type] += z.length;

    // return result
    return json;
  }

  sendMessage(type: number, body: any) {
    // obfuscate
    const j = JSON.stringify(body);
    const x = this.obfuscate(j, this.getKey(this.skew));

    // package body as GameMessage
    const g = new GameMessage();
    g.body = x;
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

  // computes key accounting for potential skew
  getKey(offset: number): string {
    const now = new Date(Date.now());

    return `${
      now.getUTCMinutes() + offset
    }^${now.getUTCHours()}|${now.getUTCDate()}*${now.getUTCFullYear()}`;
  }

  // performs a xor on a string to obfuscate / deobfuscate it
  obfuscate(text: string, key: string) {
    var result = '';

    for (var i = 0; i < text.length; i++) {
      result += String.fromCharCode(
        text.charCodeAt(i) ^ key.charCodeAt(i % key.length)
      );
    }

    return result;
  }
}
