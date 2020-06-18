import { WsService } from './ws.service';
import { ServerJoinBody } from './wsModels/join';
import { GameMessage, MessageTypes } from './wsModels/gameMessage';

class EngineSack {
  constructor() {}

  // user and session
  uid: string;
  sid: string;

  // graphics and client-server communication
  gfx: any;
  wsSvc: WsService;
}

const engineSack: EngineSack = new EngineSack();

export function clientStart(wsService: WsService, canvas: any, sid: string) {
  // store globals
  engineSack.gfx = canvas;
  engineSack.wsSvc = wsService;
  engineSack.sid = sid;

  // connect
  wsService.connect(sid, (d, ws) => {
    if (d.type === MessageTypes.Join) {
      handleJoin(d);
    } else if (d.type === MessageTypes.Update) {
      test(d);
    }
  });
}

function handleJoin(d: GameMessage) {
  // parse body
  const msg = JSON.parse(d.body) as ServerJoinBody;

  // stash userid
  engineSack.uid = msg.uid;
}

function test(d: any) {
  // clear screen
  gfxBlank();

  // debug out
  const ctx = engineSack.gfx.getContext('2d');
  ctx.fillStyle = 'red';
  ctx.font = '8px Arial';
  ctx.fillText(JSON.stringify(d), 10, 50);
}

// clears the screen
function gfxBlank() {
  const ctx = engineSack.gfx.getContext('2d');
  ctx.fillStyle = 'pink';
  ctx.fillRect(0, 0, engineSack.gfx.width, engineSack.gfx.height);
}
