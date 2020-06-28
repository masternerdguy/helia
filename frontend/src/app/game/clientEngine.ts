import { WsService } from './ws.service';
import { ServerJoinBody } from './wsModels/join';
import { GameMessage, MessageTypes } from './wsModels/gameMessage';
import { Player } from './engineModels/player';
import { System } from './engineModels/system';
import { Ship } from './engineModels/ship';
import { Camera } from './engineModels/camera';
import { ServerGlobalUpdateBody } from './wsModels/globalUpdate';
import { Backplate } from './procedural/backplate/backplate';

class EngineSack {
  constructor() {}

  // player
  player: Player;

  // camera
  camera: Camera;

  // foreground graphics
  gfx: HTMLCanvasElement;
  ctx: any;

  // backplate graphics
  backplateCanvas: HTMLCanvasElement;
  backplateRenderer: Backplate;

  // client-server communication
  wsSvc: WsService;

  // tpf
  lastFrameTime: number;
  tpf: number;
}

const engineSack: EngineSack = new EngineSack();

export function clientStart(wsService: WsService, gameCanvas: HTMLCanvasElement, backCanvas: HTMLCanvasElement, sid: string) {
  // initialize
  engineSack.player = new Player();
  engineSack.camera = new Camera(gameCanvas.width, gameCanvas.height);
  engineSack.backplateRenderer = new Backplate(backCanvas);

  // store globals
  engineSack.gfx = gameCanvas;
  engineSack.ctx = gameCanvas.getContext('2d');
  engineSack.backplateCanvas = backCanvas;
  engineSack.wsSvc = wsService;
  engineSack.player.sid = sid;

  // connect
  wsService.connect(sid, (d, ws) => {
    if (d.type === MessageTypes.Join) {
      handleJoin(d);
    } else if (d.type === MessageTypes.Update) {
      handleGlobalUpdate(d);
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
  engineSack.lastFrameTime = Date.now();
  engineSack.tpf = 0;

  setInterval(clientLoop, 20);
}

function handleGlobalUpdate(d: GameMessage) {
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

// clears the screen
function gfxBlank() {
  engineSack.ctx.fillStyle = 'pink';
  engineSack.ctx.fillRect(0, 0, engineSack.gfx.width, engineSack.gfx.height);
}

// draws the backplate for the current system
function gfxBackplate() {
  if (!engineSack.player.currentSystem.backplateImg) {
    // render backplate
    engineSack.backplateRenderer.render(
      {
        renderPointStars: true,
        renderStars: true,
        renderSun: false,
        renderNebulae: true,
        shortScale: false,
        seed: engineSack.player.currentSystem.id // quick way to get a different plate for each system
      }
    );

    // get data url and convert to image
    engineSack.player.currentSystem.backplateImg = new Image();
    engineSack.player.currentSystem.backplateImg.src = engineSack.backplateCanvas.toDataURL('image/png');
  }

  // draw system backplate
  engineSack.ctx.drawImage(engineSack.player.currentSystem.backplateImg, 0, 0);
}

function clientLoop() {
  // render
  clientRender();

  // calculate time since last frame
  engineSack.tpf = Date.now() - engineSack.lastFrameTime;

  // store frame time
  engineSack.lastFrameTime = Date.now();
}

function clientRender() {
  // blank screen
  gfxBlank();

  // draw system backplate
  gfxBackplate();

  // draw ships
  for (const sh of engineSack.player.currentSystem.ships) {
    sh.render(engineSack.ctx, engineSack.camera);
  }
}
