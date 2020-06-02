import { Injectable } from '@angular/core';
import { WebSocketSubject, webSocket } from 'rxjs/webSocket';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class WsService {
  ws: WebSocketSubject<any>;

  constructor() { }

  connect(sid: string, handler: (data: any) => void) {
    this.ws = webSocket(environment.wsUrl + 'connect');
    this.ws.asObservable().subscribe(data => {
      handler(data);
    });

    this.ws.next({sid});
  }
}
