import { WsService } from './ws.service';
import { ServerJoinBody } from './wsModels/join';
import { GameMessage, MessageTypes } from './wsModels/gameMessage';
import { Player } from './engineModels/player';
import { System } from './engineModels/system';
import { Ship } from './engineModels/ship';
import { Camera } from './engineModels/camera';
import { ServerGlobalUpdateBody } from './wsModels/globalUpdate';

class EngineSack {
  constructor() {}

  // player
  player: Player;

  // camera
  camera: Camera;

  // graphics and client-server communication
  gfx: any;
  ctx: any;
  wsSvc: WsService;
}

const engineSack: EngineSack = new EngineSack();

export function clientStart(wsService: WsService, canvas: any, sid: string) {
  // initialize
  engineSack.player = new Player();
  engineSack.camera = new Camera(canvas.width, canvas.height);

  // store globals
  engineSack.gfx = canvas;
  engineSack.ctx = canvas.getContext('2d');
  engineSack.wsSvc = wsService;
  engineSack.player.sid = sid;

  // connect
  wsService.connect(sid, (d, ws) => {
    if (d.type === MessageTypes.Join) {
      handleJoin(d);
    } else if (d.type === MessageTypes.Update) {
      handleUpdate(d);
    }
  });
}

function handleJoin(d: GameMessage) {
  // parse body
  const msg = JSON.parse(d.body) as ServerJoinBody;

  // stash world welcome
  engineSack.player.uid = msg.currentShipInfo.uid;
  engineSack.player.currentShip = new Ship(msg.currentShipInfo);
  engineSack.player.currentSystem = new System(msg.currentSystemInfo);

  // start game loop
  setInterval(clientLoop, 20);
}

function handleUpdate(d: GameMessage) {
  // parse body
  const msg = JSON.parse(d.body) as ServerGlobalUpdateBody;

  // update system
  engineSack.player.currentSystem.id = msg.currentSystemInfo.id;
  engineSack.player.currentSystem.systemName = msg.currentSystemInfo.systemName;

  // update ships
  for (const sh of msg.ships) {
    let match = false;

    // find ship in memory
    for (const sm of engineSack.player.currentSystem.ships) {
      if (sh.id === sm.id) {
        match = true;

        // sync ship in memory
        sm.sync(sh);

        // is this the player ship?
        if (sm.id === engineSack.player.currentShip.id) {
          // update player ship cache
          engineSack.player.currentShip.sync(sh);

          // update camera position to track player ship
          engineSack.camera.x = sm.x;
          engineSack.camera.y = sm.y;
        }

        // exit loop
        break;
      }
    }

    if (!match) {
      // add ship to memory
      engineSack.player.currentSystem.ships.push(new Ship(sh));
    }

    // todo: handle ship leaving or dying
  }
}

function test(d: any) {
  // clear screen
  gfxBlank();

  // debug out
  engineSack.ctx.fillStyle = 'red';
  engineSack.ctx.font = '8px Arial';
  engineSack.ctx.fillText(JSON.stringify(d), 10, 50);
}

// clears the screen
function gfxBlank() {
  engineSack.ctx.fillStyle = 'pink';
  engineSack.ctx.fillRect(0, 0, engineSack.gfx.width, engineSack.gfx.height);
}

function clientLoop() {
  // render
  clientRender();
}

function clientRender() {
  // blank screen
  gfxBlank();

  // todo: draw system backplate

  // draw ships
  for (const sh of engineSack.player.currentSystem.ships) {
    sh.render(engineSack.ctx, engineSack.camera);
  }
}
